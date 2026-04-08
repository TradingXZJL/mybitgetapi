package mybitgetapi

type PrivateRestClassicSpotTradePlaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradePlaceOrderReq
}

type PrivateRestClassicSpotTradePlaceOrderReq struct {
	Symbol                 *string `json:"symbol"`                 // String 是 交易对名称，如 BTCUSDT
	Side                   *string `json:"side"`                   // String 是 交易方向 buy/sell
	OrderType              *string `json:"orderType"`              // String 是 订单类型 limit/market
	Force                  *string `json:"force"`                  // String 是 执行策略 gtc/post_only/fok/ioc
	Price                  *string `json:"price"`                  // String 否 限价价格
	Size                   *string `json:"size"`                   // String 是 委托数量
	ClientOid              *string `json:"clientOid"`              // String 否 自定义订单ID
	TpslType               *string `json:"tpslType"`               // String 否 订单类型 normal/tpsl
	TriggerPrice           *string `json:"triggerPrice"`           // String 否 止盈止损触发价格
	RequestTime            *string `json:"requestTime"`            // String 否 请求时间，Unix毫秒时间戳
	ReceiveWindow          *string `json:"receiveWindow"`          // String 否 有效窗口期，毫秒
	StpMode                *string `json:"stpMode"`                // String 否 STP 模式
	PresetTakeProfitPrice  *string `json:"presetTakeProfitPrice"`  // String 否 止盈价格
	ExecuteTakeProfitPrice *string `json:"executeTakeProfitPrice"` // String 否 止盈执行价格
	PresetStopLossPrice    *string `json:"presetStopLossPrice"`    // String 否 止损价格
	ExecuteStopLossPrice   *string `json:"executeStopLossPrice"`   // String 否 止损执行价格
}

func (api *PrivateRestClassicSpotTradePlaceOrderAPI) Symbol(symbol string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) Side(side string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.Side = GetPointer(side)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) OrderType(orderType string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) Force(force string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.Force = GetPointer(force)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) Price(price string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.Price = GetPointer(price)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) Size(size string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.Size = GetPointer(size)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) ClientOid(clientOid string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) TpslType(tpslType string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.TpslType = GetPointer(tpslType)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) TriggerPrice(triggerPrice string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.TriggerPrice = GetPointer(triggerPrice)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) RequestTime(requestTime string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.RequestTime = GetPointer(requestTime)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) ReceiveWindow(receiveWindow string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.ReceiveWindow = GetPointer(receiveWindow)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) StpMode(stpMode string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.StpMode = GetPointer(stpMode)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) PresetTakeProfitPrice(presetTakeProfitPrice string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.PresetTakeProfitPrice = GetPointer(presetTakeProfitPrice)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) ExecuteTakeProfitPrice(executeTakeProfitPrice string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.ExecuteTakeProfitPrice = GetPointer(executeTakeProfitPrice)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) PresetStopLossPrice(presetStopLossPrice string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.PresetStopLossPrice = GetPointer(presetStopLossPrice)
	return api
}
func (api *PrivateRestClassicSpotTradePlaceOrderAPI) ExecuteStopLossPrice(executeStopLossPrice string) *PrivateRestClassicSpotTradePlaceOrderAPI {
	api.req.ExecuteStopLossPrice = GetPointer(executeStopLossPrice)
	return api
}

type PrivateRestClassicSpotTradeCancelReplaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeCancelReplaceOrderReq
}

type PrivateRestClassicSpotTradeCancelReplaceOrderReq struct {
	Symbol                 *string `json:"symbol"`                 // String 是 交易对名称
	Price                  *string `json:"price"`                  // String 是 限价价格
	Size                   *string `json:"size"`                   // String 是 委托数量
	ClientOid              *string `json:"clientOid"`              // String 否 待撤销自定义订单ID
	OrderId                *string `json:"orderId"`                // String 否 待撤销订单号
	NewClientOid           *string `json:"newClientOid"`           // String 否 新订单自定义ID
	PresetTakeProfitPrice  *string `json:"presetTakeProfitPrice"`  // String 否 止盈价格
	ExecuteTakeProfitPrice *string `json:"executeTakeProfitPrice"` // String 否 止盈执行价格
	PresetStopLossPrice    *string `json:"presetStopLossPrice"`    // String 否 止损价格
	ExecuteStopLossPrice   *string `json:"executeStopLossPrice"`   // String 否 止损执行价格
}

