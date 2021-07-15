package hms

type TokenCache interface {
	// 获取token信息，若为空则获取新的token
	TokenCache() (*TokenInfo, error)
}

type TokenInfo struct {
	Token           string `json:"token"`
	TokenExpireTime int64  `json:"token_expire_time"`
	KeyExpire       int64  `json:"key_expire"`
}
