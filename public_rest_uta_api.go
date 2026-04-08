package mybitgetapi

type PublicRestUtaAPI int

const (
	PublicRestUtaMarketInstruments     PublicRestUtaAPI = iota // GET 获取交易产品信息 查询在线交易币对信息
	PublicRestUtaMarketTickers                                 // GET 获取行情数据 查询在线交易币对行情数据
	PublicRestUtaMarketOrderBook                               // GET 获取深度数据 查询订单簿深度
	PublicRestUtaMarketCandles                                 // GET 获取K线数据
	PublicRestUtaMarketHistoryCandles                          // GET 获取历史K线数据
	PublicRestUtaMarketCurrentFundRate                         // GET 获取当前资金费率
	PublicRestUtaMarketHistoryFundRate                         // GET 获取历史资金费率
)

var PublicRestUtaAPIMap = map[PublicRestUtaAPI]string{
	PublicRestUtaMarketInstruments:     "/api/v3/market/instruments",       // GET 获取交易产品信息 查询在线交易币对信息
	PublicRestUtaMarketTickers:         "/api/v3/market/tickers",           // GET 获取行情数据 查询在线交易币对行情数据
	PublicRestUtaMarketOrderBook:       "/api/v3/market/orderbook",         // GET 获取深度数据 查询订单簿深度
	PublicRestUtaMarketCandles:         "/api/v3/market/candles",           // GET 获取K线数据
	PublicRestUtaMarketHistoryCandles:  "/api/v3/market/history-candles",   // GET 获取历史K线数据
	PublicRestUtaMarketCurrentFundRate: "/api/v3/market/current-fund-rate", // GET 获取当前资金费率
	PublicRestUtaMarketHistoryFundRate: "/api/v3/market/history-fund-rate", // GET 获取历史资金费率
}
