package mybitgetapi

const WS_CHANNEL_FILL = "fill"

func fillSubscribeArg(instType InstType, instId string) WsSubscribeArg {
	if instId == "" {
		instId = "default"
	}
	return WsSubscribeArg{
		InstType: instType.String(),
		Channel:  WS_CHANNEL_FILL,
		InstId:   instId,
	}
}

// SubscribeFill 订阅成交明细（私有 fill 频道）。登录参见 WsStreamClient.Login。
// instId 传空等价于 "default"（订阅该 instType 下的全市场）。
// 现货文档：https://www.bitget.com/zh-CN/api-doc/classic/spot/websocket/private/Fill-Channel
// 合约文档：https://www.bitget.com/zh-CN/api-doc/classic/contract/websocket/private/Fill-Channel
func (ws *PrivateWsStreamClient) SubscribeFill(instType InstType, instId string) (*Subscription[WsFill], error) {
	arg := fillSubscribeArg(instType, instId)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = instType
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsFill]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsFill, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.fillSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeFill success: args=%v", doSub.Args)
	return sub, nil
}

// UnSubscribeFill 取消订阅 fill 频道。
func (ws *PrivateWsStreamClient) UnSubscribeFill(instType InstType, instId string) error {
	arg := fillSubscribeArg(instType, instId)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = instType
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return err
	}

	ws.sendUnSubscribeSuccessToCloseChan(args)
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		key := string(keyData)
		ws.commonSubMap.Delete(key)
		ws.fillSubMap.Delete(key)
	}
	log.Infof("UnSubscribeFill success: args=%v", doSub.Args)
	return nil
}
