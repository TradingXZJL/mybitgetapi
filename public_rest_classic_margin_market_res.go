package mybitgetapi

type PublicRestClassicMarginMarketCurrenciesResRow struct {
	Symbol                    string `json:"symbol"`                    // String 交易对名称
	BaseCoin                  string `json:"baseCoin"`                  // String 基础币
	QuoteCoin                 string `json:"quoteCoin"`                 // String 计价币
	MaxCrossedLeverage        string `json:"maxCrossedLeverage"`        // String 全仓最大杠杆倍数
	MaxIsolatedLeverage       string `json:"maxIsolatedLeverage"`       // String 逐仓最大杠杆倍数
	WarningRiskRatio          string `json:"warningRiskRatio"`          // String 预警风险率
	LiquidationRiskRatio      string `json:"liquidationRiskRatio"`      // String 强平风险率
	MinTradeAmount            string `json:"minTradeAmount"`            // String 最小交易金额
	MaxTradeAmount            string `json:"maxTradeAmount"`            // String 最大交易金额
	TakerFeeRate              string `json:"takerFeeRate"`              // String 吃单费率
	MakerFeeRate              string `json:"makerFeeRate"`              // String 挂单费率
	PricePrecision            string `json:"pricePrecision"`            // String 价格精度
	QuantityPrecision         string `json:"quantityPrecision"`         // String 数量精度
	MinTradeUSDT              string `json:"minTradeUSDT"`              // String 最小USDT交易额
	IsBorrowable              bool   `json:"isBorrowable"`              // Bool 是否可借
	UserMinBorrow             string `json:"userMinBorrow"`             // String 最小借款
	Status                    string `json:"status"`                    // String 1可交易 2维护
	IsIsolatedBaseBorrowable  bool   `json:"isIsolatedBaseBorrowable"`  // Bool 逐仓左币是否可借
	IsIsolatedQuoteBorrowable bool   `json:"isIsolatedQuoteBorrowable"` // Bool 逐仓右币是否可借
	IsCrossBorrowable         bool   `json:"isCrossBorrowable"`         // Bool 全仓是否可借
}

type PublicRestClassicMarginMarketCurrenciesRes []PublicRestClassicMarginMarketCurrenciesResRow
