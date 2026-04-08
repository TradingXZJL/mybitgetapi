package mybitgetapi

type PrivateRestUtaTradePlaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradePlaceOrderReq
}

type PrivateRestUtaTradePlaceOrderReq struct {
	Category     *string `json:"category"`     // String 是 产品类型 SPOT/MARGIN/USDT-FUTURES/COIN-FUTURES/USDC-FUTURES
	Symbol       *string `json:"symbol"`       // String 是 交易对名称 例如 BTCUSDT
	Qty          *string `json:"qty"`          // String 是 下单数量
	Price        *string `json:"price"`        // String 否 委托价格（limit单必填）
	Side         *string `json:"side"`         // String 是 方向 buy/sell
	OrderType    *string `json:"orderType"`    // String 是 订单类型 limit/market
	TimeInForce  *string `json:"timeInForce"`  // String 否 有效方式 ioc/fok/gtc/post_only/rpi
	PosSide      *string `json:"posSide"`      // String 否 持仓方向 long/short（对冲模式必填）
	ClientOid    *string `json:"clientOid"`    // String 否 客户端订单ID
	ReduceOnly   *string `json:"reduceOnly"`   // String 否 只减仓标识 yes/no
	StpMode      *string `json:"stpMode"`      // String 否 自成交保护模式
	TpTriggerBy  *string `json:"tpTriggerBy"`  // String 否 止盈触发类型 market/mark
	SlTriggerBy  *string `json:"slTriggerBy"`  // String 否 止损触发类型 market/mark
	TakeProfit   *string `json:"takeProfit"`   // String 否 预设止盈触发价
	StopLoss     *string `json:"stopLoss"`     // String 否 预设止损触发价
	TpOrderType  *string `json:"tpOrderType"`  // String 否 止盈委托类型 limit/market
	SlOrderType  *string `json:"slOrderType"`  // String 否 止损委托类型 limit/market
	TpLimitPrice *string `json:"tpLimitPrice"` // String 否 止盈限价单执行价（tpOrderType=limit时有效）
	SlLimitPrice *string `json:"slLimitPrice"` // String 否 止损限价单执行价（slOrderType=limit时有效）
}

func (api *PrivateRestUtaTradePlaceOrderAPI) Category(category string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) Symbol(symbol string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) Qty(qty string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.Qty = GetPointer(qty)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) Price(price string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.Price = GetPointer(price)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) Side(side string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.Side = GetPointer(side)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) OrderType(orderType string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) TimeInForce(timeInForce string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) PosSide(posSide string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) ClientOid(clientOid string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) ReduceOnly(reduceOnly string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) StpMode(stpMode string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.StpMode = GetPointer(stpMode)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) TpTriggerBy(tpTriggerBy string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.TpTriggerBy = GetPointer(tpTriggerBy)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) SlTriggerBy(slTriggerBy string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.SlTriggerBy = GetPointer(slTriggerBy)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) TakeProfit(takeProfit string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.TakeProfit = GetPointer(takeProfit)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) StopLoss(stopLoss string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.StopLoss = GetPointer(stopLoss)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) TpOrderType(tpOrderType string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.TpOrderType = GetPointer(tpOrderType)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) SlOrderType(slOrderType string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.SlOrderType = GetPointer(slOrderType)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) TpLimitPrice(tpLimitPrice string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.TpLimitPrice = GetPointer(tpLimitPrice)
	return api
}
func (api *PrivateRestUtaTradePlaceOrderAPI) SlLimitPrice(slLimitPrice string) *PrivateRestUtaTradePlaceOrderAPI {
	api.req.SlLimitPrice = GetPointer(slLimitPrice)
	return api
}

type PrivateRestUtaTradeModifyOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeModifyOrderReq
}

type PrivateRestUtaTradeModifyOrderReq struct {
	OrderId    *string `json:"orderId"`    // String 否 订单ID（orderId/clientOid二选一）
	ClientOid  *string `json:"clientOid"`  // String 否 客户端订单ID（orderId/clientOid二选一）
	Qty        *string `json:"qty"`        // String 否 修改后数量（qty/price至少填一项）
	Price      *string `json:"price"`      // String 否 修改后价格（qty/price至少填一项）
	AutoCancel *string `json:"autoCancel"` // String 否 修改失败是否自动撤单 yes/no
	Symbol     *string `json:"symbol"`     // String 否 交易对名称
	Category   *string `json:"category"`   // String 否 产品类型
}

