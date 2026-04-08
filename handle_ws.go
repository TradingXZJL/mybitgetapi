package mybitgetapi

// ticker 频道数据项
type WsTickerData struct {
	Bid1Price    string `json:"bid1Price"`    // 买一价
	LowPrice24h  string `json:"lowPrice24h"`  // 24小时最低价
	Ask1Size     string `json:"ask1Size"`     // 卖一量
	Volume24h    string `json:"volume24h"`    // 24小时成交量
	Price24hPcnt string `json:"price24hPcnt"` // 24小时价格涨跌幅
	HighPrice24h string `json:"highPrice24h"` // 24小时最高价
	Turnover24h  string `json:"turnover24h"`  // 24小时成交额
	Bid1Size     string `json:"bid1Size"`     // 买一量
	Ask1Price    string `json:"ask1Price"`    // 卖一价
	OpenPrice24h string `json:"openPrice24h"` // 24小时前开盘价
	LastPrice    string `json:"lastPrice"`    // 最新价
	// 以下为合约参数
	IndexPrice        string `json:"indexPrice,omitempty"`        // 指数价格
	MarkPrice         string `json:"markPrice,omitempty"`         // 标记价格
	FundingRate       string `json:"fundingRate,omitempty"`       // 当前资金费率
	NextFundingTime   string `json:"nextFundingTime,omitempty"`   // 下次资金费率结算时间
	OpenInterest      string `json:"openInterest,omitempty"`      // 当前平台总持仓量
	DeliveryTime      string `json:"deliveryTime,omitempty"`      // 交割时间
	DeliveryStartTime string `json:"deliveryStartTime,omitempty"` // 交割开始时间
	DeliveryStatus    string `json:"deliveryStatus,omitempty"`    // 交割状态
}

// ticker 频道推送结构
type WsTicker struct {
	Arg    WsSubscribeArg `json:"arg"`
	Action string         `json:"action"` // snapshot / update
	Data   []WsTickerData `json:"data"`
	Ts     int64          `json:"ts"`
}

// 解析 ticker 频道消息
func handleWsTicker(data []byte) (*WsTicker, error) {
	ticker := WsTicker{}
	if err := json.Unmarshal(data, &ticker); err != nil {
		return nil, err
	}
	return &ticker, nil
}

// K线数据项
type WsCandleData struct {
	Start    string `json:"start"`    // 蜡烛开始时间（毫秒时间戳）
	Open     string `json:"open"`     // 开盘价
	Close    string `json:"close"`    // 收盘价
	High     string `json:"high"`     // 最高价
	Low      string `json:"low"`      // 最低价
	Volume   string `json:"volume"`   // 成交量（基础币）
	Turnover string `json:"turnover"` // 成交额（计价币）
}

// K线推送结构
type WsCandle struct {
	Arg    WsSubscribeArg `json:"arg"`
	Action string         `json:"action"` // snapshot / update
	Data   []WsCandleData `json:"data"`
	Ts     int64          `json:"ts"`
}

// 解析 K线频道消息
func handleWsCandle(data []byte) (*WsCandle, error) {
	candle := WsCandle{}
	if err := json.Unmarshal(data, &candle); err != nil {
		return nil, err
	}
	return &candle, nil
}

type WsBooksMiddleData struct {
	A        [][]string `json:"a"`        // 卖盘
	B        [][]string `json:"b"`        // 买盘
	Checksum int64      `json:"checksum"` // 校验和
	PSeq     int64      `json:"pseq"`     // 上一个序列号
	Seq      int64      `json:"seq"`      // 当前序列号
	Ts       string     `json:"ts"`       // 撮合时间戳
}

type WsBooksMiddle struct {
	Arg    WsSubscribeArg      `json:"arg"`
	Action string              `json:"action"` // snapshot / update
	Data   []WsBooksMiddleData `json:"data"`
	Ts     int64               `json:"ts"`
}

