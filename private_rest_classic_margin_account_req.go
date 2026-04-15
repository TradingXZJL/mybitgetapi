package mybitgetapi

type PrivateRestClassicMarginCrossAccountAssetsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountAssetsReq
}

type PrivateRestClassicMarginCrossAccountAssetsReq struct {
	Coin *string `json:"coin"` // 币种，如 USDT
}

func (api *PrivateRestClassicMarginCrossAccountAssetsAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountAssetsAPI {
	api.req.Coin = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossAccountBorrowAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountBorrowReq
}

type PrivateRestClassicMarginCrossAccountBorrowReq struct {
	Coin         *string `json:"coin"`         // 借出币种
	BorrowAmount *string `json:"borrowAmount"` // 借出数量（最长 8 位小数）
	ClientID     *string `json:"clientid"`     // 客户自定义订单 ID
}

func (api *PrivateRestClassicMarginCrossAccountBorrowAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountBorrowAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossAccountBorrowAPI) BorrowAmount(v string) *PrivateRestClassicMarginCrossAccountBorrowAPI {
	api.req.BorrowAmount = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossAccountBorrowAPI) ClientID(v string) *PrivateRestClassicMarginCrossAccountBorrowAPI {
	api.req.ClientID = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossAccountRepayAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountRepayReq
}

type PrivateRestClassicMarginCrossAccountRepayReq struct {
	Coin        *string `json:"coin"`        // 还款币种
	RepayAmount *string `json:"repayAmount"` // 还款数量（最长 8 位小数）
}

func (api *PrivateRestClassicMarginCrossAccountRepayAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountRepayAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossAccountRepayAPI) RepayAmount(v string) *PrivateRestClassicMarginCrossAccountRepayAPI {
	api.req.RepayAmount = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossAccountRiskRateAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountRiskRateReq
}

type PrivateRestClassicMarginCrossAccountRiskRateReq struct{}

type PrivateRestClassicMarginCrossAccountMaxBorrowableAmountAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountMaxBorrowableAmountReq
}

type PrivateRestClassicMarginCrossAccountMaxBorrowableAmountReq struct {
	Coin *string `json:"coin"` // 借出币种，如 BTC
}

func (api *PrivateRestClassicMarginCrossAccountMaxBorrowableAmountAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountMaxBorrowableAmountAPI {
	api.req.Coin = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossAccountMaxTransferOutAmountAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountMaxTransferOutAmountReq
}

type PrivateRestClassicMarginCrossAccountMaxTransferOutAmountReq struct {
	Coin *string `json:"coin"` // 币种名称
}

func (api *PrivateRestClassicMarginCrossAccountMaxTransferOutAmountAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountMaxTransferOutAmountAPI {
	api.req.Coin = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossAccountInterestRateAndLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountInterestRateAndLimitReq
}

type PrivateRestClassicMarginCrossAccountInterestRateAndLimitReq struct {
	Coin *string `json:"coin"` // 币种，如 BTC、ETH
}

func (api *PrivateRestClassicMarginCrossAccountInterestRateAndLimitAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountInterestRateAndLimitAPI {
	api.req.Coin = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossAccountFlashRepayAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountFlashRepayReq
}

type PrivateRestClassicMarginCrossAccountFlashRepayReq struct {
	Coin *string `json:"coin"` // 全仓还款币种；不传则全币种一键还款
}

func (api *PrivateRestClassicMarginCrossAccountFlashRepayAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountFlashRepayAPI {
	api.req.Coin = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusReq
}

type PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusReq struct {
	IdList []string `json:"idList"` // 发起还款请求的 id 集合（最大 100 个）
}

func (api *PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI) AddId(v string) *PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI {
	api.req.IdList = append(api.req.IdList, v)
	return api
}
func (api *PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI) IdList(v []string) *PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusAPI {
	api.req.IdList = v
	return api
}

type PrivateRestClassicMarginCrossRecordBorrowHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossRecordBorrowHistoryReq
}

type PrivateRestClassicMarginCrossRecordBorrowHistoryReq struct {
	LoanID     *string `json:"loanId"`     // 借款 ID（单条数据精准匹配）
	Coin       *string `json:"coin"`       // 币种
	StartTime  *string `json:"startTime"`  // 开始时间，Unix 毫秒时间戳
	EndTime    *string `json:"endTime"`    // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit      *string `json:"limit"`      // 查询条数，默认 100，最大 500
	IDLessThan *string `json:"idLessThan"` // 翻页游标，传上次返回最后一条 loanId
}

func (api *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI) LoanID(v string) *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI {
	api.req.LoanID = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI) Coin(v string) *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI) StartTime(v string) *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI) EndTime(v string) *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI) Limit(v string) *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossRecordBorrowHistoryAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossRecordRepayHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossRecordRepayHistoryReq
}

