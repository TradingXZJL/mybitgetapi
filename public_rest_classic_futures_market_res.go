package mybitgetapi

import "strconv"

type PublicRestClassicFuturesMarketVipFeeRateResRow struct {
	Level              string `json:"level"`              // String VIP等级
	DealAmount         string `json:"dealAmount"`         // String 近30日累计成交额
	AssetAmount        string `json:"assetAmount"`        // String 资产总额
	TakerFeeRate       string `json:"takerFeeRate"`       // String taker费率
	MakerFeeRate       string `json:"makerFeeRate"`       // String maker费率
	BtcWithdrawAmount  string `json:"btcWithdrawAmount"`  // String BTC提币限额
	UsdtWithdrawAmount string `json:"usdtWithdrawAmount"` // String USDT提币限额
}

type PublicRestClassicFuturesMarketVipFeeRateRes []PublicRestClassicFuturesMarketVipFeeRateResRow

type PublicRestClassicFuturesMarketInterestRateHistoryItem struct {
	Ts                 string `json:"ts"`                 // String 日期
	AnnualInterestRate string `json:"annualInterestRate"` // String 年利率
	DailyInterestRate  string `json:"dailyInterestRate"`  // String 日利率
}

type PublicRestClassicFuturesMarketInterestRateRes struct {
	Coin                    string                                                  `json:"coin"`                    // String 币种
	HistoryInterestRateList []PublicRestClassicFuturesMarketInterestRateHistoryItem `json:"historyInterestRateList"` // List 利率记录
}

type PublicRestClassicFuturesMarketMergeDepthResMiddle struct {
	Asks           [][]float64 `json:"asks"`           // Array 卖盘深度
	Bids           [][]float64 `json:"bids"`           // Array 买盘深度
	Ts             string      `json:"ts"`             // String 时间戳
	Scale          string      `json:"scale"`          // String 实际精度
	Precision      string      `json:"precision"`      // String 请求精度
	IsMaxPrecision string      `json:"isMaxPrecision"` // String 是否最大精度
}

type PublicRestClassicFuturesMarketMergeDepthRes struct {
	Asks           []Books `json:"asks"`           // Array 卖盘深度
	Bids           []Books `json:"bids"`           // Array 买盘深度
	Ts             string  `json:"ts"`             // String 时间戳
	Scale          string  `json:"scale"`          // String 实际精度
	Precision      string  `json:"precision"`      // String 请求精度
	IsMaxPrecision string  `json:"isMaxPrecision"` // String 是否最大精度
}

func (m *PublicRestClassicFuturesMarketMergeDepthResMiddle) ConvertToRes() PublicRestClassicFuturesMarketMergeDepthRes {
	asks := make([]Books, 0, len(m.Asks))
	for _, ask := range m.Asks {
		if len(ask) < 2 {
			continue
		}
		asks = append(asks, Books{
			Price:    strconv.FormatFloat(ask[0], 'f', -1, 64),
			Quantity: strconv.FormatFloat(ask[1], 'f', -1, 64),
		})
	}

	bids := make([]Books, 0, len(m.Bids))
	for _, bid := range m.Bids {
		if len(bid) < 2 {
			continue
		}
		bids = append(bids, Books{
			Price:    strconv.FormatFloat(bid[0], 'f', -1, 64),
			Quantity: strconv.FormatFloat(bid[1], 'f', -1, 64),
		})
	}

	return PublicRestClassicFuturesMarketMergeDepthRes{
		Asks:           asks,
		Bids:           bids,
		Ts:             m.Ts,
		Scale:          m.Scale,
		Precision:      m.Precision,
		IsMaxPrecision: m.IsMaxPrecision,
	}
}

