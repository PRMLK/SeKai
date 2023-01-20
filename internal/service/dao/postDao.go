package dao

import (
	"SeKai/internal/logger"
	"SeKai/internal/model"
	"SeKai/internal/model/param"
	"SeKai/internal/util"
	"errors"
	"gorm.io/gorm"
)

func NewPost(newPostParam *param.PostParam, userId uint) (uint, error) {
	post := model.Post{
		Title:         newPostParam.Title,
		Author:        userId,
		Content:       newPostParam.Content,
		PostStatus:    newPostParam.PostStatus,
		CommentStatus: newPostParam.CommentStatus,
	}
	if err := util.Datebase.Create(&post).Error; err != nil {
		logger.ServerLogger.Error(err.Error())
		return 0, errors.New("系统错误")
	} else {
		return post.ID, nil
	}
}

func GetPost(id uint) (post model.Post, err error) {
	if err = util.Datebase.Where(&model.Post{Model: gorm.Model{ID: id}}).First(&post).Error; err != nil {
		logger.ServerLogger.Debug(err)
		return model.Post{}, errors.New("找不到文章")
	}
	return post, nil
}

func UpdatePost(postId uint, postParam *param.PostParam, userId uint) error {
	if err := util.Datebase.Model(&model.Post{
		Model: gorm.Model{
			ID: postId,
		},
	}).Updates(&model.Post{
		Title:         postParam.Title,
		Author:        userId,
		Content:       postParam.Content,
		PostStatus:    postParam.PostStatus,
		CommentStatus: postParam.CommentStatus,
	}).Error; err != nil {
		logger.ServerLogger.Debug(err)
		return errors.New("系统错误")
	}
	return nil
}
