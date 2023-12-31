package model

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Comment struct {
	CommentId int64  `gorm:"primary_key"`
	UserId    int64  `gorm:"default:(-)"`
	VideoId   int64  `gorm:"default:(-)"`
	Content   string `gorm:"default:(-)"`
	CreateAt  time.Time
	DeletedAt gorm.DeletedAt
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

func (*CommentDao) DeleteCommentById(commentId int64) error {

	err := DB.Where("comment_id = ?", commentId).Delete(&Comment{}).Error
	if err != nil {
		fmt.Printf("删除失败", err)
	}

	return nil
}

func (*CommentDao) CreateComment(comment *Comment) (*Comment, error) {
	//和数据库进行操作
	result := DB.Create(&comment)

	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("model层的输出")
	//fmt.Println(comment.ID)
	return comment, nil
}

func (*CommentDao) QueryComment(videoId int64) ([]*Comment, error) {
	var comment []*Comment

	err := DB.Where("video_id = ?", videoId).Find(&comment).Error

	if err != nil {
		fmt.Println("查询Video列表失败")
		return nil, err
	}

	return comment, nil
}

func (*CommentDao) GetUserIdByCommentId(id int64) (int64, error) {
	comment := Comment{CommentId: id}
	result := DB.Where("comment_id = ?", id).First(&comment).Error
	return comment.UserId, result
}

var commentDao *CommentDao
var commentOnce sync.Once //单例模式，只生成一个commentDao实例，提高性能

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}
