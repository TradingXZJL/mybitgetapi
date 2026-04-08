package mybitgetapi

import (
	"fmt"
)

type BitgetErrorRes struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}

type BitgetTimeRes struct {
	AcceptTime   string `json:"acceptTime"`   //BG服务端接受到请求的时间
	CompleteTime string `json:"completeTime"` //BG服务端请求处理完成时间
}
type BitgetRestRes[T any] struct {
	BitgetErrorRes   //错误信息
	BitgetTimeRes    //时间戳
	Data           T `json:"data"` //请求结果
}

func handlerCommonRest[T any](data []byte, respHeaderMap map[string]string) (*BitgetRestRes[T], error) {
	res := &BitgetRestRes[T]{}
	log.Debug(string(data))

	//捕获网关出入站时间信息 结果为微秒时间戳
	if acceptTimeStr, ok := respHeaderMap["X-BG-REQUEST-ACCEPT-TIME"]; ok {
		res.AcceptTime = acceptTimeStr
	}
	if completeTimeStr, ok := respHeaderMap["X-BG-RESPONSE-COMPLETE-TIME"]; ok {
		res.CompleteTime = completeTimeStr
	}

	err := json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, res.handlerError()
}

func (err *BitgetErrorRes) handlerError() error {
	if err.Code != "00000" {
		//不为2xx的返回代码代表请求失败
		return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
	} else {
		return nil
	}
}
