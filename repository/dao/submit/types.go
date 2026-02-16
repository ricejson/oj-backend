package submit

import (
	"time"
)

// QuestionSubmit 题目提交表
type QuestionSubmit struct {
	Id         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement;comment:题目提交唯一标识" json:"id"`
	QuestionId int64     `gorm:"column:question_id;type:bigint;not null;index:idx_question_id;comment:题目id" json:"question_id"`
	Code       string    `gorm:"column:code;type:text;not null;comment:提交代码" json:"code"`
	Language   string    `gorm:"column:language;type:varchar(64);not null;comment:语言（JAVA、Go等）" json:"language"`
	Status     uint8     `gorm:"column:status;type:tinyint;default:0;comment:判题结果 1-等待中 2-失败 3-成功 4-系统错误" json:"status"`
	JudgeInfo  string    `gorm:"column:judge_info;type:varchar(128);comment:判题信息JSON字符串" json:"judge_info"`
	UserId     int64     `gorm:"column:user_id;not null;index:idx_user_id;comment:提交用户id" json:"user_id"`
	CreatedAt  time.Time `gorm:"column:create_at;type:datetime;default:CURRENT_TIMESTAMP;index:idx_create_at;comment:创建时间" json:"create_at"`
	UpdatedAt  time.Time `gorm:"column:update_at;type:datetime;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:更新时间" json:"update_at"`
	IsDelete   int8      `gorm:"column:is_delete;type:tinyint;default:0;comment:是否被删除（逻辑删除）0-未删除，1-删除" json:"is_delete"`
}

// TableName 指定表名
func (q *QuestionSubmit) TableName() string {
	return "t_question_submit"
}
