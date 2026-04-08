package mybitgetapi

type PrivateRestClassicFuturesAccountSingleAccountAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSingleAccountReq
}

type PrivateRestClassicFuturesAccountSingleAccountReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	MarginCoin  *string `json:"marginCoin"`
}

func (api *PrivateRestClassicFuturesAccountSingleAccountAPI) Symbol(symbol string) *PrivateRestClassicFuturesAccountSingleAccountAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesAccountSingleAccountAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSingleAccountAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountSingleAccountAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesAccountSingleAccountAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}

type PrivateRestClassicFuturesAccountListAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountListReq
}

type PrivateRestClassicFuturesAccountListReq struct {
	ProductType *string `json:"productType"`
}

func (api *PrivateRestClassicFuturesAccountListAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountListAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}

type PrivateRestClassicFuturesAccountSubAccountAssetsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSubAccountAssetsReq
}

type PrivateRestClassicFuturesAccountSubAccountAssetsReq struct {
	ProductType *string `json:"productType"`
}

func (api *PrivateRestClassicFuturesAccountSubAccountAssetsAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSubAccountAssetsAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}

type PrivateRestClassicFuturesAccountInterestHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountInterestHistoryReq
}

type PrivateRestClassicFuturesAccountInterestHistoryReq struct {
	Coin        *string `json:"coin"`
	ProductType *string `json:"productType"`
	IDLessThan  *string `json:"idLessThan"`
	StartTime   *string `json:"startTime"`
	EndTime     *string `json:"endTime"`
	Limit       *string `json:"limit"`
}

func (api *PrivateRestClassicFuturesAccountInterestHistoryAPI) Coin(coin string) *PrivateRestClassicFuturesAccountInterestHistoryAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
func (api *PrivateRestClassicFuturesAccountInterestHistoryAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountInterestHistoryAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountInterestHistoryAPI) IDLessThan(idLessThan string) *PrivateRestClassicFuturesAccountInterestHistoryAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}
func (api *PrivateRestClassicFuturesAccountInterestHistoryAPI) StartTime(startTime string) *PrivateRestClassicFuturesAccountInterestHistoryAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicFuturesAccountInterestHistoryAPI) EndTime(endTime string) *PrivateRestClassicFuturesAccountInterestHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicFuturesAccountInterestHistoryAPI) Limit(limit string) *PrivateRestClassicFuturesAccountInterestHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestClassicFuturesAccountMaxOpenAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountMaxOpenReq
}

type PrivateRestClassicFuturesAccountMaxOpenReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	MarginCoin  *string `json:"marginCoin"`
	PosSide     *string `json:"posSide"`
	OrderType   *string `json:"orderType"`
	OpenPrice   *string `json:"openPrice"`
}

func (api *PrivateRestClassicFuturesAccountMaxOpenAPI) Symbol(symbol string) *PrivateRestClassicFuturesAccountMaxOpenAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesAccountMaxOpenAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountMaxOpenAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountMaxOpenAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesAccountMaxOpenAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}
func (api *PrivateRestClassicFuturesAccountMaxOpenAPI) PosSide(posSide string) *PrivateRestClassicFuturesAccountMaxOpenAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}
func (api *PrivateRestClassicFuturesAccountMaxOpenAPI) OrderType(orderType string) *PrivateRestClassicFuturesAccountMaxOpenAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}
func (api *PrivateRestClassicFuturesAccountMaxOpenAPI) OpenPrice(openPrice string) *PrivateRestClassicFuturesAccountMaxOpenAPI {
	api.req.OpenPrice = GetPointer(openPrice)
	return api
}

type PrivateRestClassicFuturesAccountLiquidationPriceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountLiquidationPriceReq
}

type PrivateRestClassicFuturesAccountLiquidationPriceReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	MarginCoin  *string `json:"marginCoin"`
	PosSide     *string `json:"posSide"`
	OrderType   *string `json:"orderType"`
	OpenAmount  *string `json:"openAmount"`
	OpenPrice   *string `json:"openPrice"`
}