func (api *PrivateRestUtaTradeModifyOrderAPI) OrderId(orderId string) *PrivateRestUtaTradeModifyOrderAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestUtaTradeModifyOrderAPI) ClientOid(clientOid string) *PrivateRestUtaTradeModifyOrderAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}
func (api *PrivateRestUtaTradeModifyOrderAPI) Qty(qty string) *PrivateRestUtaTradeModifyOrderAPI {
	api.req.Qty = GetPointer(qty)
	return api
}
func (api *PrivateRestUtaTradeModifyOrderAPI) Price(price string) *PrivateRestUtaTradeModifyOrderAPI {
	api.req.Price = GetPointer(price)
	return api
}
func (api *PrivateRestUtaTradeModifyOrderAPI) AutoCancel(autoCancel string) *PrivateRestUtaTradeModifyOrderAPI {
	api.req.AutoCancel = GetPointer(autoCancel)
	return api
}
func (api *PrivateRestUtaTradeModifyOrderAPI) Symbol(symbol string) *PrivateRestUtaTradeModifyOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestUtaTradeModifyOrderAPI) Category(category string) *PrivateRestUtaTradeModifyOrderAPI {
	api.req.Category = GetPointer(category)
	return api
}

type PrivateRestUtaTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeCancelOrderReq
}

type PrivateRestUtaTradeCancelOrderReq struct {
	OrderId   *string `json:"orderId"`   // String 否 订单ID（orderId/clientOid二选一）
	ClientOid *string `json:"clientOid"` // String 否 客户端订单ID（orderId/clientOid二选一）
	Category  *string `json:"category"`  // String 否 产品类型
}

func (api *PrivateRestUtaTradeCancelOrderAPI) OrderId(orderId string) *PrivateRestUtaTradeCancelOrderAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestUtaTradeCancelOrderAPI) ClientOid(clientOid string) *PrivateRestUtaTradeCancelOrderAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}
func (api *PrivateRestUtaTradeCancelOrderAPI) Category(category string) *PrivateRestUtaTradeCancelOrderAPI {
	api.req.Category = GetPointer(category)
	return api
}

type PrivateRestUtaTradePlaceBatchAPI struct {
	client *PrivateRestClient
	req    []PrivateRestUtaTradePlaceBatchReqItem
}

type PrivateRestUtaTradePlaceBatchReqItem struct {
	Category    string  `json:"category"`              // String 是 产品类型（同一批需一致）
	Symbol      string  `json:"symbol"`                // String 是 交易对名称
	Qty         string  `json:"qty"`                   // String 是 下单数量
	Price       *string `json:"price,omitempty"`       // String 否 委托价格（limit单必填）
	Side        string  `json:"side"`                  // String 是 方向 buy/sell
	OrderType   string  `json:"orderType"`             // String 是 订单类型 limit/market
	TimeInForce *string `json:"timeInForce,omitempty"` // String 否 有效方式 ioc/fok/gtc/post_only/rpi
	PosSide     *string `json:"posSide,omitempty"`     // String 否 持仓方向 long/short
	ClientOid   *string `json:"clientOid,omitempty"`   // String 否 客户端订单ID
	StpMode     *string `json:"stpMode,omitempty"`     // String 否 自成交保护模式
}

func (api *PrivateRestUtaTradePlaceBatchAPI) AddOrder(item PrivateRestUtaTradePlaceBatchReqItem) *PrivateRestUtaTradePlaceBatchAPI {
	api.req = append(api.req, item)
	return api
}

type PrivateRestUtaTradeBatchModifyOrderAPI struct {
	client *PrivateRestClient
	req    []PrivateRestUtaTradeBatchModifyOrderReqItem
}

type PrivateRestUtaTradeBatchModifyOrderReqItem struct {
	OrderId    *string `json:"orderId,omitempty"`    // String 否 订单ID（orderId/clientOid二选一）
	ClientOid  *string `json:"clientOid,omitempty"`  // String 否 客户端订单ID（orderId/clientOid二选一）
	Qty        *string `json:"qty,omitempty"`        // String 否 修改后数量
	Price      *string `json:"price,omitempty"`      // String 否 修改后价格
	AutoCancel *string `json:"autoCancel,omitempty"` // String 否 修改失败是否自动撤单 yes/no
	Symbol     *string `json:"symbol,omitempty"`     // String 否 交易对名称
	Category   *string `json:"category,omitempty"`   // String 否 产品类型
}

func (api *PrivateRestUtaTradeBatchModifyOrderAPI) AddOrder(item PrivateRestUtaTradeBatchModifyOrderReqItem) *PrivateRestUtaTradeBatchModifyOrderAPI {
	api.req = append(api.req, item)
	return api
}

