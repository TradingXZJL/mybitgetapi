package mybitgetapi

// GET 获取杠杆利率

func (client *PrivateRestClient) NewPrivateRestClassicMarginMarketInterestRateRecord() *PrivateRestClassicMarginMarketInterestRateRecordAPI {
	return &PrivateRestClassicMarginMarketInterestRateRecordAPI{
		client: client,
		req:    &PrivateRestClassicMarginMarketInterestRateRecordReq{},
	}
}

func (api *PrivateRestClassicMarginMarketInterestRateRecordAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginMarketInterestRateRecordRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(
		REST,
		api.req,
		PrivateRestClassicAPIMap[PrivateRestClassicMarginMarketInterestRateRecord],
	)

	return bitgetCallApiWithSecret[PrivateRestClassicMarginMarketInterestRateRecordRes](api.client.c, url, NIL_REQBODY, GET)
}

