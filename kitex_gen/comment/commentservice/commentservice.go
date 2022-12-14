// Code generated by Kitex v0.3.2. DO NOT EDIT.

package commentservice

import (
	"context"
	"dousheng/kitex_gen/comment"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateComment":        kitex.NewMethodInfo(createCommentHandler, newCommentServiceCreateCommentArgs, newCommentServiceCreateCommentResult, false),
		"DeleteComment":        kitex.NewMethodInfo(deleteCommentHandler, newCommentServiceDeleteCommentArgs, newCommentServiceDeleteCommentResult, false),
		"GetCommentsByVideoId": kitex.NewMethodInfo(getCommentsByVideoIdHandler, newCommentServiceGetCommentsByVideoIdArgs, newCommentServiceGetCommentsByVideoIdResult, false),
		"CountComment":         kitex.NewMethodInfo(countCommentHandler, newCommentServiceCountCommentArgs, newCommentServiceCountCommentResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func createCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCreateCommentArgs)
	realResult := result.(*comment.CommentServiceCreateCommentResult)
	success, err := handler.(comment.CommentService).CreateComment(ctx, realArg.CreateCommentReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCreateCommentArgs() interface{} {
	return comment.NewCommentServiceCreateCommentArgs()
}

func newCommentServiceCreateCommentResult() interface{} {
	return comment.NewCommentServiceCreateCommentResult()
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceDeleteCommentArgs)
	realResult := result.(*comment.CommentServiceDeleteCommentResult)
	success, err := handler.(comment.CommentService).DeleteComment(ctx, realArg.DeleteCommentReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceDeleteCommentArgs() interface{} {
	return comment.NewCommentServiceDeleteCommentArgs()
}

func newCommentServiceDeleteCommentResult() interface{} {
	return comment.NewCommentServiceDeleteCommentResult()
}

func getCommentsByVideoIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceGetCommentsByVideoIdArgs)
	realResult := result.(*comment.CommentServiceGetCommentsByVideoIdResult)
	success, err := handler.(comment.CommentService).GetCommentsByVideoId(ctx, realArg.GetCommentsByVideoIdReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceGetCommentsByVideoIdArgs() interface{} {
	return comment.NewCommentServiceGetCommentsByVideoIdArgs()
}

func newCommentServiceGetCommentsByVideoIdResult() interface{} {
	return comment.NewCommentServiceGetCommentsByVideoIdResult()
}

func countCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCountCommentArgs)
	realResult := result.(*comment.CommentServiceCountCommentResult)
	success, err := handler.(comment.CommentService).CountComment(ctx, realArg.CountCommentReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCountCommentArgs() interface{} {
	return comment.NewCommentServiceCountCommentArgs()
}

func newCommentServiceCountCommentResult() interface{} {
	return comment.NewCommentServiceCountCommentResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateComment(ctx context.Context, createCommentReq *comment.CreateCommentReq) (r *comment.CreateCommentResp, err error) {
	var _args comment.CommentServiceCreateCommentArgs
	_args.CreateCommentReq = createCommentReq
	var _result comment.CommentServiceCreateCommentResult
	if err = p.c.Call(ctx, "CreateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, deleteCommentReq *comment.DeleteCommentReq) (r *comment.DeleteCommentResp, err error) {
	var _args comment.CommentServiceDeleteCommentArgs
	_args.DeleteCommentReq = deleteCommentReq
	var _result comment.CommentServiceDeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentsByVideoId(ctx context.Context, getCommentsByVideoIdReq *comment.GetCommentsByVideoIdReq) (r *comment.GetCommentsByVideoIdResp, err error) {
	var _args comment.CommentServiceGetCommentsByVideoIdArgs
	_args.GetCommentsByVideoIdReq = getCommentsByVideoIdReq
	var _result comment.CommentServiceGetCommentsByVideoIdResult
	if err = p.c.Call(ctx, "GetCommentsByVideoId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountComment(ctx context.Context, countCommentReq *comment.CountCommentReq) (r *comment.CountCommentResp, err error) {
	var _args comment.CommentServiceCountCommentArgs
	_args.CountCommentReq = countCommentReq
	var _result comment.CommentServiceCountCommentResult
	if err = p.c.Call(ctx, "CountComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
