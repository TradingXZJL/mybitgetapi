package mybitgetapi

type PrivateRestClassicMarginCrossAccountAssetItem struct {
	Coin        string `json:"coin"`        // 币种名称
	TotalAmount string `json:"totalAmount"` // 总金额
	Available   string `json:"available"`   // 可用金额
	Frozen      string `json:"frozen"`      // 冻结资产
	Borrow      string `json:"borrow"`      // 借款
	Interest    string `json:"interest"`    // 利息（还款时先息后本）
	Net         string `json:"net"`         // 净资产 = available + frozen - borrow - interest
	CTime       string `json:"cTime"`       // 创建时间
	UTime       string `json:"uTime"`       // 更新时间
	Coupon      string `json:"coupon"`      // 体验金
}

type PrivateRestClassicMarginCrossAccountAssetsRes []PrivateRestClassicMarginCrossAccountAssetItem

type PrivateRestClassicMarginCrossAccountBorrowRes struct {
	LoanID       string `json:"loanId"`       // 借款 ID
	Coin         string `json:"coin"`         // 借出币种
	BorrowAmount string `json:"borrowAmount"` // 借出数量
}

type PrivateRestClassicMarginCrossAccountRepayRes struct {
	Coin             string `json:"coin"`             // 币种
	RepayID          string `json:"repayId"`          // 还款 ID
	RemainDebtAmount string `json:"remainDebtAmount"` // 剩余欠款
	RepayAmount      string `json:"repayAmount"`      // 还款金额
}

type PrivateRestClassicMarginCrossAccountRiskRateRes struct {
	RiskRateRatio string `json:"riskRateRatio"` // 风险率（全仓总资产/全仓总负债）
}

type PrivateRestClassicMarginCrossAccountMaxBorrowableAmountRes struct {
	Coin                string `json:"coin"`                // 币种
	MaxBorrowableAmount string `json:"maxBorrowableAmount"` // 最大可借数量
}

type PrivateRestClassicMarginCrossAccountMaxTransferOutAmountRes struct {
	Coin                 string `json:"coin"`                 // 币种
	MaxTransferOutAmount string `json:"maxTransferOutAmount"` // 最大可转出数量
}

type PrivateRestClassicMarginCrossAccountInterestRateVipItem struct {
	Level              string `json:"level"`              // VIP 等级
	Limit              string `json:"limit"`              // VIP 限额
	DailyInterestRate  string `json:"dailyInterestRate"`  // VIP 日利率
	AnnualInterestRate string `json:"annualInterestRate"` // VIP 年利率
	DiscountRate       string `json:"discountRate"`       // VIP 折扣（1 表示无折扣）
}

type PrivateRestClassicMarginCrossAccountInterestRateAndLimitItem struct {
	Coin                string                                                    `json:"coin"`                // 币种，如 BTC、ETH
	Leverage            string                                                    `json:"leverage"`            // 杠杆倍数（如 3/5/10）
	Transferable        bool                                                      `json:"transferable"`        // 是否可转入
	Borrowable          bool                                                      `json:"borrowable"`          // 是否可借
	DailyInterestRate   string                                                    `json:"dailyInterestRate"`   // 非 VIP 日利率
	AnnualInterestRate  string                                                    `json:"annualInterestRate"`  // 非 VIP 年利率
	MaxBorrowableAmount string                                                    `json:"maxBorrowableAmount"` // 最大可借
	VipList             []PrivateRestClassicMarginCrossAccountInterestRateVipItem `json:"vipList"`             // VIP 等级利率配置
}

type PrivateRestClassicMarginCrossAccountInterestRateAndLimitRes []PrivateRestClassicMarginCrossAccountInterestRateAndLimitItem

type PrivateRestClassicMarginCrossAccountFlashRepayRes struct {
	RepayID string  `json:"repayId"` // 还款 ID
	Coin    *string `json:"coin"`    // 还款币种；全币种还款时可能为空
}

type PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusItem struct {
	RepayID string `json:"repayId"` // 还款 ID
	Status  string `json:"status"`  // 还款结果状态
}

type PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusRes []PrivateRestClassicMarginCrossAccountQueryFlashRepayStatusItem

type PrivateRestClassicMarginCrossRecordBorrowHistoryItem struct {
	LoanID       string `json:"loanId"`       // 借款 ID
	Coin         string `json:"coin"`         // 借款币种
	BorrowAmount string `json:"borrowAmount"` // 借款数量
	BorrowType   string `json:"borrowType"`   // 借款类型（auto_loan/manual_loan）
	CTime        string `json:"cTime"`        // 创建时间（毫秒时间戳）
	UTime        string `json:"uTime"`        // 更新时间（毫秒时间戳）
}

