// Code generated by goctl. DO NOT EDIT.
package types

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
