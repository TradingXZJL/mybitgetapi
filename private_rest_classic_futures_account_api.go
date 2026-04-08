package mybitgetapi

import "errors"

// GET 获取单个交易对账户信息
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSingleAccount() *PrivateRestClassicFuturesAccountSingleAccountAPI {
	return &PrivateRestClassicFuturesAccountSingleAccountAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSingleAccountReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSingleAccountAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSingleAccountRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil {
		return nil, errors.New("symbol, productType, marginCoin are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSingleAccount])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSingleAccountRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取账户列表
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountList() *PrivateRestClassicFuturesAccountListAPI {
	return &PrivateRestClassicFuturesAccountListAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountListReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountListAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountListRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountList])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountListRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取所有子账户合约资产信息
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSubAccountAssets() *PrivateRestClassicFuturesAccountSubAccountAssetsAPI {
	return &PrivateRestClassicFuturesAccountSubAccountAssetsAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSubAccountAssetsReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSubAccountAssetsAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSubAccountAssetsRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSubAccountAssets])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSubAccountAssetsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET U本位合约利息记录
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountInterestHistory() *PrivateRestClassicFuturesAccountInterestHistoryAPI {
	return &PrivateRestClassicFuturesAccountInterestHistoryAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountInterestHistoryReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountInterestHistoryAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountInterestHistoryRes], error) {
	if api.req.ProductType == nil || api.req.StartTime == nil || api.req.EndTime == nil {
		return nil, errors.New("productType, startTime, endTime are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountInterestHistory])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountInterestHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取最大可开
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountMaxOpen() *PrivateRestClassicFuturesAccountMaxOpenAPI {
	return &PrivateRestClassicFuturesAccountMaxOpenAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountMaxOpenReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountMaxOpenAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountMaxOpenRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil || api.req.PosSide == nil || api.req.OrderType == nil {
		return nil, errors.New("symbol, productType, marginCoin, posSide, orderType are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountMaxOpen])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountMaxOpenRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取强平价格
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountLiquidationPrice() *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	return &PrivateRestClassicFuturesAccountLiquidationPriceAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountLiquidationPriceReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountLiquidationPriceRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil || api.req.PosSide == nil || api.req.OrderType == nil || api.req.OpenAmount == nil {
		return nil, errors.New("symbol, productType, marginCoin, posSide, orderType, openAmount are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountLiquidationPrice])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountLiquidationPriceRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取可开数量
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountOpenCount() *PrivateRestClassicFuturesAccountOpenCountAPI {
	return &PrivateRestClassicFuturesAccountOpenCountAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountOpenCountReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountOpenCountAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountOpenCountRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil || api.req.OpenAmount == nil || api.req.OpenPrice == nil {
		return nil, errors.New("symbol, productType, marginCoin, openAmount, openPrice are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountOpenCount])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountOpenCountRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 调整杠杆
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSetLeverage() *PrivateRestClassicFuturesAccountSetLeverageAPI {
	return &PrivateRestClassicFuturesAccountSetLeverageAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSetLeverageReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSetLeverageRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil {
		return nil, errors.New("symbol, productType, marginCoin are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSetLeverage])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSetLeverageRes](api.client.c, url, reqBody, POST)
}

// POST 调整产品线杠杆
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSetAllLeverage() *PrivateRestClassicFuturesAccountSetAllLeverageAPI {
	return &PrivateRestClassicFuturesAccountSetAllLeverageAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSetAllLeverageReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSetAllLeverageAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSetAllLeverageRes], error) {
	if api.req.ProductType == nil || api.req.Leverage == nil {
		return nil, errors.New("productType, leverage are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSetAllLeverage])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSetAllLeverageRes](api.client.c, url, reqBody, POST)
}

// POST 调整保证金
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSetMargin() *PrivateRestClassicFuturesAccountSetMarginAPI {
	return &PrivateRestClassicFuturesAccountSetMarginAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSetMarginReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSetMarginAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSetMarginRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil || api.req.HoldSide == nil || api.req.Amount == nil {
		return nil, errors.New("symbol, productType, marginCoin, holdSide, amount are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSetMargin])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSetMarginRes](api.client.c, url, reqBody, POST)
}

// POST 设置U本位联合保证金模式
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSetAssetMode() *PrivateRestClassicFuturesAccountSetAssetModeAPI {
	return &PrivateRestClassicFuturesAccountSetAssetModeAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSetAssetModeReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSetAssetModeAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSetAssetModeRes], error) {
	if api.req.ProductType == nil || api.req.AssetMode == nil {
		return nil, errors.New("productType, assetMode are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSetAssetMode])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSetAssetModeRes](api.client.c, url, reqBody, POST)
}

// POST 调整保证金模式
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSetMarginMode() *PrivateRestClassicFuturesAccountSetMarginModeAPI {
	return &PrivateRestClassicFuturesAccountSetMarginModeAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSetMarginModeReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSetMarginModeAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSetMarginModeRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil || api.req.MarginMode == nil {
		return nil, errors.New("symbol, productType, marginCoin, marginMode are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSetMarginMode])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSetMarginModeRes](api.client.c, url, reqBody, POST)
}

// POST 调整单双向持仓模式
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountSetPositionMode() *PrivateRestClassicFuturesAccountSetPositionModeAPI {
	return &PrivateRestClassicFuturesAccountSetPositionModeAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountSetPositionModeReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountSetPositionModeAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountSetPositionModeRes], error) {
	if api.req.ProductType == nil || api.req.PosMode == nil {
		return nil, errors.New("productType, posMode are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountSetPositionMode])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountSetPositionModeRes](api.client.c, url, reqBody, POST)
}

// GET 获取账务记录
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountBill() *PrivateRestClassicFuturesAccountBillAPI {
	return &PrivateRestClassicFuturesAccountBillAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountBillReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountBillAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountBillRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountBill])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountBillRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取逐仓模式交易对
func (client *PrivateRestClient) NewPrivateRestClassicFuturesAccountIsolatedSymbols() *PrivateRestClassicFuturesAccountIsolatedSymbolsAPI {
	return &PrivateRestClassicFuturesAccountIsolatedSymbolsAPI{
		client: client,
		req:    &PrivateRestClassicFuturesAccountIsolatedSymbolsReq{},
	}
}

func (api *PrivateRestClassicFuturesAccountIsolatedSymbolsAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesAccountIsolatedSymbolsRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesAccountIsolatedSymbols])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesAccountIsolatedSymbolsRes](api.client.c, url, NIL_REQBODY, GET)
}
