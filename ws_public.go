package mybitgetapi

import (
	"fmt"
	"strings"
)

const (
	WS_TOPIC_TICKER       = "ticker"
	WS_TOPIC_KLINE        = "kline"
	WS_TOPIC_PUBLIC_TRADE = "publicTrade"
)

type WsBooksType string

const (
	WS_BOOKS_ALL WsBooksType = "books"
	WS_BOOKS_1   WsBooksType = "books1"
	WS_BOOKS_5   WsBooksType = "books5"
	WS_BOOKS_50  WsBooksType = "books50"
)

func (wsBooksType WsBooksType) String() string {
	return string(wsBooksType)
}

type WsCandleInterval string

const (
	WS_CANDLE_1M  WsCandleInterval = "1m"
	WS_CANDLE_3M  WsCandleInterval = "3m"
	WS_CANDLE_5M  WsCandleInterval = "5m"
	WS_CANDLE_15M WsCandleInterval = "15m"
	WS_CANDLE_30M WsCandleInterval = "30m"
	WS_CANDLE_1H  WsCandleInterval = "1H"
	WS_CANDLE_4H  WsCandleInterval = "4H"
	WS_CANDLE_6H  WsCandleInterval = "6H"
	WS_CANDLE_12H WsCandleInterval = "12H"
	WS_CANDLE_1D  WsCandleInterval = "1D"
)

func (wsCandleInterval WsCandleInterval) String() string {
	return string(wsCandleInterval)
}

// 批量订阅行情频道。
// instType: spot/usdt-futures/coin-futures/usdc-futures
func (ws *PublicWsStreamClient) SubscribeTickersMultiple(instType InstType, symbols []string) (*Subscription[WsTicker], error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("symbols is empty")
	}

	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    WS_TOPIC_TICKER,
			Symbol:   symbol,
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsTicker]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsTicker, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}

	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.tickerSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeTickers success: args=%v", doSub.Args)
	return sub, nil
}

// 订阅单个行情频道。
func (ws *PublicWsStreamClient) SubscribeTickers(instType InstType, symbol string) (*Subscription[WsTicker], error) {
	return ws.SubscribeTickersMultiple(instType, []string{symbol})
}

// 批量取消订阅行情频道。
func (ws *PublicWsStreamClient) UnSubscribeTickersMultiple(instType InstType, symbols []string) error {
	if len(symbols) == 0 {
		return fmt.Errorf("symbols is empty")
	}

	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    WS_TOPIC_TICKER,
			Symbol:   symbol,
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return err
	}

	ws.sendUnSubscribeSuccessToCloseChan(args)
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.commonSubMap.Delete(string(keyData))
		ws.tickerSubMap.Delete(string(keyData))
	}
	log.Infof("UnSubscribeTickers success: args=%v", doSub.Args)
	return nil
}

// 取消订阅单个行情频道。
func (ws *PublicWsStreamClient) UnSubscribeTickers(instType InstType, symbol string) error {
	return ws.UnSubscribeTickersMultiple(instType, []string{symbol})
}

// 批量订阅 K线频道。
func (ws *PublicWsStreamClient) SubscribeCandlesMultiple(instType InstType, symbols []string, interval WsCandleInterval) (*Subscription[WsCandle], error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("symbols is empty")
	}
	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    WS_TOPIC_KLINE,
			Symbol:   symbol,
			Interval: interval.String(),
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsCandle]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsCandle, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.candleSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeCandles success: args=%v", doSub.Args)
	return sub, nil
}

// 订阅单个 K线频道。
func (ws *PublicWsStreamClient) SubscribeCandles(instType InstType, symbol string, interval WsCandleInterval) (*Subscription[WsCandle], error) {
	return ws.SubscribeCandlesMultiple(instType, []string{symbol}, interval)
}

// 批量取消订阅 K线频道。
func (ws *PublicWsStreamClient) UnSubscribeCandlesMultiple(instType InstType, symbols []string, interval WsCandleInterval) error {
	if len(symbols) == 0 {
		return fmt.Errorf("symbols is empty")
	}
	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    WS_TOPIC_KLINE,
			Symbol:   symbol,
			Interval: interval.String(),
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return err
	}

	ws.sendUnSubscribeSuccessToCloseChan(args)
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.commonSubMap.Delete(string(keyData))
		ws.candleSubMap.Delete(string(keyData))
	}
	log.Infof("UnSubscribeCandles success: args=%v", doSub.Args)
	return nil
}

