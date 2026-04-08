package mybitgetapi

type PrivateRestUtaAccountBalanceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountBalanceReq
}

type PrivateRestUtaAccountBalanceReq struct {}

type PrivateRestUtaAccountFundingAssetsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountFundingAssetsReq
}

type PrivateRestUtaAccountFundingAssetsReq struct {
	Coin *string `json:"coin"` // String 否 币种名称
}

// String 否 币种名称
func (api *PrivateRestUtaAccountFundingAssetsAPI) Coin(coin string) *PrivateRestUtaAccountFundingAssetsAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

type PrivateRestUtaAccountSettingAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountSettingReq
}

type PrivateRestUtaAccountSettingReq struct {
}

type PrivateRestUtaAccountSetLeverageAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountSetLeverageReq
}

type PrivateRestUtaAccountSetLeverageReq struct {
	Category *string `json:"category"` // String 是 产品类型 MARGIN USDT-FUTURES COIN-FUTURES USDC-FUTURES
	Symbol   *string `json:"symbol"`   // String 否 交易对名称（合约设置杠杆时必填）
	Leverage *string `json:"leverage"` // String 是 杠杆倍数
	Coin     *string `json:"coin"`     // String 否 币种名称（杠杆现货设置杠杆时必填）
	PosSide  *string `json:"posSide"`  // String 否 持仓方向 long/short（逐仓时必填）
}

func (api *PrivateRestUtaAccountSetLeverageAPI) Category(category string) *PrivateRestUtaAccountSetLeverageAPI {
	api.req.Category = GetPointer(category)
	return api
}

func (api *PrivateRestUtaAccountSetLeverageAPI) Symbol(symbol string) *PrivateRestUtaAccountSetLeverageAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PrivateRestUtaAccountSetLeverageAPI) Leverage(leverage string) *PrivateRestUtaAccountSetLeverageAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}

func (api *PrivateRestUtaAccountSetLeverageAPI) Coin(coin string) *PrivateRestUtaAccountSetLeverageAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *PrivateRestUtaAccountSetLeverageAPI) PosSide(posSide string) *PrivateRestUtaAccountSetLeverageAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}

type PrivateRestUtaAccountSetHoldModeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountSetHoldModeReq
}

type PrivateRestUtaAccountSetHoldModeReq struct {
	HoldMode *string `json:"holdMode"` // String 是 持仓模式 one_way_mode / hedge_mode
}

func (api *PrivateRestUtaAccountSetHoldModeAPI) HoldMode(holdMode string) *PrivateRestUtaAccountSetHoldModeAPI {
	api.req.HoldMode = GetPointer(holdMode)
	return api
}

type PrivateRestUtaAccountFinancialRecordsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountFinancialRecordsReq
}

type PrivateRestUtaAccountFinancialRecordsReq struct {
	Category  *string `json:"category"`  // String 是 产品类型 SPOT MARGIN USDT-FUTURES COIN-FUTURES USDC-FUTURES OTHER
	Coin      *string `json:"coin"`      // String 否 币种名称
	Type      *string `json:"type"`      // String 否 流水类型
	StartTime *string `json:"startTime"` // String 否 开始时间 Unix毫秒时间戳
	EndTime   *string `json:"endTime"`   // String 否 结束时间 Unix毫秒时间戳
	Limit     *string `json:"limit"`     // String 否 每页数量 默认100 最大100
	Cursor    *string `json:"cursor"`    // String 否 游标
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) Category(category string) *PrivateRestUtaAccountFinancialRecordsAPI {
	api.req.Category = GetPointer(category)
	return api
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) Coin(coin string) *PrivateRestUtaAccountFinancialRecordsAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) Type(financialType string) *PrivateRestUtaAccountFinancialRecordsAPI {
	api.req.Type = GetPointer(financialType)
	return api
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) StartTime(startTime string) *PrivateRestUtaAccountFinancialRecordsAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) EndTime(endTime string) *PrivateRestUtaAccountFinancialRecordsAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) Limit(limit string) *PrivateRestUtaAccountFinancialRecordsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

