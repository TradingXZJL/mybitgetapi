package mybitgetapi

type PublicRestClassicMarginMarketCurrenciesAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicMarginMarketCurrenciesReq
}

type PublicRestClassicMarginMarketCurrenciesReq struct{}

type PublicRestClassicMarginMarketInterestRateRecordAPI struct {
	client *PublicRestClient
	req    *PublicRestClassicMarginMarketInterestRateRecordReq
}

type PublicRestClassicMarginMarketInterestRateRecordReq struct {
	Coin *string `json:"coin"` // String 是 币种名称
}

func (api *PublicRestClassicMarginMarketInterestRateRecordAPI) Coin(coin string) *PublicRestClassicMarginMarketInterestRateRecordAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
