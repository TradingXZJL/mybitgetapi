package mybitgetapi

type PrivateRestClassicFuturesTradePlaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradePlaceOrderReq
}

type PrivateRestClassicFuturesTradePlaceOrderReq struct {
	Symbol                        *string `json:"symbol"`
	ProductType                   *string `json:"productType"`
	MarginMode                    *string `json:"marginMode"`
	MarginCoin                    *string `json:"marginCoin"`
	Size                          *string `json:"size"`
	Price                         *string `json:"price"`
	Side                          *string `json:"side"`
	TradeSide                     *string `json:"tradeSide"`
	OrderType                     *string `json:"orderType"`
	Force                         *string `json:"force"`
	ClientOid                     *string `json:"clientOid"`
	ReduceOnly                    *string `json:"reduceOnly"`
	PresetStopSurplusPrice        *string `json:"presetStopSurplusPrice"`
	PresetStopLossPrice           *string `json:"presetStopLossPrice"`
	PresetStopSurplusExecutePrice *string `json:"presetStopSurplusExecutePrice"`
	PresetStopLossExecutePrice    *string `json:"presetStopLossExecutePrice"`
	StpMode                       *string `json:"stpMode"`
}

func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) Symbol(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) ProductType(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) MarginMode(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.MarginMode = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) MarginCoin(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.MarginCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) Size(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.Size = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) Price(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.Price = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) Side(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.Side = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) TradeSide(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.TradeSide = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) OrderType(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.OrderType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) Force(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.Force = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) ClientOid(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) ReduceOnly(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.ReduceOnly = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) PresetStopSurplusPrice(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.PresetStopSurplusPrice = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) PresetStopLossPrice(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.PresetStopLossPrice = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) PresetStopSurplusExecutePrice(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.PresetStopSurplusExecutePrice = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) PresetStopLossExecutePrice(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.PresetStopLossExecutePrice = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradePlaceOrderAPI) StpMode(v string) *PrivateRestClassicFuturesTradePlaceOrderAPI {
	api.req.StpMode = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeReversalAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeReversalReq
}

type PrivateRestClassicFuturesTradeReversalReq struct {
	Symbol      *string `json:"symbol"`
	MarginCoin  *string `json:"marginCoin"`
	ProductType *string `json:"productType"`
	Size        *string `json:"size"`
	Side        *string `json:"side"`
	TradeSide   *string `json:"tradeSide"`
	ClientOid   *string `json:"clientOid"`
}

func (api *PrivateRestClassicFuturesTradeReversalAPI) Symbol(v string) *PrivateRestClassicFuturesTradeReversalAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeReversalAPI) MarginCoin(v string) *PrivateRestClassicFuturesTradeReversalAPI {
	api.req.MarginCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeReversalAPI) ProductType(v string) *PrivateRestClassicFuturesTradeReversalAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeReversalAPI) Size(v string) *PrivateRestClassicFuturesTradeReversalAPI {
	api.req.Size = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeReversalAPI) Side(v string) *PrivateRestClassicFuturesTradeReversalAPI {
	api.req.Side = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeReversalAPI) TradeSide(v string) *PrivateRestClassicFuturesTradeReversalAPI {
	api.req.TradeSide = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeReversalAPI) ClientOid(v string) *PrivateRestClassicFuturesTradeReversalAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeBatchOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeBatchOrderReq
}

type PrivateRestClassicFuturesTradeBatchOrderReq struct {
	Symbol      *string                                        `json:"symbol"`
	ProductType *string                                        `json:"productType"`
	MarginCoin  *string                                        `json:"marginCoin"`
	MarginMode  *string                                        `json:"marginMode"`
	OrderList   []PrivateRestClassicFuturesTradeBatchOrderItem `json:"orderList"`
}

