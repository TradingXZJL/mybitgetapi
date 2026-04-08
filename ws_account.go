package mybitgetapi

const WS_CHANNEL_ACCOUNT = "account"

func accountSubscribeArg(instType InstType, coin string) WsSubscribeArg {
	if coin == "" {
		coin = "default"
	}
	return WsSubscribeArg{
		InstType: instType.String(),
		Channel:  WS_CHANNEL_ACCOUNT,
		Coin:     coin,
	}
}

// SubscribeAccount 订阅账户推送（私有 account 频道）。登录参见 WsStreamClient.Login。
// coin 目前只支持 "default"（传空会自动用 default）。
// 现货文档：https://www.bitget.com/zh-CN/api-doc/classic/spot/websocket/private/Account-Channel
// 合约文档：https://www.bitget.com/zh-CN/api-doc/classic/contract/websocket/private/Account-Channel
func (ws *PrivateWsStreamClient) SubscribeAccount(instType InstType, coin string) (*Subscription[WsAccount], error) {
	arg := accountSubscribeArg(instType, coin)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = instType
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsAccount]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsAccount, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}

	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.accountSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeAccount success: args=%v", doSub.Args)
	return sub, nil
}

// UnSubscribeAccount 取消订阅 account 频道。
func (ws *PrivateWsStreamClient) UnSubscribeAccount(instType InstType, coin string) error {
	arg := accountSubscribeArg(instType, coin)
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
		ws.accountSubMap.Delete(key)
	}

	log.Infof("UnSubscribeAccount success: args=%v", doSub.Args)
	return nil
}
