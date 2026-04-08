package mybitgetapi

import "errors"

// GET 获取账户信息
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountInfo() *PrivateRestClassicSpotAccountInfoAPI {
	return &PrivateRestClassicSpotAccountInfoAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountInfoReq{},
	}
}

func (api *PrivateRestClassicSpotAccountInfoAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountInfoRes], error) {
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountInfo])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取账户现货资产
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountAssets() *PrivateRestClassicSpotAccountAssetsAPI {
	return &PrivateRestClassicSpotAccountAssetsAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountAssetsReq{},
	}
}

func (api *PrivateRestClassicSpotAccountAssetsAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountAssetsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountAssets])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountAssetsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取账单流水
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountBills() *PrivateRestClassicSpotAccountBillsAPI {
	return &PrivateRestClassicSpotAccountBillsAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountBillsReq{},
	}
}

func (api *PrivateRestClassicSpotAccountBillsAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountBillsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountBills])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountBillsRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 划转
func (client *PrivateRestClient) NewPrivateRestClassicSpotWalletTransfer() *PrivateRestClassicSpotWalletTransferAPI {
	return &PrivateRestClassicSpotWalletTransferAPI{
		client: client,
		req:    &PrivateRestClassicSpotWalletTransferReq{},
	}
}

func (api *PrivateRestClassicSpotWalletTransferAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotWalletTransferRes], error) {
	if api.req.FromType == nil || api.req.ToType == nil || api.req.Amount == nil || api.req.Coin == nil {
		return nil, errors.New("fromType, toType, amount, coin are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotWalletTransfer])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotWalletTransferRes](api.client.c, url, reqBody, POST)
}

// GET 获取划转记录
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountTransferRecords() *PrivateRestClassicSpotAccountTransferRecordsAPI {
	return &PrivateRestClassicSpotAccountTransferRecordsAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountTransferRecordsReq{},
	}
}

func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountTransferRecordsRes], error) {
	if api.req.Coin == nil {
		return nil, errors.New("coin is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountTransferRecords])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountTransferRecordsRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 开启BGB抵扣
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountSwitchDeduct() *PrivateRestClassicSpotAccountSwitchDeductAPI {
	return &PrivateRestClassicSpotAccountSwitchDeductAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountSwitchDeductReq{},
	}
}

func (api *PrivateRestClassicSpotAccountSwitchDeductAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountSwitchDeductRes], error) {
	if api.req.Deduct == nil {
		return nil, errors.New("deduct is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountSwitchDeduct])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountSwitchDeductRes](api.client.c, url, reqBody, POST)
}

// GET 当前是否开启BGB抵扣
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountDeductInfo() *PrivateRestClassicSpotAccountDeductInfoAPI {
	return &PrivateRestClassicSpotAccountDeductInfoAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountDeductInfoReq{},
	}
}

func (api *PrivateRestClassicSpotAccountDeductInfoAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountDeductInfoRes], error) {
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountDeductInfo])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountDeductInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 账户升级
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountUpgrade() *PrivateRestClassicSpotAccountUpgradeAPI {
	return &PrivateRestClassicSpotAccountUpgradeAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountUpgradeReq{},
	}
}

func (api *PrivateRestClassicSpotAccountUpgradeAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountUpgradeRes], error) {
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountUpgrade])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountUpgradeRes](api.client.c, url, reqBody, POST)
}

// GET 查询升级状态
func (client *PrivateRestClient) NewPrivateRestClassicSpotAccountUpgradeStatus() *PrivateRestClassicSpotAccountUpgradeStatusAPI {
	return &PrivateRestClassicSpotAccountUpgradeStatusAPI{
		client: client,
		req:    &PrivateRestClassicSpotAccountUpgradeStatusReq{},
	}
}

func (api *PrivateRestClassicSpotAccountUpgradeStatusAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotAccountUpgradeStatusRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotAccountUpgradeStatus])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotAccountUpgradeStatusRes](api.client.c, url, NIL_REQBODY, GET)
}
