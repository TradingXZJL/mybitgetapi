package mybitgetapi

// GET 获取币种信息
func (client *PublicRestClient) NewPublicRestClassicSpotMarketCoins() *PublicRestClassicSpotMarketCoinsAPI {
	return &PublicRestClassicSpotMarketCoinsAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketCoinsReq{},
	}
}

func (api *PublicRestClassicSpotMarketCoinsAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketCoinsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketCoins])
	return bitgetCallAPI[PublicRestClassicSpotMarketCoinsRes](api.client.c, url, nil, GET)
}

// GET 获取交易对信息
func (client *PublicRestClient) NewPublicRestClassicSpotMarketSymbols() *PublicRestClassicSpotMarketSymbolsAPI {
	return &PublicRestClassicSpotMarketSymbolsAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketSymbolsReq{},
	}
}

func (api *PublicRestClassicSpotMarketSymbolsAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketSymbolsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketSymbols])
	return bitgetCallAPI[PublicRestClassicSpotMarketSymbolsRes](api.client.c, url, nil, GET)
}

// GET 获取行情信息
func (client *PublicRestClient) NewPublicRestClassicSpotMarketTickers() *PublicRestClassicSpotMarketTickersAPI {
	return &PublicRestClassicSpotMarketTickersAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketTickersReq{},
	}
}

func (api *PublicRestClassicSpotMarketTickersAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketTickersRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketTickers])
	return bitgetCallAPI[PublicRestClassicSpotMarketTickersRes](api.client.c, url, nil, GET)
}

// GET 获取合并交易深度
func (client *PublicRestClient) NewPublicRestClassicSpotMarketMergeDepth() *PublicRestClassicSpotMarketMergeDepthAPI {
	return &PublicRestClassicSpotMarketMergeDepthAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketMergeDepthReq{},
	}
}

func (api *PublicRestClassicSpotMarketMergeDepthAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketMergeDepthRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketMergeDepth])
	middleRes, err := bitgetCallAPI[PublicRestClassicSpotMarketMergeDepthResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicSpotMarketMergeDepthRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取交易深度
func (client *PublicRestClient) NewPublicRestClassicSpotMarketOrderbook() *PublicRestClassicSpotMarketOrderbookAPI {
	return &PublicRestClassicSpotMarketOrderbookAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketOrderbookReq{},
	}
}

func (api *PublicRestClassicSpotMarketOrderbookAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketOrderbookRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketOrderbook])
	middleRes, err := bitgetCallAPI[PublicRestClassicSpotMarketOrderbookResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicSpotMarketOrderbookRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取K线数据
func (client *PublicRestClient) NewPublicRestClassicSpotMarketCandles() *PublicRestClassicSpotMarketCandlesAPI {
	return &PublicRestClassicSpotMarketCandlesAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketCandlesReq{},
	}
}

func (api *PublicRestClassicSpotMarketCandlesAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketCandles])
	middleRes, err := bitgetCallAPI[PublicRestClassicSpotMarketCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicSpotMarketCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取历史K线数据
func (client *PublicRestClient) NewPublicRestClassicSpotMarketHistoryCandles() *PublicRestClassicSpotMarketHistoryCandlesAPI {
	return &PublicRestClassicSpotMarketHistoryCandlesAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketHistoryCandlesReq{},
	}
}

func (api *PublicRestClassicSpotMarketHistoryCandlesAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketHistoryCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketHistoryCandles])
	middleRes, err := bitgetCallAPI[PublicRestClassicSpotMarketHistoryCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicSpotMarketHistoryCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取最近成交数据
func (client *PublicRestClient) NewPublicRestClassicSpotMarketFills() *PublicRestClassicSpotMarketFillsAPI {
	return &PublicRestClassicSpotMarketFillsAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketFillsReq{},
	}
}

func (api *PublicRestClassicSpotMarketFillsAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketFillsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketFills])
	return bitgetCallAPI[PublicRestClassicSpotMarketFillsRes](api.client.c, url, nil, GET)
}

// GET 获取市场成交数据
func (client *PublicRestClient) NewPublicRestClassicSpotMarketFillsHistory() *PublicRestClassicSpotMarketFillsHistoryAPI {
	return &PublicRestClassicSpotMarketFillsHistoryAPI{
		client: client,
		req:    &PublicRestClassicSpotMarketFillsHistoryReq{},
	}
}

func (api *PublicRestClassicSpotMarketFillsHistoryAPI) Do() (*BitgetRestRes[PublicRestClassicSpotMarketFillsHistoryRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicSpotMarketFillsHistory])
	return bitgetCallAPI[PublicRestClassicSpotMarketFillsHistoryRes](api.client.c, url, nil, GET)
}