type PrivateRestClassicMarginCrossRecordBorrowHistoryRes struct {
	ResultList []PrivateRestClassicMarginCrossRecordBorrowHistoryItem `json:"resultList"` // 借款记录列表
	MaxID      string                                                 `json:"maxId"`      // 最大 ID
	MinID      string                                                 `json:"minId"`      // 最小 ID
}

type PrivateRestClassicMarginCrossRecordRepayHistoryItem struct {
	RepayID        string `json:"repayId"`        // 还款 ID
	Coin           string `json:"coin"`           // 还款币种
	RepayAmount    string `json:"repayAmount"`    // 还款总额
	RepayType      string `json:"repayType"`      // 还款类型（auto/manual/liq/force）
	RepayInterest  string `json:"repayInterest"`  // 还款利息
	RepayPrincipal string `json:"repayPrincipal"` // 还款本金
	CTime          string `json:"cTime"`          // 创建时间（毫秒时间戳）
	UTime          string `json:"uTime"`          // 更新时间（毫秒时间戳）
}

type PrivateRestClassicMarginCrossRecordRepayHistoryRes struct {
	ResultList []PrivateRestClassicMarginCrossRecordRepayHistoryItem `json:"resultList"` // 还款记录列表
	MaxID      string                                                `json:"maxId"`      // 最大 ID
	MinID      string                                                `json:"minId"`      // 最小 ID
}

type PrivateRestClassicMarginCrossRecordInterestHistoryItem struct {
	InterestID        string `json:"interestId"`        // 利息记录 ID
	LoanCoin          string `json:"loanCoin"`          // 借贷币种
	InterestCoin      string `json:"interestCoin"`      // 利息币种
	DailyInterestRate string `json:"dailyInterestRate"` // 日利率
	InterestAmount    string `json:"interestAmount"`    // 利息金额
	InterstType       string `json:"interstType"`       // 利息类型（first/scheduled）
	UTime             string `json:"uTime"`             // 更新时间（毫秒时间戳）
	CTime             string `json:"cTime"`             // 创建时间（毫秒时间戳）
}

type PrivateRestClassicMarginCrossRecordInterestHistoryRes struct {
	MinID      string                                                   `json:"minId"`      // 最小 ID
	MaxID      string                                                   `json:"maxId"`      // 最大 ID
	ResultList []PrivateRestClassicMarginCrossRecordInterestHistoryItem `json:"resultList"` // 利息记录列表
}

type PrivateRestClassicMarginCrossRecordLiquidationHistoryItem struct {
	LiqID        string `json:"liqId"`        // 强平 ID
	LiqStartTime string `json:"liqStartTime"` // 强平开始时间
	LiqEndTime   string `json:"liqEndTime"`   // 强平结束时间
	LiqRiskRatio string `json:"liqRiskRatio"` // 强平时风险率
	TotalAssets  string `json:"totalAssets"`  // 强平时总资产（USDT）
	TotalDebt    string `json:"totalDebt"`    // 强平时总负债（USDT）
	LiqFee       string `json:"liqFee"`       // 强平手续费
	CTime        string `json:"cTime"`        // 创建时间（毫秒时间戳）
	UTime        string `json:"uTime"`        // 更新时间（毫秒时间戳）
}

type PrivateRestClassicMarginCrossRecordLiquidationHistoryRes struct {
	ResultList []PrivateRestClassicMarginCrossRecordLiquidationHistoryItem `json:"resultList"` // 强平记录列表
	MaxID      string                                                      `json:"maxId"`      // 最大 ID
	MinID      string                                                      `json:"minId"`      // 最小 ID
}

type PrivateRestClassicMarginCrossRecordFinancialRecordItem struct {
	Coin       string `json:"coin"`       // 币种
	MarginID   string `json:"marginId"`   // 流水 ID
	Amount     string `json:"amount"`     // 流水金额
	Balance    string `json:"balance"`    // 账户余额
	Fee        string `json:"fee"`        // 手续费明细
	MarginType string `json:"marginType"` // 流水类型（transfer_in/borrow/repay/...）
	CTime      string `json:"cTime"`      // 创建时间（毫秒时间戳）
	UTime      string `json:"uTime"`      // 更新时间（毫秒时间戳）
}

type PrivateRestClassicMarginCrossRecordFinancialRecordsRes struct {
	ResultList []PrivateRestClassicMarginCrossRecordFinancialRecordItem `json:"resultList"` // 财务流水记录列表
	MaxID      string                                                   `json:"maxId"`      // 最大 ID
	MinID      string                                                   `json:"minId"`      // 最小 ID
}