// 取消订阅单个 K线频道。
func (ws *PublicWsStreamClient) UnSubscribeCandles(instType InstType, symbol string, interval WsCandleInterval) error {
	return ws.UnSubscribeCandlesMultiple(instType, []string{symbol}, interval)
}

// 批量订阅深度频道。
func (ws *PublicWsStreamClient) SubscribeBooksMultiple(instType InstType, symbols []string, booksType WsBooksType) (*Subscription[WsBooks], error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("symbols is empty")
	}
	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    booksType.String(),
			Symbol:   symbol,
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsBooks]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsBooks, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.booksSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribeBooks success: args=%v", doSub.Args)
	return sub, nil
}

// 订阅单个深度频道。
func (ws *PublicWsStreamClient) SubscribeBooks(instType InstType, symbol string, booksType WsBooksType) (*Subscription[WsBooks], error) {
	return ws.SubscribeBooksMultiple(instType, []string{symbol}, booksType)
}

// 批量取消订阅深度频道。
func (ws *PublicWsStreamClient) UnSubscribeBooksMultiple(instType InstType, symbols []string, booksType WsBooksType) error {
	if len(symbols) == 0 {
		return fmt.Errorf("symbols is empty")
	}
	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    booksType.String(),
			Symbol:   symbol,
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return err
	}

	ws.sendUnSubscribeSuccessToCloseChan(args)
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.commonSubMap.Delete(string(keyData))
		ws.booksSubMap.Delete(string(keyData))
	}
	log.Infof("UnSubscribeBooks success: args=%v", doSub.Args)
	return nil
}

// 取消订阅单个深度频道。
func (ws *PublicWsStreamClient) UnSubscribeBooks(instType InstType, symbol string, booksType WsBooksType) error {
	return ws.UnSubscribeBooksMultiple(instType, []string{symbol}, booksType)
}

// 批量订阅平台成交频道。
func (ws *PublicWsStreamClient) SubscribePublicTradesMultiple(instType InstType, symbols []string) (*Subscription[WsPublicTrade], error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("symbols is empty")
	}
	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    WS_TOPIC_PUBLIC_TRADE,
			Symbol:   symbol,
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return nil, err
	}

	sub := &Subscription[WsPublicTrade]{
		SubId:      doSub.SubId,
		Op:         WS_SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsPublicTrade, 100),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.tradesSubMap.Store(string(keyData), sub)
	}
	log.Infof("SubscribePublicTrades success: args=%v", doSub.Args)
	return sub, nil
}

// 订阅单个平台成交频道。
func (ws *PublicWsStreamClient) SubscribePublicTrades(instType InstType, symbol string) (*Subscription[WsPublicTrade], error) {
	return ws.SubscribePublicTradesMultiple(instType, []string{symbol})
}

// 批量取消订阅平台成交频道。
func (ws *PublicWsStreamClient) UnSubscribePublicTradesMultiple(instType InstType, symbols []string) error {
	if len(symbols) == 0 {
		return fmt.Errorf("symbols is empty")
	}
	args := make([]WsSubscribeArg, 0, len(symbols))
	for _, symbol := range symbols {
		args = append(args, WsSubscribeArg{
			InstType: strings.ToLower(instType.String()), // uta ws 接口的 instType 需要小写
			Topic:    WS_TOPIC_PUBLIC_TRADE,
			Symbol:   symbol,
		})
	}

	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, WS_UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	if err = ws.catchSubscribeResult(doSub); err != nil {
		return err
	}

	ws.sendUnSubscribeSuccessToCloseChan(args)
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.commonSubMap.Delete(string(keyData))
		ws.tradesSubMap.Delete(string(keyData))
	}
	log.Infof("UnSubscribePublicTrades success: args=%v", doSub.Args)
	return nil
}

// 取消订阅单个平台成交频道。
func (ws *PublicWsStreamClient) UnSubscribePublicTrades(instType InstType, symbol string) error {
	return ws.UnSubscribePublicTradesMultiple(instType, []string{symbol})
}
