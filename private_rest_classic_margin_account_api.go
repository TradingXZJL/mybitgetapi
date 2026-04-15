package mybitgetapi

import "errors"

// GET 获取全仓杠杆账户资产
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountAssets() *PrivateRestClassicMarginCrossAccountAssetsAPI {
	return &PrivateRestClassicMarginCrossAccountAssetsAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountAssetsReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountAssetsAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountAssetsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountAssets])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountAssetsRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 全仓借款
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountBorrow() *PrivateRestClassicMarginCrossAccountBorrowAPI {
	return &PrivateRestClassicMarginCrossAccountBorrowAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountBorrowReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountBorrowAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountBorrowRes], error) {
	if api.req.Coin == nil || api.req.BorrowAmount == nil {
		return nil, errors.New("coin, borrowAmount are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountBorrow])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountBorrowRes](api.client.c, url, reqBody, POST)
}

// POST 全仓还款
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountRepay() *PrivateRestClassicMarginCrossAccountRepayAPI {
	return &PrivateRestClassicMarginCrossAccountRepayAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountRepayReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountRepayAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountRepayRes], error) {
	if api.req.Coin == nil || api.req.RepayAmount == nil {
		return nil, errors.New("coin, repayAmount are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountRepay])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountRepayRes](api.client.c, url, reqBody, POST)
}

// GET 全仓风险率
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountRiskRate() *PrivateRestClassicMarginCrossAccountRiskRateAPI {
	return &PrivateRestClassicMarginCrossAccountRiskRateAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountRiskRateReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountRiskRateAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountRiskRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountRiskRate])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountRiskRateRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 全仓最大可借数量
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountMaxBorrowableAmount() *PrivateRestClassicMarginCrossAccountMaxBorrowableAmountAPI {
	return &PrivateRestClassicMarginCrossAccountMaxBorrowableAmountAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountMaxBorrowableAmountReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountMaxBorrowableAmountAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountMaxBorrowableAmountRes], error) {
	if api.req.Coin == nil {
		return nil, errors.New("coin is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountMaxBorrowableAmount])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountMaxBorrowableAmountRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 全仓最大可转出数量
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountMaxTransferOutAmount() *PrivateRestClassicMarginCrossAccountMaxTransferOutAmountAPI {
	return &PrivateRestClassicMarginCrossAccountMaxTransferOutAmountAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountMaxTransferOutAmountReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountMaxTransferOutAmountAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountMaxTransferOutAmountRes], error) {
	if api.req.Coin == nil {
		return nil, errors.New("coin is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountMaxTransferOutAmount])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountMaxTransferOutAmountRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 全仓利率与最大可借数量配置
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountInterestRateAndLimit() *PrivateRestClassicMarginCrossAccountInterestRateAndLimitAPI {
	return &PrivateRestClassicMarginCrossAccountInterestRateAndLimitAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountInterestRateAndLimitReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountInterestRateAndLimitAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountInterestRateAndLimitRes], error) {
	if api.req.Coin == nil {
		return nil, errors.New("coin is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountInterestRateAndLimit])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountInterestRateAndLimitRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 全仓一键还款
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountFlashRepay() *PrivateRestClassicMarginCrossAccountFlashRepayAPI {
	return &PrivateRestClassicMarginCrossAccountFlashRepayAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountFlashRepayReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountFlashRepayAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountFlashRepayRes], error) {
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountFlashRepay])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountFlashRepayRes](api.client.c, url, reqBody, POST)
}

// POST 获取全仓还款结果
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountQueryFlashRepayStatus() *PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI {
	return &PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusReq{IdList: make([]string, 0)},
	}
}

func (api *PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusRes], error) {
	if len(api.req.IdList) == 0 {
		return nil, errors.New("idList is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountQueryFlashRepayStatus])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusRes](api.client.c, url, reqBody, POST)
}

// GET 获取全仓借款历史
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossRecordBorrowHistory() *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI {
	return &PrivateRestClassicMarginCrossRecordBorrowHistoryAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossRecordBorrowHistoryReq{},
	}
}

func (api *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossRecordBorrowHistoryRes], error) {
	if api.req.StartTime == nil {
		return nil, errors.New("startTime is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossRecordBorrowHistory])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossRecordBorrowHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全仓还款历史
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossRecordRepayHistory() *PrivateRestClassicMarginCrossRecordRepayHistoryAPI {
	return &PrivateRestClassicMarginCrossRecordRepayHistoryAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossRecordRepayHistoryReq{},
	}
}

func (api *PrivateRestClassicMarginCrossRecordRepayHistoryAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossRecordRepayHistoryRes], error) {
	if api.req.StartTime == nil {
		return nil, errors.New("startTime is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossRecordRepayHistory])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossRecordRepayHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全仓利息历史
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossRecordInterestHistory() *PrivateRestClassicMarginCrossRecordInterestHistoryAPI {
	return &PrivateRestClassicMarginCrossRecordInterestHistoryAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossRecordInterestHistoryReq{},
	}
}

func (api *PrivateRestClassicMarginCrossRecordInterestHistoryAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossRecordInterestHistoryRes], error) {
	if api.req.StartTime == nil {
		return nil, errors.New("startTime is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossRecordInterestHistory])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossRecordInterestHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全仓强平历史
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossRecordLiquidationHistory() *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI {
	return &PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossRecordLiquidationHistoryReq{},
	}
}

func (api *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossRecordLiquidationHistoryRes], error) {
	if api.req.StartTime == nil {
		return nil, errors.New("startTime is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossRecordLiquidationHistory])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossRecordLiquidationHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全仓财务流水记录
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossRecordFinancialRecords() *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI {
	return &PrivateRestClassicMarginCrossRecordFinancialRecordsAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossRecordFinancialRecordsReq{},
	}
}

func (api *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossRecordFinancialRecordsRes], error) {
	if api.req.StartTime == nil {
		return nil, errors.New("startTime is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossRecordFinancialRecords])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossRecordFinancialRecordsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取逐仓杠杆账户资产
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountAssets() *PrivateRestClassicMarginIsolatedAccountAssetsAPI {
	return &PrivateRestClassicMarginIsolatedAccountAssetsAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountAssetsReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountAssetsAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountAssetsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountAssets])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountAssetsRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 逐仓借款
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountBorrow() *PrivateRestClassicMarginIsolatedAccountBorrowAPI {
	return &PrivateRestClassicMarginIsolatedAccountBorrowAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountBorrowReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountBorrowAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountBorrowRes], error) {
	if api.req.Symbol == nil || api.req.Coin == nil || api.req.BorrowAmount == nil {
		return nil, errors.New("symbol, coin, borrowAmount are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountBorrow])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountBorrowRes](api.client.c, url, reqBody, POST)
}

// POST 逐仓还款
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountRepay() *PrivateRestClassicMarginIsolatedAccountRepayAPI {
	return &PrivateRestClassicMarginIsolatedAccountRepayAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountRepayReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountRepayAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountRepayRes], error) {
	if api.req.RepayAmount == nil || api.req.Coin == nil || api.req.Symbol == nil {
		return nil, errors.New("repayAmount, coin, symbol are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountRepay])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountRepayRes](api.client.c, url, reqBody, POST)
}

// GET 逐仓风险率
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountRiskRate() *PrivateRestClassicMarginIsolatedAccountRiskRateAPI {
	return &PrivateRestClassicMarginIsolatedAccountRiskRateAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountRiskRateReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountRiskRateAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountRiskRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountRiskRate])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountRiskRateRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 逐仓档位梯度配置
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountTierData() *PrivateRestClassicMarginIsolatedAccountTierDataAPI {
	return &PrivateRestClassicMarginIsolatedAccountTierDataAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountTierDataReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountTierDataAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountTierDataRes], error) {
	if api.req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountTierData])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountTierDataRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 逐仓利率与最大可借数量配置
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountInterestRateAndLimit() *PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitAPI {
	return &PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitRes], error) {
	if api.req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountInterestRateAndLimit])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 逐仓最大可借数量
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountMaxBorrowableAmount() *PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountAPI {
	return &PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountRes], error) {
	if api.req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmount])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 逐仓最大可转出数量
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountMaxTransferOutAmount() *PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountAPI {
	return &PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountRes], error) {
	if api.req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmount])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 逐仓一键还款
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountFlashRepay() *PrivateRestClassicMarginIsolatedAccountFlashRepayAPI {
	return &PrivateRestClassicMarginIsolatedAccountFlashRepayAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountFlashRepayReq{SymbolList: make([]string, 0)},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountFlashRepayAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountFlashRepayRes], error) {
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountFlashRepay])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountFlashRepayRes](api.client.c, url, reqBody, POST)
}

// POST 获取逐仓还款结果
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatus() *PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI {
	return &PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusReq{IdList: make([]string, 0)},
	}
}

func (api *PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusRes], error) {
	if len(api.req.IdList) == 0 {
		return nil, errors.New("idList is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatus])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusRes](api.client.c, url, reqBody, POST)
}

// GET 全仓档位梯度配置
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossAccountTierData() *PrivateRestClassicMarginCrossAccountTierDataAPI {
	return &PrivateRestClassicMarginCrossAccountTierDataAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossAccountTierDataReq{},
	}
}

func (api *PrivateRestClassicMarginCrossAccountTierDataAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossAccountTierDataRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossAccountTierData])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossAccountTierDataRes](api.client.c, url, NIL_REQBODY, GET)
}
