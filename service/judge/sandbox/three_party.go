package sandbox

import (
	"context"
)

// ThreePartyCodeSandbox 三方代码沙箱实现
type ThreePartyCodeSandbox struct {
}

func NewThreePartyCodeSandbox() *ThreePartyCodeSandbox {
	return &ThreePartyCodeSandbox{}
}

func (s *ThreePartyCodeSandbox) ExecuteCode(ctx context.Context, req *ExecuteCodeRequest) (*ExecuteCodeResponse, error) {
	return &ExecuteCodeResponse{
		nil, nil,
	}, nil
}
