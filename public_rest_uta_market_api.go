package mybitgetapi

// GET 获取交易产品信息 查询在线交易币对信息
func (client *PublicRestClient) NewPublicRestUtaMarketInstruments() *PublicRestUtaMarketInstrumentsAPI {
	return &PublicRestUtaMarketInstrumentsAPI{
		client: client,
		req:    &PublicRestUtaMarketInstrumentsReq{},
	}
}

func (api *PublicRestUtaMarketInstrumentsAPI) Do() (*BitgetRestRes[PublicRestUtaMarketInstrumentsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestUtaAPIMap[PublicRestUtaMarketInstruments])
	return bitgetCallAPI[PublicRestUtaMarketInstrumentsRes](api.client.c, url, nil, GET)
}

// GET 获取行情数据 查询在线交易币对行情数据
func (client *PublicRestClient) NewPublicRestUtaMarketTickers() *PublicRestUtaMarketTickersAPI {
	return &PublicRestUtaMarketTickersAPI{
		client: client,
		req:    &PublicRestUtaMarketTickersReq{},
	}
}

func (api *PublicRestUtaMarketTickersAPI) Do() (*BitgetRestRes[PublicRestUtaMarketTickersRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestUtaAPIMap[PublicRestUtaMarketTickers])
	return bitgetCallAPI[PublicRestUtaMarketTickersRes](api.client.c, url, nil, GET)
}

// GET 获取深度数据 查询订单簿深度
func (client *PublicRestClient) NewPublicRestUtaMarketOrderBook() *PublicRestUtaMarketOrderBookAPI {
	return &PublicRestUtaMarketOrderBookAPI{
		client: client,
		req:    &PublicRestUtaMarketOrderBookReq{},
	}
}

func (api *PublicRestUtaMarketOrderBookAPI) Do() (*BitgetRestRes[PublicRestUtaMarketOrderBookRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestUtaAPIMap[PublicRestUtaMarketOrderBook])
	middleRes, err := bitgetCallAPI[PublicRestUtaMarketOrderBookResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestUtaMarketOrderBookRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取K线数据
func (client *PublicRestClient) NewPublicRestUtaMarketCandles() *PublicRestUtaMarketCandlesAPI {
	return &PublicRestUtaMarketCandlesAPI{
		client: client,
		req:    &PublicRestUtaMarketCandlesReq{},
	}
}

func (api *PublicRestUtaMarketCandlesAPI) Do() (*BitgetRestRes[PublicRestUtaMarketCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestUtaAPIMap[PublicRestUtaMarketCandles])
	middleRes, err := bitgetCallAPI[PublicRestUtaMarketCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestUtaMarketCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取历史K线数据
func (client *PublicRestClient) NewPublicRestUtaMarketHistoryCandles() *PublicRestUtaMarketHistoryCandlesAPI {
	return &PublicRestUtaMarketHistoryCandlesAPI{
		client: client,
		req:    &PublicRestUtaMarketHistoryCandlesReq{},
	}
}

func (api *PublicRestUtaMarketHistoryCandlesAPI) Do() (*BitgetRestRes[PublicRestUtaMarketHistoryCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestUtaAPIMap[PublicRestUtaMarketHistoryCandles])
	middleRes, err := bitgetCallAPI[PublicRestUtaMarketHistoryCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestUtaMarketHistoryCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取当前资金费率
func (client *PublicRestClient) NewPublicRestUtaMarketCurrentFundRate() *PublicRestUtaMarketCurrentFundRateAPI {
	return &PublicRestUtaMarketCurrentFundRateAPI{
		client: client,
		req:    &PublicRestUtaMarketCurrentFundRateReq{},
	}
}

func (api *PublicRestUtaMarketCurrentFundRateAPI) Do() (*BitgetRestRes[PublicRestUtaMarketCurrentFundRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestUtaAPIMap[PublicRestUtaMarketCurrentFundRate])
	return bitgetCallAPI[PublicRestUtaMarketCurrentFundRateRes](api.client.c, url, nil, GET)
}

// GET 获取历史资金费率
func (client *PublicRestClient) NewPublicRestUtaMarketHistoryFundRate() *PublicRestUtaMarketHistoryFundRateAPI {
	return &PublicRestUtaMarketHistoryFundRateAPI{
		client: client,
		req:    &PublicRestUtaMarketHistoryFundRateReq{},
	}
}

func (api *PublicRestUtaMarketHistoryFundRateAPI) Do() (*BitgetRestRes[PublicRestUtaMarketHistoryFundRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestUtaAPIMap[PublicRestUtaMarketHistoryFundRate])
	return bitgetCallAPI[PublicRestUtaMarketHistoryFundRateRes](api.client.c, url, nil, GET)
}
