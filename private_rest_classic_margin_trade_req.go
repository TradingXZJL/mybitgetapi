package mybitgetapi

type PrivateRestClassicMarginCrossTradePlaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradePlaceOrderReq
}

type PrivateRestClassicMarginCrossTradePlaceOrderReq struct {
	Symbol    *string `json:"symbol"`
	OrderType *string `json:"orderType"`
	Price     *string `json:"price"`
	LoanType  *string `json:"loanType"`
	Force     *string `json:"force"`
	BaseSize  *string `json:"baseSize"`
	QuoteSize *string `json:"quoteSize"`
	ClientOid *string `json:"clientOid"`
	Side      *string `json:"side"`
	StpMode   *string `json:"stpMode"`
}

func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) OrderType(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.OrderType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) Price(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.Price = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) LoanType(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.LoanType = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) Force(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.Force = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) BaseSize(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.BaseSize = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) QuoteSize(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.QuoteSize = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) ClientOid(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) Side(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.Side = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradePlaceOrderAPI) StpMode(v string) *PrivateRestClassicMarginCrossTradePlaceOrderAPI {
	api.req.StpMode = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradeBatchPlaceOrderReq
}

type PrivateRestClassicMarginCrossTradeBatchPlaceOrderReq struct {
	Symbol    *string                                            `json:"symbol"`
	OrderList []PrivateRestClassicMarginCrossTradeBatchOrderItem `json:"orderList"`
}

type PrivateRestClassicMarginCrossTradeBatchOrderItem struct {
	OrderType *string `json:"orderType"`
	Price     *string `json:"price"`
	LoanType  *string `json:"loanType"`
	Force     *string `json:"force"`
	BaseSize  *string `json:"baseSize"`
	QuoteSize *string `json:"quoteSize"`
	ClientOid *string `json:"clientOid"`
	Side      *string `json:"side"`
	StpMode   *string `json:"stpMode"`
}

func (api *PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI) AddOrder(item PrivateRestClassicMarginCrossTradeBatchOrderItem) *PrivateRestClassicMarginCrossTradeBatchPlaceOrderAPI {
	api.req.OrderList = append(api.req.OrderList, item)
	return api
}

type PrivateRestClassicMarginCrossTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradeCancelOrderReq
}

type PrivateRestClassicMarginCrossTradeCancelOrderReq struct {
	Symbol    *string `json:"symbol"`
	OrderId   *string `json:"orderId"`
	ClientOid *string `json:"clientOid"`
}

func (api *PrivateRestClassicMarginCrossTradeCancelOrderAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradeCancelOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeCancelOrderAPI) OrderId(v string) *PrivateRestClassicMarginCrossTradeCancelOrderAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeCancelOrderAPI) ClientOid(v string) *PrivateRestClassicMarginCrossTradeCancelOrderAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradeBatchCancelOrderReq
}

type PrivateRestClassicMarginCrossTradeBatchCancelOrderReq struct {
	Symbol      *string                                             `json:"symbol"`
	OrderIdList []PrivateRestClassicMarginCrossTradeCancelOrderItem `json:"orderIdList"`
}

type PrivateRestClassicMarginCrossTradeCancelOrderItem struct {
	OrderId   *string `json:"orderId"`
	ClientOid *string `json:"clientOid"`
}

func (api *PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI) AddOrder(item PrivateRestClassicMarginCrossTradeCancelOrderItem) *PrivateRestClassicMarginCrossTradeBatchCancelOrderAPI {
	api.req.OrderIdList = append(api.req.OrderIdList, item)
	return api
}

type PrivateRestClassicMarginCrossTradeOpenOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradeOpenOrdersReq
}

type PrivateRestClassicMarginCrossTradeOpenOrdersReq struct {
	Symbol     *string `json:"symbol"`
	OrderId    *string `json:"orderId"`
	ClientOid  *string `json:"clientOid"`
	StartTime  *string `json:"startTime"`
	EndTime    *string `json:"endTime"`
	Limit      *string `json:"limit"`
	IDLessThan *string `json:"idLessThan"`
}

func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) OrderId(v string) *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) ClientOid(v string) *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) StartTime(v string) *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) EndTime(v string) *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) Limit(v string) *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeOpenOrdersAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossTradeOpenOrdersAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossTradeHistoryOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradeHistoryOrdersReq
}

type PrivateRestClassicMarginCrossTradeHistoryOrdersReq struct {
	Symbol           *string `json:"symbol"`
	OrderId          *string `json:"orderId"`
	EnterPointSource *string `json:"enterPointSource"`
	ClientOid        *string `json:"clientOid"`
	StartTime        *string `json:"startTime"`
	EndTime          *string `json:"endTime"`
	Limit            *string `json:"limit"`
	IDLessThan       *string `json:"idLessThan"`
}

func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) OrderId(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) EnterPointSource(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.EnterPointSource = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) ClientOid(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.ClientOid = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) StartTime(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) EndTime(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) Limit(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossTradeHistoryOrdersAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossTradeFillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradeFillsReq
}

type PrivateRestClassicMarginCrossTradeFillsReq struct {
	Symbol     *string `json:"symbol"`
	OrderId    *string `json:"orderId"`
	IDLessThan *string `json:"idLessThan"`
	StartTime  *string `json:"startTime"`
	EndTime    *string `json:"endTime"`
	Limit      *string `json:"limit"`
}

func (api *PrivateRestClassicMarginCrossTradeFillsAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradeFillsAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeFillsAPI) OrderId(v string) *PrivateRestClassicMarginCrossTradeFillsAPI {
	api.req.OrderId = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeFillsAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossTradeFillsAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeFillsAPI) StartTime(v string) *PrivateRestClassicMarginCrossTradeFillsAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeFillsAPI) EndTime(v string) *PrivateRestClassicMarginCrossTradeFillsAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeFillsAPI) Limit(v string) *PrivateRestClassicMarginCrossTradeFillsAPI {
	api.req.Limit = GetPointer(v)
	return api
}

type PrivateRestClassicMarginCrossTradeLiquidationOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestClassicMarginCrossTradeLiquidationOrderReq
}

type PrivateRestClassicMarginCrossTradeLiquidationOrderReq struct {
	Type       *string `json:"type"`
	Symbol     *string `json:"symbol"`
	FromCoin   *string `json:"fromCoin"`
	ToCoin     *string `json:"toCoin"`
	StartTime  *string `json:"startTime"`
	EndTime    *string `json:"endTime"`
	Limit      *string `json:"limit"`
	IDLessThan *string `json:"idLessThan"`
}

func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) Type(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.Type = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) Symbol(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.Symbol = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) FromCoin(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.FromCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) ToCoin(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.ToCoin = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) StartTime(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.StartTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) EndTime(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.EndTime = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) Limit(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.Limit = GetPointer(v)
	return api
}
func (api *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI) IDLessThan(v string) *PrivateRestClassicMarginCrossTradeLiquidationOrderAPI {
	api.req.IDLessThan = GetPointer(v)
	return api
}
