package mybitgetapi

type Asset struct {
	Coin      string `json:"coin"`      // String 币种名称
	Equity    string `json:"equity"`    // String 币种权益
	UsdValue  string `json:"usdValue"`  // String 币种权益折算USD
	Balance   string `json:"balance"`   // String 币种余额
	Available string `json:"available"` // String 可用余额
	Debt      string `json:"debt"`      // String 负债
	Locked    string `json:"locked"`    // String 锁定余额
}

type PrivateRestUtaAccountBalanceRes struct {
	AccountEquity     string  `json:"accountEquity"`     // String 账户权益（USD）
	UsdtEquity        string  `json:"usdtEquity"`        // String 账户权益（USDT）
	BtcEquity         string  `json:"btcEquity"`         // String 账户权益（BTC）
	UnrealisedPnl     string  `json:"unrealisedPnl"`     // String 未实现盈亏（USD）
	UsdtUnrealisedPnl string  `json:"usdtUnrealisedPnl"` // String 未实现盈亏（USDT）
	BtcUnrealizedPnl  string  `json:"btcUnrealizedPnl"`  // String 未实现盈亏（BTC）
	EffEquity         string  `json:"effEquity"`         // String 有效权益（USD）
	Mmr               string  `json:"mmr"`               // String 维持保证金（USD）
	Imr               string  `json:"imr"`               // String 初始保证金（USD）
	MgnRatio          string  `json:"mgnRatio"`          // String 保证金率
	PositionMgnRatio  string  `json:"positionMgnRatio"`  // String 持仓保证金率
	Assets            []Asset `json:"assets"`            // Array 资产明细
}

type PrivateRestUtaAccountFundingAssetResRow struct {
	Coin      string `json:"coin"`      // String 币种名称
	Balance   string `json:"balance"`   // String 总余额
	Available string `json:"available"` // String 可用
	Frozen    string `json:"frozen"`    // String 冻结
}

type PrivateRestUtaAccountFundingAssetRes []PrivateRestUtaAccountFundingAssetResRow

type PrivateRestUtaAccountSettingResSymbolConfig struct {
	Category   string `json:"category"`   // String 产品类型
	Symbol     string `json:"symbol"`     // String 交易对名称
	MarginMode string `json:"marginMode"` // String 保证金模式 crossed/isolated
	Leverage   string `json:"leverage"`   // String 杠杆倍数
}

type PrivateRestUtaAccountSettingResCoinConfig struct {
	Coin     string `json:"coin"`     // String 币种
	Leverage string `json:"leverage"` // String 杠杆倍数
}

type PrivateRestUtaAccountSettingRes struct {
	Uid              string                                   `json:"uid"`              // String UID
	AccountMode      string                                   `json:"accountMode"`      // String 账户模式
	AssetMode        string                                   `json:"assetMode"`        // String 资产模式
	AccountLevel     string                                   `json:"accountLevel"`     // String 账户等级
	HoldMode         string                                   `json:"holdMode"`         // String 持仓模式
	StpMode          string                                   `json:"stpMode"`          // String STP模式
	SymbolConfigList []PrivateRestUtaAccountSettingResSymbolConfig `json:"symbolConfigList"` // Array 交易对配置
	CoinConfigList   []PrivateRestUtaAccountSettingResCoinConfig   `json:"coinConfigList"`   // Array 币种配置
}

type PrivateRestUtaAccountSetLeverageRes string

type PrivateRestUtaAccountSetHoldModeRes string

type PrivateRestUtaAccountFinancialRecordsResRow struct {
	Category string `json:"category"` // String 产品类型
	Id       string `json:"id"`       // String 记录ID
	Symbol   string `json:"symbol"`   // String 交易对名称
	Coin     string `json:"coin"`     // String 币种名称
	Type     string `json:"type"`     // String 流水类型
	Amount   string `json:"amount"`   // String 数量
	Fee      string `json:"fee"`      // String 手续费
	Balance  string `json:"balance"`  // String 余额
	Ts       string `json:"ts"`       // String 数据时间戳 Unix毫秒
}

type PrivateRestUtaAccountFinancialRecordsRes struct {
	List   []PrivateRestUtaAccountFinancialRecordsResRow `json:"list"`   // Array 数据列表
	Cursor string                                         `json:"cursor"` // String 游标
}

type PrivateRestUtaAccountFeeRateRes struct {
	MakerFeeRate string `json:"makerFeeRate"` // String Maker费率
	TakerFeeRate string `json:"takerFeeRate"` // String Taker费率
}

type PrivateRestUtaAccountMaxTransferableRes struct {
	Coin              string `json:"coin"`              // String 币种名称
	MaxTransfer       string `json:"maxTransfer"`       // String 最大可转出数量
	BorrowMaxTransfer string `json:"borrowMaxTransfer"` // String 借币最大可转出数量
}

type PrivateRestUtaAccountTransferableCoinsRes []string

type PrivateRestUtaAccountTransferRes struct {
	TransferId string `json:"transferId"` // String 划转ID
}

// 切换账户接口返回 data 为 null
type PrivateRestUtaAccountSwitchRes any

type PrivateRestUtaAccountSwitchStatusRes struct {
	Status string `json:"status"` // String 账户切换状态 process/success/fail
	Reason string `json:"reason"` // String 失败原因，仅 status=fail 时返回
}
