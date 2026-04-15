package mybitgetapi

type PrivateRestClassicAPI int

const (
	PrivateRestClassicSpotAccountInfo            PrivateRestClassicAPI = iota // GET 获取账户信息
	PrivateRestClassicSpotAccountAssets                                       // GET 获取账户现货资产
	PrivateRestClassicSpotAccountBills                                        // GET 获取账单流水
	PrivateRestClassicSpotWalletTransfer                                      // POST 划转
	PrivateRestClassicSpotAccountTransferRecords                              // GET 获取划转记录
	PrivateRestClassicSpotAccountSwitchDeduct                                 // POST 开启BGB抵扣
	PrivateRestClassicSpotAccountDeductInfo                                   // GET 当前是否开启BGB抵扣
	PrivateRestClassicSpotAccountUpgrade                                      // POST 账户升级
	PrivateRestClassicSpotAccountUpgradeStatus                                // GET 查询升级状态

	PrivateRestClassicSpotTradePlaceOrder              // POST 下单
	PrivateRestClassicSpotTradeCancelReplaceOrder      // POST 撤单再下单
	PrivateRestClassicSpotTradeBatchCancelReplaceOrder // POST 批量撤单再下单
	PrivateRestClassicSpotTradeCancelOrder             // POST 撤单
	PrivateRestClassicSpotTradeBatchOrders             // POST 批量下单
	PrivateRestClassicSpotTradeBatchCancelOrder        // POST 批量撤单
	PrivateRestClassicSpotTradeOrderInfo               // GET 获取订单详情
	PrivateRestClassicSpotTradeUnfilledOrders          // GET 获取当前委托列表
	PrivateRestClassicSpotTradeHistoryOrders           // GET 获取历史委托列表
	PrivateRestClassicSpotTradeFills                   // GET 获取成交明细

	PrivateRestClassicFuturesAccountSingleAccount    // GET 获取单个交易对账户信息
	PrivateRestClassicFuturesAccountList             // GET 获取账户列表
	PrivateRestClassicFuturesAccountSubAccountAssets // GET 获取所有子账户合约资产信息
	PrivateRestClassicFuturesAccountInterestHistory  // GET U本位合约利息记录
	PrivateRestClassicFuturesAccountMaxOpen          // GET 获取最大可开
	PrivateRestClassicFuturesAccountLiquidationPrice // GET 获取强平价格
	PrivateRestClassicFuturesAccountOpenCount        // GET 获取可开数量
	PrivateRestClassicFuturesAccountSetLeverage      // POST 调整杠杆
	PrivateRestClassicFuturesAccountSetAllLeverage   // POST 调整产品线杠杆
	PrivateRestClassicFuturesAccountSetMargin        // POST 调整保证金
	PrivateRestClassicFuturesAccountSetAssetMode     // POST 设置U本位联合保证金模式
	PrivateRestClassicFuturesAccountSetMarginMode    // POST 调整保证金模式
	PrivateRestClassicFuturesAccountSetPositionMode  // POST 调整单双向持仓模式
	PrivateRestClassicFuturesAccountBill             // GET 获取账务记录
	PrivateRestClassicFuturesAccountIsolatedSymbols  // GET 获取逐仓模式交易对

	PrivateRestClassicFuturesPositionQueryPositionLever // GET 获取仓位档位梯度配置
	PrivateRestClassicFuturesPositionSinglePosition     // GET 获取单个合约仓位信息
	PrivateRestClassicFuturesPositionAllPosition        // GET 获取全部合约仓位信息
	PrivateRestClassicFuturesPositionHistoryPosition    // GET 获取合约历史持仓列表

	PrivateRestClassicFuturesTradePlaceOrder       // POST 下单
	PrivateRestClassicFuturesTradeReversal         // POST 一键反手
	PrivateRestClassicFuturesTradeBatchOrder       // POST 批量下单
	PrivateRestClassicFuturesTradeModifyOrder      // POST 修改订单
	PrivateRestClassicFuturesTradeCancelOrder      // POST 撤单
	PrivateRestClassicFuturesTradeBatchCancelOrder // POST 批量撤单
	PrivateRestClassicFuturesTradeClosePositions   // POST 一键市价平仓
	PrivateRestClassicFuturesTradeOrderDetail      // GET 获取订单详情
	PrivateRestClassicFuturesTradeOrderFills       // GET 获取成交明细
	PrivateRestClassicFuturesTradeFillHistory      // GET 获取历史成交明细
	PrivateRestClassicFuturesTradeOrdersPending    // GET 查询当前委托
	PrivateRestClassicFuturesTradeOrdersHistory    // GET 获取历史委托
	PrivateRestClassicFuturesTradeCancelAllOrders  // POST 一键全撤

	PrivateRestClassicMarginCrossTradePlaceOrder       // POST 全仓下单
	PrivateRestClassicMarginCrossTradeBatchPlaceOrder  // POST 全仓批量下单
	PrivateRestClassicMarginCrossTradeCancelOrder      // POST 全仓撤单
	PrivateRestClassicMarginCrossTradeBatchCancelOrder // POST 全仓批量撤单
	PrivateRestClassicMarginCrossTradeOpenOrders       // GET 获取全仓当前委托
	PrivateRestClassicMarginCrossTradeHistoryOrders    // GET 获取全仓历史委托
	PrivateRestClassicMarginCrossTradeFills            // GET 获取全仓成交明细
	PrivateRestClassicMarginCrossTradeLiquidationOrder // GET 获取全仓强平订单

	PrivateRestClassicMarginCrossAccountAssets                // GET 获取全仓杠杆账户资产
	PrivateRestClassicMarginCrossAccountBorrow                // POST 全仓借款
	PrivateRestClassicMarginCrossAccountRepay                 // POST 全仓还款
	PrivateRestClassicMarginCrossAccountRiskRate              // GET 全仓风险率
	PrivateRestClassicMarginCrossAccountMaxBorrowableAmount   // GET 全仓最大可借数量
	PrivateRestClassicMarginCrossAccountMaxTransferOutAmount  // GET 全仓最大可转出数量
	PrivateRestClassicMarginCrossAccountInterestRateAndLimit  // GET 全仓利率与最大可借数量配置
	PrivateRestClassicMarginCrossAccountTierData              // GET 全仓档位梯度配置
	PrivateRestClassicMarginCrossAccountFlashRepay            // POST 全仓一键还款
	PrivateRestClassicMarginCrossAccountQueryFlashRepayStatus // POST 获取全仓还款结果
	PrivateRestClassicMarginCrossRecordBorrowHistory          // GET 获取全仓借款历史
	PrivateRestClassicMarginCrossRecordRepayHistory           // GET 获取全仓还款历史
	PrivateRestClassicMarginCrossRecordInterestHistory        // GET 获取全仓利息历史
	PrivateRestClassicMarginCrossRecordLiquidationHistory     // GET 获取全仓强平历史
	PrivateRestClassicMarginCrossRecordFinancialRecords       // GET 获取全仓财务流水记录

	PrivateRestClassicMarginIsolatedTradePlaceOrder       // POST 逐仓下单
	PrivateRestClassicMarginIsolatedTradeBatchPlaceOrder  // POST 逐仓批量下单
	PrivateRestClassicMarginIsolatedTradeCancelOrder      // POST 逐仓撤单
	PrivateRestClassicMarginIsolatedTradeBatchCancelOrder // POST 逐仓批量撤单
	PrivateRestClassicMarginIsolatedTradeOpenOrders       // GET 获取逐仓当前委托
	PrivateRestClassicMarginIsolatedTradeHistoryOrders    // GET 获取逐仓历史委托
	PrivateRestClassicMarginIsolatedTradeFills            // GET 获取逐仓成交明细
	PrivateRestClassicMarginIsolatedTradeLiquidationOrder // GET 获取逐仓强平订单

	PrivateRestClassicMarginIsolatedAccountAssets                // GET 获取逐仓杠杆账户资产
	PrivateRestClassicMarginIsolatedAccountBorrow                // POST 逐仓借款
	PrivateRestClassicMarginIsolatedAccountRepay                 // POST 逐仓还款
	PrivateRestClassicMarginIsolatedAccountRiskRate              // GET 逐仓风险率
	PrivateRestClassicMarginIsolatedAccountTierData              // GET 逐仓档位梯度配置
	PrivateRestClassicMarginIsolatedAccountInterestRateAndLimit  // GET 逐仓利率与最大可借数量配置
	PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmount   // GET 逐仓最大可借数量
	PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmount  // GET 逐仓最大可转出数量
	PrivateRestClassicMarginIsolatedAccountFlashRepay            // POST 逐仓一键还款
	PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatus // POST 获取逐仓还款结果

	PrivateRestClassicMarginMarketInterestRateRecord // GET 获取杠杆利率
)

