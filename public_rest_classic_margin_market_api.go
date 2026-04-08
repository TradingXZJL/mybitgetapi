package mybitgetapi

// GET 获取支持杠杆的所有交易对
func (client *PublicRestClient) NewPublicRestClassicMarginMarketCurrencies() *PublicRestClassicMarginMarketCurrenciesAPI {
	return &PublicRestClassicMarginMarketCurrenciesAPI{
		client: client,
		req:    &PublicRestClassicMarginMarketCurrenciesReq{},
	}
}

func (api *PublicRestClassicMarginMarketCurrenciesAPI) Do() (*BitgetRestRes[PublicRestClassicMarginMarketCurrenciesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicMarginMarketCurrencies])
	return bitgetCallAPI[PublicRestClassicMarginMarketCurrenciesRes](api.client.c, url, nil, GET)
}
