package mybitgetapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"time"

	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	//BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
	PATCH  = "PATCH"
)

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

var httpTimeout = 100 * time.Second

func SetHttpTimeout(timeout time.Duration) {
	httpTimeout = timeout
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha512(secret, data string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func HmacSha512Base64(secret, data string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func HmacSha256Base64(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

type MyBitget struct {
}

const (
	BITGET_API_HTTP = "api.bitget.com"
	IS_GZIP         = false
)

type NetType int

const (
	MAIN_NET NetType = iota
	TEST_NET
)

var NowNetType = MAIN_NET

func SetNetType(netType NetType) {
	NowNetType = netType
}

type APIType int

const (
	REST APIType = iota
	WS_PUBLIC
	WS_CLASSIC
	WS_UTA
)

type Client struct {
	ApiKey     string
	ApiSecret  string
	Passphrase string
}
type RestClient struct {
	c *Client
}
type PublicRestClient RestClient

type PrivateRestClient RestClient

func NewRestClient(apiKey string, apiSecret string, passphrase string) *RestClient {
	client := &RestClient{
		&Client{
			ApiKey:     apiKey,
			ApiSecret:  apiSecret,
			Passphrase: passphrase,
		},
	}
	return client
}

func (c *RestClient) PublicRestClient() *PublicRestClient {
	return &PublicRestClient{
		c: c.c,
	}
}

func (c *RestClient) PrivateRestClient() *PrivateRestClient {
	return &PrivateRestClient{
		c: c.c,
	}
}

type InstType string

const (
	INST_TYPE_SPOT         InstType = "SPOT"         // 现货
	INST_TYPE_MARGIN       InstType = "MARGIN"       // 杠杆
	INST_TYPE_USDT_FUTURES InstType = "USDT-FUTURES" // U本位合约
	INST_TYPE_COIN_FUTURES InstType = "COIN-FUTURES" // 币本位合约
	INST_TYPE_USDC_FUTURES InstType = "USDC-FUTURES" // USDC合约
	INST_TYPE_UTA          InstType = "UTA"          // 统一账户
)

func (wsInstType InstType) String() string {
	return string(wsInstType)
}

type MarginMode string

const (
	MARGIN_MODE_ISOLATED MarginMode = "isolated" // 逐仓
	MARGIN_MODE_CROSSED  MarginMode = "crossed"  // 全仓
)

func (marginMode MarginMode) String() string {
	return string(marginMode)
}

type PosMode string

const (
	POS_MODE_ONE_WAY PosMode = "one_way_mode" // 单向持仓
	POS_MODE_HEDGE   PosMode = "hedge_mode"   // 双向持仓
)

func (posMode PosMode) String() string {
	return string(posMode)
}

var serverTimeDelta int64 = 0

func SetServerTimeDelta(delta int64) {
	serverTimeDelta = delta
}

// 通用接口调用
func bitgetCallAPI[T any](client *Client, url url.URL, reqBody []byte, method string) (*BitgetRestRes[T], error) {
	body, respHeaderMap, err := Request(url.String(), reqBody, method, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body, respHeaderMap)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// 通用鉴权接口调用
func bitgetCallApiWithSecret[T any](client *Client, url url.URL, reqBody []byte, method string) (*BitgetRestRes[T], error) {
	timestamp := time.Now().UnixMilli() //毫秒级时间戳
	method = strings.ToUpper(method)
	requestPath := url.Path
	query := url.RawQuery
	timestampStr := strconv.FormatInt(timestamp, BIT_BASE_10)
	bodyStr := string(reqBody)

	// timestamp + method.toUpperCase() + requestPath + ("?" + queryString) + body
	var signContent strings.Builder
	signContent.WriteString(timestampStr)
	signContent.WriteString(method)
	signContent.WriteString(requestPath)
	if query != "" {
		signContent.WriteString("?")
		signContent.WriteString(query)
	}
	signContent.WriteString(bodyStr)
	signData := signContent.String()
	sign := HmacSha256Base64(client.ApiSecret, signData)

	// log.Warn("timestamp: ", timestamp)
	// log.Warn("method: ", method)
	// log.Warn("requestPath: ", requestPath)
	// log.Warn("query: ", query)
	// log.Warn("reqBody: ", bodyStr)
	// log.Warn("signData: ", signData)
	// log.Warn("sign: ", sign)

	body, respHeaderMap, err := RequestWithHeader(url.String(), reqBody, method,
		map[string]string{
			"ACCESS-KEY":        client.ApiKey,
			"ACCESS-SIGN":       sign,
			"ACCESS-TIMESTAMP":  timestampStr,
			"ACCESS-PASSPHRASE": client.Passphrase,
			"locale":            "zh-CN",
		}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body, respHeaderMap)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// URL标准封装 带路径参数
func bitgetHandlerRequestAPIWithPathQueryParam[T any](apiType APIType, request *T, name string) url.URL {
	query := bitgetHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     BITGET_API_HTTP,
		Path:     name,
		RawQuery: query,
	}
	return u
}

// URL标准封装 不带路径参数
func bitgetHandlerRequestAPIWithoutPathQueryParam(apiType APIType, name string) url.URL {
	u := url.URL{
		Scheme:   "https",
		Host:     BITGET_API_HTTP,
		Path:     name,
		RawQuery: "",
	}
	return u
}

func bitgetHandlerReq[T any](req *T) string {
	var argBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			args := make([]reflect.Value, 0)
			result := ToStringMethod.Call(args)
			argBuffer.WriteString(argName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(argBuffer.String(), "&")
}

func GateGetRestHostByAPIType(apiType APIType) string {
	switch apiType {
	case REST:
		return BITGET_API_HTTP
	}

	return ""

}
