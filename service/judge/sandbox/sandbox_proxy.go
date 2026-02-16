package sandbox

import (
	"context"

	"github.com/ricejson/gotool/logx"
)

// CodeSandboxProxy 代码沙箱代理实现（也可以叫做装饰器）
type CodeSandboxProxy struct {
	logger      logx.Logger
	codeSandbox CodeSandbox
}

func NewCodeSandboxProxy(codeSandbox CodeSandbox, logger logx.Logger) *CodeSandboxProxy {
	return &CodeSandboxProxy{
		codeSandbox: codeSandbox,
		logger:      logger,
	}
}

func (p *CodeSandboxProxy) ExecuteCode(ctx context.Context, req *ExecuteCodeRequest) (*ExecuteCodeResponse, error) {
	// 记录请求日志
	p.logger.Info("ExecuteCode req:", logx.Field{Key: "req", Value: req})
	resp, err := p.codeSandbox.ExecuteCode(ctx, req)
	if err != nil {
		p.logger.Warn("ExecuteCode req fail:", logx.Error(err))
		return nil, err
	}
	p.logger.Info("ExecuteCode resp:", logx.Field{Key: "resp", Value: resp})
	return resp, nil
}
