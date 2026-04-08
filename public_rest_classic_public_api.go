package mybitgetapi

// GET 获取服务器时间
func (client *PublicRestClient) NewPublicRestClassicPublicTime() *PublicRestClassicPublicTimeAPI {
	return &PublicRestClassicPublicTimeAPI{
		client: client,
		req:    &PublicRestClassicPublicTimeReq{},
	}
}

func (api *PublicRestClassicPublicTimeAPI) Do() (*BitgetRestRes[PublicRestClassicPublicTimeRes], error) {
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PublicRestClassicAPIMap[PublicRestPublicTime])
	return bitgetCallAPI[PublicRestClassicPublicTimeRes](api.client.c, url, nil, GET)
}
