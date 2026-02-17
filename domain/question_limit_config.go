package domain

type QuestionLimitConfig struct {
	TimeLimit   int64 `json:"time_limit"`   // 时间限制
	MemoryLimit int64 `json:"memory_limit"` // 空间限制
}
