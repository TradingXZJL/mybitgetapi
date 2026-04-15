package mybitgetapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const (
	BITGET_API_WS_HOST_MAIN       = "ws.bitget.com"
	BITGET_API_WS_HOST_TEST       = "wspap.bitget.com"
	BITGET_API_WS_PUBLIC          = "/v3/ws/public"
	BITGET_API_WS_CLASSIC_PRIVATE = "/v2/ws/private"
	BITGET_API_WS_UTA_PRIVATE     = "/v3/ws/private"
)

const (
	WS_SUBSCRIBE   = "subscribe"
	WS_UNSUBSCRIBE = "unsubscribe"
	WS_LOGIN       = "login"
)

// Bitget 私有 WS 登录验签：timestamp + "GET" + "/user/verify"，HMAC-SHA256 后 Base64。
// 与官方说明一致，时间戳为毫秒字符串（与 REST ACCESS-TIMESTAMP 一致）。
// 参考：https://www.bitget.com/zh-CN/api-doc/classic/quickStart/websocket-intro
const wsLoginSignPath = "/user/verify"

var (
	WebsocketTimeout        = 30 * time.Second
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 200 * time.Millisecond
)

var wsSubIDSeed int64 = 0

type WsActionReq struct {
	Op   string           `json:"op"`   // subscribe / unsubscribe
	Args []WsSubscribeArg `json:"args"` // 订阅参数
}

type WsSubscribeArg struct {
	InstType string `json:"instType"`           // 产品线类型 SPOT / USDT-FUTURES 等
	Topic    string `json:"topic,omitempty"`    // 公共频道名 ticker、kline 等
	Channel  string `json:"channel,omitempty"`  // 私有频道名 fill 等
	Symbol   string `json:"symbol,omitempty"`   // 公共频道交易对
	InstId   string `json:"instId,omitempty"`   // 私有频道产品 ID，如 default
	Interval string `json:"interval,omitempty"` // k 线时间粒度
	Coin     string `json:"coin,omitempty"`     // 现货账户频道维度，目前仅支持 default
}

type WsActionResult struct {
	Event string         `json:"event"`         // subscribe / unsubscribe / login / error
	Arg   WsSubscribeArg `json:"arg,omitempty"` // 订阅参数
	Code  int64          `json:"code,omitempty"`
	Msg   string         `json:"msg,omitempty"`
}

// WsLoginReq Bitget 私有频道登录请求（与 subscribe 的 JSON 结构不同）
type WsLoginReq struct {
	Op   string       `json:"op"`
	Args []WsLoginArg `json:"args"`
}

type WsLoginArg struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

// 数据流订阅标准结构体
type Subscription[T any] struct {
	SubId      int64            // 订阅ID
	Ws         *WsStreamClient  // 订阅连接
	Op         string           // 订阅方法
	Args       []WsSubscribeArg // 订阅参数
	resultChan chan T           // 接收结果
	errChan    chan error       // 接收错误
	closeChan  chan struct{}    // 接收关闭信号
}

func (sub *Subscription[T]) ResultChan() chan T {
	return sub.resultChan
}

func (sub *Subscription[T]) ErrChan() chan error {
	return sub.errChan
}

func (sub *Subscription[T]) CloseChan() chan struct{} {
	return sub.closeChan
}

