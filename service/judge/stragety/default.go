package stragety

import "github.com/ricejson/oj-backend/common/consts"

type DefaultStrategy struct {
	ctx *StragetyContext
}

func NewDefaultStrategy(ctx *StragetyContext) *DefaultStrategy {
	return &DefaultStrategy{
		ctx: ctx,
	}
}

func (s *DefaultStrategy) JudgeMessage() string {
	cases := s.ctx.cases
	output := s.ctx.output
	judgeInfo := s.ctx.judgeInfo
	limitConfig := s.ctx.limitConfig
	// 1. 校验输入和输出样例个数是否一致
	inputCases := cases.InputCases
	outputCases := cases.OutputCases
	if len(inputCases) != len(output) {
		return consts.JudgeMessageWrongAnswer
	}
	// 2. 校验每组输入输出
	for i, actual := range output {
		expect := outputCases[i]
		if expect != actual {
			return consts.JudgeMessageWrongAnswer
		}
	}
	// 3. 校验内存限制
	if judgeInfo.Memory > limitConfig.MemoryLimit {
		return consts.JudgeMessageMemoryLimit
	}
	// 4. 校验时间限制
	if judgeInfo.Time > limitConfig.TimeLimit {
		return consts.JudgeMessageTimeLimit
	}
	return consts.JudgeMessageAccept
}