func (api *PrivateRestUtaAccountFinancialRecordsAPI) Cursor(cursor string) *PrivateRestUtaAccountFinancialRecordsAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type PrivateRestUtaAccountFeeRateAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountFeeRateReq
}

type PrivateRestUtaAccountFeeRateReq struct {
	Symbol   *string `json:"symbol"`   // String 是 交易对名称 例如：BTCUSDT
	Category *string `json:"category"` // String 是 产品类型 SPOT MARGIN USDT-FUTURES COIN-FUTURES USDC-FUTURES
}

func (api *PrivateRestUtaAccountFeeRateAPI) Symbol(symbol string) *PrivateRestUtaAccountFeeRateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PrivateRestUtaAccountFeeRateAPI) Category(category string) *PrivateRestUtaAccountFeeRateAPI {
	api.req.Category = GetPointer(category)
	return api
}

type PrivateRestUtaAccountMaxTransferableAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountMaxTransferableReq
}

type PrivateRestUtaAccountMaxTransferableReq struct {
	Coin *string `json:"coin"` // String 是 币种名称
}

func (api *PrivateRestUtaAccountMaxTransferableAPI) Coin(coin string) *PrivateRestUtaAccountMaxTransferableAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

type PrivateRestUtaAccountTransferableCoinsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountTransferableCoinsReq
}

type PrivateRestUtaAccountTransferableCoinsReq struct {
	FromType *string `json:"fromType"` // String 是 转出账户类型
	ToType   *string `json:"toType"`   // String 是 转入账户类型
}

func (api *PrivateRestUtaAccountTransferableCoinsAPI) FromType(fromType string) *PrivateRestUtaAccountTransferableCoinsAPI {
	api.req.FromType = GetPointer(fromType)
	return api
}

func (api *PrivateRestUtaAccountTransferableCoinsAPI) ToType(toType string) *PrivateRestUtaAccountTransferableCoinsAPI {
	api.req.ToType = GetPointer(toType)
	return api
}

type PrivateRestUtaAccountTransferAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountTransferReq
}

type PrivateRestUtaAccountTransferReq struct {
	FromType    *string `json:"fromType"`    // String 是 转出账户类型
	ToType      *string `json:"toType"`      // String 是 转入账户类型
	Amount      *string `json:"amount"`      // String 是 划转数量
	Coin        *string `json:"coin"`        // String 是 币种 例如 BTC
	Symbol      *string `json:"symbol"`      // String 否 逐仓杠杆交易对
	AllowBorrow *string `json:"allowBorrow"` // String 否 是否允许借币划转 yes/no
}

func (api *PrivateRestUtaAccountTransferAPI) FromType(fromType string) *PrivateRestUtaAccountTransferAPI {
	api.req.FromType = GetPointer(fromType)
	return api
}

func (api *PrivateRestUtaAccountTransferAPI) ToType(toType string) *PrivateRestUtaAccountTransferAPI {
	api.req.ToType = GetPointer(toType)
	return api
}

func (api *PrivateRestUtaAccountTransferAPI) Amount(amount string) *PrivateRestUtaAccountTransferAPI {
	api.req.Amount = GetPointer(amount)
	return api
}

func (api *PrivateRestUtaAccountTransferAPI) Coin(coin string) *PrivateRestUtaAccountTransferAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *PrivateRestUtaAccountTransferAPI) Symbol(symbol string) *PrivateRestUtaAccountTransferAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PrivateRestUtaAccountTransferAPI) AllowBorrow(allowBorrow string) *PrivateRestUtaAccountTransferAPI {
	api.req.AllowBorrow = GetPointer(allowBorrow)
	return api
}

type PrivateRestUtaAccountSwitchAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountSwitchReq
}

type PrivateRestUtaAccountSwitchReq struct{}

type PrivateRestUtaAccountSwitchStatusAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaAccountSwitchStatusReq
}

type PrivateRestUtaAccountSwitchStatusReq struct{}