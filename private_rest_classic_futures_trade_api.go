package mybitgetapi

import "errors"

// POST 下单
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradePlaceOrder() *PrivateRestClassicFuturesTradePlaceOrderAPI {
	return &PrivateRestClassicFuturesTradePlaceOrderAPI{client: client, req: &PrivateRestClassicFuturesTradePlaceOrderReq{}}
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradePlaceOrderRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginMode == nil || api.req.MarginCoin == nil || api.req.Size == nil || api.req.Side == nil || api.req.OrderType == nil {
		return nil, errors.New("symbol, productType, marginMode, marginCoin, size, side, orderType are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradePlaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradePlaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 一键反手
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeReversal() *PrivateRestClassicFuturesTradeReversalAPI {
	return &PrivateRestClassicFuturesTradeReversalAPI{client: client, req: &PrivateRestClassicFuturesTradeReversalReq{}}
}
func (api *PrivateRestClassicFuturesTradeReversalAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeReversalRes], error) {
	if api.req.Symbol == nil || api.req.MarginCoin == nil || api.req.ProductType == nil || api.req.Side == nil {
		return nil, errors.New("symbol, marginCoin, productType, side are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeReversal])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeReversalRes](api.client.c, url, reqBody, POST)
}

// POST 批量下单
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeBatchOrder() *PrivateRestClassicFuturesTradeBatchOrderAPI {
	return &PrivateRestClassicFuturesTradeBatchOrderAPI{client: client, req: &PrivateRestClassicFuturesTradeBatchOrderReq{OrderList: make([]PrivateRestClassicFuturesTradeBatchOrderItem, 0)}}
}
func (api *PrivateRestClassicFuturesTradeBatchOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeBatchOrderRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.MarginCoin == nil || api.req.MarginMode == nil || len(api.req.OrderList) == 0 {
		return nil, errors.New("symbol, productType, marginCoin, marginMode and orderList are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeBatchOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeBatchOrderRes](api.client.c, url, reqBody, POST)
}

// POST 修改订单
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeModifyOrder() *PrivateRestClassicFuturesTradeModifyOrderAPI {
	return &PrivateRestClassicFuturesTradeModifyOrderAPI{client: client, req: &PrivateRestClassicFuturesTradeModifyOrderReq{}}
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeModifyOrderRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil || api.req.NewClientOid == nil {
		return nil, errors.New("symbol, productType, newClientOid are required")
	}
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeModifyOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeModifyOrderRes](api.client.c, url, reqBody, POST)
}

// POST 撤单
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeCancelOrder() *PrivateRestClassicFuturesTradeCancelOrderAPI {
	return &PrivateRestClassicFuturesTradeCancelOrderAPI{client: client, req: &PrivateRestClassicFuturesTradeCancelOrderReq{}}
}
func (api *PrivateRestClassicFuturesTradeCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeCancelOrderRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil {
		return nil, errors.New("symbol, productType are required")
	}
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeCancelOrderRes](api.client.c, url, reqBody, POST)
}

// POST 批量撤单
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeBatchCancelOrder() *PrivateRestClassicFuturesTradeBatchCancelOrderAPI {
	return &PrivateRestClassicFuturesTradeBatchCancelOrderAPI{client: client, req: &PrivateRestClassicFuturesTradeBatchCancelOrderReq{OrderIdList: make([]PrivateRestClassicFuturesTradeCancelOrderIdItem, 0)}}
}
func (api *PrivateRestClassicFuturesTradeBatchCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeBatchCancelOrderRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	if len(api.req.OrderIdList) == 0 && api.req.Symbol == nil {
		return nil, errors.New("symbol or orderIdList is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeBatchCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeBatchCancelOrderRes](api.client.c, url, reqBody, POST)
}

// POST 一键市价平仓
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeClosePositions() *PrivateRestClassicFuturesTradeClosePositionsAPI {
	return &PrivateRestClassicFuturesTradeClosePositionsAPI{client: client, req: &PrivateRestClassicFuturesTradeClosePositionsReq{}}
}
func (api *PrivateRestClassicFuturesTradeClosePositionsAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeClosePositionsRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeClosePositions])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeClosePositionsRes](api.client.c, url, reqBody, POST)
}

// GET 获取订单详情
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeOrderDetail() *PrivateRestClassicFuturesTradeOrderDetailAPI {
	return &PrivateRestClassicFuturesTradeOrderDetailAPI{client: client, req: &PrivateRestClassicFuturesTradeOrderDetailReq{}}
}
func (api *PrivateRestClassicFuturesTradeOrderDetailAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeOrderDetailRes], error) {
	if api.req.Symbol == nil || api.req.ProductType == nil {
		return nil, errors.New("symbol, productType are required")
	}
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeOrderDetail])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeOrderDetailRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取成交明细
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeOrderFills() *PrivateRestClassicFuturesTradeOrderFillsAPI {
	return &PrivateRestClassicFuturesTradeOrderFillsAPI{client: client, req: &PrivateRestClassicFuturesTradeOrderFillsReq{}}
}
func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeOrderFillsRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeOrderFills])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeOrderFillsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取历史成交明细
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeFillHistory() *PrivateRestClassicFuturesTradeFillHistoryAPI {
	return &PrivateRestClassicFuturesTradeFillHistoryAPI{client: client, req: &PrivateRestClassicFuturesTradeFillHistoryReq{}}
}
func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeFillHistoryRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeFillHistory])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeFillHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 查询当前委托
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeOrdersPending() *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	return &PrivateRestClassicFuturesTradeOrdersPendingAPI{client: client, req: &PrivateRestClassicFuturesTradeOrdersPendingReq{}}
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeOrdersPendingRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeOrdersPending])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeOrdersPendingRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取历史委托
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeOrdersHistory() *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	return &PrivateRestClassicFuturesTradeOrdersHistoryAPI{client: client, req: &PrivateRestClassicFuturesTradeOrdersHistoryReq{}}
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeOrdersHistoryRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeOrdersHistory])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeOrdersHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 一键全撤
func (client *PrivateRestClient) NewPrivateRestClassicFuturesTradeCancelAllOrders() *PrivateRestClassicFuturesTradeCancelAllOrdersAPI {
	return &PrivateRestClassicFuturesTradeCancelAllOrdersAPI{client: client, req: &PrivateRestClassicFuturesTradeCancelAllOrdersReq{}}
}
func (api *PrivateRestClassicFuturesTradeCancelAllOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicFuturesTradeCancelAllOrdersRes], error) {
	if api.req.ProductType == nil {
		return nil, errors.New("productType is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicFuturesTradeCancelAllOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicFuturesTradeCancelAllOrdersRes](api.client.c, url, reqBody, POST)
}

// POST 逐仓下单
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradePlaceOrder() *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	return &PrivateRestClassicMarginIsolatedTradePlaceOrderAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradePlaceOrderReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradePlaceOrderRes], error) {
	if api.req.Symbol == nil || api.req.OrderType == nil || api.req.LoanType == nil || api.req.Force == nil || api.req.Side == nil {
		return nil, errors.New("symbol, orderType, loanType, force, side are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradePlaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradePlaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 逐仓批量下单
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradeBatchPlaceOrder() *PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI {
	return &PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderReq{OrderList: make([]PrivateRestClassicMarginIsolatedTradeBatchOrderItem, 0)},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderRes], error) {
	if api.req.Symbol == nil || len(api.req.OrderList) == 0 {
		return nil, errors.New("symbol and orderList are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradeBatchPlaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 逐仓撤单
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradeCancelOrder() *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI {
	return &PrivateRestClassicMarginIsolatedTradeCancelOrderAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradeCancelOrderReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradeCancelOrderRes], error) {
	if api.req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}
	if api.req.OrderId == nil && api.req.ClientOid == nil {
		return nil, errors.New("orderId or clientOid is required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradeCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradeCancelOrderRes](api.client.c, url, reqBody, POST)
}

// POST 逐仓批量撤单
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradeBatchCancelOrder() *PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI {
	return &PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradeBatchCancelOrderReq{OrderIdList: make([]PrivateRestClassicMarginIsolatedTradeCancelOrderItem, 0)},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradeBatchCancelOrderRes], error) {
	if api.req.Symbol == nil || len(api.req.OrderIdList) == 0 {
		return nil, errors.New("symbol and orderIdList are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradeBatchCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradeBatchCancelOrderRes](api.client.c, url, reqBody, POST)
}

// GET 获取逐仓当前委托
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradeOpenOrders() *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI {
	return &PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradeOpenOrdersReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradeOpenOrdersRes], error) {
	if api.req.Symbol == nil || api.req.StartTime == nil {
		return nil, errors.New("symbol and startTime are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradeOpenOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradeOpenOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取逐仓历史委托
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradeHistoryOrders() *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	return &PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradeHistoryOrdersReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradeHistoryOrdersRes], error) {
	if api.req.Symbol == nil || api.req.StartTime == nil {
		return nil, errors.New("symbol and startTime are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradeHistoryOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradeHistoryOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取逐仓成交明细
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradeFills() *PrivateRestClassicMarginIsolatedTradeFillsAPI {
	return &PrivateRestClassicMarginIsolatedTradeFillsAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradeFillsReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradeFillsAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradeFillsRes], error) {
	if api.req.Symbol == nil || api.req.StartTime == nil {
		return nil, errors.New("symbol and startTime are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradeFills])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradeFillsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取逐仓强平订单
func (client *PrivateRestClient) NewPrivateRestClassicMarginIsolatedTradeLiquidationOrder() *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	return &PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI{
		client: client,
		req:    &PrivateRestClassicMarginIsolatedTradeLiquidationOrderReq{},
	}
}

func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginIsolatedTradeLiquidationOrderRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginIsolatedTradeLiquidationOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginIsolatedTradeLiquidationOrderRes](api.client.c, url, NIL_REQBODY, GET)
}
