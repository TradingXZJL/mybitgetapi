package mybitgetapi

import "errors"

// GET 获取仓位档位梯度配置
func (client *PrivateRestClient) NewPrivateRestClassicFuturesPositionQueryPositionLever() *PrivateRestClassicFuturesPositionQueryPositionLeverAPI {
	return &PrivateRestClassicFuturesPositionQueryPositionLeverAPI{
		client: client,
		req:    &PrivateRestClassicFuturesPositionQueryPositionLeverReq{},
	}
}

func (api *PrivateRestClassicFuturesPositionQueryPositionLeverAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesPositionQueryPositionLeverRes], error) {
	if api.req.ProductType == nil || api.req.Symbol == nil {
		return nil, errors.New("productType, symbol are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesPositionQueryPositionLever])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesPositionQueryPositionLeverRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取单个合约仓位信息
func (client *PrivateRestClient) NewPrivateRestClassicFuturesPositionSinglePosition() *PrivateRestClassicFuturesPositionSinglePositionAPI {
	return &PrivateRestClassicFuturesPositionSinglePositionAPI{
		client: client,
		req:    &PrivateRestClassicFuturesPositionSinglePositionReq{},
	}
}

func (api *PrivateRestClassicFuturesPositionSinglePositionAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesPositionSinglePositionRes], error) {
	if api.req.ProductType == nil || api.req.Symbol == nil || api.req.MarginCoin == nil {
		return nil, errors.New("productType, symbol, marginCoin are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesPositionSinglePosition])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesPositionSinglePositionRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全部合约仓位信息
func (client *PrivateRestClient) NewPrivateRestClassicFuturesPositionAllPosition() *PrivateRestClassicFuturesPositionAllPositionAPI {
	return &PrivateRestClassicFuturesPositionAllPositionAPI{
		client: client,
		req:    &PrivateRestClassicFuturesPositionAllPositionReq{},
	}
}

func (api *PrivateRestClassicFuturesPositionAllPositionAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesPositionAllPositionRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesPositionAllPosition])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesPositionAllPositionRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取合约历史持仓列表
func (client *PrivateRestClient) NewPrivateRestClassicFuturesPositionHistoryPosition() *PrivateRestClassicFuturesPositionHistoryPositionAPI {
	return &PrivateRestClassicFuturesPositionHistoryPositionAPI{
		client: client,
		req:    &PrivateRestClassicFuturesPositionHistoryPositionReq{},
	}
}

func (api *PrivateRestClassicFuturesPositionHistoryPositionAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesPositionHistoryPositionRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesPositionHistoryPosition])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesPositionHistoryPositionRes](api.client.c, url, NIL_REQBODY, GET)
}