type PrivateRestClassicFuturesTradeBatchOrderItem struct {
	Size                   *string `json:"size"`
	Price                  *string `json:"price"`
	Side                   *string `json:"side"`
	TradeSide              *string `json:"tradeSide"`
	OrderType              *string `json:"orderType"`
	Force                  *string `json:"force"`
	ClientOid              *string `json:"clientOid"`
	ReduceOnly             *string `json:"reduceOnly"`
	PresetStopSurplusPrice *string `json:"presetStopSurplusPrice"`
	PresetStopLossPrice    *string `json:"presetStopLossPrice"`
	StpMode                *string `json:"stpMode"`
}

func (api *PrivateRestClassicFuturesTradeBatchOrderAPI) Symbol(v string) *PrivateRestClassicFuturesTradeBatchOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeBatchOrderAPI) ProductType(v string) *PrivateRestClassicFuturesTradeBatchOrderAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeBatchOrderAPI) MarginCoin(v string) *PrivateRestClassicFuturesTradeBatchOrderAPI {
	api.req.MarginCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeBatchOrderAPI) MarginMode(v string) *PrivateRestClassicFuturesTradeBatchOrderAPI {
	api.req.MarginMode = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeBatchOrderAPI) AddOrder(item PrivateRestClassicFuturesTradeBatchOrderItem) *PrivateRestClassicFuturesTradeBatchOrderAPI {
	api.req.OrderList = append(api.req.OrderList, item)
	return api
}

type PrivateRestClassicFuturesTradeModifyOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeModifyOrderReq
}

type PrivateRestClassicFuturesTradeModifyOrderReq struct {
	OrderId                   *string `json:"orderId"`
	ClientOid                 *string `json:"clientOid"`
	Symbol                    *string `json:"symbol"`
	ProductType               *string `json:"productType"`
	NewClientOid              *string `json:"newClientOid"`
	NewSize                   *string `json:"newSize"`
	NewPrice                  *string `json:"newPrice"`
	NewPresetStopSurplusPrice *string `json:"newPresetStopSurplusPrice"`
	NewPresetStopLossPrice    *string `json:"newPresetStopLossPrice"`
}

func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) OrderId(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) ClientOid(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) Symbol(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) ProductType(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) NewClientOid(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.NewClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) NewSize(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.NewSize = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) NewPrice(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.NewPrice = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) NewPresetStopSurplusPrice(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.NewPresetStopSurplusPrice = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeModifyOrderAPI) NewPresetStopLossPrice(v string) *PrivateRestClassicFuturesTradeModifyOrderAPI {
	api.req.NewPresetStopLossPrice = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeCancelOrderReq
}

type PrivateRestClassicFuturesTradeCancelOrderReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	MarginCoin  *string `json:"marginCoin"`
	OrderId     *string `json:"orderId"`
	ClientOid   *string `json:"clientOid"`
}

func (api *PrivateRestClassicFuturesTradeCancelOrderAPI) Symbol(v string) *PrivateRestClassicFuturesTradeCancelOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeCancelOrderAPI) ProductType(v string) *PrivateRestClassicFuturesTradeCancelOrderAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeCancelOrderAPI) MarginCoin(v string) *PrivateRestClassicFuturesTradeCancelOrderAPI {
	api.req.MarginCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeCancelOrderAPI) OrderId(v string) *PrivateRestClassicFuturesTradeCancelOrderAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeCancelOrderAPI) ClientOid(v string) *PrivateRestClassicFuturesTradeCancelOrderAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeBatchCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeBatchCancelOrderReq
}

type PrivateRestClassicFuturesTradeBatchCancelOrderReq struct {
	OrderIdList []PrivateRestClassicFuturesTradeCancelOrderIdItem `json:"orderIdList"`
	Symbol      *string                                           `json:"symbol"`
	ProductType *string                                           `json:"productType"`
	MarginCoin  *string                                           `json:"marginCoin"`
}

type PrivateRestClassicFuturesTradeCancelOrderIdItem struct {
	OrderId   *string `json:"orderId"`
	ClientOid *string `json:"clientOid"`
}

