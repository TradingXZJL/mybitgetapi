package mybitgetapi

type PrivateRestClassicFuturesTradeOrderIdRes struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
}

type PrivateRestClassicFuturesTradeFailureItem struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
	Symbol    string `json:"symbol"`
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}

type PrivateRestClassicFuturesTradeBatchRes struct {
	SuccessList []PrivateRestClassicFuturesTradeOrderIdRes  `json:"successList"`
	FailureList []PrivateRestClassicFuturesTradeFailureItem `json:"failureList"`
}

type PrivateRestClassicFuturesTradePlaceOrderRes = PrivateRestClassicFuturesTradeOrderIdRes
type PrivateRestClassicFuturesTradeReversalRes = PrivateRestClassicFuturesTradeOrderIdRes
type PrivateRestClassicFuturesTradeModifyOrderRes = PrivateRestClassicFuturesTradeOrderIdRes
type PrivateRestClassicFuturesTradeCancelOrderRes = PrivateRestClassicFuturesTradeOrderIdRes
type PrivateRestClassicFuturesTradeBatchOrderRes = PrivateRestClassicFuturesTradeBatchRes
type PrivateRestClassicFuturesTradeBatchCancelOrderRes = PrivateRestClassicFuturesTradeBatchRes
type PrivateRestClassicFuturesTradeClosePositionsRes = PrivateRestClassicFuturesTradeBatchRes
type PrivateRestClassicFuturesTradeCancelAllOrdersRes = PrivateRestClassicFuturesTradeBatchRes

type PrivateRestClassicFuturesTradeOrderDetailRes struct {
	Symbol                 string `json:"symbol"`
	Size                   string `json:"size"`
	OrderId                string `json:"orderId"`
	ClientOid              string `json:"clientOid"`
	BaseVolume             string `json:"baseVolume"`
	PriceAvg               string `json:"priceAvg"`
	Fee                    string `json:"fee"`
	Price                  string `json:"price"`
	State                  string `json:"state"`
	Side                   string `json:"side"`
	Force                  string `json:"force"`
	TotalProfits           string `json:"totalProfits"`
	PosSide                string `json:"posSide"`
	MarginCoin             string `json:"marginCoin"`
	PresetStopSurplusPrice string `json:"presetStopSurplusPrice"`
	PresetStopLossPrice    string `json:"presetStopLossPrice"`
	QuoteVolume            string `json:"quoteVolume"`
	OrderType              string `json:"orderType"`
	Leverage               string `json:"leverage"`
	MarginMode             string `json:"marginMode"`
	ReduceOnly             string `json:"reduceOnly"`
	EnterPointSource       string `json:"enterPointSource"`
	TradeSide              string `json:"tradeSide"`
	PosMode                string `json:"posMode"`
	OrderSource            string `json:"orderSource"`
	CancelReason           string `json:"cancelReason"`
	CTime                  string `json:"cTime"`
	UTime                  string `json:"uTime"`
}

type PrivateRestClassicFuturesTradeFeeDetail struct {
	Deduction         string `json:"deduction"`
	FeeCoin           string `json:"feeCoin"`
	TotalDeductionFee string `json:"totalDeductionFee"`
	TotalFee          string `json:"totalFee"`
}

type PrivateRestClassicFuturesTradeFillItem struct {
	TradeId          string                                    `json:"tradeId"`
	Symbol           string                                    `json:"symbol"`
	MarginCoin       string                                    `json:"marginCoin"`
	OrderId          string                                    `json:"orderId"`
	Price            string                                    `json:"price"`
	BaseVolume       string                                    `json:"baseVolume"`
	FeeDetail        []PrivateRestClassicFuturesTradeFeeDetail `json:"feeDetail"`
	Side             string                                    `json:"side"`
	QuoteVolume      string                                    `json:"quoteVolume"`
	Profit           string                                    `json:"profit"`
	EnterPointSource string                                    `json:"enterPointSource"`
	TradeSide        string                                    `json:"tradeSide"`
	PosMode          string                                    `json:"posMode"`
	TradeScope       string                                    `json:"tradeScope"`
	CTime            string                                    `json:"cTime"`
}

