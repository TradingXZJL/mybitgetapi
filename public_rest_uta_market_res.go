package mybitgetapi

import "strconv"

type PublicRestUtaMarketInstrumentsResRow struct {
	Category                   string `json:"category"`                   // String 产品类型 SPOT 现货交易 MARGIN 杠杆 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol                     string `json:"symbol"`                     // String 交易对名称 例如：BTCUSDT
	IsRwa                      string `json:"isRwa"`                      // String 是否为RWA交易对 YES 是 NO 否
	BaseCoin                   string `json:"baseCoin"`                   // String 交易币种 例如：BTCUSDT中的BTC
	QuoteCoin                  string `json:"quoteCoin"`                  // String 计价币种 例如：BTCUSDT中的USDT
	BuyLimitPriceRatio         string `json:"buyLimitPriceRatio"`         // String 买价限价比例 买入限价与市场价格的比率，决定买单下单时的最高价格
	SellLimitPriceRatio        string `json:"sellLimitPriceRatio"`        // String 卖价限价比例 卖出限价与市场价格的比率，决定卖单下单时的最低价格
	FeeRateUpRatio             string `json:"feeRateUpRatio"`             // String 手续费上浮比例 实际费用相对于基础费用增加的百分比
	OpenCostUpRatio            string `json:"openCostUpRatio"`            // String 开仓成本上浮比例 开仓的成本相对于基础或标准成本增加的百分比
	MinOrderQty                string `json:"minOrderQty"`                // String 最小开单数量 这指的是以基础币种下单时允许的最小数量 仅适用于合约交易，现货/杠杆请参考交易规则
	MaxOrderQty                string `json:"maxOrderQty"`                // String 单笔限价单最大下单数量 这指的是以基础币种下单时允许的最大数量 仅适用于合约交易，现货/杠杆请参考交易规则
	MinOrderAmount             string `json:"minOrderAmount"`             // String 最小下单额 这指的是以计价币种下单时允许的最小数量
	PricePrecision             string `json:"pricePrecision"`             // String 价格精度 价格允许的小数位数 下单时，填写的价格必须是价格精度的整数倍
	QuantityPrecision          string `json:"quantityPrecision"`          // String 数量精度 数量允许的小数位数 下单时，填写的数量必须是数量精度的整数倍
	QuotePrecision             string `json:"quotePrecision"`             // String 市价下单精度 计价币价格允许的小数位数
	PriceMultiplier            string `json:"priceMultiplier"`            // String 价格乘数 合约下单使用, 配合pricePrecision 例如: BTCUSDT pricePrecision: 2 精度就是0.01 priceMultiplier: 0.02 那么委托价格要满足priceMultiplier的倍数, 并且小数位是 2
	QuantityMultiplier         string `json:"quantityMultiplier"`         // String 数量乘数 合约下单使用. 配合quantityPrecision 例如：BTCUSDT quantityPrecision: 2 精度就是0.01 quantityMultiplier: 0.02 那么委托数量要满足quantityMultiplier的倍数, 并且小数位是 2
	Type                       string `json:"type"`                       // String 合约类型 perpetual 永续 delivery 交割
	MaxSymbolOrderNum          string `json:"maxSymbolOrderNum"`          // String 交易对维度最大委托数量
	MaxProductOrderNum         string `json:"maxProductOrderNum"`         // String 业务线维度最大委托数量
	MaxPositionNum             string `json:"maxPositionNum"`             // String 最大持有仓位数量（交易对维度）
	Status                     string `json:"status"`                     // String 交易对状态 listed 上架(未开盘) online 正常交易/开盘 limit_open 限制交易(合约禁止开仓) limit_close 限制交易(合约禁止平仓) offline 下架/维护 restrictedAPI API限制下单
	OffTime                    string `json:"offTime"`                    // String 停盘时间 未配置则返回: ""
	LimitOpenTime              string `json:"limitOpenTime"`              // String 限制开仓时间 未配置则返回: ""; 其它值表示symbol正在/计划维护，指定时间后禁止交易
	DeliveryTime               string `json:"deliveryTime"`               // String 交割时间 仅交割合约返回
	DeliveryStartTime          string `json:"deliveryStartTime"`          // String 交割币对开始时间 仅交割合约返回
	DeliveryPeriod             string `json:"deliveryPeriod"`             // String 交割周期 this_quarter当季 next_quarter次季 仅交割合约返回
	LaunchTime                 string `json:"launchTime"`                 // String 上线时间 Unix毫秒时间戳，表示交易对的上线时间
	FundInterval               string `json:"fundInterval"`               // String 资金费结算周期 1 每小时/8 每8小时
	MinLeverage                string `json:"minLeverage"`                // String 最小杠杆
	MaxLeverage                string `json:"maxLeverage"`                // String 最大杠杆
	MaintainTime               string `json:"maintainTime"`               // String 币对维护时间（状态处于维护/即将维护时会有值） 未配置则返回: ""
	IsIsolatedBaseBorrowable   string `json:"isIsolatedBaseBorrowable"`   // String 逐仓左币是否可以借 仅适用于杠杆交易
	IsIsolatedQuotedBorrowable string `json:"isIsolatedQuotedBorrowable"` // String 逐仓右币是否可以借 仅适用于杠杆交易
	WarningRiskRatio           string `json:"warningRiskRatio"`           // String 预警风险率
	LiquidationRiskRatio       string `json:"liquidationRiskRatio"`       // String 强平风险率
	MaxCrossedLeverage         string `json:"maxCrossedLeverage"`         // String 全仓最大杠杆倍数 仅适用于杠杆交易
	MaxIsolatedLeverage        string `json:"maxIsolatedLeverage"`        // String 逐仓最大杠杆倍数 仅适用于杠杆交易
	UserMinBorrow              string `json:"userMinBorrow"`              // String 最小可借 仅适用杠杆交易
	AreaSymbol                 string `json:"areaSymbol"`                 // String 是否区域币对 YES NO 仅适用于现货交易 只有在为YES 的币对时返回该参数
	MakerFeeRate               string `json:"makerFeeRate"`               // String 挂单手续费率 小数形式，即0.0002代表万分之二
	TakerFeeRate               string `json:"takerFeeRate"`               // String 吃单手续费率 小数形式，即0.0002代表万分之二
	MaxMarketOrderQty          string `json:"maxMarketOrderQty"`          // String 单笔市价单最大下单数量 这指的是以基础币种下单时允许的最大数量
}