func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) Symbol(symbol string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) Price(price string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.Price = GetPointer(price)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) Size(size string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.Size = GetPointer(size)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) ClientOid(clientOid string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) OrderId(orderId string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) NewClientOid(newClientOid string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.NewClientOid = GetPointer(newClientOid)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) PresetTakeProfitPrice(presetTakeProfitPrice string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.PresetTakeProfitPrice = GetPointer(presetTakeProfitPrice)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) ExecuteTakeProfitPrice(executeTakeProfitPrice string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.ExecuteTakeProfitPrice = GetPointer(executeTakeProfitPrice)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) PresetStopLossPrice(presetStopLossPrice string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.PresetStopLossPrice = GetPointer(presetStopLossPrice)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelReplaceOrderAPI) ExecuteStopLossPrice(executeStopLossPrice string) *PrivateRestClassicSpotTradeCancelReplaceOrderAPI {
	api.req.ExecuteStopLossPrice = GetPointer(executeStopLossPrice)
	return api
}

type PrivateRestClassicSpotTradeBatchCancelReplaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeBatchCancelReplaceOrderReq
}

type PrivateRestClassicSpotTradeBatchCancelReplaceOrderReq struct {
	OrderList []PrivateRestClassicSpotTradeBatchCancelReplaceOrderItem `json:"orderList"` // Array 是 订单列表，最大50
}

type PrivateRestClassicSpotTradeBatchCancelReplaceOrderItem struct {
	Symbol                 *string `json:"symbol"`                 // String 是 交易对名称
	Price                  *string `json:"price"`                  // String 是 限价价格
	Size                   *string `json:"size"`                   // String 是 委托数量
	ClientOid              *string `json:"clientOid"`              // String 否 待撤销自定义订单ID
	OrderId                *string `json:"orderId"`                // String 否 待撤销订单号
	NewClientOid           *string `json:"newClientOid"`           // String 否 新订单自定义ID
	PresetTakeProfitPrice  *string `json:"presetTakeProfitPrice"`  // String 否 止盈价格
	ExecuteTakeProfitPrice *string `json:"executeTakeProfitPrice"` // String 否 止盈执行价格
	PresetStopLossPrice    *string `json:"presetStopLossPrice"`    // String 否 止损价格
	ExecuteStopLossPrice   *string `json:"executeStopLossPrice"`   // String 否 止损执行价格
}

func (api *PrivateRestClassicSpotTradeBatchCancelReplaceOrderAPI) AddOrder(item PrivateRestClassicSpotTradeBatchCancelReplaceOrderItem) *PrivateRestClassicSpotTradeBatchCancelReplaceOrderAPI {
	api.req.OrderList = append(api.req.OrderList, item)
	return api
}

type PrivateRestClassicSpotTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeCancelOrderReq
}

type PrivateRestClassicSpotTradeCancelOrderReq struct {
	Symbol    *string `json:"symbol"`    // String 否 交易对名称
	TpslType  *string `json:"tpslType"`  // String 否 订单类型 normal/tpsl
	OrderId   *string `json:"orderId"`   // String 否 订单ID（orderId/clientOid二选一）
	ClientOid *string `json:"clientOid"` // String 否 客户自定义ID（orderId/clientOid二选一）
}

func (api *PrivateRestClassicSpotTradeCancelOrderAPI) Symbol(symbol string) *PrivateRestClassicSpotTradeCancelOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelOrderAPI) TpslType(tpslType string) *PrivateRestClassicSpotTradeCancelOrderAPI {
	api.req.TpslType = GetPointer(tpslType)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelOrderAPI) OrderId(orderId string) *PrivateRestClassicSpotTradeCancelOrderAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestClassicSpotTradeCancelOrderAPI) ClientOid(clientOid string) *PrivateRestClassicSpotTradeCancelOrderAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}

type PrivateRestClassicSpotTradeBatchOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeBatchOrdersReq
}

type PrivateRestClassicSpotTradeBatchOrdersReq struct {
	Symbol    *string                                         `json:"symbol"`    // String 否 交易对名称，single模式下有效
	BatchMode *string                                         `json:"batchMode"` // String 否 single/multiple
	OrderList []PrivateRestClassicSpotTradeBatchOrdersReqItem `json:"orderList"` // Array 是 下单订单集合，最大50
}

