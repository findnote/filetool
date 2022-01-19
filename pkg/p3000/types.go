package p3000

import "errors"

// 错误码定义
var ErrNotFound = errors.New("sql: no rows in result set")
var ErrNoMoreFound = errors.New("sql: no more found in db")
var ErrNotValidP3000 = errors.New("conn: failed to connect to p3000")
var ErrCountNum = errors.New("sql: count failed in db")

type P3000Conn struct {
	conn string
}

type OrdersPostReq struct {
	ReceiveAt string `json:"rcvAt"`
	OrderType string `json:"fileName"`

	Data []interface{} `json:"data"`
}

// K-V
type HashKV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 设置实时值info
type SetHashInfo struct {
	Hash   string   `json:"hash"`
	Values []HashKV `json:"values"`
}

// 设置实时值请求
type SetHashReq struct {
	Tdrvs []SetHashInfo `json:"tdrvs"`
}

type SetHashResp struct {
	Success string `json:"success"`
	Errcode string `json:"errorCode"`
	ErrMsg  string `json:"errorMessage"`
}