// 深度推送结构（对外暴露）
type WsBooks struct {
	WsSubscribeArg
	Action   string  `json:"action"`   // snapshot / update
	Asks     []Books `json:"asks"`     // 卖盘
	Bids     []Books `json:"bids"`     // 买盘
	Checksum int64   `json:"checksum"` // 校验和
	PSeq     int64   `json:"pseq"`     // 上一个序列号
	Seq      int64   `json:"seq"`      // 当前序列号
	Ts       string  `json:"ts"`       // 撮合时间戳
	PushTs   int64   `json:"pushTs"`   // 推送时间戳
}

// 解析深度频道消息
func handleWsBooks(data []byte) (*WsBooks, error) {
	middle := WsBooksMiddle{}
	if err := json.Unmarshal(data, &middle); err != nil {
		return nil, err
	}
	if len(middle.Data) == 0 {
		return nil, nil
	}

	row := middle.Data[0]
	asks := make([]Books, 0, len(row.A))
	bids := make([]Books, 0, len(row.B))

	for _, ask := range row.A {
		if len(ask) < 2 {
			continue
		}
		asks = append(asks, Books{
			Price:    ask[0],
			Quantity: ask[1],
		})
	}
	for _, bid := range row.B {
		if len(bid) < 2 {
			continue
		}
		bids = append(bids, Books{
			Price:    bid[0],
			Quantity: bid[1],
		})
	}

	return &WsBooks{
		WsSubscribeArg: middle.Arg,
		Action:         middle.Action,
		Asks:           asks,
		Bids:           bids,
		Checksum:       row.Checksum,
		PSeq:           row.PSeq,
		Seq:            row.Seq,
		Ts:             row.Ts,
		PushTs:         middle.Ts,
	}, nil
}

// 平台成交数据项
type WsPublicTradeData struct {
	TradeId        string `json:"i"`     // 成交ID
	TradeLinkId    string `json:"L"`     // 成交关联ID
	Price          string `json:"p"`     // 成交价格
	Volume         string `json:"v"`     // 成交数量
	Side           string `json:"S"`     // buy/sell
	TradeTimestamp string `json:"T"`     // 成交时间戳
	IsRPI          string `json:"isRPI"` // yes/no
}

// 平台成交推送结构
type WsPublicTrade struct {
	Arg    WsSubscribeArg      `json:"arg"`
	Action string              `json:"action"` // snapshot / update
	Data   []WsPublicTradeData `json:"data"`
	Ts     int64               `json:"ts"`
}

// 解析平台成交频道消息
func handleWsPublicTrade(data []byte) (*WsPublicTrade, error) {
	trade := WsPublicTrade{}
	if err := json.Unmarshal(data, &trade); err != nil {
		return nil, err
	}
	return &trade, nil
}

// fill 频道（现货/合约通用）手续费明细
type WsFillFeeDetail struct {
	FeeCoin           string `json:"feeCoin"`
	Deduction         string `json:"deduction"`
	TotalDeductionFee string `json:"totalDeductionFee"`
	TotalFee          string `json:"totalFee"`
}

// fill 频道（现货/合约通用）单条推送（字段为并集，按 instType 决定哪些会出现）
type WsFillData struct {
	// common
	OrderId    string            `json:"orderId,omitempty"`
	Symbol     string            `json:"symbol,omitempty"`
	OrderType  string            `json:"orderType,omitempty"`
	Side       string            `json:"side,omitempty"`
	TradeScope string            `json:"tradeScope,omitempty"`
	FeeDetail  []WsFillFeeDetail `json:"feeDetail,omitempty"`
	CTime      string            `json:"cTime,omitempty"`
	UTime      string            `json:"uTime,omitempty"`

	// spot
	TradeId  string `json:"tradeId,omitempty"`
	PriceAvg string `json:"priceAvg,omitempty"`
	Size     string `json:"size,omitempty"`
	Amount   string `json:"amount,omitempty"`

	// futures
	ClientOid   string `json:"clientOid,omitempty"`
	PosMode     string `json:"posMode,omitempty"`
	Price       string `json:"price,omitempty"`
	BaseVolume  string `json:"baseVolume,omitempty"`
	QuoteVolume string `json:"quoteVolume,omitempty"`
	Profit      string `json:"profit,omitempty"`
	TradeSide   string `json:"tradeSide,omitempty"`
}

