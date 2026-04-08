package mybitgetapi

type PrivateRestClassicSpotTradeOrderIdRes struct {
	OrderId   string `json:"orderId"`   // String 订单ID
	ClientOid string `json:"clientOid"` // String 客户端订单ID
}

type PrivateRestClassicSpotTradeCancelReplaceOrderRes struct {
	OrderId   string `json:"orderId"`   // String 订单ID
	ClientOid string `json:"clientOid"` // String 客户端订单ID
	Success   string `json:"success"`   // String 是否成功 success/failure
	Msg       string `json:"msg"`       // String 失败原因
}

type PrivateRestClassicSpotTradeBatchResultItem struct {
	OrderId   string `json:"orderId"`   // String 订单ID
	ClientOid string `json:"clientOid"` // String 客户端订单ID
	ErrorMsg  string `json:"errorMsg"`  // String 错误信息
	ErrorCode string `json:"errorCode"` // String 错误码
}

type PrivateRestClassicSpotTradeBatchResultRes struct {
	SuccessList []PrivateRestClassicSpotTradeOrderIdRes      `json:"successList"` // Array 成功订单
	FailureList []PrivateRestClassicSpotTradeBatchResultItem `json:"failureList"` // Array 失败订单
}

type PrivateRestClassicSpotTradePlaceOrderRes = PrivateRestClassicSpotTradeOrderIdRes
type PrivateRestClassicSpotTradeCancelOrderRes = PrivateRestClassicSpotTradeOrderIdRes
type PrivateRestClassicSpotTradeBatchCancelReplaceOrderRes []PrivateRestClassicSpotTradeCancelReplaceOrderRes
type PrivateRestClassicSpotTradeBatchOrdersRes = PrivateRestClassicSpotTradeBatchResultRes
type PrivateRestClassicSpotTradeBatchCancelOrderRes = PrivateRestClassicSpotTradeBatchResultRes

type PrivateRestClassicSpotTradeOrderInfoRow struct {
	UserId           string `json:"userId"`           // String 账户ID
	Symbol           string `json:"symbol"`           // String 交易对名称
	OrderId          string `json:"orderId"`          // String 订单ID
	ClientOid        string `json:"clientOid"`        // String 客户端订单ID
	Price            string `json:"price"`            // String 委托价格
	Size             string `json:"size"`             // String 委托数量
	OrderType        string `json:"orderType"`        // String 订单类型
	Side             string `json:"side"`             // String 方向
	Status           string `json:"status"`           // String 订单状态
	PriceAvg         string `json:"priceAvg"`         // String 成交均价
	BaseVolume       string `json:"baseVolume"`       // String 成交数量（base）
	QuoteVolume      string `json:"quoteVolume"`      // String 成交总额（quote）
	EnterPointSource string `json:"enterPointSource"` // String 订单来源端
	FeeDetail        string `json:"feeDetail"`        // String 手续费明细
	OrderSource      string `json:"orderSource"`      // String 订单来源
	CancelReason     string `json:"cancelReason"`     // String 取消原因
	CTime            string `json:"cTime"`            // String 创建时间
	UTime            string `json:"uTime"`            // String 更新时间
}

type PrivateRestClassicSpotTradeOrderInfoRes []PrivateRestClassicSpotTradeOrderInfoRow

type PrivateRestClassicSpotTradeOrderRow struct {
	UserId                 string `json:"userId"`                 // String 账户ID
	Symbol                 string `json:"symbol"`                 // String 交易对名称
	OrderId                string `json:"orderId"`                // String 订单ID
	ClientOid              string `json:"clientOid"`              // String 客户端订单ID
	Price                  string `json:"price"`                  // String 委托价格
	PriceAvg               string `json:"priceAvg"`               // String 成交均价
	Size                   string `json:"size"`                   // String 委托数量
	OrderType              string `json:"orderType"`              // String 订单类型
	Side                   string `json:"side"`                   // String 方向
	Status                 string `json:"status"`                 // String 订单状态
	BasePrice              string `json:"basePrice"`              // String 成交价格
	BaseVolume             string `json:"baseVolume"`             // String 成交数量
	QuoteVolume            string `json:"quoteVolume"`            // String 成交总额
	EnterPointSource       string `json:"enterPointSource"`       // String 订单来源端
	OrderSource            string `json:"orderSource"`            // String 订单来源
	TriggerPrice           string `json:"triggerPrice"`           // String 触发价格
	TpslType               string `json:"tpslType"`               // String 订单类型 normal/tpsl
	PresetTakeProfitPrice  string `json:"presetTakeProfitPrice"`  // String 止盈触发价格
	ExecuteTakeProfitPrice string `json:"executeTakeProfitPrice"` // String 止盈执行价格
	PresetStopLossPrice    string `json:"presetStopLossPrice"`    // String 止损触发价格
	ExecuteStopLossPrice   string `json:"executeStopLossPrice"`   // String 止损执行价格
	FeeDetail              string `json:"feeDetail"`              // String 手续费明细（history接口）
	CancelReason           string `json:"cancelReason"`           // String 取消原因（history接口）
	CTime                  string `json:"cTime"`                  // String 创建时间
	UTime                  string `json:"uTime"`                  // String 更新时间
}

type PrivateRestClassicSpotTradeUnfilledOrdersRes []PrivateRestClassicSpotTradeOrderRow
type PrivateRestClassicSpotTradeHistoryOrdersRes []PrivateRestClassicSpotTradeOrderRow

type PrivateRestClassicSpotTradeFillsFeeDetail struct {
	Deduction         string `json:"deduction"`         // String 是否折扣
	FeeCoin           string `json:"feeCoin"`           // String 手续费币种
	TotalDeductionFee string `json:"totalDeductionFee"` // String 总计折扣手续费
	TotalFee          string `json:"totalFee"`          // String 总计手续费
}

type PrivateRestClassicSpotTradeFillsRow struct {
	UserId     string                                    `json:"userId"`     // String 账户ID
	Symbol     string                                    `json:"symbol"`     // String 交易对名称
	OrderId    string                                    `json:"orderId"`    // String 订单ID
	TradeId    string                                    `json:"tradeId"`    // String 成交ID
	OrderType  string                                    `json:"orderType"`  // String 订单类型
	Side       string                                    `json:"side"`       // String 方向
	PriceAvg   string                                    `json:"priceAvg"`   // String 成交价格
	Size       string                                    `json:"size"`       // String 成交数量
	Amount     string                                    `json:"amount"`     // String 成交总额
	FeeDetail  PrivateRestClassicSpotTradeFillsFeeDetail `json:"feeDetail"`  // Object 手续费明细
	TradeScope string                                    `json:"tradeScope"` // String maker/taker
	CTime      string                                    `json:"cTime"`      // String 创建时间
	UTime      string                                    `json:"uTime"`      // String 更新时间
}

type PrivateRestClassicSpotTradeFillsRes []PrivateRestClassicSpotTradeFillsRow