func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) Symbol(symbol string) *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}
func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) PosSide(posSide string) *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}
func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) OrderType(orderType string) *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}
func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) OpenAmount(openAmount string) *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	api.req.OpenAmount = GetPointer(openAmount)
	return api
}
func (api *PrivateRestClassicFuturesAccountLiquidationPriceAPI) OpenPrice(openPrice string) *PrivateRestClassicFuturesAccountLiquidationPriceAPI {
	api.req.OpenPrice = GetPointer(openPrice)
	return api
}

type PrivateRestClassicFuturesAccountOpenCountAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountOpenCountReq
}

type PrivateRestClassicFuturesAccountOpenCountReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	MarginCoin  *string `json:"marginCoin"`
	OpenAmount  *string `json:"openAmount"`
	OpenPrice   *string `json:"openPrice"`
	Leverage    *string `json:"leverage"`
}

func (api *PrivateRestClassicFuturesAccountOpenCountAPI) Symbol(symbol string) *PrivateRestClassicFuturesAccountOpenCountAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesAccountOpenCountAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountOpenCountAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountOpenCountAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesAccountOpenCountAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}
func (api *PrivateRestClassicFuturesAccountOpenCountAPI) OpenAmount(openAmount string) *PrivateRestClassicFuturesAccountOpenCountAPI {
	api.req.OpenAmount = GetPointer(openAmount)
	return api
}
func (api *PrivateRestClassicFuturesAccountOpenCountAPI) OpenPrice(openPrice string) *PrivateRestClassicFuturesAccountOpenCountAPI {
	api.req.OpenPrice = GetPointer(openPrice)
	return api
}
func (api *PrivateRestClassicFuturesAccountOpenCountAPI) Leverage(leverage string) *PrivateRestClassicFuturesAccountOpenCountAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}

type PrivateRestClassicFuturesAccountSetLeverageAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSetLeverageReq
}

type PrivateRestClassicFuturesAccountSetLeverageReq struct {
	Symbol        *string `json:"symbol"`
	ProductType   *string `json:"productType"`
	MarginCoin    *string `json:"marginCoin"`
	Leverage      *string `json:"leverage"`
	LongLeverage  *string `json:"longLeverage"`
	ShortLeverage *string `json:"shortLeverage"`
	HoldSide      *string `json:"holdSide"`
}

func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) Symbol(symbol string) *PrivateRestClassicFuturesAccountSetLeverageAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSetLeverageAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesAccountSetLeverageAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) Leverage(leverage string) *PrivateRestClassicFuturesAccountSetLeverageAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) LongLeverage(longLeverage string) *PrivateRestClassicFuturesAccountSetLeverageAPI {
	api.req.LongLeverage = GetPointer(longLeverage)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) ShortLeverage(shortLeverage string) *PrivateRestClassicFuturesAccountSetLeverageAPI {
	api.req.ShortLeverage = GetPointer(shortLeverage)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetLeverageAPI) HoldSide(holdSide string) *PrivateRestClassicFuturesAccountSetLeverageAPI {
	api.req.HoldSide = GetPointer(holdSide)
	return api
}

type PrivateRestClassicFuturesAccountSetAllLeverageAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSetAllLeverageReq
}

type PrivateRestClassicFuturesAccountSetAllLeverageReq struct {
	ProductType *string `json:"productType"`
	Leverage    *string `json:"leverage"`
}

func (api *PrivateRestClassicFuturesAccountSetAllLeverageAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSetAllLeverageAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetAllLeverageAPI) Leverage(leverage string) *PrivateRestClassicFuturesAccountSetAllLeverageAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}

type PrivateRestClassicFuturesAccountSetMarginAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSetMarginReq
}

type PrivateRestClassicFuturesAccountSetMarginReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	MarginCoin  *string `json:"marginCoin"`
	HoldSide    *string `json:"holdSide"`
	Amount      *string `json:"amount"`
}