type PrivateRestClassicMarginCrossRecordRepayHistoryReq struct {
	RepayID    *string `json:"repayId"`    // 还款 ID
	Coin       *string `json:"coin"`       // 币种
	StartTime  *string `json:"startTime"`  // 开始时间，Unix 毫秒时间戳
	EndTime    *string `json:"endTime"`    // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit      *string `json:"limit"`      // 查询条数，默认 100，最大 500
	IDLessThan *string `json:"idLessThan"` // 翻页游标，传上次返回最后一条 repayId
}

func (api *PrivateRestClassicMarginCrossRecordRepayHistoryAPI) RepayID(v string) *PrivateRestClassicMarginCrossRecordRepayHistoryAPI {
	api.req.RepayID = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordRepayHistoryAPI) Coin(v string) *PrivateRestClassicMarginCrossRecordRepayHistoryAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordRepayHistoryAPI) StartTime(v string) *PrivateRestClassicMarginCrossRecordRepayHistoryAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordRepayHistoryAPI) EndTime(v string) *PrivateRestClassicMarginCrossRecordRepayHistoryAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordRepayHistoryAPI) Limit(v string) *PrivateRestClassicMarginCrossRecordRepayHistoryAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordRepayHistoryAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossRecordRepayHistoryAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossRecordInterestHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossRecordInterestHistoryReq
}

type PrivateRestClassicMarginCrossRecordInterestHistoryReq struct {
	Coin       *string `json:"coin"`       // 币种
	StartTime  *string `json:"startTime"`  // 开始时间，Unix 毫秒时间戳
	EndTime    *string `json:"endTime"`    // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit      *string `json:"limit"`      // 查询条数，默认 100，最大 500
	IDLessThan *string `json:"idLessThan"` // 翻页游标，传上次返回最后一条 loanId
}

func (api *PrivateRestClassicMarginCrossRecordInterestHistoryAPI) Coin(v string) *PrivateRestClassicMarginCrossRecordInterestHistoryAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordInterestHistoryAPI) StartTime(v string) *PrivateRestClassicMarginCrossRecordInterestHistoryAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordInterestHistoryAPI) EndTime(v string) *PrivateRestClassicMarginCrossRecordInterestHistoryAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordInterestHistoryAPI) Limit(v string) *PrivateRestClassicMarginCrossRecordInterestHistoryAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordInterestHistoryAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossRecordInterestHistoryAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossRecordLiquidationHistoryReq
}

type PrivateRestClassicMarginCrossRecordLiquidationHistoryReq struct {
	StartTime  *string `json:"startTime"`  // 开始时间，Unix 毫秒时间戳
	EndTime    *string `json:"endTime"`    // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit      *string `json:"limit"`      // 查询条数，默认 100，最大 500
	IDLessThan *string `json:"idLessThan"` // 翻页游标，传上次返回最后一条 liqId
}

func (api *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI) StartTime(v string) *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI) EndTime(v string) *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI) Limit(v string) *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossRecordLiquidationHistoryAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossRecordFinancialRecordsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossRecordFinancialRecordsReq
}

type PrivateRestClassicMarginCrossRecordFinancialRecordsReq struct {
	MarginType *string `json:"marginType"` // 流水类型：transfer_in/transfer_out/borrow/repay/... 等
	Coin       *string `json:"coin"`       // 币种
	StartTime  *string `json:"startTime"`  // 开始时间，Unix 毫秒时间戳
	EndTime    *string `json:"endTime"`    // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit      *string `json:"limit"`      // 查询条数，默认 100，最大 500
	IDLessThan *string `json:"idLessThan"` // 翻页游标，传上次返回最后一条 marginId
}

func (api *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI) MarginType(v string) *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI {
	api.req.MarginType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI) Coin(v string) *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI) StartTime(v string) *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI) EndTime(v string) *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI) Limit(v string) *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossRecordFinancialRecordsAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountAssetsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountAssetsReq
}

type PrivateRestClassicMarginIsolatedAccountAssetsReq struct {
	Symbol *string `json:"symbol"` // 交易对，如 BTCUSDT
}

