package consts

type QuestionSubmitStatus uint8

const (
	QuestionSubmitStatusPending QuestionSubmitStatus = 1 // 等待中
	QuestionSubmitStatusRunning QuestionSubmitStatus = 2 // 运行中
	QuestionSubmitStatusFailed  QuestionSubmitStatus = 3 // 失败
	QuestionSubmitStatusSuccess QuestionSubmitStatus = 4 // 成功
	QuestionSubmitStatusError   QuestionSubmitStatus = 5 // 系统错误
)
