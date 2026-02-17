package stragety

import "github.com/ricejson/oj-backend/domain"

type StragetyContext struct {
	cases       *domain.Cases
	output      []string
	judgeInfo   *domain.JudgeInfo
	limitConfig *domain.QuestionLimitConfig
}

func NewStrategyContext(cases *domain.Cases, output []string, judgeInfo *domain.JudgeInfo, limitConfig *domain.QuestionLimitConfig) *StragetyContext {
	return &StragetyContext{
		cases:       cases,
		output:      output,
		judgeInfo:   judgeInfo,
		limitConfig: limitConfig,
	}
}