type WsStreamClient struct {
	instType InstType
	apiType  APIType
	conn     *websocket.Conn
	isClose  bool

	client *RestClient // 用于登录；重连后可再次 Login

	// lastLogin 保存最近一次成功发起的登录参数，用于断线重连后的自动再次登录。
	// 注意：这里保存的是参数快照，不依赖外部变量生命周期。
	lastLogin *WsLoginArg

	// reSubscribeMu 用于防止重连时并发重订阅（commonSubMap 会被多个 key 关联到同一个 sub）。
	reSubscribeMu *sync.Mutex

	AutoReConnectTimes int // 自动重连次数
	reconnecting       int32

	commonSubMap                MySyncMap[string, *Subscription[WsActionResult]]
	tickerSubMap                MySyncMap[string, *Subscription[WsTicker]]
	candleSubMap                MySyncMap[string, *Subscription[WsCandle]]
	booksSubMap                 MySyncMap[string, *Subscription[WsBooks]]
	tradesSubMap                MySyncMap[string, *Subscription[WsPublicTrade]]
	fillSubMap                  MySyncMap[string, *Subscription[WsFill]]
	orderSubMap                 MySyncMap[string, *Subscription[WsOrder]]
	accountSubMap               MySyncMap[string, *Subscription[WsAccount]]
	equitySubMap                MySyncMap[string, *Subscription[WsEquity]]
	positionsSubMap             MySyncMap[string, *Subscription[WsPositions]]
	marginCrossAccountSubMap    MySyncMap[string, *Subscription[WsMarginCrossAccountAssets]]
	marginIsolatedAccountSubMap MySyncMap[string, *Subscription[WsMarginIsolatedAccountAssets]]
	marginCrossOrdersSubMap     MySyncMap[string, *Subscription[WsMarginCrossOrders]]
	marginIsolatedOrdersSubMap  MySyncMap[string, *Subscription[WsMarginIsolatedOrders]]

	waitSubResult   *Subscription[WsActionResult]
	waitSubResultMu *sync.Mutex

	writeMu *sync.Mutex

	resultChan chan []byte
	errChan    chan error
}

type PublicWsStreamClient struct {
	WsStreamClient
}

func (*MyBitget) NewPublicWsStreamClient() *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: WsStreamClient{
			instType:                    "",
			apiType:                     WS_PUBLIC,
			commonSubMap:                NewMySyncMap[string, *Subscription[WsActionResult]](),
			tickerSubMap:                NewMySyncMap[string, *Subscription[WsTicker]](),
			candleSubMap:                NewMySyncMap[string, *Subscription[WsCandle]](),
			booksSubMap:                 NewMySyncMap[string, *Subscription[WsBooks]](),
			tradesSubMap:                NewMySyncMap[string, *Subscription[WsPublicTrade]](),
			fillSubMap:                  NewMySyncMap[string, *Subscription[WsFill]](),
			orderSubMap:                 NewMySyncMap[string, *Subscription[WsOrder]](),
			accountSubMap:               NewMySyncMap[string, *Subscription[WsAccount]](),
			equitySubMap:                NewMySyncMap[string, *Subscription[WsEquity]](),
			positionsSubMap:             NewMySyncMap[string, *Subscription[WsPositions]](),
			marginCrossAccountSubMap:    NewMySyncMap[string, *Subscription[WsMarginCrossAccountAssets]](),
			marginIsolatedAccountSubMap: NewMySyncMap[string, *Subscription[WsMarginIsolatedAccountAssets]](),
			marginCrossOrdersSubMap:     NewMySyncMap[string, *Subscription[WsMarginCrossOrders]](),
			marginIsolatedOrdersSubMap:  NewMySyncMap[string, *Subscription[WsMarginIsolatedOrders]](),
			waitSubResultMu:             &sync.Mutex{},
			writeMu:                     &sync.Mutex{},
			reSubscribeMu:               &sync.Mutex{},
		},
	}
}

type PrivateWsStreamClient struct {
	WsStreamClient
}

