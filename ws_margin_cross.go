package mybitgetapi

const (
	WS_CHANNEL_MARGIN_CROSS_ACCOUNT = "account-crossed"
	WS_CHANNEL_MARGIN_CROSS_ORDERS  = "orders-crossed"
)

func marginCrossAccountSubscribeArg(coin string) WsSubscribeArg {
	if coin == "" {
		coin = "default"
	}
	return WsSubscribeArg{
		InstType: INST_TYPE_MARGIN.String(),
		Channel:  WS_CHANNEL_MARGIN_CROSS_ACCOUNT,
		Coin:     coin,
	}
}

func marginCrossOrdersSubscribeArg(instId string) WsSubscribeArg {
	return WsSubscribeArg{
		InstType: INST_TYPE_MARGIN.String(),
		Channel:  WS_CHANNEL_MARGIN_CROSS_ORDERS,
		InstId:   instId,
	}
}

// SubscribeMarginCrossAccountAssets 订阅全仓杠杆账户频道（account-crossed）。
// coin 目前只支持 default（传空会自动用 default）。
// 文档：https://www.bitget.com/zh-CN/api-doc/classic/margin/cross/websocket/private/Margin-Cross-Account-Assets
func (ws *PrivateWsStreamClient) SubscribeMarginCrossAccountAssets(coin string) (*Subscription[WsMarginCrossAccountAssets], error) {
	arg := marginCrossAccountSubscribeArg(coin)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = INST_TYPE_MARGIN
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsMarginCrossAccountAssets]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsMarginCrossAccountAssets, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.marginCrossAccountSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeMarginCrossAccountAssets success: args=%v", doSub.Args)
	return sub, nil
}

func (ws *PrivateWsStreamClient) UnSubscribeMarginCrossAccountAssets(coin string) error {
	arg := marginCrossAccountSubscribeArg(coin)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = INST_TYPE_MARGIN
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return err
	}

	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		key := string(keyData)
		ws.commonSubMap.Delete(key)
		ws.marginCrossAccountSubMap.Delete(key)
	}
	log.Infof("UnSubscribeMarginCrossAccountAssets success: args=%v", doSub.Args)
	return nil
}

// SubscribeMarginCrossOrders 订阅全仓杠杆订单频道（orders-crossed）。
// 文档：https://www.bitget.com/zh-CN/api-doc/classic/margin/cross/websocket/private/Cross-Orders
func (ws *PrivateWsStreamClient) SubscribeMarginCrossOrders(instId string) (*Subscription[WsMarginCrossOrders], error) {
	arg := marginCrossOrdersSubscribeArg(instId)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = INST_TYPE_MARGIN
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsMarginCrossOrders]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsMarginCrossOrders, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.marginCrossOrdersSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeMarginCrossOrders success: args=%v", doSub.Args)
	return sub, nil
}

func (ws *PrivateWsStreamClient) UnSubscribeMarginCrossOrders(instId string) error {
	arg := marginCrossOrdersSubscribeArg(instId)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = INST_TYPE_MARGIN
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
		ws.marginCrossOrdersSubMap.Delete(key)
	}
	log.Infof("UnSubscribeMarginCrossOrders success: args=%v", doSub.Args)
	return nil
}

