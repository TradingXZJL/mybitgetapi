package mybitgetapi

import "errors"

// POST 全仓下单
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradePlaceOrder() *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	return &PrivateRestClassicMarginCrossTradePlaceOrderAPI{client: client, req: &PrivateRestClassicMarginCrossTradePlaceOrderReq{}}
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradePlaceOrderRes], error) {
	if api.req.Symbol == nil || api.req.OrderType == nil || api.req.LoanType == nil || api.req.Force == nil || api.req.Side == nil {
		return nil, errors.New("symbol, orderType, loanType, force, side are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradePlaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradePlaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 全仓批量下单
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradeBatchPlaceOrder() *PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI {
	return &PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossTradeBatchPlaceOrderReq{OrderList: make([]PrivateRestClassicMarginCrossTradeBatchOrderItem, 0)},
	}
}
func (api *PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradeBatchPlaceOrderRes], error) {
	if api.req.Symbol == nil || len(api.req.OrderList) == 0 {
		return nil, errors.New("symbol and orderList are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradeBatchPlaceOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradeBatchPlaceOrderRes](api.client.c, url, reqBody, POST)
}

// POST 全仓撤单
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradeCancelOrder() *PrivateRestClassicMarginCrossTradeCancelOrderAPI {
	return &PrivateRestClassicMarginCrossTradeCancelOrderAPI{client: client, req: &PrivateRestClassicMarginCrossTradeCancelOrderReq{}}
}
func (api *PrivateRestClassicMarginCrossTradeCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradeCancelOrderRes], error) {
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
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradeCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradeCancelOrderRes](api.client.c, url, reqBody, POST)
}

// POST 全仓批量撤单
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradeBatchCancelOrder() *PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI {
	return &PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI{
		client: client,
		req:    &PrivateRestClassicMarginCrossTradeBatchCancelOrderReq{OrderIdList: make([]PrivateRestClassicMarginCrossTradeCancelOrderItem, 0)},
	}
}
func (api *PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradeBatchCancelOrderRes], error) {
	if api.req.Symbol == nil || len(api.req.OrderIdList) == 0 {
		return nil, errors.New("symbol and orderIdList are required")
	}
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	url := bitgetHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradeBatchCancelOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradeBatchCancelOrderRes](api.client.c, url, reqBody, POST)
}

// GET 获取全仓当前委托
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradeOpenOrders() *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	return &PrivateRestClassicMarginCrossTradeOpenOrdersAPI{client: client, req: &PrivateRestClassicMarginCrossTradeOpenOrdersReq{}}
}
func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradeOpenOrdersRes], error) {
	if api.req.Symbol == nil || api.req.StartTime == nil {
		return nil, errors.New("symbol and startTime are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradeOpenOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradeOpenOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全仓历史委托
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradeHistoryOrders() *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	return &PrivateRestClassicMarginCrossTradeHistoryOrdersAPI{client: client, req: &PrivateRestClassicMarginCrossTradeHistoryOrdersReq{}}
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradeHistoryOrdersRes], error) {
	if api.req.Symbol == nil || api.req.StartTime == nil {
		return nil, errors.New("symbol and startTime are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradeHistoryOrders])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradeHistoryOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全仓成交明细
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradeFills() *PrivateRestClassicMarginCrossTradeFillsAPI {
	return &PrivateRestClassicMarginCrossTradeFillsAPI{client: client, req: &PrivateRestClassicMarginCrossTradeFillsReq{}}
}
func (api *PrivateRestClassicMarginCrossTradeFillsAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradeFillsRes], error) {
	if api.req.Symbol == nil || api.req.StartTime == nil {
		return nil, errors.New("symbol and startTime are required")
	}
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradeFills])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradeFillsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取全仓强平订单
func (client *PrivateRestClient) NewPrivateRestClassicMarginCrossTradeLiquidationOrder() *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	return &PrivateRestClassicMarginCrossTradeLiquidationOrderAPI{client: client, req: &PrivateRestClassicMarginCrossTradeLiquidationOrderReq{}}
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) Do() (*BitgetRestRes[PrivateRestClassicMarginCrossTradeLiquidationOrderRes], error) {
	url := bitgetHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestClassicAPIMap[PrivateRestClassicMarginCrossTradeLiquidationOrder])
	return bitgetCallApiWithSecret[PrivateRestClassicMarginCrossTradeLiquidationOrderRes](api.client.c, url, NIL_REQBODY, GET)
}
