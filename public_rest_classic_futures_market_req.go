package mybitgetapi

type PublicRestClassicFuturesMarketVipFeeRateAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketVipFeeRateReq
}

type PublicRestClassicFuturesMarketVipFeeRateReq struct{}

type PublicRestClassicFuturesMarketInterestRateAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketInterestRateReq
}

type PublicRestClassicFuturesMarketInterestRateReq struct {
	Coin *string `json:"coin"` // String 是 币种
}

func (api *PublicRestClassicFuturesMarketInterestRateAPI) Coin(coin string) *PublicRestClassicFuturesMarketInterestRateAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

type PublicRestClassicFuturesMarketMergeDepthAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketMergeDepthReq
}

type PublicRestClassicFuturesMarketMergeDepthReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
	Precision   *string `json:"precision"`   // String 否 精度档位 scale0/1/2/3
	Limit       *string `json:"limit"`       // String 否 档位 1/5/15/50/max
}

func (api *PublicRestClassicFuturesMarketMergeDepthAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketMergeDepthAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketMergeDepthAPI) ProductType(productType string) *PublicRestClassicFuturesMarketMergeDepthAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketMergeDepthAPI) Precision(precision string) *PublicRestClassicFuturesMarketMergeDepthAPI {
	api.req.Precision = GetPointer(precision)
	return api
}
func (api *PublicRestClassicFuturesMarketMergeDepthAPI) Limit(limit string) *PublicRestClassicFuturesMarketMergeDepthAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicFuturesMarketTickerAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketTickerReq
}

type PublicRestClassicFuturesMarketTickerReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
}

func (api *PublicRestClassicFuturesMarketTickerAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketTickerAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketTickerAPI) ProductType(productType string) *PublicRestClassicFuturesMarketTickerAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}

type PublicRestClassicFuturesMarketTickersAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketTickersReq
}

type PublicRestClassicFuturesMarketTickersReq struct {
	ProductType *string `json:"productType"` // String 是 产品类型
}

func (api *PublicRestClassicFuturesMarketTickersAPI) ProductType(productType string) *PublicRestClassicFuturesMarketTickersAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}

type PublicRestClassicFuturesMarketFillsAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketFillsReq
}

type PublicRestClassicFuturesMarketFillsReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
	Limit       *string `json:"limit"`       // String 否 查询数量，默认100，最大100
}

func (api *PublicRestClassicFuturesMarketFillsAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketFillsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketFillsAPI) ProductType(productType string) *PublicRestClassicFuturesMarketFillsAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketFillsAPI) Limit(limit string) *PublicRestClassicFuturesMarketFillsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicFuturesMarketFillsHistoryAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketFillsHistoryReq
}

type PublicRestClassicFuturesMarketFillsHistoryReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
	Limit       *string `json:"limit"`       // String 否 查询数量，默认500，最大1000
	IDLessThan  *string `json:"idLessThan"`  // String 否 分页ID
	StartTime   *string `json:"startTime"`   // String 否 开始时间
	EndTime     *string `json:"endTime"`     // String 否 结束时间
}

func (api *PublicRestClassicFuturesMarketFillsHistoryAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketFillsHistoryAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketFillsHistoryAPI) ProductType(productType string) *PublicRestClassicFuturesMarketFillsHistoryAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketFillsHistoryAPI) Limit(limit string) *PublicRestClassicFuturesMarketFillsHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PublicRestClassicFuturesMarketFillsHistoryAPI) IDLessThan(idLessThan string) *PublicRestClassicFuturesMarketFillsHistoryAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}
func (api *PublicRestClassicFuturesMarketFillsHistoryAPI) StartTime(startTime string) *PublicRestClassicFuturesMarketFillsHistoryAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PublicRestClassicFuturesMarketFillsHistoryAPI) EndTime(endTime string) *PublicRestClassicFuturesMarketFillsHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

type PublicRestClassicFuturesMarketCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketCandlesReq
}

type PublicRestClassicFuturesMarketCandlesReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
	Granularity *string `json:"granularity"` // String 是 K线粒度
	StartTime   *string `json:"startTime"`   // String 否 开始时间
	EndTime     *string `json:"endTime"`     // String 否 结束时间
	KLineType   *string `json:"kLineType"`   // String 否 K线类型 MARKET/MARK/INDEX
	Limit       *string `json:"limit"`       // String 否 默认100，最大1000
}

