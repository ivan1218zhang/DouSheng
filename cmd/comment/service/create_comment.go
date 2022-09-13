package service

import (
	"context"
	"dousheng/cmd/comment/dao"
	"dousheng/cmd/comment/repository"
)

type CreateCommentServiceImpl struct {
	ctx context.Context
}

func NewCreateCommentServiceImpl(ctx context.Context) *CreateCommentServiceImpl {
	return &CreateCommentServiceImpl{ctx: ctx}
}

func (receiver *CreateCommentServiceImpl) CreateComment(userId int64, videoId int64, content string) (int64, error) {
	commentDao := dao.GetCommentDaoInstance()
	//保存评论
	err := commentDao.Create(repository.Comment{
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	})
	if err != nil {
		return 0, err
	}
	//查找保存的评论的最后一条
	return commentDao.GetLatestCommentIdByUserId(userId)
}