type PrivateRestClassicSpotTradeBatchOrdersReqItem struct {
	Symbol                 *string `json:"symbol"`                 // String 否 交易对名称，multiple模式下必填
	Side                   *string `json:"side"`                   // String 是 交易方向 buy/sell
	OrderType              *string `json:"orderType"`              // String 是 订单类型 limit/market
	Force                  *string `json:"force"`                  // String 是 执行策略 gtc/post_only/fok/ioc
	Price                  *string `json:"price"`                  // String 否 限价价格
	Size                   *string `json:"size"`                   // String 是 委托数量
	ClientOid              *string `json:"clientOid"`              // String 否 自定义订单ID
	StpMode                *string `json:"stpMode"`                // String 否 STP 模式
	PresetTakeProfitPrice  *string `json:"presetTakeProfitPrice"`  // String 否 止盈价格
	ExecuteTakeProfitPrice *string `json:"executeTakeProfitPrice"` // String 否 止盈执行价格
	PresetStopLossPrice    *string `json:"presetStopLossPrice"`    // String 否 止损价格
	ExecuteStopLossPrice   *string `json:"executeStopLossPrice"`   // String 否 止损执行价格
}

func (api *PrivateRestClassicSpotTradeBatchOrdersAPI) Symbol(symbol string) *PrivateRestClassicSpotTradeBatchOrdersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradeBatchOrdersAPI) BatchMode(batchMode string) *PrivateRestClassicSpotTradeBatchOrdersAPI {
	api.req.BatchMode = GetPointer(batchMode)
	return api
}
func (api *PrivateRestClassicSpotTradeBatchOrdersAPI) AddOrder(item PrivateRestClassicSpotTradeBatchOrdersReqItem) *PrivateRestClassicSpotTradeBatchOrdersAPI {
	api.req.OrderList = append(api.req.OrderList, item)
	return api
}

type PrivateRestClassicSpotTradeBatchCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeBatchCancelOrderReq
}

type PrivateRestClassicSpotTradeBatchCancelOrderReq struct {
	Symbol    *string                                           `json:"symbol"`    // String 否 交易对名称，single模式下有效
	BatchMode *string                                           `json:"batchMode"` // String 否 single/multiple
	OrderList []PrivateRestClassicSpotTradeBatchCancelOrderItem `json:"orderList"` // Array 是 订单ID数组，最大50
}

type PrivateRestClassicSpotTradeBatchCancelOrderItem struct {
	Symbol    *string `json:"symbol"`    // String 否 交易对名称，multiple模式下必填
	OrderId   *string `json:"orderId"`   // String 否 orderId，与clientOid二选一
	ClientOid *string `json:"clientOid"` // String 否 clientOid，与orderId二选一
}

func (api *PrivateRestClassicSpotTradeBatchCancelOrderAPI) Symbol(symbol string) *PrivateRestClassicSpotTradeBatchCancelOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradeBatchCancelOrderAPI) BatchMode(batchMode string) *PrivateRestClassicSpotTradeBatchCancelOrderAPI {
	api.req.BatchMode = GetPointer(batchMode)
	return api
}
func (api *PrivateRestClassicSpotTradeBatchCancelOrderAPI) AddOrder(item PrivateRestClassicSpotTradeBatchCancelOrderItem) *PrivateRestClassicSpotTradeBatchCancelOrderAPI {
	api.req.OrderList = append(api.req.OrderList, item)
	return api
}

type PrivateRestClassicSpotTradeOrderInfoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeOrderInfoReq
}

type PrivateRestClassicSpotTradeOrderInfoReq struct {
	OrderId       *string `json:"orderId"`       // String 否 订单ID（orderId/clientOid二选一）
	ClientOid     *string `json:"clientOid"`     // String 否 客户自定义ID（orderId/clientOid二选一）
	RequestTime   *string `json:"requestTime"`   // String 否 请求时间
	ReceiveWindow *string `json:"receiveWindow"` // String 否 有效窗口期
}

func (api *PrivateRestClassicSpotTradeOrderInfoAPI) OrderId(orderId string) *PrivateRestClassicSpotTradeOrderInfoAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestClassicSpotTradeOrderInfoAPI) ClientOid(clientOid string) *PrivateRestClassicSpotTradeOrderInfoAPI {
	api.req.ClientOid = GetPointer(clientOid)
	return api
}
func (api *PrivateRestClassicSpotTradeOrderInfoAPI) RequestTime(requestTime string) *PrivateRestClassicSpotTradeOrderInfoAPI {
	api.req.RequestTime = GetPointer(requestTime)
	return api
}
func (api *PrivateRestClassicSpotTradeOrderInfoAPI) ReceiveWindow(receiveWindow string) *PrivateRestClassicSpotTradeOrderInfoAPI {
	api.req.ReceiveWindow = GetPointer(receiveWindow)
	return api
}

type PrivateRestClassicSpotTradeUnfilledOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeUnfilledOrdersReq
}