func (api *PublicRestClassicFuturesMarketCandlesAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketCandlesAPI) ProductType(productType string) *PublicRestClassicFuturesMarketCandlesAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketCandlesAPI) Granularity(granularity string) *PublicRestClassicFuturesMarketCandlesAPI {
	api.req.Granularity = GetPointer(granularity)
	return api
}
func (api *PublicRestClassicFuturesMarketCandlesAPI) StartTime(startTime string) *PublicRestClassicFuturesMarketCandlesAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PublicRestClassicFuturesMarketCandlesAPI) EndTime(endTime string) *PublicRestClassicFuturesMarketCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PublicRestClassicFuturesMarketCandlesAPI) KLineType(kLineType string) *PublicRestClassicFuturesMarketCandlesAPI {
	api.req.KLineType = GetPointer(kLineType)
	return api
}
func (api *PublicRestClassicFuturesMarketCandlesAPI) Limit(limit string) *PublicRestClassicFuturesMarketCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicFuturesMarketHistoryCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketHistoryCandlesReq
}

type PublicRestClassicFuturesMarketHistoryCandlesReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
	Granularity *string `json:"granularity"` // String 是 K线粒度
	StartTime   *string `json:"startTime"`   // String 否 开始时间
	EndTime     *string `json:"endTime"`     // String 否 结束时间
	Limit       *string `json:"limit"`       // String 否 默认100，最大200
}

func (api *PublicRestClassicFuturesMarketHistoryCandlesAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketHistoryCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryCandlesAPI) ProductType(productType string) *PublicRestClassicFuturesMarketHistoryCandlesAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryCandlesAPI) Granularity(granularity string) *PublicRestClassicFuturesMarketHistoryCandlesAPI {
	api.req.Granularity = GetPointer(granularity)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryCandlesAPI) StartTime(startTime string) *PublicRestClassicFuturesMarketHistoryCandlesAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryCandlesAPI) EndTime(endTime string) *PublicRestClassicFuturesMarketHistoryCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryCandlesAPI) Limit(limit string) *PublicRestClassicFuturesMarketHistoryCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicFuturesMarketHistoryIndexCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketHistoryIndexCandlesReq
}

type PublicRestClassicFuturesMarketHistoryIndexCandlesReq = PublicRestClassicFuturesMarketHistoryCandlesReq

func (api *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI) ProductType(productType string) *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI) Granularity(granularity string) *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI {
	api.req.Granularity = GetPointer(granularity)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI) StartTime(startTime string) *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI) EndTime(endTime string) *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI) Limit(limit string) *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicFuturesMarketHistoryMarkCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketHistoryMarkCandlesReq
}

type PublicRestClassicFuturesMarketHistoryMarkCandlesReq = PublicRestClassicFuturesMarketHistoryCandlesReq

func (api *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI) ProductType(productType string) *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI) Granularity(granularity string) *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI {
	api.req.Granularity = GetPointer(granularity)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI) StartTime(startTime string) *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI) EndTime(endTime string) *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI) Limit(limit string) *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestClassicFuturesMarketCurrentFundRateAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketCurrentFundRateReq
}

type PublicRestClassicFuturesMarketCurrentFundRateReq struct {
	Symbol      *string `json:"symbol"`      // String 否 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
}

func (api *PublicRestClassicFuturesMarketCurrentFundRateAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketCurrentFundRateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketCurrentFundRateAPI) ProductType(productType string) *PublicRestClassicFuturesMarketCurrentFundRateAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}

type PublicRestClassicFuturesMarketHistoryFundRateAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketHistoryFundRateReq
}

type PublicRestClassicFuturesMarketHistoryFundRateReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
	PageSize    *string `json:"pageSize"`    // String 否 每页数量，默认20，最大100
	PageNo      *string `json:"pageNo"`      // String 否 页码
}

func (api *PublicRestClassicFuturesMarketHistoryFundRateAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketHistoryFundRateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryFundRateAPI) ProductType(productType string) *PublicRestClassicFuturesMarketHistoryFundRateAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryFundRateAPI) PageSize(pageSize string) *PublicRestClassicFuturesMarketHistoryFundRateAPI {
	api.req.PageSize = GetPointer(pageSize)
	return api
}
func (api *PublicRestClassicFuturesMarketHistoryFundRateAPI) PageNo(pageNo string) *PublicRestClassicFuturesMarketHistoryFundRateAPI {
	api.req.PageNo = GetPointer(pageNo)
	return api
}

type PublicRestClassicFuturesMarketFundingTimeAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketFundingTimeReq
}

type PublicRestClassicFuturesMarketFundingTimeReq struct {
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
}

func (api *PublicRestClassicFuturesMarketFundingTimeAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketFundingTimeAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketFundingTimeAPI) ProductType(productType string) *PublicRestClassicFuturesMarketFundingTimeAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}

type PublicRestClassicFuturesMarketContractsAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicFuturesMarketContractsReq
}

type PublicRestClassicFuturesMarketContractsReq struct {
	Symbol      *string `json:"symbol"`      // String 否 交易币对
	ProductType *string `json:"productType"` // String 是 产品类型
}

func (api *PublicRestClassicFuturesMarketContractsAPI) Symbol(symbol string) *PublicRestClassicFuturesMarketContractsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PublicRestClassicFuturesMarketContractsAPI) ProductType(productType string) *PublicRestClassicFuturesMarketContractsAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