type PrivateRestUtaTradeCancelBatchAPI struct {
	client *PrivateRestClient
	req    []PrivateRestUtaTradeCancelBatchReqOrder
}

type PrivateRestUtaTradeCancelBatchReqOrder struct {
	OrderId   *string `json:"orderId,omitempty"`   // String 否 订单ID（orderId/clientOid二选一）
	ClientOid *string `json:"clientOid,omitempty"` // String 否 客户端订单ID（orderId/clientOid二选一）
	Category  *string `json:"category"`            // String 是 产品类型
	Symbol    *string `json:"symbol"`              // String 是 交易对名称
}

func (api *PrivateRestUtaTradeCancelBatchAPI) AddOrder(item PrivateRestUtaTradeCancelBatchReqOrder) *PrivateRestUtaTradeCancelBatchAPI {
	api.req = append(api.req, item)
	return api
}

type PrivateRestUtaTradeCancelSymbolOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeCancelSymbolOrderReq
}

type PrivateRestUtaTradeCancelSymbolOrderReq struct {
	Category *string `json:"category"` // String 是 产品类型 SPOT/MARGIN/USDT-FUTURES/COIN-FUTURES/USDC-FUTURES
	Symbol   *string `json:"symbol"`   // String 否 交易对名称（为空时撤该产品全部挂单）
}

func (api *PrivateRestUtaTradeCancelSymbolOrderAPI) Category(category string) *PrivateRestUtaTradeCancelSymbolOrderAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaTradeCancelSymbolOrderAPI) Symbol(symbol string) *PrivateRestUtaTradeCancelSymbolOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PrivateRestUtaTradeClosePositionsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeClosePositionsReq
}

type PrivateRestUtaTradeClosePositionsReq struct {
	Category *string `json:"category"` // String 是 产品类型 USDT-FUTURES/COIN-FUTURES/USDC-FUTURES
	Symbol   *string `json:"symbol"`   // String 否 交易对名称（为空时平该产品全部仓位）
	PosSide  *string `json:"posSide"`  // String 否 持仓方向 long/short
}

func (api *PrivateRestUtaTradeClosePositionsAPI) Category(category string) *PrivateRestUtaTradeClosePositionsAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaTradeClosePositionsAPI) Symbol(symbol string) *PrivateRestUtaTradeClosePositionsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestUtaTradeClosePositionsAPI) PosSide(posSide string) *PrivateRestUtaTradeClosePositionsAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}

type PrivateRestUtaTradeOrderInfoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeOrderInfoReq
}

type PrivateRestUtaTradeOrderInfoReq struct {
	OrderId   *string `json:"orderId"`   // String 否 订单ID（orderId/clientOid二选一）
	ClientOid *string `json:"clientOid"` // String 否 客户端订单ID（orderId/clientOid二选一）
}

func (api *PrivateRestUtaTradeOrderInfoAPI) OrderId(orderId string) *PrivateRestUtaTradeOrderInfoAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestUtaTradeOrderInfoAPI) ClientOid(clientOid string) *PrivateRestUtaTradeOrderInfoAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}

type PrivateRestUtaTradeUnfilledOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeUnfilledOrdersReq
}

type PrivateRestUtaTradeUnfilledOrdersReq struct {
	Category  *string `json:"category"`  // String 否 产品类型
	Symbol    *string `json:"symbol"`    // String 否 交易对名称
	StartTime *string `json:"startTime"` // String 否 开始时间（毫秒）
	EndTime   *string `json:"endTime"`   // String 否 结束时间（毫秒）
	Limit     *string `json:"limit"`     // String 否 每页数量 默认100 最大100
	Cursor    *string `json:"cursor"`    // String 否 游标
}

func (api *PrivateRestUtaTradeUnfilledOrdersAPI) Category(category string) *PrivateRestUtaTradeUnfilledOrdersAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaTradeUnfilledOrdersAPI) Symbol(symbol string) *PrivateRestUtaTradeUnfilledOrdersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestUtaTradeUnfilledOrdersAPI) StartTime(startTime string) *PrivateRestUtaTradeUnfilledOrdersAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestUtaTradeUnfilledOrdersAPI) EndTime(endTime string) *PrivateRestUtaTradeUnfilledOrdersAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestUtaTradeUnfilledOrdersAPI) Limit(limit string) *PrivateRestUtaTradeUnfilledOrdersAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestUtaTradeUnfilledOrdersAPI) Cursor(cursor string) *PrivateRestUtaTradeUnfilledOrdersAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type PrivateRestUtaTradeHistoryOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeHistoryOrdersReq
}

