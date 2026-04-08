package mybitgetapi

const WS_CHANNEL_ORDERS = "orders"

func ordersSubscribeArg(instType InstType, instId string) WsSubscribeArg {
	if instId == "" {
		instId = "default"
	}
	return WsSubscribeArg{
		InstType: instType.String(),
		Channel:  WS_CHANNEL_ORDERS,
		InstId:   instId,
	}
}

// SubscribeOrders 订阅订单推送（私有 orders 频道）。登录参见 WsStreamClient.Login。
// instId 传空等价于 "default"（全交易对）。现货文档：https://www.bitget.com/zh-CN/api-doc/classic/spot/websocket/private/Order-Channel
func (ws *PrivateWsStreamClient) SubscribeOrders(instType InstType, instId string) (*Subscription[WsOrder], error) {
	arg := ordersSubscribeArg(instType, instId)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = instType
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsOrder]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsOrder, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.orderSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeOrders success: args=%v", doSub.Args)
	return sub, nil
}

// UnSubscribeOrders 取消订阅 orders 频道。
func (ws *PrivateWsStreamClient) UnSubscribeOrders(instType InstType, instId string) error {
	arg := ordersSubscribeArg(instType, instId)
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
		ws.orderSubMap.Delete(key)
	}
	log.Infof("UnSubscribeOrders success: args=%v", doSub.Args)
	return nil
}

