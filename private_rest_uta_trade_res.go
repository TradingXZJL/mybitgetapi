package mybitgetapi

type PrivateRestUtaTradeOrderIdRes struct {
	OrderId   string `json:"orderId"`   // String 订单ID
	ClientOid string `json:"clientOid"` // String 客户端订单ID
}

type PrivateRestUtaTradeBatchOrderIdRes struct {
	OrderId   string `json:"orderId"`   // String 订单ID
	ClientOid string `json:"clientOid"` // String 客户端订单ID
	Code      string `json:"code"`      // String 单条错误码
	Msg       string `json:"msg"`       // String 单条错误信息
}

type PrivateRestUtaTradePlaceOrderRes = PrivateRestUtaTradeOrderIdRes
type PrivateRestUtaTradeModifyOrderRes = PrivateRestUtaTradeOrderIdRes
type PrivateRestUtaTradeCancelOrderRes = PrivateRestUtaTradeOrderIdRes
type PrivateRestUtaTradePlaceBatchRes []PrivateRestUtaTradeBatchOrderIdRes
type PrivateRestUtaTradeBatchModifyOrderRes []PrivateRestUtaTradeBatchOrderIdRes
type PrivateRestUtaTradeCancelBatchRes []PrivateRestUtaTradeBatchOrderIdRes

type PrivateRestUtaTradeCancelSymbolOrderItem struct {
	OrderId   string `json:"orderId"`   // String 订单ID
	ClientOid string `json:"clientOid"` // String 客户端订单ID
	Code      string `json:"code"`      // String 错误码
	Msg       string `json:"msg"`       // String 错误信息
}

type PrivateRestUtaTradeCancelSymbolOrderRes struct {
	List []PrivateRestUtaTradeCancelSymbolOrderItem `json:"list"` // Array 结果列表
}

type PrivateRestUtaTradeClosePositionsRes = PrivateRestUtaTradeCancelSymbolOrderRes

type PrivateRestUtaTradeFeeDetail struct {
	FeeCoin string `json:"feeCoin"` // String 手续费币种
	Fee     string `json:"fee"`     // String 手续费
}

type PrivateRestUtaTradeOrderInfoRes struct {
	OrderId      string                         `json:"orderId"`      // String 订单ID
	ClientOid    string                         `json:"clientOid"`    // String 客户端订单ID
	Category     string                         `json:"category"`     // String 产品类型
	Symbol       string                         `json:"symbol"`       // String 交易对名称
	OrderType    string                         `json:"orderType"`    // String 订单类型 limit/market
	Side         string                         `json:"side"`         // String 方向 buy/sell
	Price        string                         `json:"price"`        // String 委托价格
	Qty          string                         `json:"qty"`          // String 委托数量
	Amount       string                         `json:"amount"`       // String 委托金额（现货/杠杆）
	CumExecQty   string                         `json:"cumExecQty"`   // String 累计成交数量
	CumExecValue string                         `json:"cumExecValue"` // String 累计成交金额
	AvgPrice     string                         `json:"avgPrice"`     // String 平均成交价
	TimeInForce  string                         `json:"timeInForce"`  // String 有效方式
	OrderStatus  string                         `json:"orderStatus"`  // String 订单状态
	PosSide      string                         `json:"posSide"`      // String 持仓方向
	HoldMode     string                         `json:"holdMode"`     // String 持仓模式 one_way_mode/hedge_mode
	TradeSide    string                         `json:"tradeSide"`    // String 交易方向 open/close
	StpMode      string                         `json:"stpMode"`      // String 自成交保护模式
	ReduceOnly   string                         `json:"reduceOnly"`   // String 只减仓标识 YES/NO
	FeeDetail    []PrivateRestUtaTradeFeeDetail `json:"feeDetail"`    // Array 手续费明细
	CancelReason string                         `json:"cancelReason"` // String 撤单原因
	ExecType     string                         `json:"execType"`     // String 执行类型
	CreatedTime  string                         `json:"createdTime"`  // String 创建时间（毫秒）
	UpdatedTime  string                         `json:"updatedTime"`  // String 更新时间（毫秒）
}