func (api *PrivateRestClassicFuturesTradeBatchCancelOrderAPI) Symbol(v string) *PrivateRestClassicFuturesTradeBatchCancelOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeBatchCancelOrderAPI) ProductType(v string) *PrivateRestClassicFuturesTradeBatchCancelOrderAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeBatchCancelOrderAPI) MarginCoin(v string) *PrivateRestClassicFuturesTradeBatchCancelOrderAPI {
	api.req.MarginCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeBatchCancelOrderAPI) AddOrder(item PrivateRestClassicFuturesTradeCancelOrderIdItem) *PrivateRestClassicFuturesTradeBatchCancelOrderAPI {
	api.req.OrderIdList = append(api.req.OrderIdList, item)
	return api
}

type PrivateRestClassicFuturesTradeClosePositionsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeClosePositionsReq
}

type PrivateRestClassicFuturesTradeClosePositionsReq struct {
	Symbol      *string `json:"symbol"`
	HoldSide    *string `json:"holdSide"`
	ProductType *string `json:"productType"`
}

func (api *PrivateRestClassicFuturesTradeClosePositionsAPI) Symbol(v string) *PrivateRestClassicFuturesTradeClosePositionsAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeClosePositionsAPI) HoldSide(v string) *PrivateRestClassicFuturesTradeClosePositionsAPI {
	api.req.HoldSide = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeClosePositionsAPI) ProductType(v string) *PrivateRestClassicFuturesTradeClosePositionsAPI {
	api.req.ProductType = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeOrderDetailAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeOrderDetailReq
}

type PrivateRestClassicFuturesTradeOrderDetailReq struct {
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	OrderId     *string `json:"orderId"`
	ClientOid   *string `json:"clientOid"`
}

func (api *PrivateRestClassicFuturesTradeOrderDetailAPI) Symbol(v string) *PrivateRestClassicFuturesTradeOrderDetailAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderDetailAPI) ProductType(v string) *PrivateRestClassicFuturesTradeOrderDetailAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderDetailAPI) OrderId(v string) *PrivateRestClassicFuturesTradeOrderDetailAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderDetailAPI) ClientOid(v string) *PrivateRestClassicFuturesTradeOrderDetailAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeOrderFillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeOrderFillsReq
}

type PrivateRestClassicFuturesTradeOrderFillsReq struct {
	OrderId     *string `json:"orderId"`
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	IDLessThan  *string `json:"idLessThan"`
	StartTime   *string `json:"startTime"`
	EndTime     *string `json:"endTime"`
	Limit       *string `json:"limit"`
}

func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) OrderId(v string) *PrivateRestClassicFuturesTradeOrderFillsAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) Symbol(v string) *PrivateRestClassicFuturesTradeOrderFillsAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) ProductType(v string) *PrivateRestClassicFuturesTradeOrderFillsAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) IDLessThan(v string) *PrivateRestClassicFuturesTradeOrderFillsAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) StartTime(v string) *PrivateRestClassicFuturesTradeOrderFillsAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) EndTime(v string) *PrivateRestClassicFuturesTradeOrderFillsAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrderFillsAPI) Limit(v string) *PrivateRestClassicFuturesTradeOrderFillsAPI {
	api.req.Limit = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeFillHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeFillHistoryReq
}

type PrivateRestClassicFuturesTradeFillHistoryReq = PrivateRestClassicFuturesTradeOrderFillsReq

func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) OrderId(v string) *PrivateRestClassicFuturesTradeFillHistoryAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) Symbol(v string) *PrivateRestClassicFuturesTradeFillHistoryAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) ProductType(v string) *PrivateRestClassicFuturesTradeFillHistoryAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) IDLessThan(v string) *PrivateRestClassicFuturesTradeFillHistoryAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) StartTime(v string) *PrivateRestClassicFuturesTradeFillHistoryAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) EndTime(v string) *PrivateRestClassicFuturesTradeFillHistoryAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeFillHistoryAPI) Limit(v string) *PrivateRestClassicFuturesTradeFillHistoryAPI {
	api.req.Limit = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeOrdersPendingAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeOrdersPendingReq
}

