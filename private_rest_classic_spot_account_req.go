package mybitgetapi

type PrivateRestClassicSpotAccountInfoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountInfoReq
}

type PrivateRestClassicSpotAccountInfoReq struct{}

type PrivateRestClassicSpotAccountAssetsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountAssetsReq
}

type PrivateRestClassicSpotAccountAssetsReq struct {
	Coin      *string `json:"coin"`      // String 否 币种名称，如 USDT
	AssetType *string `json:"assetType"` // String 否 资产类型 hold_only/all
}

func (api *PrivateRestClassicSpotAccountAssetsAPI) Coin(coin string) *PrivateRestClassicSpotAccountAssetsAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *PrivateRestClassicSpotAccountAssetsAPI) AssetType(assetType string) *PrivateRestClassicSpotAccountAssetsAPI {
	api.req.AssetType = GetPointer(assetType)
	return api
}

type PrivateRestClassicSpotAccountBillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountBillsReq
}

type PrivateRestClassicSpotAccountBillsReq struct {
	Coin         *string `json:"coin"`         // String 否 币种名称
	GroupType    *string `json:"groupType"`    // String 否 账单类型
	BusinessType *string `json:"businessType"` // String 否 业务类型
	StartTime    *string `json:"startTime"`    // String 否 开始时间
	EndTime      *string `json:"endTime"`      // String 否 结束时间
	Limit        *string `json:"limit"`        // String 否 数量，默认100，最大500
	IDLessThan   *string `json:"idLessThan"`   // String 否 分页游标 billId
}

func (api *PrivateRestClassicSpotAccountBillsAPI) Coin(coin string) *PrivateRestClassicSpotAccountBillsAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
func (api *PrivateRestClassicSpotAccountBillsAPI) GroupType(groupType string) *PrivateRestClassicSpotAccountBillsAPI {
	api.req.GroupType = GetPointer(groupType)
	return api
}
func (api *PrivateRestClassicSpotAccountBillsAPI) BusinessType(businessType string) *PrivateRestClassicSpotAccountBillsAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}
func (api *PrivateRestClassicSpotAccountBillsAPI) StartTime(startTime string) *PrivateRestClassicSpotAccountBillsAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicSpotAccountBillsAPI) EndTime(endTime string) *PrivateRestClassicSpotAccountBillsAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicSpotAccountBillsAPI) Limit(limit string) *PrivateRestClassicSpotAccountBillsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestClassicSpotAccountBillsAPI) IDLessThan(idLessThan string) *PrivateRestClassicSpotAccountBillsAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}

type PrivateRestClassicSpotWalletTransferAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotWalletTransferReq
}

type PrivateRestClassicSpotWalletTransferReq struct {
	FromType  *string `json:"fromType"`  // String 是 转出账户类型
	ToType    *string `json:"toType"`    // String 是 转入账户类型
	Amount    *string `json:"amount"`    // String 是 划转数量
	Coin      *string `json:"coin"`      // String 是 币种
	Symbol    *string `json:"symbol"`    // String 否 逐仓交易对
	ClientOid *string `json:"clientOid"` // String 否 客户自定义ID
}

func (api *PrivateRestClassicSpotWalletTransferAPI) FromType(fromType string) *PrivateRestClassicSpotWalletTransferAPI {
	api.req.FromType = GetPointer(fromType)
	return api
}
func (api *PrivateRestClassicSpotWalletTransferAPI) ToType(toType string) *PrivateRestClassicSpotWalletTransferAPI {
	api.req.ToType = GetPointer(toType)
	return api
}
func (api *PrivateRestClassicSpotWalletTransferAPI) Amount(amount string) *PrivateRestClassicSpotWalletTransferAPI {
	api.req.Amount = GetPointer(amount)
	return api
}
func (api *PrivateRestClassicSpotWalletTransferAPI) Coin(coin string) *PrivateRestClassicSpotWalletTransferAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
func (api *PrivateRestClassicSpotWalletTransferAPI) Symbol(symbol string) *PrivateRestClassicSpotWalletTransferAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotWalletTransferAPI) ClientOid(clientOid string) *PrivateRestClassicSpotWalletTransferAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}

type PrivateRestClassicSpotAccountTransferRecordsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountTransferRecordsReq
}

type PrivateRestClassicSpotAccountTransferRecordsReq struct {
	Coin       *string `json:"coin"`       // String 是 币种名称
	FromType   *string `json:"fromType"`   // String 否 划转账户类型
	StartTime  *string `json:"startTime"`  // String 否 开始时间
	EndTime    *string `json:"endTime"`    // String 否 结束时间
	ClientOid  *string `json:"clientOid"`  // String 否 客户自定义订单ID
	PageNum    *string `json:"pageNum"`    // String 否 已废弃分页参数
	Limit      *string `json:"limit"`      // String 否 数量，默认100，最大500
	IDLessThan *string `json:"idLessThan"` // String 否 分页ID
}

func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) Coin(coin string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) FromType(fromType string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.FromType = GetPointer(fromType)
	return api
}
func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) StartTime(startTime string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) EndTime(endTime string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) ClientOid(clientOid string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}
func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) PageNum(pageNum string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.PageNum = GetPointer(pageNum)
	return api
}
func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) Limit(limit string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestClassicSpotAccountTransferRecordsAPI) IDLessThan(idLessThan string) *PrivateRestClassicSpotAccountTransferRecordsAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}

type PrivateRestClassicSpotAccountSwitchDeductAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountSwitchDeductReq
}

type PrivateRestClassicSpotAccountSwitchDeductReq struct {
	Deduct *string `json:"deduct"` // String 是 是否开启 on/off
}

func (api *PrivateRestClassicSpotAccountSwitchDeductAPI) Deduct(deduct string) *PrivateRestClassicSpotAccountSwitchDeductAPI {
	api.req.Deduct = GetPointer(deduct)
	return api
}

type PrivateRestClassicSpotAccountDeductInfoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountDeductInfoReq
}

type PrivateRestClassicSpotAccountDeductInfoReq struct{}

type PrivateRestClassicSpotAccountUpgradeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountUpgradeReq
}

type PrivateRestClassicSpotAccountUpgradeReq struct {
	SubUid *string `json:"subUid"` // String 否 子账户用户ID
}

func (api *PrivateRestClassicSpotAccountUpgradeAPI) SubUid(subUid string) *PrivateRestClassicSpotAccountUpgradeAPI {
	api.req.SubUid = GetPointer(subUid)
	return api
}

type PrivateRestClassicSpotAccountUpgradeStatusAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotAccountUpgradeStatusReq
}

type PrivateRestClassicSpotAccountUpgradeStatusReq struct {
	SubUid *string `json:"subUid"` // String 否 子账户用户ID
}

func (api *PrivateRestClassicSpotAccountUpgradeStatusAPI) SubUid(subUid string) *PrivateRestClassicSpotAccountUpgradeStatusAPI {
	api.req.SubUid = GetPointer(subUid)
	return api
}
