package mybitgetapi

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/robfig/cron/v3"
)

type RestProxy struct {
	ProxyUrl string //代理的协议IP端口URL
	Weight   ProxyWeight
}

type ProxyWeight struct {
	RemainWeight int  //剩余可用权重
	IsLimited    bool //是否已被限制
}

func (w *ProxyWeight) restore() {
	w.RemainWeight = 20
	w.IsLimited = false
}

var proxyList = []*RestProxy{}

var UseProxy = false
var WsUseProxy = false

func GetCurrentProxyList() []*RestProxy {
	return proxyList
}

func SetUseProxy(useProxy bool, proxyUrls ...string) {
	UseProxy = useProxy
	var newProxyList []*RestProxy
	for _, proxyUrl := range proxyUrls {
		newProxyList = append(newProxyList, &RestProxy{
			ProxyUrl: proxyUrl,
			Weight: ProxyWeight{
				RemainWeight: 20,
			},
		})
	}
	proxyList = newProxyList
}

func SetWsUseProxy(useProxy bool) error {
	if !UseProxy {
		return errors.New("please set UseProxy first")
	}
	WsUseProxy = useProxy
	return nil
}

func isUseProxy() bool {
	return UseProxy
}

func init() {
	c := cron.New(cron.WithSeconds())
	//每2秒权重清零，状态恢复
	_, err := c.AddFunc("*/2 * * * * *", func() {
		for _, proxy := range proxyList {
			proxy.Weight.restore()
		}
	})
	if err != nil {
		log.Error(err)
		return
	}
	c.Start()
}

// 获取最佳代理
func getBestProxyAndWeight() (*RestProxy, *ProxyWeight) {
	var maxWeightProxy *RestProxy
	var maxWeight *ProxyWeight
	for _, proxy := range proxyList {
		proxyWeight := &proxy.Weight
		if proxyWeight.IsLimited {
			continue
		}
		if maxWeightProxy == nil {
			maxWeightProxy = proxy
			maxWeight = proxyWeight
			continue
		}
		if proxyWeight.RemainWeight > maxWeight.RemainWeight {
			maxWeightProxy = proxy
			maxWeight = proxyWeight
		}
	}

	return maxWeightProxy, maxWeight
}

// 获取随机代理
func getRandomProxy() (*RestProxy, error) {
	length := len(proxyList)
	if length == 0 {
		return nil, errors.New("proxyList is empty")
	}

	return proxyList[rand.Intn(length)], nil
}

// Request 发送请求
func Request(url string, reqBody []byte, method string, isGzip bool) ([]byte, map[string]string, error) {
	return RequestWithHeader(url, reqBody, method, map[string]string{}, isGzip)
}

func RequestWithHeader(urlStr string, reqBody []byte, method string, headerMap map[string]string, isGzip bool) ([]byte, map[string]string, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}
	log.Debug("header: ", req.Header)
	log.Debug("reqURL: ", req.URL.String())
	if len(reqBody) > 0 {
		log.Debug("reqBody: ", string(reqBody))
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}

	var currentProxy *RestProxy
	var currentProxyWeight *ProxyWeight
	if UseProxy {
		currentProxy, currentProxyWeight = getBestProxyAndWeight()
		if currentProxy == nil || currentProxyWeight.RemainWeight <= 0 {
			currentProxyWeight.IsLimited = true
			return nil, nil, errors.New("all proxy ip weight limit reached")
		}

		url_i := url.URL{}
		bestProxy, _ := url_i.Parse(currentProxy.ProxyUrl)

		reqProxy := &http.Transport{}
		reqProxy.Proxy = http.ProxyURL(bestProxy)                        // set proxy
		reqProxy.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //set ssl

		client.Transport = reqProxy
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, nil, err
		}
	}
	data, err := io.ReadAll(body)
	// log.Debug(string(data))
	log.Debug(string(data))
	log.Debug(resp.Header)
	if isUseProxy() {
		if resp.StatusCode == http.StatusForbidden {
			currentProxyWeight.RemainWeight = 0
			currentProxyWeight.IsLimited = true
			return data, nil, errors.New("proxy ip is limited")
		}
		// 限频：请求过于频繁会返回 429 Too Many Requests
		// 文档：公共接口 1 秒最多 20 请求；授权接口按 apikey 维度限频。
		// 这里将 429 视为“当前代理暂时不可用”，交由上层重试/切换代理。
		if resp.StatusCode == http.StatusTooManyRequests {
			currentProxyWeight.RemainWeight = 0
			currentProxyWeight.IsLimited = true
			return data, nil, errors.New("too many requests (429)")
		}
		// 回填权重（兼容不同交易所/网关头）
		// - Binance 风格：X-Mbx-Used-Remain-Limit: 19 / [19]
		// - Gate 风格：X-Gate-RateLimit-Requests-Remain: 123
		headerToRemain := func(v string) (int, error) {
			v = strings.TrimSpace(v)
			v = strings.TrimPrefix(v, "[")
			v = strings.TrimSuffix(v, "]")
			return strconv.Atoi(v)
		}

		remainHeader := resp.Header.Get("X-Mbx-Used-Remain-Limit")
		if remainHeader != "" {
			remainWeight, err := headerToRemain(remainHeader)
			if err != nil {
				log.Error(err)
			} else if remainWeight < currentProxyWeight.RemainWeight {
				currentProxyWeight.RemainWeight = remainWeight
			}
		}

		// 回填是否接口权重已用完
		if currentProxyWeight.RemainWeight <= 0 {
			currentProxyWeight.IsLimited = true
		}
	}
	return data, headersToMapSingleValue(resp.Header), err
}

func headersToMapSingleValue(header http.Header) map[string]string {
	result := make(map[string]string)
	for key, values := range header {
		if len(values) > 0 {
			result[key] = values[0] // 只取第一个值
		}
	}
	return result
}
