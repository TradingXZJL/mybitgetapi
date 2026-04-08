package mybitgetapi

type PrivateRestAPI int

const (
	PrivateRestUtaAccountBalance           PrivateRestAPI = iota // GET 获取账户余额
	PrivateRestUtaAccountFundingAssets                           // GET 获取资金账户资产
	PrivateRestUtaAccountSetting                                 // GET 获取账户设置
	PrivateRestUtaAccountSetLeverage                             // POST 调整杠杆
	PrivateRestUtaAccountSetHoldMode                             // POST 调整持仓模式
	PrivateRestUtaAccountFinancialRecords                        // GET 获取资金流水
	PrivateRestUtaAccountFeeRate                                 // GET 获取账户费率
	PrivateRestUtaAccountMaxTransferable                         // GET 获取最大可转出
	PrivateRestUtaAccountTransferableCoins                       // GET 获取可划转币种
	PrivateRestUtaAccountTransfer                                // POST 资金划转
	PrivateRestUtaAccountSwitch                                  // POST 切换账户（切换至 classic 账户）
	PrivateRestUtaAccountSwitchStatus                             // GET 获取切换状态

	PrivateRestUtaTradePlaceOrder        // POST 下单
	PrivateRestUtaTradeModifyOrder       // POST 修改订单
	PrivateRestUtaTradeCancelOrder       // POST 撤单
	PrivateRestUtaTradePlaceBatch        // POST 批量下单
	PrivateRestUtaTradeBatchModifyOrder  // POST 批量改单
	PrivateRestUtaTradeCancelBatch       // POST 批量撤单
	PrivateRestUtaTradeCancelSymbolOrder // POST 一键撤单
	PrivateRestUtaTradeClosePositions    // POST 一键平仓
	PrivateRestUtaTradeOrderInfo         // GET 查询订单详情
	PrivateRestUtaTradeUnfilledOrders    // GET 查询当前委托
	PrivateRestUtaTradeHistoryOrders     // GET 查询历史委托
	PrivateRestUtaTradeFills             // GET 查询成交明细
	PrivateRestUtaPositionCurrent        // GET 查询当前持仓
	PrivateRestUtaPositionHistory        // GET 查询历史持仓
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	PrivateRestUtaAccountBalance:           "/api/v3/account/assets",             // GET 获取账户余额
	PrivateRestUtaAccountFundingAssets:     "/api/v3/account/funding-assets",     // GET 获取资金账户资产
	PrivateRestUtaAccountSetting:           "/api/v3/account/settings",           // GET 获取账户设置
	PrivateRestUtaAccountSetLeverage:       "/api/v3/account/set-leverage",       // POST 调整杠杆
	PrivateRestUtaAccountSetHoldMode:       "/api/v3/account/set-hold-mode",      // POST 调整持仓模式
	PrivateRestUtaAccountFinancialRecords:  "/api/v3/account/financial-records",  // GET 获取资金流水
	PrivateRestUtaAccountFeeRate:           "/api/v3/account/fee-rate",           // GET 获取账户费率
	PrivateRestUtaAccountMaxTransferable:   "/api/v3/account/max-transferable",   // GET 获取最大可转出
	PrivateRestUtaAccountTransferableCoins: "/api/v3/account/transferable-coins", // GET 获取可划转币种
	PrivateRestUtaAccountTransfer:          "/api/v3/account/transfer",           // POST 资金划转
	PrivateRestUtaAccountSwitch:            "/api/v3/account/switch",             // POST 切换账户（切换至 classic 账户）
	PrivateRestUtaAccountSwitchStatus:      "/api/v3/account/switch-status",      // GET 获取切换状态

	PrivateRestUtaTradePlaceOrder:        "/api/v3/trade/place-order",         // POST 下单
	PrivateRestUtaTradeModifyOrder:       "/api/v3/trade/modify-order",        // POST 修改订单
	PrivateRestUtaTradeCancelOrder:       "/api/v3/trade/cancel-order",        // POST 撤单
	PrivateRestUtaTradePlaceBatch:        "/api/v3/trade/place-batch",         // POST 批量下单
	PrivateRestUtaTradeBatchModifyOrder:  "/api/v3/trade/batch-modify-order",  // POST 批量改单
	PrivateRestUtaTradeCancelBatch:       "/api/v3/trade/cancel-batch",        // POST 批量撤单
	PrivateRestUtaTradeCancelSymbolOrder: "/api/v3/trade/cancel-symbol-order", // POST 一键撤单
	PrivateRestUtaTradeClosePositions:    "/api/v3/trade/close-positions",     // POST 一键平仓
	PrivateRestUtaTradeOrderInfo:         "/api/v3/trade/order-info",          // GET 查询订单详情
	PrivateRestUtaTradeUnfilledOrders:    "/api/v3/trade/unfilled-orders",     // GET 查询当前委托
	PrivateRestUtaTradeHistoryOrders:     "/api/v3/trade/history-orders",      // GET 查询历史委托
	PrivateRestUtaTradeFills:             "/api/v3/trade/fills",               // GET 查询成交明细
	PrivateRestUtaPositionCurrent:        "/api/v3/position/current-position", // GET 查询当前持仓
	PrivateRestUtaPositionHistory:        "/api/v3/position/history-position", // GET 查询历史持仓
}
