namespace go comment


include "model.thrift"

/*对视频进行评论
@param videoId 视频ID
@param content 评论内容
@param userId 用户ID
*/
struct CommentVideoReq {
    1: required i64 videoId,
    2: required string content,
    3: required i64 userId
}

struct CommentVideoResp {
    1:required model.BaseResp baseResp
}

/*对评论进行回复
@param commentId 评论ID
@param content 回复内容
@param userId 用户ID
*/
struct ReplyCommentReq {
    1: required i64 commentId,
    2: required string content,
    3: required i64 userId
}

struct ReplyCommentResp {
    1:required model.BaseResp baseResp
}

/*获取视频评论列表
*/
struct GetVideoCommentListReq {
    1: required i64 videoId,
    2: required i32 page,
    3: required i32 pageSize
}

struct GetVideoCommentListResp {
    1:required model.BaseResp baseResp,
    2:required list<model.Comment> commentList
}

/*获取评论回复列表
*/
struct GetCommentReplyListReq {
    1: required i64 commentId,
    2: required i32 page,
    3: required i32 pageSize
}

struct GetCommentReplyListResp {
    1:required model.BaseResp baseResp,
    2:required list<model.Comment> commentList
}

struct DeleteCommentReq {
    1: required i64 commentId,
    2: required i64 userId
}

struct DeleteCommentResp {
    1:required model.BaseResp baseResp
}

service Comment{
    CommentVideoResp CommentVideo(1:CommentVideoReq req),
    ReplyCommentResp ReplyComment(1:ReplyCommentReq req),
    GetVideoCommentListResp GetVideoCommentList(1:GetVideoCommentListReq req),
    GetCommentReplyListResp GetCommentReplyList(1:GetCommentReplyListReq req),
    DeleteCommentResp DeleteComment(1:DeleteCommentReq req)
}