var PrivateRestClassicAPIMap = map[PrivateRestClassicAPI]string{
	PrivateRestClassicSpotAccountInfo:            "/api/v2/spot/account/info",            // GET 获取账户信息
	PrivateRestClassicSpotAccountAssets:          "/api/v2/spot/account/assets",          // GET 获取账户现货资产
	PrivateRestClassicSpotAccountBills:           "/api/v2/spot/account/bills",           // GET 获取账单流水
	PrivateRestClassicSpotWalletTransfer:         "/api/v2/spot/wallet/transfer",         // POST 划转
	PrivateRestClassicSpotAccountTransferRecords: "/api/v2/spot/account/transferRecords", // GET 获取划转记录
	PrivateRestClassicSpotAccountSwitchDeduct:    "/api/v2/spot/account/switch-deduct",   // POST 开启BGB抵扣
	PrivateRestClassicSpotAccountDeductInfo:      "/api/v2/spot/account/deduct-info",     // GET 当前是否开启BGB抵扣
	PrivateRestClassicSpotAccountUpgrade:         "/api/v2/spot/account/upgrade",         // POST 账户升级
	PrivateRestClassicSpotAccountUpgradeStatus:   "/api/v2/spot/account/upgrade-status",  // GET 查询升级状态

	PrivateRestClassicSpotTradePlaceOrder:              "/api/v2/spot/trade/place-order",                // POST 下单
	PrivateRestClassicSpotTradeCancelReplaceOrder:      "/api/v2/spot/trade/cancel-replace-order",       // POST 撤单再下单
	PrivateRestClassicSpotTradeBatchCancelReplaceOrder: "/api/v2/spot/trade/batch-cancel-replace-order", // POST 批量撤单再下单
	PrivateRestClassicSpotTradeCancelOrder:             "/api/v2/spot/trade/cancel-order",               // POST 撤单
	PrivateRestClassicSpotTradeBatchOrders:             "/api/v2/spot/trade/batch-orders",               // POST 批量下单
	PrivateRestClassicSpotTradeBatchCancelOrder:        "/api/v2/spot/trade/batch-cancel-order",         // POST 批量撤单
	PrivateRestClassicSpotTradeOrderInfo:               "/api/v2/spot/trade/orderInfo",                  // GET 获取订单详情
	PrivateRestClassicSpotTradeUnfilledOrders:          "/api/v2/spot/trade/unfilled-orders",            // GET 获取当前委托列表
	PrivateRestClassicSpotTradeHistoryOrders:           "/api/v2/spot/trade/history-orders",             // GET 获取历史委托列表
	PrivateRestClassicSpotTradeFills:                   "/api/v2/spot/trade/fills",                      // GET 获取成交明细

	PrivateRestClassicFuturesAccountSingleAccount:    "/api/v2/mix/account/account",            // GET 获取单个交易对账户信息
	PrivateRestClassicFuturesAccountList:             "/api/v2/mix/account/accounts",           // GET 获取账户列表
	PrivateRestClassicFuturesAccountSubAccountAssets: "/api/v2/mix/account/sub-account-assets", // GET 获取所有子账户合约资产信息
	PrivateRestClassicFuturesAccountInterestHistory:  "/api/v2/mix/account/interest-history",   // GET U本位合约利息记录
	PrivateRestClassicFuturesAccountMaxOpen:          "/api/v2/mix/account/max-open",           // GET 获取最大可开
	PrivateRestClassicFuturesAccountLiquidationPrice: "/api/v2/mix/account/liq-price",          // GET 获取强平价格
	PrivateRestClassicFuturesAccountOpenCount:        "/api/v2/mix/account/open-count",         // GET 获取可开数量
	PrivateRestClassicFuturesAccountSetLeverage:      "/api/v2/mix/account/set-leverage",       // POST 调整杠杆
	PrivateRestClassicFuturesAccountSetAllLeverage:   "/api/v2/mix/account/set-all-leverage",   // POST 调整产品线杠杆
	PrivateRestClassicFuturesAccountSetMargin:        "/api/v2/mix/account/set-margin",         // POST 调整保证金
	PrivateRestClassicFuturesAccountSetAssetMode:     "/api/v2/mix/account/set-asset-mode",     // POST 设置U本位联合保证金模式
	PrivateRestClassicFuturesAccountSetMarginMode:    "/api/v2/mix/account/set-margin-mode",    // POST 调整保证金模式
	PrivateRestClassicFuturesAccountSetPositionMode:  "/api/v2/mix/account/set-position-mode",  // POST 调整单双向持仓模式
	PrivateRestClassicFuturesAccountBill:             "/api/v2/mix/account/bill",               // GET 获取账务记录
	PrivateRestClassicFuturesAccountIsolatedSymbols:  "/api/v2/mix/account/isolated-symbols",   // GET 获取逐仓模式交易对

	PrivateRestClassicFuturesPositionQueryPositionLever: "/api/v2/mix/market/query-position-lever", // GET 获取仓位档位梯度配置
	PrivateRestClassicFuturesPositionSinglePosition:     "/api/v2/mix/position/single-position",    // GET 获取单个合约仓位信息
	PrivateRestClassicFuturesPositionAllPosition:        "/api/v2/mix/position/all-position",       // GET 获取全部合约仓位信息
	PrivateRestClassicFuturesPositionHistoryPosition:    "/api/v2/mix/position/history-position",   // GET 获取合约历史持仓列表

	PrivateRestClassicFuturesTradePlaceOrder:       "/api/v2/mix/order/place-order",         // POST 下单
	PrivateRestClassicFuturesTradeReversal:         "/api/v2/mix/order/click-backhand",      // POST 一键反手
	PrivateRestClassicFuturesTradeBatchOrder:       "/api/v2/mix/order/batch-place-order",   // POST 批量下单
	PrivateRestClassicFuturesTradeModifyOrder:      "/api/v2/mix/order/modify-order",        // POST 修改订单
	PrivateRestClassicFuturesTradeCancelOrder:      "/api/v2/mix/order/cancel-order",        // POST 撤单
	PrivateRestClassicFuturesTradeBatchCancelOrder: "/api/v2/mix/order/batch-cancel-orders", // POST 批量撤单
	PrivateRestClassicFuturesTradeClosePositions:   "/api/v2/mix/order/close-positions",     // POST 一键市价平仓
	PrivateRestClassicFuturesTradeOrderDetail:      "/api/v2/mix/order/detail",              // GET 获取订单详情
	PrivateRestClassicFuturesTradeOrderFills:       "/api/v2/mix/order/fills",               // GET 获取成交明细
	PrivateRestClassicFuturesTradeFillHistory:      "/api/v2/mix/order/fill-history",        // GET 获取历史成交明细
	PrivateRestClassicFuturesTradeOrdersPending:    "/api/v2/mix/order/orders-pending",      // GET 查询当前委托
	PrivateRestClassicFuturesTradeOrdersHistory:    "/api/v2/mix/order/orders-history",      // GET 获取历史委托
	PrivateRestClassicFuturesTradeCancelAllOrders:  "/api/v2/mix/order/cancel-all-orders",   // POST 一键全撤

	PrivateRestClassicMarginCrossTradePlaceOrder:       "/api/v2/margin/crossed/place-order",        // POST 全仓下单
	PrivateRestClassicMarginCrossTradeBatchPlaceOrder:  "/api/v2/margin/crossed/batch-place-order",  // POST 全仓批量下单
	PrivateRestClassicMarginCrossTradeCancelOrder:      "/api/v2/margin/crossed/cancel-order",       // POST 全仓撤单
	PrivateRestClassicMarginCrossTradeBatchCancelOrder: "/api/v2/margin/crossed/batch-cancel-order", // POST 全仓批量撤单
	PrivateRestClassicMarginCrossTradeOpenOrders:       "/api/v2/margin/crossed/open-orders",        // GET 获取全仓当前委托
	PrivateRestClassicMarginCrossTradeHistoryOrders:    "/api/v2/margin/crossed/history-orders",     // GET 获取全仓历史委托
	PrivateRestClassicMarginCrossTradeFills:            "/api/v2/margin/crossed/fills",              // GET 获取全仓成交明细
	PrivateRestClassicMarginCrossTradeLiquidationOrder: "/api/v2/margin/crossed/liquidation-order",  // GET 获取全仓强平订单

	PrivateRestClassicMarginCrossAccountAssets:                "/api/v2/margin/crossed/account/assets",                   // GET 获取全仓杠杆账户资产
	PrivateRestClassicMarginCrossAccountBorrow:                "/api/v2/margin/crossed/account/borrow",                   // POST 全仓借款
	PrivateRestClassicMarginCrossAccountRepay:                 "/api/v2/margin/crossed/account/repay",                    // POST 全仓还款
	PrivateRestClassicMarginCrossAccountRiskRate:              "/api/v2/margin/crossed/account/risk-rate",                // GET 全仓风险率
	PrivateRestClassicMarginCrossAccountMaxBorrowableAmount:   "/api/v2/margin/crossed/account/max-borrowable-amount",    // GET 全仓最大可借数量
	PrivateRestClassicMarginCrossAccountMaxTransferOutAmount:  "/api/v2/margin/crossed/account/max-transfer-out-amount",  // GET 全仓最大可转出数量
	PrivateRestClassicMarginCrossAccountInterestRateAndLimit:  "/api/v2/margin/crossed/interest-rate-and-limit",          // GET 全仓利率与最大可借数量配置
	PrivateRestClassicMarginCrossAccountTierData:              "/api/v2/margin/crossed/tier-data",                        // GET 全仓档位梯度配置
	PrivateRestClassicMarginCrossAccountFlashRepay:            "/api/v2/margin/crossed/account/flash-repay",              // POST 全仓一键还款
	PrivateRestClassicMarginCrossAccountQueryFlashRepayStatus: "/api/v2/margin/crossed/account/query-flash-repay-status", // POST 获取全仓还款结果
	PrivateRestClassicMarginCrossRecordBorrowHistory:          "/api/v2/margin/crossed/borrow-history",                   // GET 获取全仓借款历史
	PrivateRestClassicMarginCrossRecordRepayHistory:           "/api/v2/margin/crossed/repay-history",                    // GET 获取全仓还款历史
	PrivateRestClassicMarginCrossRecordInterestHistory:        "/api/v2/margin/crossed/interest-history",                 // GET 获取全仓利息历史
	PrivateRestClassicMarginCrossRecordLiquidationHistory:     "/api/v2/margin/crossed/liquidation-history",              // GET 获取全仓强平历史
	PrivateRestClassicMarginCrossRecordFinancialRecords:       "/api/v2/margin/crossed/financial-records",                // GET 获取全仓财务流水记录

	PrivateRestClassicMarginIsolatedTradePlaceOrder:       "/api/v2/margin/isolated/place-order",        // POST 逐仓下单
	PrivateRestClassicMarginIsolatedTradeBatchPlaceOrder:  "/api/v2/margin/isolated/batch-place-order",  // POST 逐仓批量下单
	PrivateRestClassicMarginIsolatedTradeCancelOrder:      "/api/v2/margin/isolated/cancel-order",       // POST 逐仓撤单
	PrivateRestClassicMarginIsolatedTradeBatchCancelOrder: "/api/v2/margin/isolated/batch-cancel-order", // POST 逐仓批量撤单
	PrivateRestClassicMarginIsolatedTradeOpenOrders:       "/api/v2/margin/isolated/open-orders",        // GET 获取逐仓当前委托
	PrivateRestClassicMarginIsolatedTradeHistoryOrders:    "/api/v2/margin/isolated/history-orders",     // GET 获取逐仓历史委托
	PrivateRestClassicMarginIsolatedTradeFills:            "/api/v2/margin/isolated/fills",              // GET 获取逐仓成交明细
	PrivateRestClassicMarginIsolatedTradeLiquidationOrder: "/api/v2/margin/isolated/liquidation-order",  // GET 获取逐仓强平订单

	PrivateRestClassicMarginIsolatedAccountAssets:                "/api/v2/margin/isolated/account/assets",                   // GET 获取逐仓杠杆账户资产
	PrivateRestClassicMarginIsolatedAccountBorrow:                "/api/v2/margin/isolated/account/borrow",                   // POST 逐仓借款
	PrivateRestClassicMarginIsolatedAccountRepay:                 "/api/v2/margin/isolated/account/repay",                    // POST 逐仓还款
	PrivateRestClassicMarginIsolatedAccountRiskRate:              "/api/v2/margin/isolated/account/risk-rate",                // GET 逐仓风险率
	PrivateRestClassicMarginIsolatedAccountTierData:              "/api/v2/margin/isolated/tier-data",                        // GET 逐仓档位梯度配置
	PrivateRestClassicMarginIsolatedAccountInterestRateAndLimit:  "/api/v2/margin/isolated/interest-rate-and-limit",          // GET 逐仓利率与最大可借数量配置
	PrivateRestClassicMarginIsolatedAccountMaxBorrowableAmount:   "/api/v2/margin/isolated/account/max-borrowable-amount",    // GET 逐仓最大可借数量
	PrivateRestClassicMarginIsolatedAccountMaxTransferOutAmount:  "/api/v2/margin/isolated/account/max-transfer-out-amount",  // GET 逐仓最大可转出数量
	PrivateRestClassicMarginIsolatedAccountFlashRepay:            "/api/v2/margin/isolated/account/flash-repay",              // POST 逐仓一键还款
	PrivateRestClassicMarginIsolatedAccountQueryFlashRepayStatus: "/api/v2/margin/isolated/account/query-flash-repay-status", // POST 获取逐仓还款结果

	PrivateRestClassicMarginMarketInterestRateRecord: "/api/v2/margin/interest-rate-record", // GET 获取杠杆利率
}