func (*MyBitget) NewPrivateWsStreamClient(apiType APIType) *PrivateWsStreamClient {

	return &PrivateWsStreamClient{
		WsStreamClient: WsStreamClient{
			instType:                    "",
			apiType:                     apiType,
			commonSubMap:                NewMySyncMap[string, *Subscription[WsActionResult]](),
			tickerSubMap:                NewMySyncMap[string, *Subscription[WsTicker]](),
			candleSubMap:                NewMySyncMap[string, *Subscription[WsCandle]](),
			booksSubMap:                 NewMySyncMap[string, *Subscription[WsBooks]](),
			tradesSubMap:                NewMySyncMap[string, *Subscription[WsPublicTrade]](),
			fillSubMap:                  NewMySyncMap[string, *Subscription[WsFill]](),
			orderSubMap:                 NewMySyncMap[string, *Subscription[WsOrder]](),
			accountSubMap:               NewMySyncMap[string, *Subscription[WsAccount]](),
			equitySubMap:                NewMySyncMap[string, *Subscription[WsEquity]](),
			positionsSubMap:             NewMySyncMap[string, *Subscription[WsPositions]](),
			marginCrossAccountSubMap:    NewMySyncMap[string, *Subscription[WsMarginCrossAccountAssets]](),
			marginIsolatedAccountSubMap: NewMySyncMap[string, *Subscription[WsMarginIsolatedAccountAssets]](),
			marginCrossOrdersSubMap:     NewMySyncMap[string, *Subscription[WsMarginCrossOrders]](),
			marginIsolatedOrdersSubMap:  NewMySyncMap[string, *Subscription[WsMarginIsolatedOrders]](),
			waitSubResultMu:             &sync.Mutex{},
			writeMu:                     &sync.Mutex{},
			reSubscribeMu:               &sync.Mutex{},
		},
	}
}

func subscribe[T any](ws *WsStreamClient, op string, args []WsSubscribeArg) (*Subscription[T], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitSubResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}

	ws.waitSubResultMu.Lock()

	subscribeReq := WsActionReq{
		Op:   op,
		Args: args,
	}
	data, err := json.Marshal(subscribeReq)
	if err != nil {
		ws.waitSubResultMu.Unlock()
		return nil, err
	}
	log.Debugf("send ws msg: %s", string(data))

	ws.writeMu.Lock()
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	ws.writeMu.Unlock()
	if err != nil {
		ws.waitSubResultMu.Unlock()
		return nil, err
	}

	id := atomic.AddInt64(&wsSubIDSeed, 1)
	result := &Subscription[T]{
		SubId:      id,
		Op:         op,
		Args:       args,
		resultChan: make(chan T, 50),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         ws,
	}
	return result, nil
}

// loginWithArg 发送登录帧并返回用于接收 login 应答的 Subscription（需配合 CatchLoginResult）。
func (ws *WsStreamClient) loginWithArg(arg WsLoginArg) (*Subscription[WsActionResult], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitSubResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}

	ws.waitSubResultMu.Lock()

	req := WsLoginReq{
		Op:   WS_LOGIN,
		Args: []WsLoginArg{arg},
	}
	data, err := json.Marshal(req)
	if err != nil {
		ws.waitSubResultMu.Unlock()
		return nil, err
	}
	log.Debugf("send ws msg: %s", string(data))

	ws.writeMu.Lock()
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	ws.writeMu.Unlock()
	if err != nil {
		ws.waitSubResultMu.Unlock()
		return nil, err
	}

	id := atomic.AddInt64(&wsSubIDSeed, 1)
	result := &Subscription[WsActionResult]{
		SubId:      id,
		Op:         WS_LOGIN,
		Args:       nil,
		resultChan: make(chan WsActionResult, 10),
		errChan:    make(chan error, 5),
		closeChan:  make(chan struct{}, 1),
		Ws:         ws,
	}
	return result, nil
}

// CatchLoginResult 等待登录应答（event=login 且 code=0）。
func (ws *WsStreamClient) CatchLoginResult(sub *Subscription[WsActionResult]) error {
	ws.waitSubResult = sub
	defer ws.DeferSub()

	select {
	case err := <-sub.ErrChan():
		return fmt.Errorf("ws login error: %w", err)
	case res := <-sub.ResultChan():
		if res.Code != 0 {
			return fmt.Errorf("ws login failed: code=%d msg=%s", res.Code, res.Msg)
		}
		if !strings.EqualFold(res.Event, "login") {
			return fmt.Errorf("ws login unexpected event: %s", res.Event)
		}
		log.Infof("websocket login success")
		return nil
	}
}

