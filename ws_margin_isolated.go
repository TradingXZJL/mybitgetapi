package mybitgetapi

const (
	WS_CHANNEL_MARGIN_ISOLATED_ACCOUNT = "account-isolated"
	WS_CHANNEL_MARGIN_ISOLATED_ORDERS  = "orders-isolated"
)

func marginIsolatedAccountSubscribeArg(coin string) WsSubscribeArg {
	if coin == "" {
		coin = "default"
	}
	return WsSubscribeArg{
		InstType: INST_TYPE_MARGIN.String(),
		Channel:  WS_CHANNEL_MARGIN_ISOLATED_ACCOUNT,
		Coin:     coin,
	}
}

func marginIsolatedOrdersSubscribeArg(instId string) WsSubscribeArg {
	return WsSubscribeArg{
		InstType: INST_TYPE_MARGIN.String(),
		Channel:  WS_CHANNEL_MARGIN_ISOLATED_ORDERS,
		InstId:   instId,
	}
}

// SubscribeMarginIsolatedAccountAssets 订阅逐仓杠杆账户频道（account-isolated）。
// coin 目前只支持 default（传空会自动用 default）。
// 文档：https://www.bitget.com/zh-CN/api-doc/classic/margin/isolated/websocket/private/Margin-isolated-account-assets
func (ws *PrivateWsStreamClient) SubscribeMarginIsolatedAccountAssets(coin string) (*Subscription[WsMarginIsolatedAccountAssets], error) {
	arg := marginIsolatedAccountSubscribeArg(coin)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = INST_TYPE_MARGIN
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsMarginIsolatedAccountAssets]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsMarginIsolatedAccountAssets, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.marginIsolatedAccountSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeMarginIsolatedAccountAssets success: args=%v", doSub.Args)
	return sub, nil
}

func (ws *PrivateWsStreamClient) UnSubscribeMarginIsolatedAccountAssets(coin string) error {
	arg := marginIsolatedAccountSubscribeArg(coin)
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
		ws.marginIsolatedAccountSubMap.Delete(key)
	}
	log.Infof("UnSubscribeMarginIsolatedAccountAssets success: args=%v", doSub.Args)
	return nil
}

// SubscribeMarginIsolatedOrders 订阅逐仓杠杆订单频道（orders-isolated）。
// 文档：https://www.bitget.com/zh-CN/api-doc/classic/margin/isolated/websocket/private/Isolate-Orders
func (ws *PrivateWsStreamClient) SubscribeMarginIsolatedOrders(instId string) (*Subscription[WsMarginIsolatedOrders], error) {
	arg := marginIsolatedOrdersSubscribeArg(instId)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = INST_TYPE_MARGIN
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsMarginIsolatedOrders]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsMarginIsolatedOrders, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.marginIsolatedOrdersSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeMarginIsolatedOrders success: args=%v", doSub.Args)
	return sub, nil
}

func (ws *PrivateWsStreamClient) UnSubscribeMarginIsolatedOrders(instId string) error {
	arg := marginIsolatedOrdersSubscribeArg(instId)
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
		ws.marginIsolatedOrdersSubMap.Delete(key)
	}
	log.Infof("UnSubscribeMarginIsolatedOrders success: args=%v", doSub.Args)
	return nil
}

