package mybitgetapi

import "errors"

// POST 下单
func (client *PrivateRestClient) NewPrivateRestUtaTradePlaceOrder() *PrivateRestUtaTradePlaceOrderAPI {
	return &PrivateRestUtaTradePlaceOrderAPI{
		client: client,
		req:    &PrivateRestUtaTradePlaceOrderReq{},
	}
}

func (api *PrivateRestUtaTradePlaceOrderAPI) Do() (*BitgetRestRes[PrivateRestUtaTradePlaceOrderRes], error) {
	if api.req.Category == nil || api.req.Symbol == nil || api.req.Qty == nil || api.req.Side == nil || api.req.OrderType == nil {
		return nil, errors.New("category, symbol, qty, side, orderType are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradePlaceOrder])
	return bitgetCallApiWithSecret[PrivateRestUtaTradePlaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 修改订单
func (client *PrivateRestClient) NewPrivateRestUtaTradeModifyOrder() *PrivateRestUtaTradeModifyOrderAPI {
	return &PrivateRestUtaTradeModifyOrderAPI{
		client: client,
		req:    &PrivateRestUtaTradeModifyOrderReq{},
	}
}

func (api *PrivateRestUtaTradeModifyOrderAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeModifyOrderRes], error) {
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	if api.req.Qty == nil && api.req.Price == nil {
		return nil, errors.New("qty or price is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradeModifyOrder])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeModifyOrderRes](api.client.c, url, reqBody, POST)
}

// POST 撤单
func (client *PrivateRestClient) NewPrivateRestUtaTradeCancelOrder() *PrivateRestUtaTradeCancelOrderAPI {
	return &PrivateRestUtaTradeCancelOrderAPI{
		client: client,
		req:    &PrivateRestUtaTradeCancelOrderReq{},
	}
}

func (api *PrivateRestUtaTradeCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeCancelOrderRes], error) {
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradeCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeCancelOrderRes](api.client.c, url, reqBody, POST)
}

// POST 批量下单
func (client *PrivateRestClient) NewPrivateRestUtaTradePlaceBatch() *PrivateRestUtaTradePlaceBatchAPI {
	return &PrivateRestUtaTradePlaceBatchAPI{
		client: client,
		req:    make([]PrivateRestUtaTradePlaceBatchReqItem, 0),
	}
}

func (api *PrivateRestUtaTradePlaceBatchAPI) Do() (*BitgetRestRes[PrivateRestUtaTradePlaceBatchRes], error) {
	if len(api.req) == 0 {
		return nil, errors.New("at least one order is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradePlaceBatch])
	return bitgetCallApiWithSecret[PrivateRestUtaTradePlaceBatchRes](api.client.c, url, reqBody, POST)
}

// POST 批量改单
func (client *PrivateRestClient) NewPrivateRestUtaTradeBatchModifyOrder() *PrivateRestUtaTradeBatchModifyOrderAPI {
	return &PrivateRestUtaTradeBatchModifyOrderAPI{
		client: client,
		req:    make([]PrivateRestUtaTradeBatchModifyOrderReqItem, 0),
	}
}

func (api *PrivateRestUtaTradeBatchModifyOrderAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeBatchModifyOrderRes], error) {
	if len(api.req) == 0 {
		return nil, errors.New("at least one order is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradeBatchModifyOrder])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeBatchModifyOrderRes](api.client.c, url, reqBody, POST)
}

// POST 批量撤单
func (client *PrivateRestClient) NewPrivateRestUtaTradeCancelBatch() *PrivateRestUtaTradeCancelBatchAPI {
	return &PrivateRestUtaTradeCancelBatchAPI{
		client: client,
		req:    make([]PrivateRestUtaTradeCancelBatchReqOrder, 0),
	}
}

func (api *PrivateRestUtaTradeCancelBatchAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeCancelBatchRes], error) {
	if len(api.req) == 0 {
		return nil, errors.New("at least one order is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradeCancelBatch])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeCancelBatchRes](api.client.c, url, reqBody, POST)
}

// POST 一键撤单
func (client *PrivateRestClient) NewPrivateRestUtaTradeCancelSymbolOrder() *PrivateRestUtaTradeCancelSymbolOrderAPI {
	return &PrivateRestUtaTradeCancelSymbolOrderAPI{
		client: client,
		req:    &PrivateRestUtaTradeCancelSymbolOrderReq{},
	}
}

func (api *PrivateRestUtaTradeCancelSymbolOrderAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeCancelSymbolOrderRes], error) {
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradeCancelSymbolOrder])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeCancelSymbolOrderRes](api.client.c, url, reqBody, POST)
}

// POST 一键平仓
func (client *PrivateRestClient) NewPrivateRestUtaTradeClosePositions() *PrivateRestUtaTradeClosePositionsAPI {
	return &PrivateRestUtaTradeClosePositionsAPI{
		client: client,
		req:    &PrivateRestUtaTradeClosePositionsReq{},
	}
}

func (api *PrivateRestUtaTradeClosePositionsAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeClosePositionsRes], error) {
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestUtaTradeClosePositions])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeClosePositionsRes](api.client.c, url, reqBody, POST)
}

