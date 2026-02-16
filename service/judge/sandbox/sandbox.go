package sandbox

import (
	"context"

	"github.com/ricejson/oj-backend/domain"
)

const (
	CodeSandBoxTypeExample    = "example"
	CodeSandBoxTypeRemote     = "remote"
	CodeSandBoxTypeThreeParty = "three-party"
)

type ExecuteCodeRequest struct {
	Code         string   `json:"code"`          // 代码
	Language     string   `json:"language"`      // 编程语言
	InputSamples []string `json:"input_samples"` // 输入样例
}

type ExecuteCodeResponse struct {
	OutputResults []string          `json:"output_results"` // 输出结果
	JudgeInfo     *domain.JudgeInfo `json:"judge_info"`     // 判题信息
}

// CodeSandbox 代码沙箱
type CodeSandbox interface {
	ExecuteCode(ctx context.Context, req *ExecuteCodeRequest) (*ExecuteCodeResponse, error)
}

// NewInstance 工厂方法获取代码沙箱实例
func NewInstance(typ string) CodeSandbox {
	switch typ {
	case CodeSandBoxTypeExample:
		return NewExampleCodeSandbox()
	case CodeSandBoxTypeRemote:
		return NewRemoteCodeSandbox()
	case CodeSandBoxTypeThreeParty:
		return NewThreePartyCodeSandbox()
	}
	// 默认用样例代码沙箱
	return NewExampleCodeSandbox()
}