// fill 频道（现货/合约通用）推送
type WsFill struct {
	Arg    WsSubscribeArg `json:"arg"`
	Action string         `json:"action"`
	Data   []WsFillData   `json:"data"`
	Ts     int64          `json:"ts"`
}

// 兼容旧类型名
type WsSpotFill = WsFill
type WsFuturesFill = WsFill

func handleWsFill(data []byte) (*WsFill, error) {
	fill := WsFill{}
	if err := json.Unmarshal(data, &fill); err != nil {
		return nil, err
	}
	return &fill, nil
}

// 现货私有订单（orders 频道）手续费明细
type WsOrderFeeDetail struct {
	FeeCoin string `json:"feeCoin"`
	Fee     string `json:"fee"`
}

// 现货私有 orders 频道单条推送
type WsOrderData struct {
	InstId           string             `json:"instId"`
	OrderId          string             `json:"orderId"`
	ClientOid        string             `json:"clientOid"`
	Price            string             `json:"price,omitempty"`
	Size             string             `json:"size,omitempty"`
	NewSize          string             `json:"newSize,omitempty"`
	Notional         string             `json:"notional,omitempty"`
	OrderType        string             `json:"orderType,omitempty"` // 文档示例字段
	OrdType          string             `json:"ordType,omitempty"`   // 文档参数表字段（兼容）
	Force            string             `json:"force,omitempty"`
	Side             string             `json:"side,omitempty"`
	FillPrice        string             `json:"fillPrice,omitempty"`
	TradeId          string             `json:"tradeId,omitempty"`
	BaseVolume       string             `json:"baseVolume,omitempty"`
	QuoteVolume      string             `json:"quoteVolume,omitempty"`
	FillTime         string             `json:"fillTime,omitempty"`
	FillFee          string             `json:"fillFee,omitempty"`
	FillFeeCoin      string             `json:"fillFeeCoin,omitempty"`
	TradeScope       string             `json:"tradeScope,omitempty"`
	AccBaseVolume    string             `json:"accBaseVolume,omitempty"`
	FillNotionalUsd  string             `json:"fillNotionalUsd,omitempty"`
	PriceAvg         string             `json:"priceAvg,omitempty"`
	Status           string             `json:"status,omitempty"`
	CancelReason     string             `json:"cancelReason,omitempty"`
	CTime            string             `json:"cTime,omitempty"`
	UTime            string             `json:"uTime,omitempty"`
	StpMode          string             `json:"stpMode,omitempty"`
	FeeDetail        []WsOrderFeeDetail `json:"feeDetail,omitempty"`
	EnterPointSource string             `json:"enterPointSource,omitempty"`

	// futures only
	Leverage                      string `json:"leverage,omitempty"`
	PosMode                       string `json:"posMode,omitempty"`
	MarginMode                    string `json:"marginMode,omitempty"`
	TradeSide                     string `json:"tradeSide,omitempty"`
	PosSide                       string `json:"posSide,omitempty"`
	ReduceOnly                    string `json:"reduceOnly,omitempty"`
	MarginCoin                    string `json:"marginCoin,omitempty"`
	NotionalUsd                   string `json:"notionalUsd,omitempty"`
	TotalProfits                  string `json:"totalProfits,omitempty"`
	Pnl                           string `json:"pnl,omitempty"`
	PresetStopLossType            string `json:"presetStopLossType,omitempty"`
	PresetStopSurplusType         string `json:"presetStopSurplusType,omitempty"`
	PresetStopLossPrice           string `json:"presetStopLossPrice,omitempty"`
	PresetStopSurplusPrice        string `json:"presetStopSurplusPrice,omitempty"`
	PresetStopLossExecutePrice    string `json:"presetStopLossExecutePrice,omitempty"`
	PresetStopSurplusExecutePrice string `json:"presetStopSurplusExecutePrice,omitempty"`
}

// 现货私有 orders 频道推送
type WsOrder struct {
	Arg    WsSubscribeArg `json:"arg"`
	Action string         `json:"action"`
	Data   []WsOrderData  `json:"data"`
	Ts     int64          `json:"ts"`
}

