package utils

import (
	"context"
	"dousheng/cmd/api/controller/vo"
	"dousheng/cmd/api/rpc"
	"dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/favorite"
	"dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/video"
	"dousheng/pkg/conf"
	"dousheng/pkg/middleware"
	"fmt"
	"log"
)

func PackageVideo(userId int64, v *video.Video) (*vo.Video, error) {
	isFollow, err := rpc.IsFollow(context.Background(), &relation.IsFollowReq{
		UserId:       userId,
		TargetUserId: v.Author.Id,
	})
	if err != nil {
		return nil, err
	}
	followCount, err := rpc.CountFollows(context.Background(), &relation.CountFollowsReq{UserId: v.Author.Id})
	if err != nil {
		return nil, err
	}
	followerCount, err := rpc.CountFollowers(context.Background(), &relation.CountFollowersReq{UserId: v.Author.Id})
	if err != nil {
		return nil, err
	}
	isFavorite, err := rpc.IsFavorite(context.Background(), &favorite.IsFavoriteReq{
		UserId:  userId,
		VideoId: v.Id,
	})
	if err != nil {
		return nil, err
	}
	favoriteCount, err := rpc.CountFavorite(context.Background(), &favorite.CountFavoriteReq{VideoId: v.Id})
	if err != nil {
		return nil, err
	}
	commentCount, err := rpc.CountComment(context.Background(), &comment.CountCommentReq{VideoId: v.Id})
	if err != nil {
		return nil, err
	}
	return &vo.Video{
		Id: v.Id,
		Author: vo.User{
			Id:            v.Author.Id,
			Name:          v.Author.Name,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      isFollow,
		},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
	}, nil
}

func PackageUser(userId int64, targetUserId int64, targetUserName string) (*vo.User, error) {
	followCount, err := rpc.CountFollows(context.Background(), &relation.CountFollowsReq{UserId: targetUserId})
	if err != nil {
		return nil, err
	}
	followerCount, err := rpc.CountFollowers(context.Background(), &relation.CountFollowersReq{UserId: targetUserId})
	if err != nil {
		return nil, err
	}
	isFollow, err := rpc.IsFollow(context.Background(), &relation.IsFollowReq{
		UserId:       userId,
		TargetUserId: targetUserId,
	})
	if err != nil {
		return nil, err
	}
	return &vo.User{
		Id:            targetUserId,
		Name:          targetUserName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}, nil
}

func PublishGoroutine(finalName string, userId int64) {
	//保存视频到cdn
	err := middleware.Upload(finalName)
	if err != nil {
		log.Println(err.Error())
		fmt.Printf("[RedoVideo]:%s:%d\n", finalName, userId)
		return
	}
	//保存视频封面到cdn
	coverName := "cover_" + finalName + ".png"
	err = middleware.SavePicture(finalName, coverName)
	if err != nil {
		log.Println(err.Error())
		fmt.Printf("[RedoVideo]:%s:%d\n", finalName, userId)
		return
	}
	err = middleware.Upload(coverName)
	if err != nil {
		log.Println(err.Error())
		fmt.Printf("[RedoVideo]:%s:%d\n", finalName, userId)
		return
	}
	//将视频信息插入数据库
	err = rpc.CreateVideo(context.Background(), &video.CreateVideoReq{
		UserId:   userId,
		PlayUrl:  conf.CDN.Url + finalName,
		CoverUrl: conf.CDN.Url + coverName,
	})
	if err != nil {
		log.Println(err.Error())
		fmt.Printf("[RedoVideo]:%s:%d\n", finalName, userId)
		return
	}
}
