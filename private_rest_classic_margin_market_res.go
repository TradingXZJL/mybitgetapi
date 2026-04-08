package mybitgetapi

// 杠杆利率（interest-rate-record）返回数据
type PrivateRestClassicMarginMarketInterestRateRecordRes struct {
	Coin               string `json:"coin"`               // String 币种名称
	DailyInterestRate  string `json:"dailyInterestRate"`  // String 日利率
	AnnualInterestRate string `json:"annualInterestRate"` // String 年利率
	UpdatedTime        string `json:"updatedTime"`        // String 数据更新时间
}
