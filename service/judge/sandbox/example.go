package sandbox

import (
	"context"

	"github.com/ricejson/oj-backend/domain"
)

const (
	DefaultMemory = 10
	DefaultTime   = 20
)

// ExampleCodeSandbox 样例代码沙箱实现
type ExampleCodeSandbox struct {
}

func NewExampleCodeSandbox() *ExampleCodeSandbox {
	return &ExampleCodeSandbox{}
}

func (s *ExampleCodeSandbox) ExecuteCode(ctx context.Context, req *ExecuteCodeRequest) (*ExecuteCodeResponse, error) {
	return &ExecuteCodeResponse{
		OutputResults: req.InputSamples,
		JudgeInfo: &domain.JudgeInfo{
			Memory: DefaultMemory,
			Time:   DefaultTime,
		},
	}, nil
}
