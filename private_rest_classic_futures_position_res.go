package mybitgetapi

type PrivateRestClassicFuturesPositionQueryPositionLeverResRow struct {
	Symbol         string `json:"symbol"`         // String 币对名
	Level          string `json:"level"`          // String 阶梯档位
	StartUnit      string `json:"startUnit"`      // String 价值下限
	EndUnit        string `json:"endUnit"`        // String 价值上限
	Leverage       string `json:"leverage"`       // String 杠杆倍数
	KeepMarginRate string `json:"keepMarginRate"` // String 维持保证金率
}

type PrivateRestClassicFuturesPositionQueryPositionLeverRes []PrivateRestClassicFuturesPositionQueryPositionLeverResRow

type PrivateRestClassicFuturesPositionInfoRow struct {
	Symbol           string  `json:"symbol"`           // String 币对名称
	MarginCoin       string  `json:"marginCoin"`       // String 保证金币种
	HoldSide         string  `json:"holdSide"`         // String 持仓方向
	OpenDelegateSize string  `json:"openDelegateSize"` // String 待成交数量
	MarginSize       string  `json:"marginSize"`       // String 保证金数量
	Available        string  `json:"available"`        // String 仓位可用
	Locked           string  `json:"locked"`           // String 仓位冻结
	Total            string  `json:"total"`            // String 仓位总数量
	Leverage         string  `json:"leverage"`         // String 杠杆倍数
	AchievedProfits  string  `json:"achievedProfits"`  // String 已实现盈亏
	OpenPriceAvg     string  `json:"openPriceAvg"`     // String 平均开仓价
	MarginMode       string  `json:"marginMode"`       // String 保证金模式
	PosMode          string  `json:"posMode"`          // String 持仓模式
	UnrealizedPL     string  `json:"unrealizedPL"`     // String 未实现盈亏
	LiquidationPrice string  `json:"liquidationPrice"` // String 预估强平价
	KeepMarginRate   string  `json:"keepMarginRate"`   // String 仓位档位维持保证金率
	MarkPrice        string  `json:"markPrice"`        // String 标记价格
	MarginRatio      string  `json:"marginRatio"`      // String 维持保证金率
	BreakEvenPrice   string  `json:"breakEvenPrice"`   // String 盈亏平衡价
	TotalFee         string  `json:"totalFee"`         // String 累计资金费
	DeductedFee      string  `json:"deductedFee"`      // String 已扣手续费
	Grant            *string `json:"grant"`            // String 合约空投券金额
	AssetMode        *string `json:"assetMode"`        // String 资产模式
	AutoMargin       *string `json:"autoMargin"`       // String 自动追加保证金
	TakeProfit       *string `json:"takeProfit"`       // String 止盈价格
	StopLoss         *string `json:"stopLoss"`         // String 止损价格
	TakeProfitId     *string `json:"takeProfitId"`     // String 止盈订单ID
	StopLossId       *string `json:"stopLossId"`       // String 止损订单ID
	CTime            string  `json:"cTime"`            // String 创建时间
	UTime            string  `json:"uTime"`            // String 更新时间
}

type PrivateRestClassicFuturesPositionSinglePositionRes []PrivateRestClassicFuturesPositionInfoRow
type PrivateRestClassicFuturesPositionAllPositionRes []PrivateRestClassicFuturesPositionInfoRow

type PrivateRestClassicFuturesPositionHistoryPositionListRow struct {
	PositionId    string `json:"positionId"`    // String ID
	Symbol        string `json:"symbol"`        // String 币对名称
	MarginCoin    string `json:"marginCoin"`    // String 保证金币种
	HoldSide      string `json:"holdSide"`      // String 持仓方向
	PosMode       string `json:"posMode"`       // String 持仓模式
	OpenAvgPrice  string `json:"openAvgPrice"`  // String 开仓均价
	CloseAvgPrice string `json:"closeAvgPrice"` // String 平仓均价
	MarginMode    string `json:"marginMode"`    // String 保证金模式
	OpenTotalPos  string `json:"openTotalPos"`  // String 累计开仓数量
	CloseTotalPos string `json:"closeTotalPos"` // String 累计平仓数量
	Pnl           string `json:"pnl"`           // String 已实现盈亏
	NetProfit     string `json:"netProfit"`     // String 净盈亏
	TotalFunding  string `json:"totalFunding"`  // String 累计资金费用
	OpenFee       string `json:"openFee"`       // String 开仓手续费
	CloseFee      string `json:"closeFee"`      // String 平仓手续费
	CTime         string `json:"cTime"`         // String 创建时间
	UTime         string `json:"uTime"`         // String 更新时间
}

type PrivateRestClassicFuturesPositionHistoryPositionRes struct {
	List  []PrivateRestClassicFuturesPositionHistoryPositionListRow `json:"list"`  // Array 历史持仓列表
	EndId string                                                    `json:"endId"` // String 最后一条数据ID
}
