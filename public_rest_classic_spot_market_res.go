package mybitgetapi

import "strconv"

type PublicRestClassicSpotMarketCoinsResChain struct {
	Chain             string `json:"chain"`             // String 链名称
	NeedTag           string `json:"needTag"`           // String 是否需要 tag
	Withdrawable      string `json:"withdrawable"`      // String 是否可提现
	Rechargeable      string `json:"rechargeable"`      // String 是否可充值
	WithdrawFee       string `json:"withdrawFee"`       // String 提现手续费
	ExtraWithdrawFee  string `json:"extraWithdrawFee"`  // String 额外收取比例，0.1 表示 10%
	DepositConfirm    string `json:"depositConfirm"`    // String 充值确认块数
	WithdrawConfirm   string `json:"withdrawConfirm"`   // String 提现确认块数
	MinDepositAmount  string `json:"minDepositAmount"`  // String 最小充值数
	MinWithdrawAmount string `json:"minWithdrawAmount"` // String 最小提现数
	BrowserURL        string `json:"browserUrl"`        // String 区块浏览器地址
	ContractAddress   string `json:"contractAddress"`   // String 币种合约地址
	WithdrawStep      string `json:"withdrawStep"`      // String 提币步长
	WithdrawMinScale  string `json:"withdrawMinScale"`  // String 提币数量精度
	Congestion        string `json:"congestion"`        // String 链网络拥堵情况
}

type PublicRestClassicSpotMarketCoinsResRow struct {
	CoinID   string                                     `json:"coinId"`   // String 币种 ID
	Coin     string                                     `json:"coin"`     // String 币种名称
	Transfer string                                     `json:"transfer"` // String 是否可以划转
	Chains   []PublicRestClassicSpotMarketCoinsResChain `json:"chains"`   // Array 支持的链列表
}

type PublicRestClassicSpotMarketCoinsRes []PublicRestClassicSpotMarketCoinsResRow

type PublicRestClassicSpotMarketSymbolsResRow struct {
	Symbol              string `json:"symbol"`              // String 交易对名称
	BaseCoin            string `json:"baseCoin"`            // String 基础币
	QuoteCoin           string `json:"quoteCoin"`           // String 计价币
	MinTradeAmount      string `json:"minTradeAmount"`      // String 最小交易数量（文档标注废弃）
	MaxTradeAmount      string `json:"maxTradeAmount"`      // String 最大交易数量（文档标注废弃）
	TakerFeeRate        string `json:"takerFeeRate"`        // String 默认吃单手续费率
	MakerFeeRate        string `json:"makerFeeRate"`        // String 默认挂单手续费率
	PricePrecision      string `json:"pricePrecision"`      // String 价格精度
	QuantityPrecision   string `json:"quantityPrecision"`   // String 数量精度
	QuotePrecision      string `json:"quotePrecision"`      // String 右币精度
	Status              string `json:"status"`              // String 上架状态
	MinTradeUSDT        string `json:"minTradeUSDT"`        // String 最小 USDT 交易额
	BuyLimitPriceRatio  string `json:"buyLimitPriceRatio"`  // String 买价限价比例
	SellLimitPriceRatio string `json:"sellLimitPriceRatio"` // String 卖价限价比例
	AreaSymbol          string `json:"areaSymbol"`          // String 区域币对标记
	OrderQuantity       string `json:"orderQuantity"`       // String 当前币对最多允许订单数量
	OpenTime            string `json:"openTime"`            // String 该字段已废弃
	OffTime             string `json:"offTime"`             // String 下架时间
}

type PublicRestClassicSpotMarketSymbolsRes []PublicRestClassicSpotMarketSymbolsResRow

type PublicRestClassicSpotMarketTickersResRow struct {
	Symbol       string `json:"symbol"`       // String 交易对名称
	High24h      string `json:"high24h"`      // String 24小时最高价
	Open         string `json:"open"`         // String 24小时开盘价
	Low24h       string `json:"low24h"`       // String 24小时最低价
	LastPr       string `json:"lastPr"`       // String 最新成交价
	QuoteVolume  string `json:"quoteVolume"`  // String 计价币成交额
	BaseVolume   string `json:"baseVolume"`   // String 基础币成交额
	UsdtVolume   string `json:"usdtVolume"`   // String USDT 成交额
	BidPr        string `json:"bidPr"`        // String 买一价
	AskPr        string `json:"askPr"`        // String 卖一价
	BidSz        string `json:"bidSz"`        // String 买一量
	AskSz        string `json:"askSz"`        // String 卖一量
	OpenUTC      string `json:"openUtc"`      // String UTC0 开盘价
	Ts           string `json:"ts"`           // String 时间戳
	ChangeUTC24h string `json:"changeUtc24h"` // String UTC0 时涨跌幅
	Change24h    string `json:"change24h"`    // String 24小时涨跌幅
}

type PublicRestClassicSpotMarketTickersRes []PublicRestClassicSpotMarketTickersResRow

type PublicRestClassicSpotMarketMergeDepthResMiddle struct {
	Asks           [][]float64 `json:"asks"`           // Array 卖盘深度
	Bids           [][]float64 `json:"bids"`           // Array 买盘深度
	Ts             string      `json:"ts"`             // String 撮合引擎时间戳
	Scale          string      `json:"scale"`          // String 实际精度值
	Precision      string      `json:"precision"`      // String 当前档位
	IsMaxPrecision string      `json:"isMaxPrecision"` // String 是否最大精度
}