func handleWsOrder(data []byte) (*WsOrder, error) {
	order := WsOrder{}
	if err := json.Unmarshal(data, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

// 私有账户（account 频道）单条推送（现货/合约通用，字段为并集）
type WsAccountData struct {
	// spot
	Coin           string `json:"coin,omitempty"`
	Locked         string `json:"locked,omitempty"`
	LimitAvailable string `json:"limitAvailable,omitempty"`
	UTime          string `json:"uTime,omitempty"`

	// contract
	MarginCoin          string `json:"marginCoin,omitempty"`
	MaxOpenPosAvailable string `json:"maxOpenPosAvailable,omitempty"`
	MaxTransferOut      string `json:"maxTransferOut,omitempty"`
	Equity              string `json:"equity,omitempty"`
	UsdtEquity          string `json:"usdtEquity,omitempty"`
	CrossedRiskRate     string `json:"crossedRiskRate,omitempty"`
	UnrealizedPL        string `json:"unrealizedPL,omitempty"`
	UnionTotalMargin    string `json:"unionTotalMargin,omitempty"`
	UnionAvailable      string `json:"unionAvailable,omitempty"`
	UnionMm             string `json:"unionMm,omitempty"`
	AssetsMode          string `json:"assetsMode,omitempty"`

	// common
	Available string `json:"available,omitempty"`
	Frozen    string `json:"frozen,omitempty"`
}

// 私有 account 频道推送（现货/合约通用）
type WsAccount struct {
	Arg    WsSubscribeArg  `json:"arg"`
	Action string          `json:"action"`
	Data   []WsAccountData `json:"data"`
	Ts     int64           `json:"ts"`
}

func handleWsAccount(data []byte) (*WsAccount, error) {
	account := WsAccount{}
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, err
	}
	return &account, nil
}

// 合约私有权益（equity 频道）推送
type WsEquityData struct {
	UsdtEquity       string `json:"usdtEquity"`
	BtcEquity        string `json:"btcEquity"`
	UsdtUnrealized   string `json:"usdtUnrealized"`
	UnionTotalMargin string `json:"unionTotalMargin"`
	UnionAvailable   string `json:"unionAvailable"`
	UnionMm          string `json:"unionMm"`
}

type WsEquity struct {
	Arg    WsSubscribeArg `json:"arg"`
	Action string         `json:"action"`
	Data   []WsEquityData `json:"data"`
	Ts     int64          `json:"ts"`
}

func handleWsEquity(data []byte) (*WsEquity, error) {
	equity := WsEquity{}
	if err := json.Unmarshal(data, &equity); err != nil {
		return nil, err
	}
	return &equity, nil
}

// 合约私有持仓（positions 频道）单条推送
type WsPositionsData struct {
	PosId              string `json:"posId"`
	InstId             string `json:"instId"`
	MarginCoin         string `json:"marginCoin"`
	MarginSize         string `json:"marginSize"`
	MarginMode         string `json:"marginMode"`
	HoldSide           string `json:"holdSide"`
	PosMode            string `json:"posMode"`
	Total              string `json:"total"`
	Available          string `json:"available"`
	Frozen             string `json:"frozen"`
	OpenPriceAvg       string `json:"openPriceAvg"`
	Leverage           int64  `json:"leverage"`
	AchievedProfits    string `json:"achievedProfits"`
	UnrealizedPL       string `json:"unrealizedPL"`
	UnrealizedPLR      string `json:"unrealizedPLR"`
	LiquidationPrice   string `json:"liquidationPrice"`
	KeepMarginRate     string `json:"keepMarginRate"`
	IsolatedMarginRate string `json:"isolatedMarginRate,omitempty"`
	MarginRate         string `json:"marginRate"`
	CTime              string `json:"cTime"`
	BreakEvenPrice     string `json:"breakEvenPrice"`
	TotalFee           string `json:"totalFee"`
	DeductedFee        string `json:"deductedFee"`
	MarkPrice          string `json:"markPrice"`
	AssetMode          string `json:"assetMode"`
	UTime              string `json:"uTime"`
	AutoMargin         string `json:"autoMargin"`
}

// 合约私有 positions 频道推送
type WsPositions struct {
	Arg    WsSubscribeArg    `json:"arg"`
	Action string            `json:"action"`
	Data   []WsPositionsData `json:"data"`
	Ts     int64             `json:"ts"`
}

func handleWsPositions(data []byte) (*WsPositions, error) {
	pos := WsPositions{}
	if err := json.Unmarshal(data, &pos); err != nil {
		return nil, err
	}
	return &pos, nil
}

// --- Margin: cross / isolated ---

type WsMarginAccountAssetData struct {
	UTime     string `json:"uTime"`
	Id        string `json:"id"`
	Coin      string `json:"coin"`
	Symbol    string `json:"symbol,omitempty"` // isolated 才有
	Available string `json:"available"`
	Borrow    string `json:"borrow"`
	Frozen    string `json:"frozen"`
	Interest  string `json:"interest"`
	Coupon    string `json:"coupon"`
}

type WsMarginCrossAccountAssets struct {
	Arg    WsSubscribeArg             `json:"arg"`
	Action string                     `json:"action"`
	Data   []WsMarginAccountAssetData `json:"data"`
	Ts     int64                      `json:"ts"`
}

func handleWsMarginCrossAccountAssets(data []byte) (*WsMarginCrossAccountAssets, error) {
	res := WsMarginCrossAccountAssets{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type WsMarginIsolatedAccountAssets struct {
	Arg    WsSubscribeArg             `json:"arg"`
	Action string                     `json:"action"`
	Data   []WsMarginAccountAssetData `json:"data"`
	Ts     int64                      `json:"ts"`
}

func handleWsMarginIsolatedAccountAssets(data []byte) (*WsMarginIsolatedAccountAssets, error) {
	res := WsMarginIsolatedAccountAssets{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type WsMarginOrderFeeDetail struct {
	FeeCoin           string `json:"feeCoin"`
	Deduction         string `json:"deduction"`
	TotalDeductionFee string `json:"totalDeductionFee"`
	TotalFee          string `json:"totalFee"`
}

type WsMarginOrderData struct {
	Force            string                   `json:"force,omitempty"`
	OrderType        string                   `json:"orderType,omitempty"`
	Price            string                   `json:"price,omitempty"`
	QuoteSize        string                   `json:"quoteSize,omitempty"`
	Side             string                   `json:"side,omitempty"`
	FeeDetail        []WsMarginOrderFeeDetail `json:"feeDetail,omitempty"`
	EnterPointSource string                   `json:"enterPointSource,omitempty"`
	Status           string                   `json:"status,omitempty"`
	BaseSize         string                   `json:"baseSize,omitempty"`
	CTime            string                   `json:"cTime,omitempty"`
	UTime            string                   `json:"uTime,omitempty"`
	StpMode          string                   `json:"stpMode,omitempty"`
	ClientOid        string                   `json:"clientOid,omitempty"`
	FillPrice        string                   `json:"fillPrice,omitempty"`
	BaseVolume       string                   `json:"baseVolume,omitempty"`
	FillTotalAmount  string                   `json:"fillTotalAmount,omitempty"`
	LoanType         string                   `json:"loanType,omitempty"`
	OrderId          string                   `json:"orderId,omitempty"`
}

type WsMarginCrossOrders struct {
	Arg    WsSubscribeArg      `json:"arg"`
	Action string              `json:"action"`
	Data   []WsMarginOrderData `json:"data"`
	Ts     int64               `json:"ts"`
}

func handleWsMarginCrossOrders(data []byte) (*WsMarginCrossOrders, error) {
	res := WsMarginCrossOrders{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type WsMarginIsolatedOrders struct {
	Arg    WsSubscribeArg      `json:"arg"`
	Action string              `json:"action"`
	Data   []WsMarginOrderData `json:"data"`
	Ts     int64               `json:"ts"`
}

func handleWsMarginIsolatedOrders(data []byte) (*WsMarginIsolatedOrders, error) {
	res := WsMarginIsolatedOrders{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
