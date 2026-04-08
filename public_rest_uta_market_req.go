package mybitgetapi

type PublicRestUtaMarketInstrumentsAPI struct {
	client *PublicRestClient
	req    *PublicRestUtaMarketInstrumentsReq
}

type PublicRestUtaMarketInstrumentsReq struct {
	Category *string `json:"category"` // String 是 产品类型 SPOT 现货交易 MARGIN 杠杆 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol   *string `json:"symbol"`   // String 否 交易对名称 例如：BTCUSDT
}

// String 是 产品类型 SPOT 现货交易 MARGIN 杠杆 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
func (api *PublicRestUtaMarketInstrumentsAPI) Category(category string) *PublicRestUtaMarketInstrumentsAPI {
	api.req.Category = GetPointer(category)
	return api
}

// String 否 交易对名称 例如：BTCUSDT
func (api *PublicRestUtaMarketInstrumentsAPI) Symbol(symbol string) *PublicRestUtaMarketInstrumentsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PublicRestUtaMarketTickersAPI struct {
	client *PublicRestClient
	req    *PublicRestUtaMarketTickersReq
}

type PublicRestUtaMarketTickersReq struct {
	Category *string `json:"category"` // String 是 产品类型 SPOT 现货交易 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol   *string `json:"symbol"`   // String 否 交易对名称 例如：BTCUSDT
}

// String 是 产品类型 SPOT 现货交易 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
func (api *PublicRestUtaMarketTickersAPI) Category(category string) *PublicRestUtaMarketTickersAPI {
	api.req.Category = GetPointer(category)
	return api
}

// String 否 交易对名称 例如：BTCUSDT
func (api *PublicRestUtaMarketTickersAPI) Symbol(symbol string) *PublicRestUtaMarketTickersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PublicRestUtaMarketOrderBookAPI struct {
	client *PublicRestClient
	req    *PublicRestUtaMarketOrderBookReq
}

type PublicRestUtaMarketOrderBookReq struct {
	Category *string `json:"category"` // String 是 产品类型 SPOT 现货交易 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol   *string `json:"symbol"`   // String 是 交易对名称 例如：BTCUSDT
	Limit    *string `json:"limit"`    // String 否 深度档位 默认5 最大200
}

// String 是 产品类型 SPOT 现货交易 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
func (api *PublicRestUtaMarketOrderBookAPI) Category(category string) *PublicRestUtaMarketOrderBookAPI {
	api.req.Category = GetPointer(category)
	return api
}

// String 是 交易对名称 例如：BTCUSDT
func (api *PublicRestUtaMarketOrderBookAPI) Symbol(symbol string) *PublicRestUtaMarketOrderBookAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// String 否 深度档位 默认5 最大200
func (api *PublicRestUtaMarketOrderBookAPI) Limit(limit string) *PublicRestUtaMarketOrderBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestUtaMarketCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestUtaMarketCandlesReq
}

type PublicRestUtaMarketCandlesReq struct {
	Category  *string `json:"category"`  // String 是 产品类型 SPOT 现货交易 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol    *string `json:"symbol"`    // String 是 交易对名称 例如：BTCUSDT
	Interval  *string `json:"interval"`  // String 是 K线粒度 例如：1m 5m 1H 1D
	StartTime *string `json:"startTime"` // String 否 开始时间 Unix毫秒时间戳
	EndTime   *string `json:"endTime"`   // String 否 结束时间 Unix毫秒时间戳
	Type      *string `json:"type"`      // String 否 K线类型 market mark index premium 默认market
	Limit     *string `json:"limit"`     // String 否 每页数量 默认1000 最大100
}

func (api *PublicRestUtaMarketCandlesAPI) Category(category string) *PublicRestUtaMarketCandlesAPI {
	api.req.Category = GetPointer(category)
	return api
}

func (api *PublicRestUtaMarketCandlesAPI) Symbol(symbol string) *PublicRestUtaMarketCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestUtaMarketCandlesAPI) Interval(interval string) *PublicRestUtaMarketCandlesAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

func (api *PublicRestUtaMarketCandlesAPI) StartTime(startTime string) *PublicRestUtaMarketCandlesAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

func (api *PublicRestUtaMarketCandlesAPI) EndTime(endTime string) *PublicRestUtaMarketCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

func (api *PublicRestUtaMarketCandlesAPI) Type(candleType string) *PublicRestUtaMarketCandlesAPI {
	api.req.Type = GetPointer(candleType)
	return api
}

func (api *PublicRestUtaMarketCandlesAPI) Limit(limit string) *PublicRestUtaMarketCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestUtaMarketHistoryCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestUtaMarketHistoryCandlesReq
}

type PublicRestUtaMarketHistoryCandlesReq struct {
	Category  *string `json:"category"`  // String 是 产品类型 SPOT 现货交易 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol    *string `json:"symbol"`    // String 是 交易对名称 例如：BTCUSDT
	Interval  *string `json:"interval"`  // String 是 K线粒度 例如：1m 5m 1H 1D
	StartTime *string `json:"startTime"` // String 否 开始时间 Unix毫秒时间戳
	EndTime   *string `json:"endTime"`   // String 否 结束时间 Unix毫秒时间戳
	Type      *string `json:"type"`      // String 否 K线类型 market mark index premium 默认market
	Limit     *string `json:"limit"`     // String 否 每页数量 默认100 最大100
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) Category(category string) *PublicRestUtaMarketHistoryCandlesAPI {
	api.req.Category = GetPointer(category)
	return api
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) Symbol(symbol string) *PublicRestUtaMarketHistoryCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) Interval(interval string) *PublicRestUtaMarketHistoryCandlesAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) StartTime(startTime string) *PublicRestUtaMarketHistoryCandlesAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) EndTime(endTime string) *PublicRestUtaMarketHistoryCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) Type(candleType string) *PublicRestUtaMarketHistoryCandlesAPI {
	api.req.Type = GetPointer(candleType)
	return api
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) Limit(limit string) *PublicRestUtaMarketHistoryCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestUtaMarketCurrentFundRateAPI struct {
	client *PublicRestClient
	req    *PublicRestUtaMarketCurrentFundRateReq
}

type PublicRestUtaMarketCurrentFundRateReq struct {
	Symbol *string `json:"symbol"` // String 是 交易对名称 例如：BTCUSDT
}

func (api *PublicRestUtaMarketCurrentFundRateAPI) Symbol(symbol string) *PublicRestUtaMarketCurrentFundRateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PublicRestUtaMarketHistoryFundRateAPI struct {
	client *PublicRestClient
	req    *PublicRestUtaMarketHistoryFundRateReq
}

type PublicRestUtaMarketHistoryFundRateReq struct {
	Category *string `json:"category"` // String 是 产品类型 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol   *string `json:"symbol"`   // String 是 交易对名称 例如：BTCUSDT
	Cursor   *string `json:"cursor"`   // String 否 页码 默认1 最大100
	Limit    *string `json:"limit"`    // String 否 每页数量 默认20 最大100
}

func (api *PublicRestUtaMarketHistoryFundRateAPI) Category(category string) *PublicRestUtaMarketHistoryFundRateAPI {
	api.req.Category = GetPointer(category)
	return api
}

func (api *PublicRestUtaMarketHistoryFundRateAPI) Symbol(symbol string) *PublicRestUtaMarketHistoryFundRateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestUtaMarketHistoryFundRateAPI) Cursor(cursor string) *PublicRestUtaMarketHistoryFundRateAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

func (api *PublicRestUtaMarketHistoryFundRateAPI) Limit(limit string) *PublicRestUtaMarketHistoryFundRateAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