type PrivateRestUtaTradeHistoryOrdersReq struct {
	Category  *string `json:"category"`  // String 是 产品类型
	Symbol    *string `json:"symbol"`    // String 否 交易对名称
	StartTime *string `json:"startTime"` // String 否 开始时间（毫秒）
	EndTime   *string `json:"endTime"`   // String 否 结束时间（毫秒）
	Limit     *string `json:"limit"`     // String 否 每页数量 默认100 最大100
	Cursor    *string `json:"cursor"`    // String 否 游标
}

func (api *PrivateRestUtaTradeHistoryOrdersAPI) Category(category string) *PrivateRestUtaTradeHistoryOrdersAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaTradeHistoryOrdersAPI) Symbol(symbol string) *PrivateRestUtaTradeHistoryOrdersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestUtaTradeHistoryOrdersAPI) StartTime(startTime string) *PrivateRestUtaTradeHistoryOrdersAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestUtaTradeHistoryOrdersAPI) EndTime(endTime string) *PrivateRestUtaTradeHistoryOrdersAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestUtaTradeHistoryOrdersAPI) Limit(limit string) *PrivateRestUtaTradeHistoryOrdersAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestUtaTradeHistoryOrdersAPI) Cursor(cursor string) *PrivateRestUtaTradeHistoryOrdersAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type PrivateRestUtaTradeFillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaTradeFillsReq
}

type PrivateRestUtaTradeFillsReq struct {
	Category  *string `json:"category"`  // String 否 产品类型
	OrderId   *string `json:"orderId"`   // String 否 订单ID
	StartTime *string `json:"startTime"` // String 否 开始时间（毫秒）
	EndTime   *string `json:"endTime"`   // String 否 结束时间（毫秒）
	Limit     *string `json:"limit"`     // String 否 每页数量 默认100 最大100
	Cursor    *string `json:"cursor"`    // String 否 游标
}

func (api *PrivateRestUtaTradeFillsAPI) Category(category string) *PrivateRestUtaTradeFillsAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaTradeFillsAPI) OrderId(orderId string) *PrivateRestUtaTradeFillsAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestUtaTradeFillsAPI) StartTime(startTime string) *PrivateRestUtaTradeFillsAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestUtaTradeFillsAPI) EndTime(endTime string) *PrivateRestUtaTradeFillsAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestUtaTradeFillsAPI) Limit(limit string) *PrivateRestUtaTradeFillsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestUtaTradeFillsAPI) Cursor(cursor string) *PrivateRestUtaTradeFillsAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type PrivateRestUtaPositionCurrentAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUtaPositionCurrentReq
}

type PrivateRestUtaPositionCurrentReq struct {
	Category *string `json:"category"` // String 是 产品类型 USDT-FUTURES/COIN-FUTURES/USDC-FUTURES
	Symbol   *string `json:"symbol"`   // String 否 交易对名称
	PosSide  *string `json:"posSide"`  // String 否 持仓方向 long/short
}

func (api *PrivateRestUtaPositionCurrentAPI) Category(category string) *PrivateRestUtaPositionCurrentAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaPositionCurrentAPI) Symbol(symbol string) *PrivateRestUtaPositionCurrentAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestUtaPositionCurrentAPI) PosSide(posSide string) *PrivateRestUtaPositionCurrentAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}

type PrivateRestUtaPositionHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestPositionUtaHistoryReq
}

type PrivateRestPositionUtaHistoryReq struct {
	Category  *string `json:"category"`  // String 是 产品类型 USDT-FUTURES/COIN-FUTURES/USDC-FUTURES
	Symbol    *string `json:"symbol"`    // String 否 交易对名称
	StartTime *string `json:"startTime"` // String 否 开始时间（毫秒）
	EndTime   *string `json:"endTime"`   // String 否 结束时间（毫秒）
	Limit     *string `json:"limit"`     // String 否 每页数量 默认100 最大100
	Cursor    *string `json:"cursor"`    // String 否 游标
}

func (api *PrivateRestUtaPositionHistoryAPI) Category(category string) *PrivateRestUtaPositionHistoryAPI {
	api.req.Category = GetPointer(category)
	return api
}
func (api *PrivateRestUtaPositionHistoryAPI) Symbol(symbol string) *PrivateRestUtaPositionHistoryAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestUtaPositionHistoryAPI) StartTime(startTime string) *PrivateRestUtaPositionHistoryAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestUtaPositionHistoryAPI) EndTime(endTime string) *PrivateRestUtaPositionHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestUtaPositionHistoryAPI) Limit(limit string) *PrivateRestUtaPositionHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestUtaPositionHistoryAPI) Cursor(cursor string) *PrivateRestUtaPositionHistoryAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}