type PrivateRestClassicFuturesTradeOrderFillsRes struct {
	FillList []PrivateRestClassicFuturesTradeFillItem `json:"fillList"`
	EndId    string                                   `json:"endId"`
}

type PrivateRestClassicFuturesTradeFillHistoryRes = PrivateRestClassicFuturesTradeOrderFillsRes

type PrivateRestClassicFuturesTradeEntrustedItem struct {
	Symbol                        string `json:"symbol"`
	Size                          string `json:"size"`
	OrderId                       string `json:"orderId"`
	ClientOid                     string `json:"clientOid"`
	BaseVolume                    string `json:"baseVolume"`
	Fee                           string `json:"fee"`
	Price                         string `json:"price"`
	PriceAvg                      string `json:"priceAvg"`
	Status                        string `json:"status"`
	Side                          string `json:"side"`
	Force                         string `json:"force"`
	TotalProfits                  string `json:"totalProfits"`
	PosSide                       string `json:"posSide"`
	MarginCoin                    string `json:"marginCoin"`
	QuoteVolume                   string `json:"quoteVolume"`
	Leverage                      string `json:"leverage"`
	MarginMode                    string `json:"marginMode"`
	ReduceOnly                    string `json:"reduceOnly"`
	EnterPointSource              string `json:"enterPointSource"`
	TradeSide                     string `json:"tradeSide"`
	PosMode                       string `json:"posMode"`
	OrderType                     string `json:"orderType"`
	OrderSource                   string `json:"orderSource"`
	LiqPrice                      string `json:"liqPrice"`
	PosAvg                        string `json:"posAvg"`
	CTime                         string `json:"cTime"`
	UTime                         string `json:"uTime"`
	PresetStopSurplusPrice        string `json:"presetStopSurplusPrice"`
	PresetStopSurplusType         string `json:"presetStopSurplusType"`
	PresetStopSurplusExecutePrice string `json:"presetStopSurplusExecutePrice"`
	PresetStopLossPrice           string `json:"presetStopLossPrice"`
	PresetStopLossType            string `json:"presetStopLossType"`
	PresetStopLossExecutePrice    string `json:"presetStopLossExecutePrice"`
}

type PrivateRestClassicFuturesTradeOrdersListRes struct {
	EntrustedList []PrivateRestClassicFuturesTradeEntrustedItem `json:"entrustedList"`
	EndId         string                                        `json:"endId"`
}

type PrivateRestClassicFuturesTradeOrdersPendingRes = PrivateRestClassicFuturesTradeOrdersListRes
type PrivateRestClassicFuturesTradeOrdersHistoryRes = PrivateRestClassicFuturesTradeOrdersListRes

type PrivateRestClassicMarginIsolatedTradeOrderIdRes struct {
	OrderId   string `json:"orderId"`   // 订单 ID
	ClientOid string `json:"clientOid"` // 自定义 ID
}

type PrivateRestClassicMarginIsolatedTradeBatchFailureItem struct {
	OrderId   string `json:"orderId"`   // 订单 ID
	ClientOid string `json:"clientOid"` // 自定义 ID
	ErrorMsg  string `json:"errorMsg"`  // 错误信息
	ErrorCode string `json:"errorCode"` // 错误码
}

type PrivateRestClassicMarginIsolatedTradeBatchRes struct {
	SuccessList []PrivateRestClassicMarginIsolatedTradeOrderIdRes       `json:"successList"` // 成功列表
	FailureList []PrivateRestClassicMarginIsolatedTradeBatchFailureItem `json:"failureList"` // 失败列表
}