func (api *PrivateRestClassicMarginIsolatedAccountAssetsAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountAssetsAPI {
	api.req.Symbol = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountBorrowAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountBorrowReq
}

type PrivateRestClassicMarginIsolatedAccountBorrowReq struct {
	Symbol       *string `json:"symbol"`       // 借出交易对，如 BTCUSDT
	Coin         *string `json:"coin"`         // 借出币种，如 BTC
	BorrowAmount *string `json:"borrowAmount"` // 借出数量（最长 8 位小数）
	ClientOid    *string `json:"clientOid"`    // 客户自定义订单 ID
}

func (api *PrivateRestClassicMarginIsolatedAccountBorrowAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountBorrowAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountBorrowAPI) Coin(v string) *PrivateRestClassicMarginIsolatedAccountBorrowAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountBorrowAPI) BorrowAmount(v string) *PrivateRestClassicMarginIsolatedAccountBorrowAPI {
	api.req.BorrowAmount = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountBorrowAPI) ClientOid(v string) *PrivateRestClassicMarginIsolatedAccountBorrowAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountRepayAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountRepayReq
}

type PrivateRestClassicMarginIsolatedAccountRepayReq struct {
	RepayAmount *string `json:"repayAmount"` // 还款数量（最长 8 位小数）
	Coin        *string `json:"coin"`        // 还款币种
	Symbol      *string `json:"symbol"`      // 还款交易对，如 BTCUSDT
	ClientOid   *string `json:"clientOid"`   // 客户自定义订单 ID
}

func (api *PrivateRestClassicMarginIsolatedAccountRepayAPI) RepayAmount(v string) *PrivateRestClassicMarginIsolatedAccountRepayAPI {
	api.req.RepayAmount = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountRepayAPI) Coin(v string) *PrivateRestClassicMarginIsolatedAccountRepayAPI {
	api.req.Coin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountRepayAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountRepayAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountRepayAPI) ClientOid(v string) *PrivateRestClassicMarginIsolatedAccountRepayAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountRiskRateAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountRiskRateReq
}

type PrivateRestClassicMarginIsolatedAccountRiskRateReq struct {
	Symbol   *string `json:"symbol"`   // 交易对，如 BTCUSDT
	PageNum  *string `json:"pageNum"`  // 页码，默认 1
	PageSize *string `json:"pageSize"` // 每页大小，默认 100，最大 500
}

func (api *PrivateRestClassicMarginIsolatedAccountRiskRateAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountRiskRateAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountRiskRateAPI) PageNum(v string) *PrivateRestClassicMarginIsolatedAccountRiskRateAPI {
	api.req.PageNum = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountRiskRateAPI) PageSize(v string) *PrivateRestClassicMarginIsolatedAccountRiskRateAPI {
	api.req.PageSize = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountTierDataAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountTierDataReq
}

type PrivateRestClassicMarginIsolatedAccountTierDataReq struct {
	Symbol *string `json:"symbol"` // 交易对，如 BTCUSDT
}

func (api *PrivateRestClassicMarginIsolatedAccountTierDataAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountTierDataAPI {
	api.req.Symbol = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitReq
}

type PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitReq struct {
	Symbol *string `json:"symbol"` // 交易对，如 BTCUSDT
}

func (api *PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitAPI {
	api.req.Symbol = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountReq
}

type PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountReq struct {
	Symbol *string `json:"symbol"` // 交易对名称，如 BTCUSDT
}

func (api *PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountAPI {
	api.req.Symbol = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountReq
}

type PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountReq struct {
	Symbol *string `json:"symbol"` // 交易对，如 BTCUSDT
}

func (api *PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountAPI {
	api.req.Symbol = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedAccountFlashRepayAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountFlashRepayReq
}

type PrivateRestClassicMarginIsolatedAccountFlashRepayReq struct {
	SymbolList []string `json:"symbolList"` // 逐仓交易对数组；不传默认全部交易对，最多 100 个
}

func (api *PrivateRestClassicMarginIsolatedAccountFlashRepayAPI) AddSymbol(v string) *PrivateRestClassicMarginIsolatedAccountFlashRepayAPI {
	api.req.SymbolList = append(api.req.SymbolList, v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountFlashRepayAPI) SymbolList(v []string) *PrivateRestClassicMarginIsolatedAccountFlashRepayAPI {
	api.req.SymbolList = v
	return api
}

type PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusReq
}

type PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusReq struct {
	IdList []string `json:"idList"` // 逐仓还款 ID 列表，单次最多 100 条
}

func (api *PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI) AddId(v string) *PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI {
	api.req.IdList = append(api.req.IdList, v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI) IdList(v []string) *PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusAPI {
	api.req.IdList = v
	return api
}

type PrivateRestClassicMarginCrossAccountTierDataAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossAccountTierDataReq
}

type PrivateRestClassicMarginCrossAccountTierDataReq struct {
	Coin *string `json:"coin"` // 币种，如 BTC
}

func (api *PrivateRestClassicMarginCrossAccountTierDataAPI) Coin(v string) *PrivateRestClassicMarginCrossAccountTierDataAPI {
	api.req.Coin = GetPointer(v)
	return api
}