// GET 查询订单详情
func (client *PrivateRestClient) NewPrivateRestUtaTradeOrderInfo() *PrivateRestUtaTradeOrderInfoAPI {
	return &PrivateRestUtaTradeOrderInfoAPI{
		client: client,
		req:    &PrivateRestUtaTradeOrderInfoReq{},
	}
}

func (api *PrivateRestUtaTradeOrderInfoAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeOrderInfoRes], error) {
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaTradeOrderInfo])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeOrderInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 查询当前委托
func (client *PrivateRestClient) NewPrivateRestUtaTradeUnfilledOrders() *PrivateRestUtaTradeUnfilledOrdersAPI {
	return &PrivateRestUtaTradeUnfilledOrdersAPI{
		client: client,
		req:    &PrivateRestUtaTradeUnfilledOrdersReq{},
	}
}

func (api *PrivateRestUtaTradeUnfilledOrdersAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeUnfilledOrdersRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaTradeUnfilledOrders])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeUnfilledOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 查询历史委托
func (client *PrivateRestClient) NewPrivateRestUtaTradeHistoryOrders() *PrivateRestUtaTradeHistoryOrdersAPI {
	return &PrivateRestUtaTradeHistoryOrdersAPI{
		client: client,
		req:    &PrivateRestUtaTradeHistoryOrdersReq{},
	}
}

func (api *PrivateRestUtaTradeHistoryOrdersAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeHistoryOrdersRes], error) {
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaTradeHistoryOrders])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeHistoryOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 查询成交明细
func (client *PrivateRestClient) NewPrivateRestUtaTradeFills() *PrivateRestUtaTradeFillsAPI {
	return &PrivateRestUtaTradeFillsAPI{
		client: client,
		req:    &PrivateRestUtaTradeFillsReq{},
	}
}

func (api *PrivateRestUtaTradeFillsAPI) Do() (*BitgetRestRes[PrivateRestUtaTradeFillsRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaTradeFills])
	return bitgetCallApiWithSecret[PrivateRestUtaTradeFillsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 查询当前持仓
func (client *PrivateRestClient) NewPrivateRestUtaPositionCurrent() *PrivateRestUtaPositionCurrentAPI {
	return &PrivateRestUtaPositionCurrentAPI{
		client: client,
		req:    &PrivateRestUtaPositionCurrentReq{},
	}
}

func (api *PrivateRestUtaPositionCurrentAPI) Do() (*BitgetRestRes[PrivateRestUtaPositionCurrentRes], error) {
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaPositionCurrent])
	return bitgetCallApiWithSecret[PrivateRestUtaPositionCurrentRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 查询历史持仓
func (client *PrivateRestClient) NewPrivateRestUtaPositionHistory() *PrivateRestUtaPositionHistoryAPI {
	return &PrivateRestUtaPositionHistoryAPI{
		client: client,
		req:    &PrivateRestPositionUtaHistoryReq{},
	}
}

func (api *PrivateRestUtaPositionHistoryAPI) Do() (*BitgetRestRes[PrivateRestUtaPositionHistoryRes], error) {
	if api.req.Category == nil {
		return nil, errors.New("category is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUtaPositionHistory])
	return bitgetCallApiWithSecret[PrivateRestUtaPositionHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}