// Login 使用 REST 密钥对私有 WS 进行鉴权，须在 OpenConn 之后、订阅私有频道之前调用。
func (ws *WsStreamClient) Login(client *RestClient) error {
	if client == nil || client.c == nil {
		return fmt.Errorf("rest client is nil")
	}
	ws.client = client

	ts := time.Now().UnixMilli()
	tsStr := strconv.FormatInt(ts, 10)
	preSign := tsStr + GET + wsLoginSignPath
	log.Warnf("ws login: preSign: %s", preSign)
	sign := HmacSha256Base64(client.c.ApiSecret, preSign)

	arg := WsLoginArg{
		ApiKey:     client.c.ApiKey,
		Passphrase: client.c.Passphrase,
		Timestamp:  tsStr,
		Sign:       sign,
	}
	// 保存快照以便重连后再次登录
	ws.lastLogin = &WsLoginArg{
		ApiKey:     arg.ApiKey,
		Passphrase: arg.Passphrase,
		Timestamp:  arg.Timestamp,
		Sign:       arg.Sign,
	}
	sub, err := ws.loginWithArg(arg)
	if err != nil {
		return err
	}
	return ws.CatchLoginResult(sub)
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte, 100)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error, 20)
	}
	apiURL := handlerWsStreamRequestApi(ws.apiType)

	// 首次 OpenConn 负责启动 handleResult；后续重连只需要替换 conn，
	// 避免重复启动多个 handleResult goroutine 去消费同一组 channel。
	isFirst := ws.conn == nil

	conn, err := wsStreamServe(apiURL, ws.resultChan, ws.errChan, ws.writeMu)
	if err != nil {
		return err
	}
	ws.conn = conn
	ws.isClose = false
	if isFirst {
		ws.handleResult(ws.resultChan, ws.errChan)
	}
	log.Infof("open websocket success: %s", apiURL)
	return nil
}

