package dao

import (
	"context"
	"github.com/abuziming/dousheng_demo/pkg/errno"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func (Comment) TableName() string {
	return "comments"
}

// NewComment creates a new Comment
func NewComment(ctx context.Context, comment *Comment) error {
	err := Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 1. 新增评论数据
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 comment count
		res := tx.Model(&Video{}).Where("ID = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

// DelComment deletes a comment from the database.
func DelComment(ctx context.Context, commentID int64, vid int64) error {
	err := Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment := new(Comment)
		if err := tx.First(&comment, commentID).Error; err != nil {
			return err
		}

		// 1. 删除评论数据
		// 因为 Comment中包含了gorm.Model所以拥有软删除能力
		// 而tx.Unscoped().Delete()将永久删除记录
		err := tx.Unscoped().Delete(&comment).Error
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 comment count
		res := tx.Model(&Video{}).Where("ID = ?", vid).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

// GetVideoComments returns a list of video comments.
func GetVideoComments(ctx context.Context, vid string) ([]*Comment, error) {
	var comments []*Comment
	num, _ := strconv.ParseInt(vid, 10, 64)
	log.Println(num)
	err := Db.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoId: num}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func GetUserIDByVideoID(videoID int64) (int64, error) {
	var video Video
	if err := Db.Select("user_id").First(&video, videoID).Error; err != nil {
		return 0, err
	}
	return video.UserId, nil
}
