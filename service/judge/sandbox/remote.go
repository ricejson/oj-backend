package sandbox

import (
	"context"
)

// RemoteCodeSandbox 远程代码沙箱实现
type RemoteCodeSandbox struct {
}

func NewRemoteCodeSandbox() *RemoteCodeSandbox {
	return &RemoteCodeSandbox{}
}

func (s *RemoteCodeSandbox) ExecuteCode(ctx context.Context, req *ExecuteCodeRequest) (*ExecuteCodeResponse, error) {
	return &ExecuteCodeResponse{
		nil, nil,
	}, nil
}
