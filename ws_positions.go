package mybitgetapi

import "fmt"

const WS_CHANNEL_POSITIONS = "positions"

func positionsSubscribeArg(instType InstType) WsSubscribeArg {
	return WsSubscribeArg{
		InstType: instType.String(),
		Channel:  WS_CHANNEL_POSITIONS,
		InstId:   "default",
	}
}

// SubscribePositions 订阅合约持仓频道（私有 positions 频道）。登录参见 WsStreamClient.Login。
// 文档：https://www.bitget.com/zh-CN/api-doc/classic/contract/websocket/private/Positions-Channel
func (ws *PrivateWsStreamClient) SubscribePositions(instType InstType) (*Subscription[WsPositions], error) {
	switch instType {
	case INST_TYPE_USDT_FUTURES, INST_TYPE_COIN_FUTURES, INST_TYPE_USDC_FUTURES:
	default:
		return nil, fmt.Errorf("invalid futures instType: %s", instType)
	}

	arg := positionsSubscribeArg(instType)
	args := []WsSubscribeArg{arg}

	ws.WsStreamClient.instType = instType
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsPositions]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsPositions, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	
	for _, a := range args {
		keyData, _ := json.Marshal(&a)
		ws.positionsSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribePositions success: args=%v", doSub.Args)
	return sub, nil
}

// UnSubscribePositions 取消订阅合约持仓频道。
func (ws *PrivateWsStreamClient) UnSubscribePositions(instType InstType) error {
	switch instType {
	case INST_TYPE_USDT_FUTURES, INST_TYPE_COIN_FUTURES, INST_TYPE_USDC_FUTURES:
	default:
		return fmt.Errorf("invalid futures instType: %s", instType)
	}

	arg := positionsSubscribeArg(instType)
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
		ws.positionsSubMap.Delete(key)
	}
	log.Infof("UnSubscribePositions success: args=%v", doSub.Args)
	return nil
}
