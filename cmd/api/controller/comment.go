package controller

import (
	"context"
	"dousheng/cmd/api/controller/vo"
	"dousheng/cmd/api/rpc"
	"dousheng/cmd/api/utils"
	"dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	vo.Response
	CommentList []vo.Comment `json:"comment_list,omitempty"`
}
type CommentActionResponse struct {
	vo.Response
	Comment vo.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoIdData := c.Query("video_id")
	actionType := c.Query("action_type")
	videoId, err := strconv.ParseInt(videoIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "video_id error"})
		return
	}
	//鉴权
	userId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	//评论
	comm := vo.Comment{}
	if actionType == "1" {
		content := c.Query("comment_text")
		comm.Id, err = rpc.CreateComment(context.Background(), &comment.CreateCommentReq{
			UserId:  userId,
			VideoId: videoId,
			Content: content,
		})
		comm.Content = content
		u, err := rpc.GetUserById(context.Background(), &user.GetUserByIdReq{Id: userId})
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		u2, err := utils.PackageUser(u.Id, u.Id, u.Name)
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		comm.User = vo.User{
			Id:            u2.Id,
			Name:          u2.Name,
			FollowCount:   u2.FollowCount,
			FollowerCount: u2.FollowerCount,
			IsFollow:      u2.IsFollow,
		}
		var timeLayoutStr = "01-02"
		comm.CreateDate = time.Now().Format(timeLayoutStr)
	}
	//删除评论
	if actionType == "2" {
		commentIdData := c.Query("comment_id")
		commentId, err := strconv.ParseInt(commentIdData, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "comment_id error"})
			return
		}
		err = rpc.DeleteComment(context.Background(), &comment.DeleteCommentReq{
			VideoId:   videoId,
			CommentId: commentId,
		})
	}
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, CommentActionResponse{
		Response: vo.Response{StatusCode: 0},
		Comment:  comm,
	})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	//鉴权
	userId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		userId = -1
	}
	videoIdData := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "video_id error"})
		return
	}
	comments, err := rpc.GetCommentsByVideoId(context.Background(), &comment.GetCommentsByVideoIdReq{VideoId: videoId})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	result := make([]vo.Comment, len(comments))
	for i, co := range comments {
		u, err := utils.PackageUser(userId, co.User.Id, co.User.Name)
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		result[i] = vo.Comment{
			Id:         co.Id,
			User:       *u,
			Content:    co.Content,
			CreateDate: co.CreateDate,
		}
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    vo.Response{StatusCode: 0},
		CommentList: result,
	})
}
