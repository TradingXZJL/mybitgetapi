package mybitgetapi

type PrivateRestClassicSpotAccountInfoRes struct {
	UserId      string   `json:"userId"`      // String 用户ID
	InviterId   string   `json:"inviterId"`   // String 邀请人UID
	IPs         string   `json:"ips"`         // String IP白名单
	Authorities []string `json:"authorities"` // Array 权限列表
	ParentId    int64    `json:"parentId"`    // Int 母账户UID
	TraderType  string   `json:"traderType"`  // String trader/not_trader
	ChannelCode string   `json:"channelCode"` // String 渠道邀请码
	Channel     string   `json:"channel"`     // String 渠道
	RegisTime   string   `json:"regisTime"`   // String 注册时间
}

type PrivateRestClassicSpotAccountAssetsRow struct {
	Coin           string `json:"coin"`           // String 币种名称
	Available      string `json:"available"`      // String 可用资产
	Frozen         string `json:"frozen"`         // String 冻结资产
	Locked         string `json:"locked"`         // String 锁仓资产
	LimitAvailable string `json:"limitAvailable"` // String 受限可用
	UTime          string `json:"uTime"`          // String 更新时间
}

type PrivateRestClassicSpotAccountAssetsRes []PrivateRestClassicSpotAccountAssetsRow

type PrivateRestClassicSpotAccountBillsRow struct {
	CTime        string `json:"cTime"`        // String 创建时间
	Coin         string `json:"coin"`         // String 币种名称
	GroupType    string `json:"groupType"`    // String 账单类型
	BusinessType string `json:"businessType"` // String 业务类型
	Size         string `json:"size"`         // String 数量
	Balance      string `json:"balance"`      // String 划转后资产
	Fees         string `json:"fees"`         // String 手续费
	BillId       string `json:"billId"`       // String 账单流水ID
}

type PrivateRestClassicSpotAccountBillsRes []PrivateRestClassicSpotAccountBillsRow

type PrivateRestClassicSpotWalletTransferRes struct {
	TransferId string `json:"transferId"` // String 划转ID
	ClientOid  string `json:"clientOid"`  // String 客户自定义ID
}

type PrivateRestClassicSpotAccountTransferRecordsRow struct {
	Coin       string `json:"coin"`       // String 币种名称
	Status     string `json:"status"`     // String 划转状态
	ToType     string `json:"toType"`     // String 转入账户类型
	ToSymbol   string `json:"toSymbol"`   // String 转入账户交易对
	FromType   string `json:"fromType"`   // String 转出账户类型
	FromSymbol string `json:"fromSymbol"` // String 转出账户交易对
	Size       string `json:"size"`       // String 数量
	Ts         string `json:"ts"`         // String 划转时间
	ClientOid  string `json:"clientOid"`  // String 客户自定义订单ID
	TransferId string `json:"transferId"` // String 划转订单ID
}

type PrivateRestClassicSpotAccountTransferRecordsRes []PrivateRestClassicSpotAccountTransferRecordsRow

type PrivateRestClassicSpotAccountSwitchDeductRes bool

type PrivateRestClassicSpotAccountDeductInfoRow struct {
	Deduct string `json:"deduct"` // String 是否开启 on/off
}

type PrivateRestClassicSpotAccountDeductInfoRes PrivateRestClassicSpotAccountDeductInfoRow

// 升级账户接口返回 data 为 null
type PrivateRestClassicSpotAccountUpgradeRes any

type PrivateRestClassicSpotAccountUpgradeStatusRes struct {
	Status string `json:"status"` // String 账户升级状态 process/success/fail
	Reason string `json:"reason"` // String 失败原因，仅 status=fail 时返回
}
