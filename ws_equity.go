package mybitgetapi

const WS_CHANNEL_EQUITY = "equity"

func equitySubscribeArg(instType InstType) WsSubscribeArg {
	return WsSubscribeArg{
		InstType: instType.String(),
		Channel:  WS_CHANNEL_EQUITY,
		InstId:   "default",
	}
}

// SubscribeEquity 订阅合约权益频道（私有 equity 频道）。登录参见 WsStreamClient.Login。
// 文档：https://www.bitget.com/zh-CN/api-doc/classic/contract/websocket/private/Equity-Channel
func (ws *PrivateWsStreamClient) SubscribeEquity(instType InstType) (*Subscription[WsEquity], error) {
	arg := equitySubscribeArg(instType)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = instType
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsEquity]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsEquity, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.equitySubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeEquity success: args=%v", doSub.Args)
	return sub, nil
}

// UnSubscribeEquity 取消订阅合约权益频道。
func (ws *PrivateWsStreamClient) UnSubscribeEquity(instType InstType) error {
	arg := equitySubscribeArg(instType)
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
		ws.equitySubMap.Delete(key)
	}
	log.Infof("UnSubscribeEquity success: args=%v", doSub.Args)
	return nil
}