type PublicRestClassicFuturesMarketTickerResRow struct {
	Symbol            string `json:"symbol"`            // String 币对名称
	LastPr            string `json:"lastPr"`            // String 最新成交价
	AskPr             string `json:"askPr"`             // String 卖一价
	BidPr             string `json:"bidPr"`             // String 买一价
	BidSz             string `json:"bidSz"`             // String 买一量
	AskSz             string `json:"askSz"`             // String 卖一量
	High24h           string `json:"high24h"`           // String 24小时最高价
	Low24h            string `json:"low24h"`            // String 24小时最低价
	Ts                string `json:"ts"`                // String 时间戳
	Change24h         string `json:"change24h"`         // String 24小时涨跌幅
	BaseVolume        string `json:"baseVolume"`        // String 基础币成交量
	QuoteVolume       string `json:"quoteVolume"`       // String 计价币成交量
	UsdtVolume        string `json:"usdtVolume"`        // String USDT成交量
	OpenUTC           string `json:"openUtc"`           // String UTC开盘价
	ChangeUTC24h      string `json:"changeUtc24h"`      // String UTC 24h涨跌幅
	IndexPrice        string `json:"indexPrice"`        // String 指数价格
	FundingRate       string `json:"fundingRate"`       // String 资金费率
	HoldingAmount     string `json:"holdingAmount"`     // String 持仓数量
	DeliveryStartTime string `json:"deliveryStartTime"` // String 交割开始时间
	DeliveryTime      string `json:"deliveryTime"`      // String 交割时间
	DeliveryStatus    string `json:"deliveryStatus"`    // String 交割状态
	Open24h           string `json:"open24h"`           // String 开盘价
	MarkPrice         string `json:"markPrice"`         // String 标记价格
}

type PublicRestClassicFuturesMarketTickerRes []PublicRestClassicFuturesMarketTickerResRow
type PublicRestClassicFuturesMarketTickersRes []PublicRestClassicFuturesMarketTickerResRow

type PublicRestClassicFuturesMarketFillsResRow struct {
	TradeID string `json:"tradeId"` // String 成交ID
	Price   string `json:"price"`   // String 价格
	Size    string `json:"size"`    // String 数量
	Side    string `json:"side"`    // String 方向
	Ts      string `json:"ts"`      // String 时间戳
	Symbol  string `json:"symbol"`  // String 交易对
}

type PublicRestClassicFuturesMarketFillsRes []PublicRestClassicFuturesMarketFillsResRow
type PublicRestClassicFuturesMarketFillsHistoryRes []PublicRestClassicFuturesMarketFillsResRow

type PublicRestClassicFuturesMarketCandlesResMiddle [][]string
type PublicRestClassicFuturesMarketCandlesRes []Kline

func (m *PublicRestClassicFuturesMarketCandlesResMiddle) ConvertToRes() PublicRestClassicFuturesMarketCandlesRes {
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
	return PublicRestClassicFuturesMarketCandlesRes(klines)
}

type PublicRestClassicFuturesMarketHistoryCandlesResMiddle [][]string
type PublicRestClassicFuturesMarketHistoryCandlesRes []Kline

func (m *PublicRestClassicFuturesMarketHistoryCandlesResMiddle) ConvertToRes() PublicRestClassicFuturesMarketHistoryCandlesRes {
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
	return PublicRestClassicFuturesMarketHistoryCandlesRes(klines)
}

type PublicRestClassicFuturesMarketHistoryIndexCandlesResMiddle [][]string
type PublicRestClassicFuturesMarketHistoryIndexCandlesRes []Kline

func (m *PublicRestClassicFuturesMarketHistoryIndexCandlesResMiddle) ConvertToRes() PublicRestClassicFuturesMarketHistoryIndexCandlesRes {
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
	return PublicRestClassicFuturesMarketHistoryIndexCandlesRes(klines)
}

type PublicRestClassicFuturesMarketHistoryMarkCandlesResMiddle [][]string
type PublicRestClassicFuturesMarketHistoryMarkCandlesRes []Kline

func (m *PublicRestClassicFuturesMarketHistoryMarkCandlesResMiddle) ConvertToRes() PublicRestClassicFuturesMarketHistoryMarkCandlesRes {
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
	return PublicRestClassicFuturesMarketHistoryMarkCandlesRes(klines)
}

type PublicRestClassicFuturesMarketCurrentFundRateResRow struct {
	Symbol              string `json:"symbol"`              // String 交易对
	FundingRate         string `json:"fundingRate"`         // String 当前资金费率
	FundingRateInterval string `json:"fundingRateInterval"` // String 结算周期（小时）
	NextUpdate          string `json:"nextUpdate"`          // String 下次更新时间
	MinFundingRate      string `json:"minFundingRate"`      // String 资金费率下限
	MaxFundingRate      string `json:"maxFundingRate"`      // String 资金费率上限
}