type PrivateRestClassicMarginIsolatedTradeFeeDetail struct {
	Deduction         string `json:"deduction"`         // 是否折扣
	FeeCoin           string `json:"feeCoin"`           // 手续费币种
	TotalDeductionFee string `json:"totalDeductionFee"` // 总计折扣手续费
	TotalFee          string `json:"totalFee"`          // 总计手续费
}

type PrivateRestClassicMarginIsolatedTradeFillItem struct {
	OrderId    string                                         `json:"orderId"`    // 订单号
	TradeId    string                                         `json:"tradeId"`    // 成交明细 ID
	OrderType  string                                         `json:"orderType"`  // 订单类型：limit/market
	Side       string                                         `json:"side"`       // 交易方向：buy/sell
	PriceAvg   string                                         `json:"priceAvg"`   // 成交价格
	Size       string                                         `json:"size"`       // 成交数量
	Amount     string                                         `json:"amount"`     // 成交金额
	TradeScope string                                         `json:"tradeScope"` // 交易者标识：taker/maker
	CTime      string                                         `json:"cTime"`      // 创建时间
	UTime      string                                         `json:"uTime"`      // 更新时间
	FeeDetail  PrivateRestClassicMarginIsolatedTradeFeeDetail `json:"feeDetail"`  // 手续费明细
}

type PrivateRestClassicMarginIsolatedAccountAssetItem struct {
	Symbol      string `json:"symbol"`      // 交易对名称
	Coin        string `json:"coin"`        // 币种名称
	TotalAmount string `json:"totalAmount"` // 总金额
	Available   string `json:"available"`   // 可用金额
	Frozen      string `json:"frozen"`      // 冻结资产
	Borrow      string `json:"borrow"`      // 借款
	Interest    string `json:"interest"`    // 利息（还款时先息后本）
	Net         string `json:"net"`         // 净资产 = available + frozen - borrow - interest
	Coupon      string `json:"coupon"`      // 体验金
	CTime       string `json:"cTime"`       // 创建时间
	UTime       string `json:"uTime"`       // 更新时间
}

type PrivateRestClassicMarginIsolatedAccountAssetsRes []PrivateRestClassicMarginIsolatedAccountAssetItem

type PrivateRestClassicMarginIsolatedAccountBorrowRes struct {
	LoanID       string `json:"loanId"`       // 借款 ID
	Symbol       string `json:"symbol"`       // 借出交易对
	Coin         string `json:"coin"`         // 借出币种
	BorrowAmount string `json:"borrowAmount"` // 借出数量
}

type PrivateRestClassicMarginIsolatedAccountRepayRes struct {
	RemainDebtAmount string `json:"remainDebtAmount"` // 剩余欠款
	Symbol           string `json:"symbol"`           // 交易对
	RepayID          string `json:"repayId"`          // 还款 ID
	Coin             string `json:"coin"`             // 币种
	RepayAmount      string `json:"repayAmount"`      // 还款金额
}

type PrivateRestClassicMarginIsolatedAccountRiskRateItem struct {
	Symbol        string `json:"symbol"`        // 交易对
	RiskRateRatio string `json:"riskRateRatio"` // 风险率（逐仓总资产/逐仓总负债）
}

type PrivateRestClassicMarginIsolatedAccountRiskRateRes []PrivateRestClassicMarginIsolatedAccountRiskRateItem

type PrivateRestClassicMarginIsolatedAccountTierDataItem struct {
	Tier                     string `json:"tier"`                     // 档位
	Symbol                   string `json:"symbol"`                   // 交易对
	Leverage                 string `json:"leverage"`                 // 有效杠杆倍数
	BaseCoin                 string `json:"baseCoin"`                 // 基础币
	QuoteCoin                string `json:"quoteCoin"`                // 计价币
	BaseMaxBorrowableAmount  string `json:"baseMaxBorrowableAmount"`  // 基础币最大可借
	QuoteMaxBorrowableAmount string `json:"quoteMaxBorrowableAmount"` // 计价币最大可借
	MaintainMarginRate       string `json:"maintainMarginRate"`       // 维持保证金率
	InitRate                 string `json:"initRate"`                 // 初始保证金率
}

type PrivateRestClassicMarginIsolatedAccountTierDataRes []PrivateRestClassicMarginIsolatedAccountTierDataItem