type PrivateRestUtaTradeOrderListItem struct {
	OrderId      string                         `json:"orderId"`      // String 订单ID
	ClientOid    string                         `json:"clientOid"`    // String 客户端订单ID
	Category     string                         `json:"category"`     // String 产品类型
	Symbol       string                         `json:"symbol"`       // String 交易对名称
	OrderType    string                         `json:"orderType"`    // String 订单类型 limit/market
	Side         string                         `json:"side"`         // String 方向 buy/sell
	Price        string                         `json:"price"`        // String 委托价格
	Qty          string                         `json:"qty"`          // String 委托数量
	Amount       string                         `json:"amount"`       // String 委托金额（现货/杠杆）
	CumExecQty   string                         `json:"cumExecQty"`   // String 累计成交数量
	CumExecValue string                         `json:"cumExecValue"` // String 累计成交金额
	AvgPrice     string                         `json:"avgPrice"`     // String 平均成交价
	TimeInForce  string                         `json:"timeInForce"`  // String 有效方式
	OrderStatus  string                         `json:"orderStatus"`  // String 订单状态
	PosSide      string                         `json:"posSide"`      // String 持仓方向
	HoldMode     string                         `json:"holdMode"`     // String 持仓模式 one_way_mode/hedge_mode
	DelegateType string                         `json:"delegateType"` // String 委托类型
	ReduceOnly   string                         `json:"reduceOnly"`   // String 只减仓标识 YES/NO
	StpMode      string                         `json:"stpMode"`      // String 自成交保护模式
	FeeDetail    []PrivateRestUtaTradeFeeDetail `json:"feeDetail"`    // Array 手续费明细
	CancelReason string                         `json:"cancelReason"` // String 撤单原因
	ExecType     string                         `json:"execType"`     // String 执行类型
	CreatedTime  string                         `json:"createdTime"`  // String 创建时间（毫秒）
	UpdatedTime  string                         `json:"updatedTime"`  // String 更新时间（毫秒）
}

type PrivateRestUtaTradeOrderListRes struct {
	List   []PrivateRestUtaTradeOrderListItem `json:"list"`   // Array 订单列表
	Cursor string                             `json:"cursor"` // String 游标
}

type PrivateRestUtaTradeUnfilledOrdersRes = PrivateRestUtaTradeOrderListRes
type PrivateRestUtaTradeHistoryOrdersRes = PrivateRestUtaTradeOrderListRes

type PrivateRestUtaTradeFillItem struct {
	ExecId      string                         `json:"execId"`      // String 成交ID
	ExecLinkId  string                         `json:"execLinkId"`  // String 成交关联ID
	OrderId     string                         `json:"orderId"`     // String 订单ID
	ClientOid   string                         `json:"clientOid"`   // String 客户端订单ID
	Category    string                         `json:"category"`    // String 产品类型
	Symbol      string                         `json:"symbol"`      // String 交易对名称
	OrderType   string                         `json:"orderType"`   // String 订单类型
	Side        string                         `json:"side"`        // String 成交方向 buy/sell
	ExecPrice   string                         `json:"execPrice"`   // String 成交价格
	ExecQty     string                         `json:"execQty"`     // String 成交数量
	ExecValue   string                         `json:"execValue"`   // String 成交金额
	TradeScope  string                         `json:"tradeScope"`  // String 成交角色 maker/taker
	TradeSide   string                         `json:"tradeSide"`   // String 交易方向 open/close
	FeeDetail   []PrivateRestUtaTradeFeeDetail `json:"feeDetail"`   // Array 手续费明细
	CreatedTime string                         `json:"createdTime"` // String 创建时间（毫秒）
	UpdatedTime string                         `json:"updatedTime"` // String 更新时间（毫秒）
	ExecPnl     string                         `json:"execPnl"`     // String 平仓盈亏
	IsRPI       string                         `json:"isRPI"`       // String 是否RPI成交 yes/no
}

type PrivateRestUtaTradeFillsRes struct {
	List   []PrivateRestUtaTradeFillItem `json:"list"`   // Array 成交列表
	Cursor string                        `json:"cursor"` // String 游标
}

