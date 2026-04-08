package mybitgetapi

import "errors"

// POST 下单
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradePlaceOrder() *PrivateRestClassicSpotTradePlaceOrderAPI {
	return &PrivateRestClassicSpotTradePlaceOrderAPI{
		client: client,
		req:    &PrivateRestClassicSpotTradePlaceOrderReq{},
	}
}

func (api *PrivateRestClassicSpotTradePlaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradePlaceOrderRes], error) {
	if api.req.Symbol == nil || api.req.Side == nil || api.req.OrderType == nil || api.req.Force == nil || api.req.Size == nil {
		return nil, errors.New("symbol, side, orderType, force, size are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradePlaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradePlaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 撤单再下单
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeCancelReplaceOrder() *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	return &PrivateRestClassicSpotTradeCancelReplaceOrderAPI{
		client: client,
		req:    &PrivateRestClassicSpotTradeCancelReplaceOrderReq{},
	}
}

func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeCancelReplaceOrderRes], error) {
	if api.req.Symbol == nil || api.req.Price == nil || api.req.Size == nil {
		return nil, errors.New("symbol, price, size are required")
	}
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeCancelReplaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeCancelReplaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 批量撤单再下单
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeBatchCancelReplaceOrder() *PrivateRestClassicSpotTradeBatchCancelReplaceOrderAPI {
	return &PrivateRestClassicSpotTradeBatchCancelReplaceOrderAPI{
		client: client,
		req: &PrivateRestClassicSpotTradeBatchCancelReplaceOrderReq{
			OrderList: make([]PrivateRestClassicSpotTradeBatchCancelReplaceOrderItem, 0),
		},
	}
}

func (api *PrivateRestClassicSpotTradeBatchCancelReplaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeBatchCancelReplaceOrderRes], error) {
	if len(api.req.OrderList) == 0 {
		return nil, errors.New("at least one order is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeBatchCancelReplaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeBatchCancelReplaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 撤单
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeCancelOrder() *PrivateRestClassicSpotTradeCancelOrderAPI {
	return &PrivateRestClassicSpotTradeCancelOrderAPI{
		client: client,
		req:    &PrivateRestClassicSpotTradeCancelOrderReq{},
	}
}

func (api *PrivateRestClassicSpotTradeCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeCancelOrderRes], error) {
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeCancelOrderRes](api.client.c, url, reqBody, POST)
}

// POST 批量下单
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeBatchOrders() *PrivateRestClassicSpotTradeBatchOrdersAPI {
	return &PrivateRestClassicSpotTradeBatchOrdersAPI{
		client: client,
		req: &PrivateRestClassicSpotTradeBatchOrdersReq{
			OrderList: make([]PrivateRestClassicSpotTradeBatchOrdersReqItem, 0),
		},
	}
}

func (api *PrivateRestClassicSpotTradeBatchOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeBatchOrdersRes], error) {
	if len(api.req.OrderList) == 0 {
		return nil, errors.New("at least one order is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeBatchOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeBatchOrdersRes](api.client.c, url, reqBody, POST)
}

// POST 批量撤单
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeBatchCancelOrder() *PrivateRestClassicSpotTradeBatchCancelOrderAPI {
	return &PrivateRestClassicSpotTradeBatchCancelOrderAPI{
		client: client,
		req: &PrivateRestClassicSpotTradeBatchCancelOrderReq{
			OrderList: make([]PrivateRestClassicSpotTradeBatchCancelOrderItem, 0),
		},
	}
}

func (api *PrivateRestClassicSpotTradeBatchCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeBatchCancelOrderRes], error) {
	if len(api.req.OrderList) == 0 {
		return nil, errors.New("at least one order is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeBatchCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeBatchCancelOrderRes](api.client.c, url, reqBody, POST)
}

// GET 获取订单详情
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeOrderInfo() *PrivateRestClassicSpotTradeOrderInfoAPI {
	return &PrivateRestClassicSpotTradeOrderInfoAPI{
		client: client,
		req:    &PrivateRestClassicSpotTradeOrderInfoReq{},
	}
}

func (api *PrivateRestClassicSpotTradeOrderInfoAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeOrderInfoRes], error) {
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeOrderInfo])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeOrderInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取当前委托列表
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeUnfilledOrders() *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	return &PrivateRestClassicSpotTradeUnfilledOrdersAPI{
		client: client,
		req:    &PrivateRestClassicSpotTradeUnfilledOrdersReq{},
	}
}

func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeUnfilledOrdersRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeUnfilledOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeUnfilledOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取历史委托列表
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeHistoryOrders() *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	return &PrivateRestClassicSpotTradeHistoryOrdersAPI{
		client: client,
		req:    &PrivateRestClassicSpotTradeHistoryOrdersReq{},
	}
}

func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeHistoryOrdersRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeHistoryOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeHistoryOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取成交明细
func (client *PrivateRestClient) NewPrivateRestClassicSpotTradeFills() *PrivateRestClassicSpotTradeFillsAPI {
	return &PrivateRestClassicSpotTradeFillsAPI{
		client: client,
		req:    &PrivateRestClassicSpotTradeFillsReq{},
	}
}

func (api *PrivateRestClassicSpotTradeFillsAPI) Do() (*BitgetRestRes[PrivateRestClassicSpotTradeFillsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicSpotTradeFills])
	return bitgetCallApiWithSecret[PrivateRestClassicSpotTradeFillsRes](api.client.c, url, NIL_REQBODY, GET)
}