type PublicRestUtaMarketInstrumentsRes []PublicRestUtaMarketInstrumentsResRow

type PublicRestUtaMarketTickersResRow struct {
	Category          string `json:"category"`          // String 产品类型 SPOT 现货交易 USDT-FUTURES U本位合约 COIN-FUTURES 币本位合约 USDC-FUTURES USDC合约
	Symbol            string `json:"symbol"`            // String 交易对名称 例如：BTCUSDT
	LastPrice         string `json:"lastPrice"`         // String 最新成交价
	OpenPrice24h      string `json:"openPrice24h"`      // String 24小时前开盘价
	HighPrice24h      string `json:"highPrice24h"`      // String 24小时最高价
	LowPrice24h       string `json:"lowPrice24h"`       // String 24小时最低价
	Ask1Price         string `json:"ask1Price"`         // String 最优卖价
	Bid1Price         string `json:"bid1Price"`         // String 最优买价
	Bid1Size          string `json:"bid1Size"`          // String 最优买量
	Ask1Size          string `json:"ask1Size"`          // String 最优卖量
	Price24hPcnt      string `json:"price24hPcnt"`      // String 24小时涨跌幅比例
	Volume24h         string `json:"volume24h"`         // String 24小时成交量
	Turnover24h       string `json:"turnover24h"`       // String 24小时成交额
	IndexPrice        string `json:"indexPrice"`        // String 指数价格 仅合约返回
	MarkPrice         string `json:"markPrice"`         // String 标记价格 仅合约返回
	FundingRate       string `json:"fundingRate"`       // String 资金费率 仅合约返回
	OpenInterest      string `json:"openInterest"`      // String 持仓量 仅合约返回
	DeliveryStartTime string `json:"deliveryStartTime"` // String 交割开始时间 仅交割合约返回
	DeliveryTime      string `json:"deliveryTime"`      // String 交割时间 仅交割合约返回
	DeliveryStatus    string `json:"deliveryStatus"`    // String 交割状态 仅交割合约返回
	Ts                string `json:"ts"`                // String 数据产生时间 Unix毫秒时间戳
}

type PublicRestUtaMarketTickersRes []PublicRestUtaMarketTickersResRow

type Books struct {
	Price    string `json:"price"`    // String 价格
	Quantity string `json:"quantity"` // String 数量
}

type PublicRestUtaMarketOrderBookResMiddle struct {
	A  [][]float64 `json:"a"`  // Array 卖盘深度 按价格升序
	B  [][]float64 `json:"b"`  // Array 买盘深度 按价格降序
	Ts string      `json:"ts"` // String 数据产生时间 Unix毫秒时间戳
}

type PublicRestUtaMarketOrderBookRes struct {
	Asks []Books `json:"asks"` // Array 卖盘深度 按价格升序
	Bids []Books `json:"bids"` // Array 买盘深度 按价格降序
	Ts   string  `json:"ts"`   // String 数据产生时间 Unix毫秒时间戳
}