type PrivateRestClassicFuturesTradeOrdersPendingReq struct {
	OrderId     *string `json:"orderId"`
	ClientOid   *string `json:"clientOid"`
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	Status      *string `json:"status"`
	IDLessThan  *string `json:"idLessThan"`
	StartTime   *string `json:"startTime"`
	EndTime     *string `json:"endTime"`
	Limit       *string `json:"limit"`
}

func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) OrderId(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) ClientOid(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) Symbol(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) ProductType(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) Status(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.Status = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) IDLessThan(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) StartTime(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) EndTime(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersPendingAPI) Limit(v string) *PrivateRestClassicFuturesTradeOrdersPendingAPI {
	api.req.Limit = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeOrdersHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeOrdersHistoryReq
}

type PrivateRestClassicFuturesTradeOrdersHistoryReq struct {
	OrderId     *string `json:"orderId"`
	ClientOid   *string `json:"clientOid"`
	Symbol      *string `json:"symbol"`
	ProductType *string `json:"productType"`
	IDLessThan  *string `json:"idLessThan"`
	OrderSource *string `json:"orderSource"`
	EndTime     *string `json:"endTime"`
	StartTime   *string `json:"startTime"`
	Limit       *string `json:"limit"`
}

func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) OrderId(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) ClientOid(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) Symbol(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) ProductType(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) IDLessThan(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) OrderSource(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.OrderSource = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) EndTime(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) StartTime(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeOrdersHistoryAPI) Limit(v string) *PrivateRestClassicFuturesTradeOrdersHistoryAPI {
	api.req.Limit = GetPointer(v)
	return api
}

type PrivateRestClassicFuturesTradeCancelAllOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicFuturesTradeCancelAllOrdersReq
}

type PrivateRestClassicFuturesTradeCancelAllOrdersReq struct {
	ProductType   *string `json:"productType"`
	MarginCoin    *string `json:"marginCoin"`
	RequestTime   *string `json:"requestTime"`
	ReceiveWindow *string `json:"receiveWindow"`
}

func (api *PrivateRestClassicFuturesTradeCancelAllOrdersAPI) ProductType(v string) *PrivateRestClassicFuturesTradeCancelAllOrdersAPI {
	api.req.ProductType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeCancelAllOrdersAPI) MarginCoin(v string) *PrivateRestClassicFuturesTradeCancelAllOrdersAPI {
	api.req.MarginCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeCancelAllOrdersAPI) RequestTime(v string) *PrivateRestClassicFuturesTradeCancelAllOrdersAPI {
	api.req.RequestTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicFuturesTradeCancelAllOrdersAPI) ReceiveWindow(v string) *PrivateRestClassicFuturesTradeCancelAllOrdersAPI {
	api.req.ReceiveWindow = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedTradePlaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradePlaceOrderReq
}

type PrivateRestClassicMarginIsolatedTradePlaceOrderReq struct {
	Symbol    *string `json:"symbol"`    // 交易对，如 BTCUSDT
	OrderType *string `json:"orderType"` // 订单类型：limit/market
	Price     *string `json:"price"`     // 价格（限价单传入）
	LoanType  *string `json:"loanType"`  // 杠杆订单模式：normal/autoLoan/autoRepay/autoLoanAndRepay
	Force     *string `json:"force"`     // 订单策略：gtc/post_only/fok/ioc
	BaseSize  *string `json:"baseSize"`  // limit 必填，market sell 必填；基础币数量
	QuoteSize *string `json:"quoteSize"` // market buy 必填；计价币数量
	ClientOid *string `json:"clientOid"` // 自定义 ID
	Side      *string `json:"side"`      // 交易方向：buy/sell
	StpMode   *string `json:"stpMode"`   // STP 模式：none/cancel_taker/cancel_maker/cancel_both
}