type PrivateRestClassicSpotTradeUnfilledOrdersReq struct {
	Symbol        *string `json:"symbol"`        // String 否 交易对名称
	StartTime     *string `json:"startTime"`     // String 否 查询开始时间
	EndTime       *string `json:"endTime"`       // String 否 查询结束时间
	IDLessThan    *string `json:"idLessThan"`    // String 否 上次最后一条 orderId
	Limit         *string `json:"limit"`         // String 否 查询条数 默认100 最大100
	OrderId       *string `json:"orderId"`       // String 否 订单ID
	TpslType      *string `json:"tpslType"`      // String 否 订单类型 normal/tpsl
	RequestTime   *string `json:"requestTime"`   // String 否 请求时间
	ReceiveWindow *string `json:"receiveWindow"` // String 否 有效窗口期
}

func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) Symbol(symbol string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) StartTime(startTime string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) EndTime(endTime string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) IDLessThan(idLessThan string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) Limit(limit string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) OrderId(orderId string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) TpslType(tpslType string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.TpslType = GetPointer(tpslType)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) RequestTime(requestTime string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.RequestTime = GetPointer(requestTime)
	return api
}
func (api *PrivateRestClassicSpotTradeUnfilledOrdersAPI) ReceiveWindow(receiveWindow string) *PrivateRestClassicSpotTradeUnfilledOrdersAPI {
	api.req.ReceiveWindow = GetPointer(receiveWindow)
	return api
}

type PrivateRestClassicSpotTradeHistoryOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeHistoryOrdersReq
}

type PrivateRestClassicSpotTradeHistoryOrdersReq struct {
	Symbol        *string `json:"symbol"`        // String 否 交易对名称
	StartTime     *string `json:"startTime"`     // String 否 查询开始时间
	EndTime       *string `json:"endTime"`       // String 否 查询结束时间
	IDLessThan    *string `json:"idLessThan"`    // String 否 上次最后一条 orderId
	Limit         *string `json:"limit"`         // String 否 查询条数 默认100 最大100
	OrderId       *string `json:"orderId"`       // String 否 订单ID
	TpslType      *string `json:"tpslType"`      // String 否 订单类型 normal/tpsl
	RequestTime   *string `json:"requestTime"`   // String 否 请求时间
	ReceiveWindow *string `json:"receiveWindow"` // String 否 有效窗口期
}

func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) Symbol(symbol string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) StartTime(startTime string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) EndTime(endTime string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) IDLessThan(idLessThan string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) Limit(limit string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) OrderId(orderId string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) TpslType(tpslType string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.TpslType = GetPointer(tpslType)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) RequestTime(requestTime string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.RequestTime = GetPointer(requestTime)
	return api
}
func (api *PrivateRestClassicSpotTradeHistoryOrdersAPI) ReceiveWindow(receiveWindow string) *PrivateRestClassicSpotTradeHistoryOrdersAPI {
	api.req.ReceiveWindow = GetPointer(receiveWindow)
	return api
}

type PrivateRestClassicSpotTradeFillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicSpotTradeFillsReq
}

type PrivateRestClassicSpotTradeFillsReq struct {
	Symbol     *string `json:"symbol"`     // String 否 交易对名称
	OrderId    *string `json:"orderId"`    // String 否 订单ID
	StartTime  *string `json:"startTime"`  // String 否 查询开始时间
	EndTime    *string `json:"endTime"`    // String 否 查询结束时间
	Limit      *string `json:"limit"`      // String 否 查询条数 默认100 最大100
	IDLessThan *string `json:"idLessThan"` // String 否 请求此 tradeId 之前的数据
}

func (api *PrivateRestClassicSpotTradeFillsAPI) Symbol(symbol string) *PrivateRestClassicSpotTradeFillsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
func (api *PrivateRestClassicSpotTradeFillsAPI) OrderId(orderId string) *PrivateRestClassicSpotTradeFillsAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}
func (api *PrivateRestClassicSpotTradeFillsAPI) StartTime(startTime string) *PrivateRestClassicSpotTradeFillsAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}
func (api *PrivateRestClassicSpotTradeFillsAPI) EndTime(endTime string) *PrivateRestClassicSpotTradeFillsAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}
func (api *PrivateRestClassicSpotTradeFillsAPI) Limit(limit string) *PrivateRestClassicSpotTradeFillsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
func (api *PrivateRestClassicSpotTradeFillsAPI) IDLessThan(idLessThan string) *PrivateRestClassicSpotTradeFillsAPI {
	api.req.IDLessThan = GetPointer(idLessThan)
	return api
}