func (m *PublicRestUtaMarketOrderBookResMiddle) ConvertToRes() PublicRestUtaMarketOrderBookRes {
	asks := make([]Books, 0, len(m.A))
	for _, ask := range m.A {
		if len(ask) < 2 {
			continue
		}
		asks = append(asks, Books{
			Price:    strconv.FormatFloat(ask[0], 'f', -1, 64),
			Quantity: strconv.FormatFloat(ask[1], 'f', -1, 64),
		})
	}

	bids := make([]Books, 0, len(m.B))
	for _, bid := range m.B {
		if len(bid) < 2 {
			continue
		}
		bids = append(bids, Books{
			Price:    strconv.FormatFloat(bid[0], 'f', -1, 64),
			Quantity: strconv.FormatFloat(bid[1], 'f', -1, 64),
		})
	}

	return PublicRestUtaMarketOrderBookRes{
		Asks: asks,
		Bids: bids,
		Ts:   m.Ts,
	}
}

type Kline struct {
	Period     string `json:"period,omitempty"`     // String 周期（该接口通常不返回）
	OpenTime   string `json:"openTime"`             // String 开盘时间/时间戳
	CloseTime  string `json:"closeTime,omitempty"`  // String 收盘时间（该接口通常不返回）
	OpenPrice  string `json:"openPrice"`            // String 开盘价
	ClosePrice string `json:"closePrice"`           // String 收盘价
	HighPrice  string `json:"highPrice"`            // String 最高价
	LowPrice   string `json:"lowPrice"`             // String 最低价
	FillQty    string `json:"fillQty,omitempty"`    // String 成交量（base coin）
	FillAmount string `json:"fillAmount,omitempty"` // String 成交额（quote coin）
}

type PublicRestUtaMarketCandlesResMiddle [][]string

type PublicRestUtaMarketCandlesRes []Kline

func (m *PublicRestUtaMarketCandlesResMiddle) ConvertToRes() PublicRestUtaMarketCandlesRes {
	klines := make([]Kline, 0, len(*m))
	for _, row := range *m {
		// Bitget candles数组: [0]ts [1]open [2]high [3]low [4]close [5]volume [6]turnover
		klines = append(klines, Kline{
			OpenTime:   getRowValue(row, 0),
			OpenPrice:  getRowValue(row, 1),
			HighPrice:  getRowValue(row, 2),
			LowPrice:   getRowValue(row, 3),
			ClosePrice: getRowValue(row, 4),
			FillQty:    getRowValue(row, 5),
			FillAmount: getRowValue(row, 6),
		})
	}
	return PublicRestUtaMarketCandlesRes(klines)
}

type PublicRestUtaMarketHistoryCandlesResMiddle [][]string

type PublicRestUtaMarketHistoryCandlesRes []Kline

func (m *PublicRestUtaMarketHistoryCandlesResMiddle) ConvertToRes() PublicRestUtaMarketHistoryCandlesRes {
	klines := make([]Kline, 0, len(*m))
	for _, row := range *m {
		klines = append(klines, Kline{
			OpenTime:   getRowValue(row, 0),
			OpenPrice:  getRowValue(row, 1),
			HighPrice:  getRowValue(row, 2),
			LowPrice:   getRowValue(row, 3),
			ClosePrice: getRowValue(row, 4),
			FillQty:    getRowValue(row, 5),
			FillAmount: getRowValue(row, 6),
		})
	}
	return PublicRestUtaMarketHistoryCandlesRes(klines)
}

func getRowValue(row []string, index int) string {
	if len(row) <= index {
		return ""
	}
	return row[index]
}

type PublicRestUtaMarketCurrentFundRateResRow struct {
	Symbol              string `json:"symbol"`              // String 交易对名称
	FundingRate         string `json:"fundingRate"`         // String 当前资金费率
	FundingRateInterval string `json:"fundingRateInterval"` // String 资金费结算周期 小时
	NextUpdate          string `json:"nextUpdate"`          // String 下次更新时间 Unix毫秒时间戳
	MinFundingRate      string `json:"minFundingRate"`      // String 资金费率下限
	MaxFundingRate      string `json:"maxFundingRate"`      // String 资金费率上限
}

type PublicRestUtaMarketCurrentFundRateRes []PublicRestUtaMarketCurrentFundRateResRow

type PublicRestUtaMarketHistoryFundRateResRow struct {
	Symbol               string `json:"symbol"`               // String 交易对名称
	FundingRate          string `json:"fundingRate"`          // String 历史资金费率
	FundingRateTimestamp string `json:"fundingRateTimestamp"` // String 资金费率时间 Unix毫秒时间戳
}

type PublicRestUtaMarketHistoryFundRateRes struct {
	ResultList []PublicRestUtaMarketHistoryFundRateResRow `json:"resultList"` // List 数据列表
}