func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) OrderType(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.OrderType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) Price(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.Price = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) LoanType(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.LoanType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) Force(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.Force = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) BaseSize(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.BaseSize = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) QuoteSize(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.QuoteSize = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) ClientOid(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) Side(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.Side = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI) StpMode(v string) *PrivateRestClassicMarginIsolatedTradePlaceOrderAPI {
	api.req.StpMode = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderReq
}

type PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderReq struct {
	Symbol    *string                                               `json:"symbol"`    // 交易对，如 BTCUSDT
	OrderList []PrivateRestClassicMarginIsolatedTradeBatchOrderItem `json:"orderList"` // 订单列表
}

type PrivateRestClassicMarginIsolatedTradeBatchOrderItem struct {
	OrderType *string `json:"orderType"` // 订单类型：limit/market
	Price     *string `json:"price"`     // 价格（限价单传入）
	LoanType  *string `json:"loanType"`  // 杠杆订单模式：normal/autoLoan/autoRepay/autoLoanAndRepay
	Force     *string `json:"force"`     // 订单策略：gtc/post_only/fok/ioc
	BaseSize  *string `json:"baseSize"`  // limit 必填，market sell 必填；基础币数量
	QuoteSize *string `json:"quoteSize"` // market buy 必填；计价币数量
	ClientOid *string `json:"clientOid"` // 自定义 ID
	Side      *string `json:"side"`      // 交易方向：buy/sell
	StpMode   *string `json:"stpMode"`   // STP 模式：none/cancel_taker/cancel_maker/cancel_both
}

func (api *PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI) AddOrder(item PrivateRestClassicMarginIsolatedTradeBatchOrderItem) *PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderAPI {
	api.req.OrderList = append(api.req.OrderList, item)
	return api
}

type PrivateRestClassicMarginIsolatedTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradeCancelOrderReq
}

type PrivateRestClassicMarginIsolatedTradeCancelOrderReq struct {
	Symbol    *string `json:"symbol"`    // 交易对，如 BTCUSDT
	OrderId   *string `json:"orderId"`   // 订单 ID（orderId 与 clientOid 至少传一个）
	ClientOid *string `json:"clientOid"` // 自定义 ID（orderId 与 clientOid 至少传一个）
}

func (api *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI) OrderId(v string) *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI) ClientOid(v string) *PrivateRestClassicMarginIsolatedTradeCancelOrderAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradeBatchCancelOrderReq
}

type PrivateRestClassicMarginIsolatedTradeBatchCancelOrderReq struct {
	Symbol      *string                                                `json:"symbol"`      // 交易对，如 BTCUSDT
	OrderIdList []PrivateRestClassicMarginIsolatedTradeCancelOrderItem `json:"orderIdList"` // 撤单列表
}

type PrivateRestClassicMarginIsolatedTradeCancelOrderItem struct {
	OrderId   *string `json:"orderId"`   // 订单 ID（orderId 与 clientOid 至少传一个）
	ClientOid *string `json:"clientOid"` // 自定义 ID（orderId 与 clientOid 至少传一个）
}

func (api *PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI) AddOrder(item PrivateRestClassicMarginIsolatedTradeCancelOrderItem) *PrivateRestClassicMarginIsolatedTradeBatchCancelOrderAPI {
	api.req.OrderIdList = append(api.req.OrderIdList, item)
	return api
}

type PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradeOpenOrdersReq
}

type PrivateRestClassicMarginIsolatedTradeOpenOrdersReq struct {
	Symbol    *string `json:"symbol"`    // 交易对，如 BTCUSDT
	OrderId   *string `json:"orderId"`   // 订单 ID
	ClientOid *string `json:"clientOid"` // 客户自定义 ID
	StartTime *string `json:"startTime"` // 开始时间，Unix 毫秒时间戳
	EndTime   *string `json:"endTime"`   // 结束时间，Unix 毫秒时间戳
	Limit     *string `json:"limit"`     // 查询条数，默认 100，最大 500
}