type PublicRestClassicFuturesMarketCurrentFundRateRes []PublicRestClassicFuturesMarketCurrentFundRateResRow

type PublicRestClassicFuturesMarketHistoryFundRateResRow struct {
	Symbol      string `json:"symbol"`      // String 交易对
	FundingRate string `json:"fundingRate"` // String 资金费率
	FundingTime string `json:"fundingTime"` // String 结算时间
}

type PublicRestClassicFuturesMarketHistoryFundRateRes []PublicRestClassicFuturesMarketHistoryFundRateResRow

type PublicRestClassicFuturesMarketFundingTimeResRow struct {
	Symbol          string `json:"symbol"`          // String 交易对
	NextFundingTime string `json:"nextFundingTime"` // String 下次结算时间
	RatePeriod      string `json:"ratePeriod"`      // String 结算周期
}

type PublicRestClassicFuturesMarketFundingTimeRes []PublicRestClassicFuturesMarketFundingTimeResRow

type PublicRestClassicFuturesMarketContractsResRow struct {
	Symbol              string   `json:"symbol"`              // String 产品名称
	BaseCoin            string   `json:"baseCoin"`            // String 基础币
	QuoteCoin           string   `json:"quoteCoin"`           // String 计价币
	BuyLimitPriceRatio  string   `json:"buyLimitPriceRatio"`  // String 买价限价比例
	SellLimitPriceRatio string   `json:"sellLimitPriceRatio"` // String 卖价限价比例
	FeeRateUpRatio      string   `json:"feeRateUpRatio"`      // String 手续费上浮比例
	MakerFeeRate        string   `json:"makerFeeRate"`        // String maker费率
	TakerFeeRate        string   `json:"takerFeeRate"`        // String taker费率
	OpenCostUpRatio     string   `json:"openCostUpRatio"`     // String 开仓成本上浮比例
	SupportMarginCoins  []string `json:"supportMarginCoins"`  // List 支持保证金币种
	MinTradeNum         string   `json:"minTradeNum"`         // String 最小开单数量
	PriceEndStep        string   `json:"priceEndStep"`        // String 价格步长
	VolumePlace         string   `json:"volumePlace"`         // String 数量小数位
	PricePlace          string   `json:"pricePlace"`          // String 价格小数位
	SizeMultiplier      string   `json:"sizeMultiplier"`      // String 数量乘数
	SymbolType          string   `json:"symbolType"`          // String 合约类型
	MinTradeUSDT        string   `json:"minTradeUSDT"`        // String 最小USDT交易额
	MaxSymbolOrderNum   string   `json:"maxSymbolOrderNum"`   // String symbol维度最大订单数
	MaxProductOrderNum  string   `json:"maxProductOrderNum"`  // String 产品维度最大订单数
	MaxPositionNum      string   `json:"maxPositionNum"`      // String 最大仓位数
	SymbolStatus        string   `json:"symbolStatus"`        // String 交易对状态
	OffTime             string   `json:"offTime"`             // String 停盘时间
	LimitOpenTime       string   `json:"limitOpenTime"`       // String 可开仓时间
	DeliveryTime        string   `json:"deliveryTime"`        // String 交割时间
	DeliveryStartTime   string   `json:"deliveryStartTime"`   // String 交割开始时间
	DeliveryPeriod      string   `json:"deliveryPeriod"`      // String 交割周期
	LaunchTime          string   `json:"launchTime"`          // String 上架时间
	FundInterval        string   `json:"fundInterval"`        // String 资金费结算周期
	MinLever            string   `json:"minLever"`            // String 最小杠杆
	MaxLever            string   `json:"maxLever"`            // String 最大杠杆
	PosLimit            string   `json:"posLimit"`            // String 持仓限制（废弃）
	MaintainTime        string   `json:"maintainTime"`        // String 维护时间
	MaxMarketOrderQty   string   `json:"maxMarketOrderQty"`   // String 市价单最大数量
	MaxOrderQty         string   `json:"maxOrderQty"`         // String 限价单最大数量
	IsRwa               string   `json:"isRwa"`               // String 是否RWA
	OpenTime            string   `json:"openTime"`            // String 已废弃字段
}

type PublicRestClassicFuturesMarketContractsRes []PublicRestClassicFuturesMarketContractsResRow
