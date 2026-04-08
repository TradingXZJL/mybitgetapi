package mybitgetapi

type PublicRestClassicAPI int

const (
	PublicRestPublicTime PublicRestClassicAPI = iota // 获取服务器时间

	PublicRestClassicSpotMarketCoins          // GET 获取币种信息
	PublicRestClassicSpotMarketSymbols        // GET 获取交易对信息
	PublicRestClassicSpotMarketTickers        // GET 获取行情信息
	PublicRestClassicSpotMarketMergeDepth     // GET 获取合并交易深度
	PublicRestClassicSpotMarketOrderbook      // GET 获取交易深度
	PublicRestClassicSpotMarketCandles        // GET 获取K线数据
	PublicRestClassicSpotMarketHistoryCandles // GET 获取历史K线数据
	PublicRestClassicSpotMarketFills          // GET 获取最近成交数据
	PublicRestClassicSpotMarketFillsHistory   // GET 获取市场成交数据

	PublicRestClassicFuturesMarketVipFeeRate          // GET 获取合约VIP费率
	PublicRestClassicFuturesMarketInterestRate        // GET 获取利率记录
	PublicRestClassicFuturesMarketMergeDepth          // GET 获取合并深度
	PublicRestClassicFuturesMarketTicker              // GET 获取单个交易对行情
	PublicRestClassicFuturesMarketTickers             // GET 获取全部交易对行情
	PublicRestClassicFuturesMarketFills               // GET 获取最近成交
	PublicRestClassicFuturesMarketFillsHistory        // GET 获取历史成交
	PublicRestClassicFuturesMarketCandles             // GET 获取K线数据
	PublicRestClassicFuturesMarketHistoryCandles      // GET 获取历史K线
	PublicRestClassicFuturesMarketHistoryIndexCandles // GET 获取指数历史K线
	PublicRestClassicFuturesMarketHistoryMarkCandles  // GET 获取标记价格历史K线
	PublicRestClassicFuturesMarketCurrentFundRate     // GET 获取当前资金费率
	PublicRestClassicFuturesMarketHistoryFundRate     // GET 获取历史资金费率
	PublicRestClassicFuturesMarketFundingTime         // GET 获取下次资金费结算时间
	PublicRestClassicFuturesMarketContracts           // GET 获取合约信息

	PublicRestClassicMarginMarketCurrencies // GET 获取支持杠杆的所有交易对
)

var PublicRestClassicAPIMap = map[PublicRestClassicAPI]string{
	PublicRestPublicTime: "/api/v2/public/time", // GET 获取服务器时间

	PublicRestClassicSpotMarketCoins:          "/api/v2/spot/public/coins",           // GET 获取币种信息
	PublicRestClassicSpotMarketSymbols:        "/api/v2/spot/public/symbols",         // GET 获取交易对信息
	PublicRestClassicSpotMarketTickers:        "/api/v2/spot/market/tickers",         // GET 获取行情信息
	PublicRestClassicSpotMarketMergeDepth:     "/api/v2/spot/market/merge-depth",     // GET 获取合并交易深度
	PublicRestClassicSpotMarketOrderbook:      "/api/v2/spot/market/orderbook",       // GET 获取交易深度
	PublicRestClassicSpotMarketCandles:        "/api/v2/spot/market/candles",         // GET 获取K线数据
	PublicRestClassicSpotMarketHistoryCandles: "/api/v2/spot/market/history-candles", // GET 获取历史K线数据
	PublicRestClassicSpotMarketFills:          "/api/v2/spot/market/fills",           // GET 获取最近成交数据
	PublicRestClassicSpotMarketFillsHistory:   "/api/v2/spot/market/fills-history",   // GET 获取市场成交数据

	PublicRestClassicFuturesMarketVipFeeRate:          "/api/v2/mix/market/vip-fee-rate",                // GET 获取合约VIP费率
	PublicRestClassicFuturesMarketInterestRate:        "/api/v2/mix/market/union-interest-rate-history", // GET 获取利率记录
	PublicRestClassicFuturesMarketMergeDepth:          "/api/v2/mix/market/merge-depth",                 // GET 获取合并深度
	PublicRestClassicFuturesMarketTicker:              "/api/v2/mix/market/ticker",                      // GET 获取单个交易对行情
	PublicRestClassicFuturesMarketTickers:             "/api/v2/mix/market/tickers",                     // GET 获取全部交易对行情
	PublicRestClassicFuturesMarketFills:               "/api/v2/mix/market/fills",                       // GET 获取最近成交
	PublicRestClassicFuturesMarketFillsHistory:        "/api/v2/mix/market/fills-history",               // GET 获取历史成交
	PublicRestClassicFuturesMarketCandles:             "/api/v2/mix/market/candles",                     // GET 获取K线数据
	PublicRestClassicFuturesMarketHistoryCandles:      "/api/v2/mix/market/history-candles",             // GET 获取历史K线
	PublicRestClassicFuturesMarketHistoryIndexCandles: "/api/v2/mix/market/history-index-candles",       // GET 获取指数历史K线
	PublicRestClassicFuturesMarketHistoryMarkCandles:  "/api/v2/mix/market/history-mark-candles",        // GET 获取标记价格历史K线
	PublicRestClassicFuturesMarketCurrentFundRate:     "/api/v2/mix/market/current-fund-rate",           // GET 获取当前资金费率
	PublicRestClassicFuturesMarketHistoryFundRate:     "/api/v2/mix/market/history-fund-rate",           // GET 获取历史资金费率
	PublicRestClassicFuturesMarketFundingTime:         "/api/v2/mix/market/funding-time",                // GET 获取下次资金费结算时间
	PublicRestClassicFuturesMarketContracts:           "/api/v2/mix/market/contracts",                   // GET 获取合约信息

	PublicRestClassicMarginMarketCurrencies: "/api/v2/margin/currencies", // GET 获取支持杠杆的所有交易对
}