func (ws *WsStreamClient) Close() error {
	ws.isClose = true
	if ws.conn != nil {
		_ = ws.conn.Close()
		ws.conn = nil
	}

	ws.sendWsCloseToAllSub()
	if ws.waitSubResult != nil {
		ws.waitSubResult.errChan <- fmt.Errorf("websocket is closed")
	}
	return nil
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result WsActionResult) {
	if ws.waitSubResult == nil {
		return
	}
	if result.Code != 0 {
		ws.waitSubResult.errChan <- fmt.Errorf("ws action error:[code:%d][msg:%s]", result.Code, result.Msg)
		return
	}
	ws.waitSubResult.resultChan <- result
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(args []WsSubscribeArg) {
	// 取消订阅时：
	// 同一个 Subscription 可能被多个 key 引用（如 *Multiple 订阅多个 symbol）。
	// 因此必须在 Delete(key) 后检查该 sub 是否仍被该频道 map 的其它 key 引用，
	// 只有“完全没有引用”时才向 CloseChan 发送关闭信号。
	tryClose := func(ch *chan struct{}) {
		if ch == nil || *ch == nil {
			return
		}
		select {
		case *ch <- struct{}{}:
		default:
		}
		*ch = nil
	}

	for _, arg := range args {
		data, _ := json.Marshal(&arg)
		key := string(data)
		if sub, ok := ws.tickerSubMap.Load(key); ok {
			ws.tickerSubMap.Delete(key)
			stillRef := false
			ws.tickerSubMap.Range(func(_ string, v *Subscription[WsTicker]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.candleSubMap.Load(key); ok {
			ws.candleSubMap.Delete(key)
			stillRef := false
			ws.candleSubMap.Range(func(_ string, v *Subscription[WsCandle]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.booksSubMap.Load(key); ok {
			ws.booksSubMap.Delete(key)
			stillRef := false
			ws.booksSubMap.Range(func(_ string, v *Subscription[WsBooks]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.tradesSubMap.Load(key); ok {
			ws.tradesSubMap.Delete(key)
			stillRef := false
			ws.tradesSubMap.Range(func(_ string, v *Subscription[WsPublicTrade]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.fillSubMap.Load(key); ok {
			ws.fillSubMap.Delete(key)
			stillRef := false
			ws.fillSubMap.Range(func(_ string, v *Subscription[WsFill]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.orderSubMap.Load(key); ok {
			ws.orderSubMap.Delete(key)
			stillRef := false
			ws.orderSubMap.Range(func(_ string, v *Subscription[WsOrder]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.accountSubMap.Load(key); ok {
			ws.accountSubMap.Delete(key)
			stillRef := false
			ws.accountSubMap.Range(func(_ string, v *Subscription[WsAccount]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.equitySubMap.Load(key); ok {
			ws.equitySubMap.Delete(key)
			stillRef := false
			ws.equitySubMap.Range(func(_ string, v *Subscription[WsEquity]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.positionsSubMap.Load(key); ok {
			ws.positionsSubMap.Delete(key)
			stillRef := false
			ws.positionsSubMap.Range(func(_ string, v *Subscription[WsPositions]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.marginCrossAccountSubMap.Load(key); ok {
			ws.marginCrossAccountSubMap.Delete(key)
			stillRef := false
			ws.marginCrossAccountSubMap.Range(func(_ string, v *Subscription[WsMarginCrossAccountAssets]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.marginIsolatedAccountSubMap.Load(key); ok {
			ws.marginIsolatedAccountSubMap.Delete(key)
			stillRef := false
			ws.marginIsolatedAccountSubMap.Range(func(_ string, v *Subscription[WsMarginIsolatedAccountAssets]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.marginCrossOrdersSubMap.Load(key); ok {
			ws.marginCrossOrdersSubMap.Delete(key)
			stillRef := false
			ws.marginCrossOrdersSubMap.Range(func(_ string, v *Subscription[WsMarginCrossOrders]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
		if sub, ok := ws.marginIsolatedOrdersSubMap.Load(key); ok {
			ws.marginIsolatedOrdersSubMap.Delete(key)
			stillRef := false
			ws.marginIsolatedOrdersSubMap.Range(func(_ string, v *Subscription[WsMarginIsolatedOrders]) bool {
				if v == sub {
					stillRef = true
					return false
				}
				return true
			})
			if !stillRef {
				tryClose(&sub.closeChan)
			}
		}
	}
}

func (ws *WsStreamClient) sendWsCloseToAllSub() {
	args := []WsSubscribeArg{}
	ws.commonSubMap.Range(func(key string, _ *Subscription[WsActionResult]) bool {
		arg := WsSubscribeArg{}
		if err := json.Unmarshal([]byte(key), &arg); err != nil {
			return true
		}
		args = append(args, arg)
		return true
	})
	ws.sendUnSubscribeSuccessToCloseChan(args)
}

// reSubscribeForReconnect 用于断线重连后，按 commonSubMap 中保存的订阅信息重新订阅。
// commonSubMap 的 key 是 arg JSON；同一个 Subscription 可能被多个 key 指向，因此需要去重。
func (ws *WsStreamClient) reSubscribeForReconnect() error {
	if ws == nil {
		return fmt.Errorf("ws is nil")
	}
	if ws.reSubscribeMu == nil {
		ws.reSubscribeMu = &sync.Mutex{}
	}

	ws.reSubscribeMu.Lock()
	defer ws.reSubscribeMu.Unlock()

	visited := map[*Subscription[WsActionResult]]bool{}
	snapshot := make([]*Subscription[WsActionResult], 0, 16)
	ws.commonSubMap.Range(func(_ string, sub *Subscription[WsActionResult]) bool {
		if sub == nil || visited[sub] {
			return true
		}
		visited[sub] = true
		snapshot = append(snapshot, sub)
		return true
	})

	isDoReSubscribe := map[int64]bool{}
	var wErr error
	for _, sub := range snapshot {
		if _, ok := isDoReSubscribe[sub.SubId]; ok {
			continue
		}

		reSub, err := subscribe[WsActionResult](ws, sub.Op, sub.Args)
		if err != nil {
			log.Error(err)
			wErr = err
			continue
		}
		err = ws.catchSubscribeResult(reSub)
		if err != nil {
			log.Error(err)
			wErr = err
			continue
		}
		log.Infof("reSubscribe Success: args:%v", reSub.Args)

		sub.SubId = reSub.SubId
		isDoReSubscribe[sub.SubId] = true
		time.Sleep(SUBSCRIBE_INTERVAL_TIME)
	}
	return wErr
}

func (ws *WsStreamClient) DeferSub() {
	if ws.waitSubResult == nil {
		ws.waitSubResultMu.Unlock()
		return
	}
	for _, arg := range ws.waitSubResult.Args {
		keyData, _ := json.Marshal(&arg)
		ws.commonSubMap.Store(string(keyData), ws.waitSubResult)
	}
	ws.waitSubResult = nil
	ws.waitSubResultMu.Unlock()
}

func (ws *WsStreamClient) catchSubscribeResult(sub *Subscription[WsActionResult]) error {
	ws.waitSubResult = sub
	defer sub.Ws.DeferSub()

	want := map[string]bool{}
	for _, arg := range sub.Args {
		keyData, _ := json.Marshal(&arg)
		want[string(keyData)] = false
	}

	for {
		select {
		case err := <-sub.ErrChan():
			return fmt.Errorf("ws subscribe action error: %w", err)
		case subResult := <-sub.ResultChan():
			keyData, _ := json.Marshal(subResult.Arg)
			if _, ok := want[string(keyData)]; ok {
				want[string(keyData)] = true
			}
			allDone := true
			for _, done := range want {
				if !done {
					allDone = false
					break
				}
			}
			if allDone {
				return nil
			}
		}
	}
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					log.Error("errChan is closed")
					return
				}
				log.Error(err)
				//错误处理 重连等
				//ws标记为非关闭 且返回错误包含EOF、close、reset时自动重连
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					// 1) 收到断线错误后立刻尝试 OpenConn，失败则循环重试
					// 2) OpenConn 成功后异步执行 Login（如需）和 reSubscribe
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(1500 * time.Millisecond)
						err = ws.OpenConn()
					}
					ws.AutoReConnectTimes += 1

					go func() {
						// 重新登陆（仅私有 WS 需要）
						if ws.lastLogin != nil && ws.client != nil {
							err = ws.Login(ws.client)
							for err != nil {
								time.Sleep(1500 * time.Millisecond)
								err = ws.Login(ws.client)
							}
						}

						// 重新订阅
						err = ws.reSubscribeForReconnect()
						if err != nil {
							log.Error(err)
						}
					}()
				} else {
					continue
				}
			case data, ok := <-resultChan:
				if !ok {
					return
				}

				msg := string(data)
				if msg == "pong" {
					continue
				}

				log.Debugf("ws msg: %s", msg)
				if strings.Contains(msg, "\"event\"") {
					result := WsActionResult{}
					if err := json.Unmarshal(data, &result); err != nil {
						log.Error(err)
						continue
					}
					ws.sendSubscribeResultToChan(result)
					continue
				}

				// Public
				if strings.Contains(msg, "\"topic\":\"ticker\"") {
					ticker, err := handleWsTicker(data)
					if err != nil {
						log.Error(err)
						continue
					}
					keyData, _ := json.Marshal(ticker.Arg)
					if sub, ok := ws.tickerSubMap.Load(string(keyData)); ok {
						sub.resultChan <- *ticker
					}
					continue
				}
				if strings.Contains(msg, "\"topic\":\"kline\"") {
					candle, err := handleWsCandle(data)
					if err != nil {
						log.Error(err)
						continue
					}
					keyData, _ := json.Marshal(candle.Arg)
					if sub, ok := ws.candleSubMap.Load(string(keyData)); ok {
						sub.resultChan <- *candle
					}
					continue
				}
				if strings.Contains(msg, "\"topic\":\"books") {
					books, err := handleWsBooks(data)
					if err != nil {
						log.Error(err)
						continue
					}
					if books == nil {
						continue
					}
					keyData, _ := json.Marshal(books.WsSubscribeArg)
					if sub, ok := ws.booksSubMap.Load(string(keyData)); ok {
						sub.resultChan <- *books
					}
					continue
				}
				if strings.Contains(msg, "\"topic\":\"publicTrade\"") {
					trade, err := handleWsPublicTrade(data)
					if err != nil {
						log.Error(err)
						continue
					}
					keyData, _ := json.Marshal(trade.Arg)
					if sub, ok := ws.tradesSubMap.Load(string(keyData)); ok {
						sub.resultChan <- *trade
					}
					continue
				}

				// Private
				if strings.Contains(msg, "\"channel\":\"fill\"") {
					fill, err := handleWsFill(data)
					if err != nil {
						log.Error(err)
						continue
					}
					keyData, _ := json.Marshal(fill.Arg)
					if sub, ok := ws.fillSubMap.Load(string(keyData)); ok {
						sub.resultChan <- *fill
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"orders\"") {
					switch ws.instType {
					case INST_TYPE_SPOT, INST_TYPE_USDT_FUTURES, INST_TYPE_COIN_FUTURES, INST_TYPE_USDC_FUTURES:
						order, err := handleWsOrder(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(order.Arg)
						if sub, ok := ws.orderSubMap.Load(string(keyData)); ok {
							sub.resultChan <- *order
						}
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"account\"") {
					switch ws.instType {
					case INST_TYPE_SPOT, INST_TYPE_USDT_FUTURES, INST_TYPE_COIN_FUTURES, INST_TYPE_USDC_FUTURES:
						account, err := handleWsAccount(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(account.Arg)
						if sub, ok := ws.accountSubMap.Load(string(keyData)); ok {
							sub.resultChan <- *account
						}
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"equity\"") {
					switch ws.instType {
					case INST_TYPE_USDT_FUTURES, INST_TYPE_COIN_FUTURES, INST_TYPE_USDC_FUTURES:
						equity, err := handleWsEquity(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(equity.Arg)
						if sub, ok := ws.equitySubMap.Load(string(keyData)); ok {
							sub.resultChan <- *equity
						}
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"positions\"") {
					switch ws.instType {
					case INST_TYPE_USDT_FUTURES, INST_TYPE_COIN_FUTURES, INST_TYPE_USDC_FUTURES:
						pos, err := handleWsPositions(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(pos.Arg)
						if sub, ok := ws.positionsSubMap.Load(string(keyData)); ok {
							sub.resultChan <- *pos
						}
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"account-crossed\"") {
					switch ws.instType {
					case INST_TYPE_MARGIN:
						acct, err := handleWsMarginCrossAccountAssets(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(acct.Arg)
						if sub, ok := ws.marginCrossAccountSubMap.Load(string(keyData)); ok {
							sub.resultChan <- *acct
						}
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"account-isolated\"") {
					switch ws.instType {
					case INST_TYPE_MARGIN:
						acct, err := handleWsMarginIsolatedAccountAssets(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(acct.Arg)
						if sub, ok := ws.marginIsolatedAccountSubMap.Load(string(keyData)); ok {
							sub.resultChan <- *acct
						}
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"orders-crossed\"") {
					switch ws.instType {
					case INST_TYPE_MARGIN:
						ord, err := handleWsMarginCrossOrders(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(ord.Arg)
						if sub, ok := ws.marginCrossOrdersSubMap.Load(string(keyData)); ok {
							sub.resultChan <- *ord
						}
					}
					continue
				}
				if strings.Contains(msg, "\"channel\":\"orders-isolated\"") {
					switch ws.instType {
					case INST_TYPE_MARGIN:
						ord, err := handleWsMarginIsolatedOrders(data)
						if err != nil {
							log.Error(err)
							continue
						}
						keyData, _ := json.Marshal(ord.Arg)
						if sub, ok := ws.marginIsolatedOrdersSubMap.Load(string(keyData)); ok {
							sub.resultChan <- *ord
						}
					}
					continue
				}
			}
		}
	}()
}

func (sub *Subscription[T]) Unsubscribe() error {
	unSub, err := subscribe[WsActionResult](sub.Ws, WS_UNSUBSCRIBE, sub.Args)
	if err != nil {
		return err
	}
	if err = sub.Ws.catchSubscribeResult(unSub); err != nil {
		return err
	}
	sub.Ws.sendUnSubscribeSuccessToCloseChan(unSub.Args)
	for _, arg := range unSub.Args {
		data, _ := json.Marshal(&arg)
		key := string(data)
		sub.Ws.commonSubMap.Delete(key)
		sub.Ws.fillSubMap.Delete(key)
		sub.Ws.orderSubMap.Delete(key)
		sub.Ws.accountSubMap.Delete(key)
		sub.Ws.equitySubMap.Delete(key)
		sub.Ws.positionsSubMap.Delete(key)
		sub.Ws.marginCrossAccountSubMap.Delete(key)
		sub.Ws.marginIsolatedAccountSubMap.Delete(key)
		sub.Ws.marginCrossOrdersSubMap.Delete(key)
		sub.Ws.marginIsolatedOrdersSubMap.Delete(key)
	}
	return nil
}

// 标准订阅方法
func wsStreamServe(api string, resultChan chan []byte, errChan chan error, writeMu *sync.Mutex) (*websocket.Conn, error) {
	dialer := websocket.DefaultDialer
	if WsUseProxy {
		proxy, _ := getBestProxyAndWeight()
		if proxy == nil {
			return nil, errors.New("no proxy available")
		}
		proxyURL, err := url.Parse(proxy.ProxyUrl)
		if err != nil {
			return nil, err
		}
		dialer = &websocket.Dialer{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	c, _, err := dialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout, writeMu)
		}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- message
		}
	}()
	return c, nil
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(apiType APIType) string {
	host := BITGET_API_WS_HOST_MAIN
	if NowNetType == TEST_NET {
		host = BITGET_API_WS_HOST_TEST
	}
	u := url.URL{
		Scheme: "wss",
		Host:   host,
		Path:   getWsApi(apiType),
	}
	return u.String()
}

// 获取数据流请求Path
func getWsApi(apiType APIType) string {
	switch apiType {
	case WS_PUBLIC:
		return BITGET_API_WS_PUBLIC
	case WS_CLASSIC:
		return BITGET_API_WS_CLASSIC_PRIVATE
	case WS_UTA:
		return BITGET_API_WS_UTA_PRIVATE
	default:
		log.Error("apiType error: ", apiType)
		return ""
	}
}

// 发送ping消息以检查连接稳定性
func keepAlive(c *websocket.Conn, timeout time.Duration, writeMu *sync.Mutex) {
	ticker := time.NewTicker(timeout)
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			writeMu.Lock()
			err := c.WriteMessage(websocket.TextMessage, []byte("ping"))
			writeMu.Unlock()
			if err != nil {
				return
			}
		}
	}()
}
