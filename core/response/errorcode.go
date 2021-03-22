package response

const (
	COMMON  = 9999
	SUCCESS = COMMON + iota // 成功
	UNKNOWN                 // 未知错误
	FAILED                  // 失败
)
