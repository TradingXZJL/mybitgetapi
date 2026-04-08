package mybitgetapi

type PublicRestClassicSpotMarketCoinsAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketCoinsReq
}

type PublicRestClassicSpotMarketCoinsReq struct {
	Coin *string `json:"coin"` // String 否 币种名称，如 BTC；不填默认返回全部币种信息
}

// String 否 币种名称，如 BTC；不填默认返回全部币种信息
func (api *PublicRestClassicSpotMarketCoinsAPI) Coin(coin string) *PublicRestClassicSpotMarketCoinsAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

type PublicRestClassicSpotMarketSymbolsAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketSymbolsReq
}

type PublicRestClassicSpotMarketSymbolsReq struct {
	Symbol *string `json:"symbol"` // String 否 交易对名称，如 BTCUSDT；不填默认返回全部交易对信息
}

func (api *PublicRestClassicSpotMarketSymbolsAPI) Symbol(symbol string) *PublicRestClassicSpotMarketSymbolsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PublicRestClassicSpotMarketTickersAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketTickersReq
}

type PublicRestClassicSpotMarketTickersReq struct {
	Symbol *string `json:"symbol"` // String 否 交易对名称，如 BTCUSDT；不填默认返回全部交易对信息
}

func (api *PublicRestClassicSpotMarketTickersAPI) Symbol(symbol string) *PublicRestClassicSpotMarketTickersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PublicRestClassicSpotMarketMergeDepthAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketMergeDepthReq
}

type PublicRestClassicSpotMarketMergeDepthReq struct {
	Symbol    *string `json:"symbol"`    // String 是 交易对名称
	Precision *string `json:"precision"` // String 否 深度精度档位，scale0/scale1/scale2/scale3
	Limit     *string `json:"limit"`     // String 否 深度档位，1/5/15/50/max
}

func (api *PublicRestClassicSpotMarketMergeDepthAPI) Symbol(symbol string) *PublicRestClassicSpotMarketMergeDepthAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestClassicSpotMarketMergeDepthAPI) Precision(precision string) *PublicRestClassicSpotMarketMergeDepthAPI {
	api.req.Precision = GetPointer(precision)
	return api
}

func (api *PublicRestClassicSpotMarketMergeDepthAPI) Limit(limit string) *PublicRestClassicSpotMarketMergeDepthAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicSpotMarketOrderbookAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketOrderbookReq
}

type PublicRestClassicSpotMarketOrderbookReq struct {
	Symbol *string `json:"symbol"` // String 是 交易对名称
	Type   *string `json:"type"`   // String 否 深度类型，step0/step1/step2/step3/step4/step5
	Limit  *string `json:"limit"`  // String 否 查询条数，默认150，最大150
}

func (api *PublicRestClassicSpotMarketOrderbookAPI) Symbol(symbol string) *PublicRestClassicSpotMarketOrderbookAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestClassicSpotMarketOrderbookAPI) Type(depthType string) *PublicRestClassicSpotMarketOrderbookAPI {
	api.req.Type = GetPointer(depthType)
	return api
}

func (api *PublicRestClassicSpotMarketOrderbookAPI) Limit(limit string) *PublicRestClassicSpotMarketOrderbookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicSpotMarketCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketCandlesReq
}

type PublicRestClassicSpotMarketCandlesReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易对名称，如 BTCUSDT
	Granularity *string `json:"granularity"` // String 是 K线粒度，如 1min/5min/1h/1day/1week/1M
	StartTime   *string `json:"startTime"`   // String 否 开始时间，Unix毫秒时间戳
	EndTime     *string `json:"endTime"`     // String 否 结束时间，Unix毫秒时间戳
	Limit       *string `json:"limit"`       // String 否 查询条数，默认100，最大1000
}

func (api *PublicRestClassicSpotMarketCandlesAPI) Symbol(symbol string) *PublicRestClassicSpotMarketCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestClassicSpotMarketCandlesAPI) Granularity(granularity string) *PublicRestClassicSpotMarketCandlesAPI {
	api.req.Granularity = GetPointer(granularity)
	return api
}

func (api *PublicRestClassicSpotMarketCandlesAPI) StartTime(startTime string) *PublicRestClassicSpotMarketCandlesAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

func (api *PublicRestClassicSpotMarketCandlesAPI) EndTime(endTime string) *PublicRestClassicSpotMarketCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

func (api *PublicRestClassicSpotMarketCandlesAPI) Limit(limit string) *PublicRestClassicSpotMarketCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicSpotMarketHistoryCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketHistoryCandlesReq
}

type PublicRestClassicSpotMarketHistoryCandlesReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易对名称，如 BTCUSDT
	Granularity *string `json:"granularity"` // String 是 K线粒度，如 1min/5min/1h/1day/1week/1M
	EndTime     *string `json:"endTime"`     // String 是 结束时间，Unix毫秒时间戳
	Limit       *string `json:"limit"`       // String 否 查询条数，默认100，最大200
}

func (api *PublicRestClassicSpotMarketHistoryCandlesAPI) Symbol(symbol string) *PublicRestClassicSpotMarketHistoryCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestClassicSpotMarketHistoryCandlesAPI) Granularity(granularity string) *PublicRestClassicSpotMarketHistoryCandlesAPI {
	api.req.Granularity = GetPointer(granularity)
	return api
}

func (api *PublicRestClassicSpotMarketHistoryCandlesAPI) EndTime(endTime string) *PublicRestClassicSpotMarketHistoryCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

func (api *PublicRestClassicSpotMarketHistoryCandlesAPI) Limit(limit string) *PublicRestClassicSpotMarketHistoryCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicSpotMarketFillsAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketFillsReq
}

type PublicRestClassicSpotMarketFillsReq struct {
	Symbol *string `json:"symbol"` // String 是 交易对名称，如 BTCUSDT
	Limit  *string `json:"limit"`  // String 否 查询条数，默认100，最大500
}

func (api *PublicRestClassicSpotMarketFillsAPI) Symbol(symbol string) *PublicRestClassicSpotMarketFillsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestClassicSpotMarketFillsAPI) Limit(limit string) *PublicRestClassicSpotMarketFillsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicSpotMarketFillsHistoryAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicSpotMarketFillsHistoryReq
}

type PublicRestClassicSpotMarketFillsHistoryReq struct {
	Symbol     *string `json:"symbol"`     // String 是 交易对名称，如 BTCUSDT
	Limit      *string `json:"limit"`      // String 否 查询条数，默认500，最大1000
	IDLessThan *string `json:"idLessThan"` // String 否 返回 tradeId 小于该值的数据
	StartTime  *string `json:"startTime"`  // String 否 开始时间，Unix毫秒时间戳
	EndTime    *string `json:"endTime"`    // String 否 结束时间，Unix毫秒时间戳
}

func (api *PublicRestClassicSpotMarketFillsHistoryAPI) Symbol(symbol string) *PublicRestClassicSpotMarketFillsHistoryAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PublicRestClassicSpotMarketFillsHistoryAPI) Limit(limit string) *PublicRestClassicSpotMarketFillsHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

func (api *PublicRestClassicSpotMarketFillsHistoryAPI) IDLessThan(idLessThan string) *PublicRestClassicSpotMarketFillsHistoryAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}

func (api *PublicRestClassicSpotMarketFillsHistoryAPI) StartTime(startTime string) *PublicRestClassicSpotMarketFillsHistoryAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

func (api *PublicRestClassicSpotMarketFillsHistoryAPI) EndTime(endTime string) *PublicRestClassicSpotMarketFillsHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