func (api *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI) OrderId(v string) *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI) ClientOid(v string) *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI) StartTime(v string) *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI) EndTime(v string) *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI) Limit(v string) *PrivateRestClassicMarginIsolatedTradeOpenOrdersAPI {
	api.req.Limit = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradeHistoryOrdersReq
}

type PrivateRestClassicMarginIsolatedTradeHistoryOrdersReq struct {
	Symbol           *string `json:"symbol"`           // 交易对，如 BTCUSDT
	OrderId          *string `json:"orderId"`          // 订单 ID
	EnterPointSource *string `json:"enterPointSource"` // 订单来源：WEB/IOS/ANDROID/API/SYS
	ClientOid        *string `json:"clientOid"`        // 客户自定义 ID
	StartTime        *string `json:"startTime"`        // 开始时间，Unix 毫秒时间戳
	EndTime          *string `json:"endTime"`          // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit            *string `json:"limit"`            // 查询条数，默认 100，最大 500
	IDLessThan       *string `json:"idLessThan"`       // 翻页游标，传上次返回最小 orderId
}

func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) OrderId(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) EnterPointSource(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.EnterPointSource = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) ClientOid(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) StartTime(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) EndTime(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) Limit(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI) IDLessThan(v string) *PrivateRestClassicMarginIsolatedTradeHistoryOrdersAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedTradeFillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradeFillsReq
}

type PrivateRestClassicMarginIsolatedTradeFillsReq struct {
	Symbol     *string `json:"symbol"`     // 交易对，如 BTCUSDT
	OrderId    *string `json:"orderId"`    // 订单 ID
	IDLessThan *string `json:"idLessThan"` // 撮合 ID 翻页游标，传上次返回最后一条 fillId
	StartTime  *string `json:"startTime"`  // 开始时间，Unix 毫秒时间戳
	EndTime    *string `json:"endTime"`    // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit      *string `json:"limit"`      // 查询条数，默认 100，最大 500
}

func (api *PrivateRestClassicMarginIsolatedTradeFillsAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradeFillsAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeFillsAPI) OrderId(v string) *PrivateRestClassicMarginIsolatedTradeFillsAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeFillsAPI) IDLessThan(v string) *PrivateRestClassicMarginIsolatedTradeFillsAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeFillsAPI) StartTime(v string) *PrivateRestClassicMarginIsolatedTradeFillsAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeFillsAPI) EndTime(v string) *PrivateRestClassicMarginIsolatedTradeFillsAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeFillsAPI) Limit(v string) *PrivateRestClassicMarginIsolatedTradeFillsAPI {
	api.req.Limit = GetPointer(v)
	return api
}

type PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginIsolatedTradeLiquidationOrderReq
}

type PrivateRestClassicMarginIsolatedTradeLiquidationOrderReq struct {
	Type       *string `json:"type"`       // 强平类型：swap/place_order（默认 place_order）
	Symbol     *string `json:"symbol"`     // 交易对（type=place_order 时生效）
	FromCoin   *string `json:"fromCoin"`   // 兑换来源币种（type=swap 时生效）
	ToCoin     *string `json:"toCoin"`     // 兑换目标币种（type=swap 时生效）
	StartTime  *string `json:"startTime"`  // 开始时间，Unix 毫秒时间戳
	EndTime    *string `json:"endTime"`    // 结束时间，Unix 毫秒时间戳（与开始时间最大间隔 90 天）
	Limit      *string `json:"limit"`      // 查询条数，默认 100，最大 500
	IDLessThan *string `json:"idLessThan"` // 翻页游标，传上次返回 endId
}

func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) Type(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.Type = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) Symbol(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) FromCoin(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.FromCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) ToCoin(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.ToCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) StartTime(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) EndTime(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) Limit(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI) IDLessThan(v string) *PrivateRestClassicMarginIsolatedTradeLiquidationOrderAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