func (api *PrivateRestClassicFuturesAccountSetMarginAPI) Symbol(symbol string) *PrivateRestClassicFuturesAccountSetMarginAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetMarginAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSetMarginAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetMarginAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesAccountSetMarginAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetMarginAPI) HoldSide(holdSide string) *PrivateRestClassicFuturesAccountSetMarginAPI {
	api.req.HoldSide = GetPointer(holdSide)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetMarginAPI) Amount(amount string) *PrivateRestClassicFuturesAccountSetMarginAPI {
	api.req.Amount = GetPointer(amount)
	return api
}

type PrivateRestClassicFuturesAccountSetAssetModeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSetAssetModeReq
}

type PrivateRestClassicFuturesAccountSetAssetModeReq struct {
	ProductType *string `json:"productType"`
	AssetMode   *string `json:"assetMode"`
}

func (api *PrivateRestClassicFuturesAccountSetAssetModeAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSetAssetModeAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetAssetModeAPI) AssetMode(assetMode string) *PrivateRestClassicFuturesAccountSetAssetModeAPI {
	api.req.AssetMode = GetPointer(assetMode)
	return api
}

type PrivateRestClassicFuturesAccountSetMarginModeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSetMarginModeReq
}

type PrivateRestClassicFuturesAccountSetMarginModeReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	MarginCoin  *string `json:"marginCoin"`
	MarginMode  *string `json:"marginMode"`
}

func (api *PrivateRestClassicFuturesAccountSetMarginModeAPI) Symbol(symbol string) *PrivateRestClassicFuturesAccountSetMarginModeAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetMarginModeAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSetMarginModeAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetMarginModeAPI) MarginCoin(marginCoin string) *PrivateRestClassicFuturesAccountSetMarginModeAPI {
	api.req.MarginCoin = GetPointer(marginCoin)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetMarginModeAPI) MarginMode(marginMode string) *PrivateRestClassicFuturesAccountSetMarginModeAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

type PrivateRestClassicFuturesAccountSetPositionModeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountSetPositionModeReq
}

type PrivateRestClassicFuturesAccountSetPositionModeReq struct {
	ProductType *string `json:"productType"`
	PosMode     *string `json:"posMode"`
}

func (api *PrivateRestClassicFuturesAccountSetPositionModeAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountSetPositionModeAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountSetPositionModeAPI) PosMode(posMode string) *PrivateRestClassicFuturesAccountSetPositionModeAPI {
	api.req.PosMode = GetPointer(posMode)
	return api
}

type PrivateRestClassicFuturesAccountBillAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountBillReq
}

type PrivateRestClassicFuturesAccountBillReq struct {
	ProductType  *string `json:"productType"`
	Coin         *string `json:"coin"`
	BusinessType *string `json:"businessType"`
	OnlyFunding  *string `json:"onlyFunding"`
	IDLessThan   *string `json:"idLessThan"`
	StartTime    *string `json:"startTime"`
	EndTime      *string `json:"endTime"`
	Limit        *string `json:"limit"`
}

func (api *PrivateRestClassicFuturesAccountBillAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
func (api *PrivateRestClassicFuturesAccountBillAPI) Coin(coin string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
func (api *PrivateRestClassicFuturesAccountBillAPI) BusinessType(businessType string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}
func (api *PrivateRestClassicFuturesAccountBillAPI) OnlyFunding(onlyFunding string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.OnlyFunding = GetPointer(onlyFunding)
	return api
}
func (api *PrivateRestClassicFuturesAccountBillAPI) IDLessThan(idLessThan string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}
func (api *PrivateRestClassicFuturesAccountBillAPI) StartTime(startTime string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicFuturesAccountBillAPI) EndTime(endTime string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicFuturesAccountBillAPI) Limit(limit string) *PrivateRestClassicFuturesAccountBillAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestClassicFuturesAccountIsolatedSymbolsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesAccountIsolatedSymbolsReq
}

type PrivateRestClassicFuturesAccountIsolatedSymbolsReq struct {
	ProductType *string `json:"productType"`
}

func (api *PrivateRestClassicFuturesAccountIsolatedSymbolsAPI) ProductType(productType string) *PrivateRestClassicFuturesAccountIsolatedSymbolsAPI {
	api.req.ProductType = GetPointer(productType)
	return api
}
