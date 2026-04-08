package mybitgetapi

type PrivateRestClassicMarginMarketInterestRateRecordAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginMarketInterestRateRecordReq
}

// 杠杆利率（interest-rate-record）请求参数
type PrivateRestClassicMarginMarketInterestRateRecordReq struct {
	Coin *string `json:"coin"` // String 是 币种名称
}

func (api *PrivateRestClassicMarginMarketInterestRateRecordAPI) Coin(coin string) *PrivateRestClassicMarginMarketInterestRateRecordAPI {
	api.req.Coin = GetPointer(coin)
	return api
}
