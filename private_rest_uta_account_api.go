package mybitgetapi

import "errors"

// GET 获取账户余额
func (client *PrivateRestClient) NewPrivateRestUtaAccountBalance() *PrivateRestUtaAccountBalanceAPI {
	return &PrivateRestUtaAccountBalanceAPI{
		client: client,
		req:    &PrivateRestUtaAccountBalanceReq{},
	}
}

func (api *PrivateRestUtaAccountBalanceAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountBalanceRes], error) {
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaAccountBalance])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取资金账户资产（仅返回有资产的币种；可选 coin 筛选）
func (client *PrivateRestClient) NewPrivateRestUtaAccountFundingAssets() *PrivateRestUtaAccountFundingAssetsAPI {
	return &PrivateRestUtaAccountFundingAssetsAPI{
		client: client,
		req:    &PrivateRestUtaAccountFundingAssetsReq{},
	}
}

func (api *PrivateRestUtaAccountFundingAssetsAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountFundingAssetRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaAccountFundingAssets])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountFundingAssetRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取账户设置
func (client *PrivateRestClient) NewPrivateRestUtaAccountSetting() *PrivateRestUtaAccountSettingAPI {
	return &PrivateRestUtaAccountSettingAPI{
		client: client,
		req:    &PrivateRestUtaAccountSettingReq{},
	}
}

func (api *PrivateRestUtaAccountSettingAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountSettingRes], error) {
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaAccountSetting])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountSettingRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 调整杠杆
func (client *PrivateRestClient) NewPrivateRestUtaAccountSetLeverage() *PrivateRestUtaAccountSetLeverageAPI {
	return &PrivateRestUtaAccountSetLeverageAPI{
		client: client,
		req:    &PrivateRestUtaAccountSetLeverageReq{},
	}
}

func (api *PrivateRestUtaAccountSetLeverageAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountSetLeverageRes], error) {
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	if api.req.Leverage == nil {
		return nil, errors.New("leverage is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaAccountSetLeverage])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountSetLeverageRes](api.client.c, url, reqBody, POST)
}

// POST 调整持仓模式
func (client *PrivateRestClient) NewPrivateRestUtaAccountSetHoldMode() *PrivateRestUtaAccountSetHoldModeAPI {
	return &PrivateRestUtaAccountSetHoldModeAPI{
		client: client,
		req:    &PrivateRestUtaAccountSetHoldModeReq{},
	}
}

func (api *PrivateRestUtaAccountSetHoldModeAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountSetHoldModeRes], error) {
	if api.req.HoldMode == nil {
		return nil, errors.New("holdMode is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaAccountSetHoldMode])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountSetHoldModeRes](api.client.c, url, reqBody, POST)
}

// GET 获取资金流水
func (client *PrivateRestClient) NewPrivateRestUtaAccountFinancialRecords() *PrivateRestUtaAccountFinancialRecordsAPI {
	return &PrivateRestUtaAccountFinancialRecordsAPI{
		client: client,
		req:    &PrivateRestUtaAccountFinancialRecordsReq{},
	}
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountFinancialRecordsRes], error) {
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaAccountFinancialRecords])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountFinancialRecordsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取账户费率
func (client *PrivateRestClient) NewPrivateRestUtaAccountFeeRate() *PrivateRestUtaAccountFeeRateAPI {
	return &PrivateRestUtaAccountFeeRateAPI{
		client: client,
		req:    &PrivateRestUtaAccountFeeRateReq{},
	}
}

func (api *PrivateRestUtaAccountFeeRateAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountFeeRateRes], error) {
	if api.req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaAccountFeeRate])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountFeeRateRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取最大可转出
func (client *PrivateRestClient) NewPrivateRestUtaAccountMaxTransferable() *PrivateRestUtaAccountMaxTransferableAPI {
	return &PrivateRestUtaAccountMaxTransferableAPI{
		client: client,
		req:    &PrivateRestUtaAccountMaxTransferableReq{},
	}
}

func (api *PrivateRestUtaAccountMaxTransferableAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountMaxTransferableRes], error) {
	if api.req.Coin == nil {
		return nil, errors.New("coin is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaAccountMaxTransferable])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountMaxTransferableRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取可划转币种
func (client *PrivateRestClient) NewPrivateRestUtaAccountTransferableCoins() *PrivateRestUtaAccountTransferableCoinsAPI {
	return &PrivateRestUtaAccountTransferableCoinsAPI{
		client: client,
		req:    &PrivateRestUtaAccountTransferableCoinsReq{},
	}
}

func (api *PrivateRestUtaAccountTransferableCoinsAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountTransferableCoinsRes], error) {
	if api.req.FromType == nil {
		return nil, errors.New("fromType is required")
	}
	if api.req.ToType == nil {
		return nil, errors.New("toType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaAccountTransferableCoins])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountTransferableCoinsRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 资金划转
func (client *PrivateRestClient) NewPrivateRestUtaAccountTransfer() *PrivateRestUtaAccountTransferAPI {
	return &PrivateRestUtaAccountTransferAPI{
		client: client,
		req:    &PrivateRestUtaAccountTransferReq{},
	}
}

func (api *PrivateRestUtaAccountTransferAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountTransferRes], error) {
	if api.req.FromType == nil {
		return nil, errors.New("fromType is required")
	}
	if api.req.ToType == nil {
		return nil, errors.New("toType is required")
	}
	if api.req.Amount == nil {
		return nil, errors.New("amount is required")
	}
	if api.req.Coin == nil {
		return nil, errors.New("coin is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaAccountTransfer])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountTransferRes](api.client.c, url, reqBody, POST)
}

// POST 切换账户（切换至 classic 账户）
func (client *PrivateRestClient) NewPrivateRestUtaAccountSwitch() *PrivateRestUtaAccountSwitchAPI {
	return &PrivateRestUtaAccountSwitchAPI{
		client: client,
		req:    &PrivateRestUtaAccountSwitchReq{},
	}
}

func (api *PrivateRestUtaAccountSwitchAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountSwitchRes], error) {
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaAccountSwitch])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountSwitchRes](api.client.c, url, reqBody, POST)
}

// GET 获取切换状态
func (client *PrivateRestClient) NewPrivateRestUtaAccountSwitchStatus() *PrivateRestUtaAccountSwitchStatusAPI {
	return &PrivateRestUtaAccountSwitchStatusAPI{
		client: client,
		req:    &PrivateRestUtaAccountSwitchStatusReq{},
	}
}

func (api *PrivateRestUtaAccountSwitchStatusAPI) Do() (*BitgetRestRes[PrivateRestUtaAccountSwitchStatusRes], error) {
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaAccountSwitchStatus])
	return bitgetCallApiWithSecret[PrivateRestUtaAccountSwitchStatusRes](api.client.c, url, NIL_REQBODY, GET)
}