type PublicRestClassicSpotMarketMergeDepthRes struct {
	Asks           []Books `json:"asks"`           // Array 卖盘深度
	Bids           []Books `json:"bids"`           // Array 买盘深度
	Ts             string  `json:"ts"`             // String 撮合引擎时间戳
	Scale          string  `json:"scale"`          // String 实际精度值
	Precision      string  `json:"precision"`      // String 当前档位
	IsMaxPrecision string  `json:"isMaxPrecision"` // String 是否最大精度
}

func (m *PublicRestClassicSpotMarketMergeDepthResMiddle) ConvertToRes() PublicRestClassicSpotMarketMergeDepthRes {
	asks := make([]Books, 0, len(m.Asks))
	for _, ask := range m.Asks {
		asks = append(asks, Books{
			Price:    strconv.FormatFloat(ask[0], 'f', -1, 64),
			Quantity: strconv.FormatFloat(ask[1], 'f', -1, 64),
		})
	}

	bids := make([]Books, 0, len(m.Bids))
	for _, bid := range m.Bids {
		bids = append(bids, Books{
			Price:    strconv.FormatFloat(bid[0], 'f', -1, 64),
			Quantity: strconv.FormatFloat(bid[1], 'f', -1, 64),
		})
	}

	return PublicRestClassicSpotMarketMergeDepthRes{
		Asks:           asks,
		Bids:           bids,
		Ts:             m.Ts,
		Scale:          m.Scale,
		Precision:      m.Precision,
		IsMaxPrecision: m.IsMaxPrecision,
	}
}

type PublicRestClassicSpotMarketOrderbookResMiddle struct {
	Asks [][]string `json:"asks"` // Array 卖盘深度
	Bids [][]string `json:"bids"` // Array 买盘深度
	Ts   string     `json:"ts"`   // String 撮合引擎时间戳
}

type PublicRestClassicSpotMarketOrderbookRes struct {
	Asks []Books `json:"asks"` // Array 卖盘深度
	Bids []Books `json:"bids"` // Array 买盘深度
	Ts   string  `json:"ts"`   // String 撮合引擎时间戳
}

func (m *PublicRestClassicSpotMarketOrderbookResMiddle) ConvertToRes() PublicRestClassicSpotMarketOrderbookRes {
	asks := make([]Books, 0, len(m.Asks))
	for _, ask := range m.Asks {
		asks = append(asks, Books{
			Price:    getRowValue(ask, 0),
			Quantity: getRowValue(ask, 1),
		})
	}

	bids := make([]Books, 0, len(m.Bids))
	for _, bid := range m.Bids {
		bids = append(bids, Books{
			Price:    getRowValue(bid, 0),
			Quantity: getRowValue(bid, 1),
		})
	}

	return PublicRestClassicSpotMarketOrderbookRes{
		Asks: asks,
		Bids: bids,
		Ts:   m.Ts,
	}
}

type PublicRestClassicSpotMarketCandlesResMiddle [][]string

type PublicRestClassicSpotMarketCandlesRes []Kline

func (m *PublicRestClassicSpotMarketCandlesResMiddle) ConvertToRes() PublicRestClassicSpotMarketCandlesRes {
	klines := make([]Kline, 0, len(*m))
	for _, row := range *m {
		klines = append(klines, Kline{
			OpenTime:   getRowValue(row, 0),
			OpenPrice:  getRowValue(row, 1),
			HighPrice:  getRowValue(row, 2),
			LowPrice:   getRowValue(row, 3),
			ClosePrice: getRowValue(row, 4),
			FillQty:    getRowValue(row, 5),
			FillAmount: getRowValue(row, 7),
		})
	}
	return PublicRestClassicSpotMarketCandlesRes(klines)
}

type PublicRestClassicSpotMarketHistoryCandlesResMiddle [][]string

type PublicRestClassicSpotMarketHistoryCandlesRes []Kline

func (m *PublicRestClassicSpotMarketHistoryCandlesResMiddle) ConvertToRes() PublicRestClassicSpotMarketHistoryCandlesRes {
	klines := make([]Kline, 0, len(*m))
	for _, row := range *m {
		klines = append(klines, Kline{
			OpenTime:   getRowValue(row, 0),
			OpenPrice:  getRowValue(row, 1),
			HighPrice:  getRowValue(row, 2),
			LowPrice:   getRowValue(row, 3),
			ClosePrice: getRowValue(row, 4),
			FillQty:    getRowValue(row, 5),
			FillAmount: getRowValue(row, 7),
		})
	}
	return PublicRestClassicSpotMarketHistoryCandlesRes(klines)
}

type PublicRestClassicSpotMarketFillsResRow struct {
	Symbol  string `json:"symbol"`  // String 交易对名称
	TradeID string `json:"tradeId"` // String 成交单 ID
	Side    string `json:"side"`    // String 成交方向 Buy/Sell
	Price   string `json:"price"`   // String 成交价格
	Size    string `json:"size"`    // String 成交数量
	Ts      string `json:"ts"`      // String 成交时间戳
}

type PublicRestClassicSpotMarketFillsRes []PublicRestClassicSpotMarketFillsResRow

type PublicRestClassicSpotMarketFillsHistoryRes []PublicRestClassicSpotMarketFillsResRow