type PrivateRestClassicMarginIsolatedAccountInterestRateVipItem struct {
	Level                string `json:"level"`                // VIP 等级
	DailyInterestRate    string `json:"dailyInterestRate"`    // VIP 日利率
	Limit                string `json:"limit"`                // VIP 限额
	AnnuallyInterestRate string `json:"annuallyInterestRate"` // VIP 年利率
	DiscountRate         string `json:"discountRate"`         // VIP 折扣（1 表示无折扣）
}

type PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitItem struct {
	Symbol                    string                                                       `json:"symbol"`                    // 交易对
	Leverage                  string                                                       `json:"leverage"`                  // 杠杆倍数
	BaseCoin                  string                                                       `json:"baseCoin"`                  // 基础币
	BaseTransferable          bool                                                         `json:"baseTransferable"`          // 基础币是否可转入
	BaseBorrowable            bool                                                         `json:"baseBorrowable"`            // 基础币是否可借
	BaseDailyInterestRate     string                                                       `json:"baseDailyInterestRate"`     // 基础币非 VIP 日利率
	BaseAnnuallyInterestRate  string                                                       `json:"baseAnnuallyInterestRate"`  // 基础币非 VIP 年利率
	BaseMaxBorrowableAmount   string                                                       `json:"baseMaxBorrowableAmount"`   // 基础币最大可借
	BaseVipList               []PrivateRestClassicMarginIsolatedAccountInterestRateVipItem `json:"baseVipList"`               // 基础币 VIP 利率配置
	QuoteCoin                 string                                                       `json:"quoteCoin"`                 // 计价币
	QuoteTransferable         bool                                                         `json:"quoteTransferable"`         // 计价币是否可转入
	QuoteBorrowable           bool                                                         `json:"quoteBorrowable"`           // 计价币是否可借
	QuoteDailyInterestRate    string                                                       `json:"quoteDailyInterestRate"`    // 计价币非 VIP 日利率
	QuoteAnnuallyInterestRate string                                                       `json:"quoteAnnuallyInterestRate"` // 计价币非 VIP 年利率
	QuoteMaxBorrowableAmount  string                                                       `json:"quoteMaxBorrowableAmount"`  // 计价币最大可借
	QuoteList                 []PrivateRestClassicMarginIsolatedAccountInterestRateVipItem `json:"quoteList"`                 // 计价币 VIP 利率配置
}

type PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitRes []PrivateRestClassicMarginIsolatedAccountInterestRateAndLimitItem

type PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmountRes struct {
	Symbol                   string `json:"symbol"`                   // 交易对
	BaseCoin                 string `json:"baseCoin"`                 // 左币
	BaseCoinMaxBorrowAmount  string `json:"baseCoinMaxBorrowAmount"`  // 左币最大可借数量
	QuoteCoin                string `json:"quoteCoin"`                // 右币
	QuoteCoinMaxBorrowAmount string `json:"quoteCoinMaxBorrowAmount"` // 右币最大可借数量
}

type PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmountRes struct {
	BaseCoin                      string `json:"baseCoin"`                      // 基础币币种
	Symbol                        string `json:"symbol"`                        // 交易对
	BaseCoinMaxTransferOutAmount  string `json:"baseCoinMaxTransferOutAmount"`  // 最大基础币可转出数量
	QuoteCoin                     string `json:"quoteCoin"`                     // 计价币币种
	QuoteCoinMaxTransferOutAmount string `json:"quoteCoinMaxTransferOutAmount"` // 最大计价币可转出数量
}

type PrivateRestClassicMarginIsolatedAccountFlashRepayItem struct {
	RepayID string `json:"repayId"` // 还款 ID
	Symbol  string `json:"symbol"`  // 杠杆交易对
	Result  string `json:"result"`  // 请求结果：success/failure
}

type PrivateRestClassicMarginIsolatedAccountFlashRepayRes []PrivateRestClassicMarginIsolatedAccountFlashRepayItem

type PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusItem struct {
	RepayID string `json:"repayId"` // 还款 ID
	Status  string `json:"status"`  // 还款结果状态
}

type PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusRes []PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatusItem

type PrivateRestClassicMarginCrossAccountTierDataResRow struct {
	Tier                string `json:"tier"`                // 档位
	Leverage            string `json:"leverage"`            // 有效杠杆倍数，全局默认3倍
	Coin                string `json:"coin"`                // 币种
	MaxBorrowableAmount string `json:"maxBorrowableAmount"` // 最大可借
	MaintainMarginRate  string `json:"maintainMarginRate"`  // 维持保证金率
}
type PrivateRestClassicMarginCrossAccountTierDataRes []PrivateRestClassicMarginCrossAccountTierDataResRow