type PrivateRestUtaPositionCurrentItem struct {
	Category         string `json:"category"`         // String 产品类型
	Symbol           string `json:"symbol"`           // String 交易对名称
	MarginCoin       string `json:"marginCoin"`       // String 保证金币种
	HoldMode         string `json:"holdMode"`         // String 持仓模式 one_way_mode/hedge_mode
	PosSide          string `json:"posSide"`          // String 持仓方向 long/short
	MarginMode       string `json:"marginMode"`       // String 保证金模式 crossed/isolated
	PositionBalance  string `json:"positionBalance"`  // String 持仓保证金
	Available        string `json:"available"`        // String 可用仓位
	Frozen           string `json:"frozen"`           // String 冻结仓位
	Total            string `json:"total"`            // String 总仓位
	Leverage         string `json:"leverage"`         // String 杠杆倍数
	CurRealisedPnl   string `json:"curRealisedPnl"`   // String 当前已实现盈亏
	AvgPrice         string `json:"avgPrice"`         // String 持仓均价
	PositionStatus   string `json:"positionStatus"`   // String 持仓状态
	UnrealisedPnl    string `json:"unrealisedPnl"`    // String 未实现盈亏
	LiquidationPrice string `json:"liquidationPrice"` // String 预估强平价
	Mmr              string `json:"mmr"`              // String 维持保证金率
	ProfitRate       string `json:"profitRate"`       // String 收益率
	MarkPrice        string `json:"markPrice"`        // String 标记价格
	BreakEvenPrice   string `json:"breakEvenPrice"`   // String 保本价
	TotalFunding     string `json:"totalFunding"`     // String 累计资金费
	OpenFeeTotal     string `json:"openFeeTotal"`     // String 开仓累计手续费
	CloseFeeTotal    string `json:"closeFeeTotal"`    // String 平仓累计手续费
	CreatedTime      string `json:"createdTime"`      // String 创建时间（毫秒）
	UpdatedTime      string `json:"updatedTime"`      // String 更新时间（毫秒）
}

type PrivateRestUtaPositionCurrentRes struct {
	List []PrivateRestUtaPositionCurrentItem `json:"list"` // Array 当前持仓列表
}

type PrivateRestUtaPositionHistoryItem struct {
	PositionId     string `json:"positionId"`     // String 持仓ID
	Category       string `json:"category"`       // String 产品类型
	Symbol         string `json:"symbol"`         // String 交易对名称
	MarginCoin     string `json:"marginCoin"`     // String 保证金币种
	HoldMode       string `json:"holdMode"`       // String 持仓模式 one_way_mode/hedge_mode
	PosSide        string `json:"posSide"`        // String 持仓方向 long/short
	MarginMode     string `json:"marginMode"`     // String 保证金模式 crossed/isolated
	OpenPriceAvg   string `json:"openPriceAvg"`   // String 开仓均价
	ClosePriceAvg  string `json:"closePriceAvg"`  // String 平仓均价
	OpenTotalPos   string `json:"openTotalPos"`   // String 累计开仓数量
	CloseTotalPos  string `json:"closeTotalPos"`  // String 累计平仓数量
	CumRealisedPnl string `json:"cumRealisedPnl"` // String 累计已实现盈亏（不含费用）
	NetProfit      string `json:"netProfit"`      // String 净盈亏（含费用和资金费）
	TotalFunding   string `json:"totalFunding"`   // String 累计资金费
	OpenFeeTotal   string `json:"openFeeTotal"`   // String 开仓累计手续费
	CloseFeeTotal  string `json:"closeFeeTotal"`  // String 平仓累计手续费
	CreatedTime    string `json:"createdTime"`    // String 创建时间（毫秒）
	UpdatedTime    string `json:"updatedTime"`    // String 更新时间（毫秒）
}

type PrivateRestUtaPositionHistoryRes struct {
	List   []PrivateRestUtaPositionHistoryItem `json:"list"`   // Array 历史持仓列表
	Cursor string                              `json:"cursor"` // String 游标
}
