package submit

import (
	"context"

	"gorm.io/gorm"
)

type QuestionSubmitDAO struct {
	db *gorm.DB
}

func NewQuestionSubmitDAO(db *gorm.DB) *QuestionSubmitDAO {
	return &QuestionSubmitDAO{
		db: db,
	}
}

// Insert 插入一条题目提交记录
func (dao *QuestionSubmitDAO) Insert(ctx context.Context, questionSubmit QuestionSubmit) (int64, error) {
	err := dao.db.WithContext(ctx).Create(&questionSubmit).Error
	if err != nil {
		return 0, err
	}
	return questionSubmit.Id, nil
}

// FindById 查询id查询题目
func (dao *QuestionSubmitDAO) FindById(ctx context.Context, id int64) (QuestionSubmit, error) {
	var questionSubmit QuestionSubmit
	err := dao.db.WithContext(ctx).
		Model(&QuestionSubmit{}).
		Where("id = ?", id).
		Find(&questionSubmit).Error
	return questionSubmit, err
}
