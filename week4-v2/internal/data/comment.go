package data

import (
	"context"
	"time"
	"week4-v2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type Comment struct {
	ID           int32  `gorm :"type:int(11)"`
	Content      string `gorm:"longtext"`
	ContentType  int8   `gorm:"type:int(1)"`
	ContentTitle string `gorm:"type:varchar(225)"`
	ContentID    int32  `gorm:"type:int(11)"`
	Status       int8   `gorm:"type:int(1)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Like         int64 `gorm:"type:int(11)"`
}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{data: data, log: log.NewHelper(logger)}
}

func (repo *commentRepo) CreateComment(ctx context.Context, co *biz.Comment) (int32, error) {
	repo.data.db.AutoMigrate(&Comment{})
	comment := &Comment{Content: co.Content, ContentType: co.ContentType,
		ContentTitle: co.ContentTitle, ContentID: co.ContentId}
	result := repo.data.db.Create(comment)
	if result.Error != nil {
		return 0, result.Error
	}
	return comment.ID, nil
}

func (repo *commentRepo) GetComment(ctx context.Context, id int32) (*biz.Comment, error) {
	co := &Comment{ID: id}
	res := repo.data.db.First(co)
	if res.Error != nil {
		return &biz.Comment{}, res.Error
	}
	return &biz.Comment{
		ID:           co.ID,
		Content:      co.Content,
		ContentId:    co.ContentID,
		ContentType:  co.ContentType,
		ContentTitle: co.ContentTitle,
		Status:       co.Status,
		CreatedAt:    co.CreatedAt,
		UpdatedAt:    co.UpdatedAt,
	}, nil
}
