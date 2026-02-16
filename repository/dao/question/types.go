package question

import "time"

type Question struct {
	Id          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement;comment:题目唯一标识" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(64);not null;comment:题目标题" json:"title"`
	Content     string    `gorm:"column:content;type:text;null;comment:题目内容" json:"content"`
	Tags        string    `gorm:"column:tags;type:varchar(64);null;comment:题目标签JSON字符串（比如简单、二叉树）" json:"tags"`
	Cases       string    `gorm:"column:cases;type:varchar(128);null;comment:用例JSON字符串（比如输入用例、输出用例）" json:"cases"`
	LimitConfig string    `gorm:"column:limit_config;type:varchar(128);null;comment:限制配置JSON字符串（比如时间限制，内存限制）" json:"limit_config"`
	Answer      string    `gorm:"column:answer;type:text;null;comment:题解" json:"answer"`
	UserId      int64     `gorm:"column:user_id;type:bigint;not null;comment:创建用户id" json:"user_id"`
	CreateAt    time.Time `gorm:"column:create_at;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"`
	UpdateAt    time.Time `gorm:"column:update_at;type:datetime;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:更新时间" json:"update_at"`
	IsDeleted   uint8     `gorm:"column:is_deleted;type:tinyint;default:0;comment:是否被删除（逻辑删除）" json:"is_deleted"`
}

// TableName 指定表名
func (q *Question) TableName() string {
	return "t_question"
}
