package mybitgetapi

type PrivateRestClassicMarginCrossTradeOrderIdRes struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
}

type PrivateRestClassicMarginCrossTradeBatchFailureItem struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}

type PrivateRestClassicMarginCrossTradeBatchRes struct {
	SuccessList []PrivateRestClassicMarginCrossTradeOrderIdRes       `json:"successList"`
	FailureList []PrivateRestClassicMarginCrossTradeBatchFailureItem `json:"failureList"`
}

type PrivateRestClassicMarginCrossTradePlaceOrderRes = PrivateRestClassicMarginCrossTradeOrderIdRes
type PrivateRestClassicMarginCrossTradeCancelOrderRes = PrivateRestClassicMarginCrossTradeOrderIdRes
type PrivateRestClassicMarginCrossTradeBatchPlaceOrderRes = PrivateRestClassicMarginCrossTradeBatchRes
type PrivateRestClassicMarginCrossTradeBatchCancelOrderRes = PrivateRestClassicMarginCrossTradeBatchRes

type PrivateRestClassicMarginCrossTradeOrderItem struct {
	OrderId          string `json:"orderId"`
	Symbol           string `json:"symbol"`
	OrderType        string `json:"orderType"`
	EnterPointSource string `json:"enterPointSource"`
	ClientOid        string `json:"clientOid"`
	LoanType         string `json:"loanType"`
	Price            string `json:"price"`
	Side             string `json:"side"`
	Status           string `json:"status"`
	BaseSize         string `json:"baseSize"`
	QuoteSize        string `json:"quoteSize"`
	PriceAvg         string `json:"priceAvg"`
	Size             string `json:"size"`
	Amount           string `json:"amount"`
	Force            string `json:"force"`
	CTime            string `json:"cTime"`
	UTime            string `json:"uTime"`
}

type PrivateRestClassicMarginCrossTradeOpenOrdersRes struct {
	OrderList []PrivateRestClassicMarginCrossTradeOrderItem `json:"orderList"`
	MaxID     string                                        `json:"maxId"`
	MinID     string                                        `json:"minId"`
}

type PrivateRestClassicMarginCrossTradeHistoryOrdersRes = PrivateRestClassicMarginCrossTradeOpenOrdersRes

type PrivateRestClassicMarginCrossTradeFeeDetail struct {
	Deduction         string `json:"deduction"`
	FeeCoin           string `json:"feeCoin"`
	TotalDeductionFee string `json:"totalDeductionFee"`
	TotalFee          string `json:"totalFee"`
}

type PrivateRestClassicMarginCrossTradeFillItem struct {
	OrderId    string                                      `json:"orderId"`
	TradeId    string                                      `json:"tradeId"`
	OrderType  string                                      `json:"orderType"`
	Side       string                                      `json:"side"`
	PriceAvg   string                                      `json:"priceAvg"`
	Size       string                                      `json:"size"`
	Amount     string                                      `json:"amount"`
	TradeScope string                                      `json:"tradeScope"`
	CTime      string                                      `json:"cTime"`
	UTime      string                                      `json:"uTime"`
	FeeDetail  PrivateRestClassicMarginCrossTradeFeeDetail `json:"feeDetail"`
}

type PrivateRestClassicMarginCrossTradeFillsRes struct {
	Fills []PrivateRestClassicMarginCrossTradeFillItem `json:"fills"`
	MaxID string                                       `json:"maxId"`
	MinID string                                       `json:"minId"`
}

type PrivateRestClassicMarginCrossTradeLiquidationOrderItem struct {
	Symbol    string `json:"symbol"`
	OrderType string `json:"orderType"`
	Side      string `json:"side"`
	PriceAvg  string `json:"priceAvg"`
	Price     string `json:"price"`
	FillSize  string `json:"fillSize"`
	Size      string `json:"size"`
	Amount    string `json:"amount"`
	OrderId   string `json:"orderId"`
	FromCoin  string `json:"fromCoin"`
	ToCoin    string `json:"toCoin"`
	FromSize  string `json:"fromSize"`
	ToSize    string `json:"toSize"`
	CTime     string `json:"cTime"`
	UTime     string `json:"uTime"`
}

type PrivateRestClassicMarginCrossTradeLiquidationOrderRes struct {
	ResultList []PrivateRestClassicMarginCrossTradeLiquidationOrderItem `json:"resultList"`
	IDLessThan string                                                   `json:"idLessThan"`
}
