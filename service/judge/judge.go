package judge

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ricejson/oj-backend/common/consts"
	"github.com/ricejson/oj-backend/domain"
	"github.com/ricejson/oj-backend/repository/dao/question"
	"github.com/ricejson/oj-backend/repository/dao/submit"
	"github.com/ricejson/oj-backend/service/judge/sandbox"
	"github.com/ricejson/oj-backend/service/judge/stragety"
)

// JudgeService 判题服务
type JudgeService struct {
	questionDAO       *question.QuestionDAO
	questionSubmitDAO *submit.QuestionSubmitDAO
	sandbox           sandbox.CodeSandbox
}

func NewJudgeService(questionDAO *question.QuestionDAO, questionSubmitDAO *submit.QuestionSubmitDAO, sandbox sandbox.CodeSandbox) *JudgeService {
	return &JudgeService{
		questionDAO:       questionDAO,
		questionSubmitDAO: questionSubmitDAO,
		sandbox:           sandbox,
	}
}

func (s *JudgeService) DoJudge(ctx context.Context, submitId int64) (*domain.JudgeInfo, error) {
	// 1. 根据提交id查询提交记录和题目记录
	questionSubmit, err := s.questionSubmitDAO.FindById(ctx, submitId)
	if err != nil {
		return nil, err
	}
	questionId := questionSubmit.QuestionId
	question, err := s.questionDAO.FindById(ctx, questionId)
	if err != nil {
		return nil, err
	}
	// 2.判断判题状态是否为等待中
	if consts.QuestionSubmitStatus(questionSubmit.Status) != consts.QuestionSubmitStatusPending {
		return nil, errors.New("题目已在判题中")
	}
	// 3.将状态修改为运行中
	updateQuestionSubmit := submit.QuestionSubmit{
		Id:     questionSubmit.Id,
		Status: uint8(consts.QuestionSubmitStatusRunning),
	}
	_, err = s.questionSubmitDAO.UpdateById(ctx, updateQuestionSubmit)
	if err != nil {
		return nil, err
	}
	// 4. 交给判题机处理
	var cases *domain.Cases
	casesStr := question.Cases
	err = json.Unmarshal([]byte(casesStr), &cases)
	if err != nil {
		return nil, err
	}
	resp, executeErr := s.sandbox.ExecuteCode(ctx, &sandbox.ExecuteCodeRequest{
		Code:         questionSubmit.Code,
		Language:     questionSubmit.Language,
		InputSamples: cases.InputCases,
	})
	// 5. 完成判题信息更新
	var limitConfig *domain.QuestionLimitConfig
	limitConfigStr := question.LimitConfig
	err = json.Unmarshal([]byte(limitConfigStr), &limitConfig)
	if err != nil {
		return nil, err
	}
	updateQuestionSubmit = submit.QuestionSubmit{
		Id: questionSubmit.Id,
	}
	if executeErr != nil {
		updateQuestionSubmit.Status = uint8(consts.QuestionSubmitStatusFailed)
	} else {
		updateQuestionSubmit.Status = uint8(consts.QuestionSubmitStatusSuccess)
		judgeInfo := resp.JudgeInfo
		judgeInfo.Message = judgeMessage(cases, resp.OutputResults, judgeInfo, limitConfig, questionSubmit.Language)
		bytes, _ := json.Marshal(resp.JudgeInfo)
		updateQuestionSubmit.JudgeInfo = string(bytes)
	}
	_, _ = s.questionSubmitDAO.UpdateById(ctx, updateQuestionSubmit)
	questionSubmit, err = s.questionSubmitDAO.FindById(ctx, submitId)
	if err != nil {
		return nil, err
	}
	var judgeInfo *domain.JudgeInfo
	judgeInfoStr := questionSubmit.JudgeInfo
	err = json.Unmarshal([]byte(judgeInfoStr), &judgeInfo)
	if err != nil {
		return nil, err
	}
	return judgeInfo, nil
}

// judgeMessage 策略模式获取message执行信息
func judgeMessage(cases *domain.Cases, output []string, judgeInfo *domain.JudgeInfo, limitConfig *domain.QuestionLimitConfig, language string) string {
	ctx := stragety.NewStrategyContext(cases, output, judgeInfo, limitConfig)
	var s stragety.JudgeStrategy = stragety.NewDefaultStrategy(ctx)
	switch language {
	case consts.JavaLanguage:
		s = stragety.NewJavaLanguageStrategy(ctx)
	}
	return s.JudgeMessage()
}
