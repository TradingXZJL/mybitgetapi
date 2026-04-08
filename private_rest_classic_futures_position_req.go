package mybitgetapi

type PrivateRestClassicFuturesPositionQueryPositionLeverAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesPositionQueryPositionLeverReq
}

type PrivateRestClassicFuturesPositionQueryPositionLeverReq struct {
	ProductType *string `json:"productType"` // String 是 产品类型
	Symbol      *string `json:"symbol"`      // String 是 交易币对
}

func (api *PrivateRestClassicFuturesPositionQueryPositionLeverAPI) ProductType(productType string) *PrivateRestClassicFuturesPositionQueryPositionLeverAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesPositionQueryPositionLeverAPI) Symbol(symbol string) *PrivateRestClassicFuturesPositionQueryPositionLeverAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PrivateRestClassicFuturesPositionSinglePositionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesPositionSinglePositionReq
}

type PrivateRestClassicFuturesPositionSinglePositionReq struct {
	ProductType *string `json:"productType"` // String 是 产品类型
	Symbol      *string `json:"symbol"`      // String 是 交易币对
	MarginCoin  *string `json:"marginCoin"`  // String 是 保证金币种
}

func (api *PrivateRestClassicFuturesPositionSinglePositionAPI) ProductType(productType string) *PrivateRestClassicFuturesPositionSinglePositionAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesPositionSinglePositionAPI) Symbol(symbol string) *PrivateRestClassicFuturesPositionSinglePositionAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesPositionSinglePositionAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesPositionSinglePositionAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}

type PrivateRestClassicFuturesPositionAllPositionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesPositionAllPositionReq
}

type PrivateRestClassicFuturesPositionAllPositionReq struct {
	ProductType *string `json:"productType"` // String 是 产品类型
	MarginCoin  *string `json:"marginCoin"`  // String 否 保证金币种
}

func (api *PrivateRestClassicFuturesPositionAllPositionAPI) ProductType(productType string) *PrivateRestClassicFuturesPositionAllPositionAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesPositionAllPositionAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesPositionAllPositionAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}

type PrivateRestClassicFuturesPositionHistoryPositionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesPositionHistoryPositionReq
}

type PrivateRestClassicFuturesPositionHistoryPositionReq struct {
	Symbol      *string `json:"symbol"`      // String 否 交易币对
	ProductType *string `json:"productType"` // String 否 产品类型（默认USDT-FUTURES）
	IDLessThan  *string `json:"idLessThan"`  // String 否 分页ID
	StartTime   *string `json:"startTime"`   // String 否 开始时间
	EndTime     *string `json:"endTime"`     // String 否 结束时间
	Limit       *string `json:"limit"`       // String 否 默认20，最大100
}

func (api *PrivateRestClassicFuturesPositionHistoryPositionAPI) Symbol(symbol string) *PrivateRestClassicFuturesPositionHistoryPositionAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesPositionHistoryPositionAPI) ProductType(productType string) *PrivateRestClassicFuturesPositionHistoryPositionAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesPositionHistoryPositionAPI) IDLessThan(idLessThan string) *PrivateRestClassicFuturesPositionHistoryPositionAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}
func (api *PrivateRestClassicFuturesPositionHistoryPositionAPI) StartTime(startTime string) *PrivateRestClassicFuturesPositionHistoryPositionAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicFuturesPositionHistoryPositionAPI) EndTime(endTime string) *PrivateRestClassicFuturesPositionHistoryPositionAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicFuturesPositionHistoryPositionAPI) Limit(limit string) *PrivateRestClassicFuturesPositionHistoryPositionAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