type PrivateRestClassicMarginIsolatedTradePlaceOrderRes = PrivateRestClassicMarginIsolatedTradeOrderIdRes
type PrivateRestClassicMarginIsolatedTradeCancelOrderRes = PrivateRestClassicMarginIsolatedTradeOrderIdRes
type PrivateRestClassicMarginIsolatedTradeBatchPlaceOrderRes = PrivateRestClassicMarginIsolatedTradeBatchRes
type PrivateRestClassicMarginIsolatedTradeBatchCancelOrderRes = PrivateRestClassicMarginIsolatedTradeBatchRes

type PrivateRestClassicMarginIsolatedTradeOrderItem struct {
	OrderId          string `json:"orderId"`          // 订单号
	Symbol           string `json:"symbol"`           // 交易对名称
	OrderType        string `json:"orderType"`        // 订单类型：limit/market
	EnterPointSource string `json:"enterPointSource"` // 订单来源
	ClientOid        string `json:"clientOid"`        // 客户自定义 ID
	LoanType         string `json:"loanType"`         // 杠杆订单模式
	Price            string `json:"price"`            // 委托价格
	Side             string `json:"side"`             // 交易方向：buy/sell
	Status           string `json:"status"`           // 订单状态
	BaseSize         string `json:"baseSize"`         // 基础币数量
	QuoteSize        string `json:"quoteSize"`        // 计价币数量
	PriceAvg         string `json:"priceAvg"`         // 成交价格
	Size             string `json:"size"`             // 成交数量
	Amount           string `json:"amount"`           // 成交金额
	Force            string `json:"force"`            // 订单策略
	CTime            string `json:"cTime"`            // 创建时间
	UTime            string `json:"uTime"`            // 更新时间
}

type PrivateRestClassicMarginIsolatedTradeOpenOrdersRes struct {
	OrderList []PrivateRestClassicMarginIsolatedTradeOrderItem `json:"orderList"` // 委托列表
	MaxID     string                                           `json:"maxId"`     // 本次查询最大 ID
	MinID     string                                           `json:"minId"`     // 本次查询最小 ID
}

type PrivateRestClassicMarginIsolatedTradeHistoryOrdersRes = PrivateRestClassicMarginIsolatedTradeOpenOrdersRes

type PrivateRestClassicMarginIsolatedTradeLiquidationOrderItem struct {
	Symbol    string `json:"symbol"`    // 交易对（type=place_order）
	OrderType string `json:"orderType"` // 订单类型（type=place_order）
	Side      string `json:"side"`      // 交易方向（type=place_order）
	PriceAvg  string `json:"priceAvg"`  // 成交价格（type=place_order）
	Price     string `json:"price"`     // 委托价格（type=place_order）
	FillSize  string `json:"fillSize"`  // 成交数量（type=place_order）
	Size      string `json:"size"`      // 委托数量（type=place_order）
	Amount    string `json:"amount"`    // 成交金额（type=place_order）
	OrderId   string `json:"orderId"`   // 订单号
	FromCoin  string `json:"fromCoin"`  // 兑换来源币种（type=swap）
	ToCoin    string `json:"toCoin"`    // 兑换目标币种（type=swap）
	FromSize  string `json:"fromSize"`  // 兑换来源币种数量（type=swap）
	ToSize    string `json:"toSize"`    // 兑换目标币种数量（type=swap）
	CTime     string `json:"cTime"`     // 创建时间
	UTime     string `json:"uTime"`     // 更新时间
}

type PrivateRestClassicMarginIsolatedTradeLiquidationOrderRes struct {
	ResultList []PrivateRestClassicMarginIsolatedTradeLiquidationOrderItem `json:"resultList"` // 强平订单列表
	IDLessThan string                                                      `json:"idLessThan"` // 翻页游标
}

type PrivateRestClassicMarginIsolatedTradeFillsRes struct {
	Fills []PrivateRestClassicMarginIsolatedTradeFillItem `json:"fills"` // 成交明细列表
	MaxID string                                          `json:"maxId"` // 本次查询最大 ID
	MinID string                                          `json:"minId"` // 本次查询最小 ID
}