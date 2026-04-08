package mybitgetapi

// GET 获取合约VIP费率
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketVipFeeRate() *PublicRestClassicFuturesMarketVipFeeRateAPI {
	return &PublicRestClassicFuturesMarketVipFeeRateAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketVipFeeRateReq{},
	}
}

func (api *PublicRestClassicFuturesMarketVipFeeRateAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketVipFeeRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketVipFeeRate])
	return bitgetCallAPI[PublicRestClassicFuturesMarketVipFeeRateRes](api.client.c, url, nil, GET)
}

// GET 获取利率记录
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketInterestRate() *PublicRestClassicFuturesMarketInterestRateAPI {
	return &PublicRestClassicFuturesMarketInterestRateAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketInterestRateReq{},
	}
}

func (api *PublicRestClassicFuturesMarketInterestRateAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketInterestRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketInterestRate])
	return bitgetCallAPI[PublicRestClassicFuturesMarketInterestRateRes](api.client.c, url, nil, GET)
}

// GET 获取合并深度
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketMergeDepth() *PublicRestClassicFuturesMarketMergeDepthAPI {
	return &PublicRestClassicFuturesMarketMergeDepthAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketMergeDepthReq{},
	}
}

func (api *PublicRestClassicFuturesMarketMergeDepthAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketMergeDepthRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketMergeDepth])
	middleRes, err := bitgetCallAPI[PublicRestClassicFuturesMarketMergeDepthResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicFuturesMarketMergeDepthRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取单个交易对行情
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketTicker() *PublicRestClassicFuturesMarketTickerAPI {
	return &PublicRestClassicFuturesMarketTickerAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketTickerReq{},
	}
}

func (api *PublicRestClassicFuturesMarketTickerAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketTickerRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketTicker])
	return bitgetCallAPI[PublicRestClassicFuturesMarketTickerRes](api.client.c, url, nil, GET)
}

// GET 获取全部交易对行情
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketTickers() *PublicRestClassicFuturesMarketTickersAPI {
	return &PublicRestClassicFuturesMarketTickersAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketTickersReq{},
	}
}

func (api *PublicRestClassicFuturesMarketTickersAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketTickersRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketTickers])
	return bitgetCallAPI[PublicRestClassicFuturesMarketTickersRes](api.client.c, url, nil, GET)
}

// GET 获取最近成交
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketFills() *PublicRestClassicFuturesMarketFillsAPI {
	return &PublicRestClassicFuturesMarketFillsAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketFillsReq{},
	}
}

func (api *PublicRestClassicFuturesMarketFillsAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketFillsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketFills])
	return bitgetCallAPI[PublicRestClassicFuturesMarketFillsRes](api.client.c, url, nil, GET)
}

// GET 获取历史成交
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketFillsHistory() *PublicRestClassicFuturesMarketFillsHistoryAPI {
	return &PublicRestClassicFuturesMarketFillsHistoryAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketFillsHistoryReq{},
	}
}

func (api *PublicRestClassicFuturesMarketFillsHistoryAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketFillsHistoryRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketFillsHistory])
	return bitgetCallAPI[PublicRestClassicFuturesMarketFillsHistoryRes](api.client.c, url, nil, GET)
}

// GET 获取K线数据
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketCandles() *PublicRestClassicFuturesMarketCandlesAPI {
	return &PublicRestClassicFuturesMarketCandlesAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketCandlesReq{},
	}
}

func (api *PublicRestClassicFuturesMarketCandlesAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketCandles])
	middleRes, err := bitgetCallAPI[PublicRestClassicFuturesMarketCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicFuturesMarketCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取历史K线
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketHistoryCandles() *PublicRestClassicFuturesMarketHistoryCandlesAPI {
	return &PublicRestClassicFuturesMarketHistoryCandlesAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketHistoryCandlesReq{},
	}
}

func (api *PublicRestClassicFuturesMarketHistoryCandlesAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketHistoryCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketHistoryCandles])
	middleRes, err := bitgetCallAPI[PublicRestClassicFuturesMarketHistoryCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicFuturesMarketHistoryCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取指数历史K线
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketHistoryIndexCandles() *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI {
	return &PublicRestClassicFuturesMarketHistoryIndexCandlesAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketHistoryIndexCandlesReq{},
	}
}

func (api *PublicRestClassicFuturesMarketHistoryIndexCandlesAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketHistoryIndexCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketHistoryIndexCandles])
	middleRes, err := bitgetCallAPI[PublicRestClassicFuturesMarketHistoryIndexCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicFuturesMarketHistoryIndexCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取标记价格历史K线
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketHistoryMarkCandles() *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI {
	return &PublicRestClassicFuturesMarketHistoryMarkCandlesAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketHistoryMarkCandlesReq{},
	}
}

func (api *PublicRestClassicFuturesMarketHistoryMarkCandlesAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketHistoryMarkCandlesRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketHistoryMarkCandles])
	middleRes, err := bitgetCallAPI[PublicRestClassicFuturesMarketHistoryMarkCandlesResMiddle](api.client.c, url, nil, GET)
	if err != nil {
		return nil, err
	}
	return &BitgetRestRes[PublicRestClassicFuturesMarketHistoryMarkCandlesRes]{
		BitgetErrorRes: middleRes.BitgetErrorRes,
		BitgetTimeRes:  middleRes.BitgetTimeRes,
		Data:           middleRes.Data.ConvertToRes(),
	}, nil
}

// GET 获取当前资金费率
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketCurrentFundRate() *PublicRestClassicFuturesMarketCurrentFundRateAPI {
	return &PublicRestClassicFuturesMarketCurrentFundRateAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketCurrentFundRateReq{},
	}
}

func (api *PublicRestClassicFuturesMarketCurrentFundRateAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketCurrentFundRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketCurrentFundRate])
	return bitgetCallAPI[PublicRestClassicFuturesMarketCurrentFundRateRes](api.client.c, url, nil, GET)
}

// GET 获取历史资金费率
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketHistoryFundRate() *PublicRestClassicFuturesMarketHistoryFundRateAPI {
	return &PublicRestClassicFuturesMarketHistoryFundRateAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketHistoryFundRateReq{},
	}
}

func (api *PublicRestClassicFuturesMarketHistoryFundRateAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketHistoryFundRateRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketHistoryFundRate])
	return bitgetCallAPI[PublicRestClassicFuturesMarketHistoryFundRateRes](api.client.c, url, nil, GET)
}

// GET 获取下次资金费结算时间
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketFundingTime() *PublicRestClassicFuturesMarketFundingTimeAPI {
	return &PublicRestClassicFuturesMarketFundingTimeAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketFundingTimeReq{},
	}
}

func (api *PublicRestClassicFuturesMarketFundingTimeAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketFundingTimeRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketFundingTime])
	return bitgetCallAPI[PublicRestClassicFuturesMarketFundingTimeRes](api.client.c, url, nil, GET)
}

// GET 获取合约信息
func (client *PublicRestClient) NewPublicRestClassicFuturesMarketContracts() *PublicRestClassicFuturesMarketContractsAPI {
	return &PublicRestClassicFuturesMarketContractsAPI{
		client: client,
		req:    &PublicRestClassicFuturesMarketContractsReq{},
	}
}

func (api *PublicRestClassicFuturesMarketContractsAPI) Do() (*BitgetRestRes[PublicRestClassicFuturesMarketContractsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestClassicAPIMap[PublicRestClassicFuturesMarketContracts])
	return bitgetCallAPI[PublicRestClassicFuturesMarketContractsRes](api.client.c, url, nil, GET)
}
