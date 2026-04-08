package mybitgetapi

type PrivateRestClassicFuturesAccountAssetCommon struct {
	MarginCoin           string  `json:"marginCoin"`
	Locked               string  `json:"locked"`
	Available            string  `json:"available"`
	CrossedMaxAvailable  string  `json:"crossedMaxAvailable"`
	IsolatedMaxAvailable string  `json:"isolatedMaxAvailable"`
	MaxTransferOut       string  `json:"maxTransferOut"`
	AccountEquity        string  `json:"accountEquity"`
	UsdtEquity           string  `json:"usdtEquity"`
	BtcEquity            string  `json:"btcEquity"`
	CrossedRiskRate      string  `json:"crossedRiskRate"`
	UnrealizedPL         string  `json:"unrealizedPL"`
	Coupon               *string `json:"coupon"`
	CrossedUnrealizedPL  *string `json:"crossedUnrealizedPL"`
	IsolatedUnrealizedPL *string `json:"isolatedUnrealizedPL"`
	Grant                *string `json:"grant"`
	AssetMode            *string `json:"assetMode"`
	IsolatedMargin       *string `json:"isolatedMargin"`
	CrossedMargin        *string `json:"crossedMargin"`
}

type PrivateRestClassicFuturesAccountSingleAccountRes struct {
	PrivateRestClassicFuturesAccountAssetCommon
	CrossedMarginLeverage int    `json:"crossedMarginLeverage"`
	IsolatedLongLever     int    `json:"isolatedLongLever"`
	IsolatedShortLever    int    `json:"isolatedShortLever"`
	MarginMode            string `json:"marginMode"`
	PosMode               string `json:"posMode"`
}

type PrivateRestClassicFuturesAccountUnionAssetItem struct {
	Coin      string `json:"coin"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
}

type PrivateRestClassicFuturesAccountListRow struct {
	PrivateRestClassicFuturesAccountAssetCommon
	UnionTotalMargin string                                           `json:"unionTotalMargin"`
	UnionTotalMagin  string                                           `json:"unionTotalMagin"`
	UnionAvailable   string                                           `json:"unionAvailable"`
	UnionMm          string                                           `json:"unionMm"`
	AssetList        []PrivateRestClassicFuturesAccountUnionAssetItem `json:"assetList"`
}

type PrivateRestClassicFuturesAccountListRes []PrivateRestClassicFuturesAccountListRow

type PrivateRestClassicFuturesAccountSubAccountAssetsRow struct {
	UserId    int64                                         `json:"userId"`
	AssetList []PrivateRestClassicFuturesAccountAssetCommon `json:"assetList"`
}

type PrivateRestClassicFuturesAccountSubAccountAssetsRes []PrivateRestClassicFuturesAccountSubAccountAssetsRow

type PrivateRestClassicFuturesAccountInterestHistoryItem struct {
	Coin              string `json:"coin"`
	Liability         string `json:"liability"`
	InterestFreeLimit string `json:"interestFreeLimit"`
	InterestLimit     string `json:"interestLimit"`
	HourInterestRate  string `json:"hourInterestRate"`
	Interest          string `json:"interest"`
	CTime             string `json:"cTime"`
}

type PrivateRestClassicFuturesAccountInterestHistoryRes struct {
	NextSettleTime string                                                `json:"nextSettleTime"`
	BorrowAmount   string                                                `json:"borrowAmount"`
	BorrowLimit    string                                                `json:"borrowLimit"`
	InterestList   []PrivateRestClassicFuturesAccountInterestHistoryItem `json:"interestList"`
	EndId          string                                                `json:"endId"`
}

type PrivateRestClassicFuturesAccountMaxOpenRes struct {
	MaxOpen string `json:"maxOpen"`
}

type PrivateRestClassicFuturesAccountLiquidationPriceRes struct {
	LiqPrice string `json:"liqPrice"`
}

type PrivateRestClassicFuturesAccountOpenCountRes struct {
	Size string `json:"size"`
}

type PrivateRestClassicFuturesAccountSetLeverageRes struct {
	Symbol              string `json:"symbol"`
	MarginCoin          string `json:"marginCoin"`
	LongLeverage        string `json:"longLeverage"`
	ShortLeverage       string `json:"shortLeverage"`
	CrossMarginLeverage string `json:"crossMarginLeverage"`
	MarginMode          string `json:"marginMode"`
}

type PrivateRestClassicFuturesAccountSetAllLeverageRes string
type PrivateRestClassicFuturesAccountSetMarginRes string
type PrivateRestClassicFuturesAccountSetAssetModeRes string
type PrivateRestClassicFuturesAccountSetMarginModeRes = PrivateRestClassicFuturesAccountSetLeverageRes

type PrivateRestClassicFuturesAccountSetPositionModeRes struct {
	PosMode string `json:"posMode"`
}

type PrivateRestClassicFuturesAccountBillItem struct {
	BillId       string `json:"billId"`
	Symbol       string `json:"symbol"`
	Amount       string `json:"amount"`
	Fee          string `json:"fee"`
	FeeByCoupon  string `json:"feeByCoupon"`
	BusinessType string `json:"businessType"`
	Coin         string `json:"coin"`
	Balance      string `json:"balance"`
	CTime        string `json:"cTime"`
}

type PrivateRestClassicFuturesAccountBillRes struct {
	Bills []PrivateRestClassicFuturesAccountBillItem `json:"bills"`
	EndId string                                     `json:"endId"`
}

type PrivateRestClassicFuturesAccountIsolatedSymbolsRow struct {
	Symbol     string `json:"symbol"`
	MarginMode string `json:"marginMode"`
}

type PrivateRestClassicFuturesAccountIsolatedSymbolsRes []PrivateRestClassicFuturesAccountIsolatedSymbolsRow
