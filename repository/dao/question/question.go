package question

import (
	"context"

	"gorm.io/gorm"
)

type QuestionDAO struct {
	db *gorm.DB
}

func NewQuestionDAO(db *gorm.DB) *QuestionDAO {
	return &QuestionDAO{
		db: db,
	}
}

// Insert 插入一条题目记录
func (dao *QuestionDAO) Insert(ctx context.Context, question Question) (int64, error) {
	err := dao.db.WithContext(ctx).Create(&question).Error
	if err != nil {
		return 0, err
	}
	return question.Id, nil
}

// FindById 查询id查询题目
func (dao *QuestionDAO) FindById(ctx context.Context, id int64) (Question, error) {
	var question Question
	err := dao.db.WithContext(ctx).
		Model(&Question{}).
		Where("id = ?", id).
		Find(&question).Error
	return question, err
}
