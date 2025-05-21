namespace go comment

include "model.thrift"

/*对视频进行评论
@param videoId 视频ID
@param content 评论内容
@param userId 用户ID
*/
struct CommentVideoReq {
    1: required string videoId,
    2: required string content,
}

struct CommentVideoResp {
    1:required string commentId
}

/*对评论进行回复
@param commentId 评论ID
@param content 回复内容
@param userId 用户ID
*/
struct ReplyCommentReq {
    1: required string commentId,
    2: required string content,
}

struct ReplyCommentResp {
    1:required string commentId
}

//获取评论列表
struct GetCommentListReq {
    1: required string Id,
    2: required i64 page,
    3: required i64 pageSize
    4: required i64 CommentTpye//1:视频评论 2:回复评论
}

struct GetCommentListResp {
    1:required list<model.Comment> commentList
}

struct DeleteCommentReq {
    1: required string commentId
}

struct DeleteCommentResp {
}

service CommentService {
    CommentVideoResp CommentVideo(1:CommentVideoReq req),
    ReplyCommentResp ReplyComment(1:ReplyCommentReq req),
    GetCommentListResp GetCommentList(1:GetCommentListReq req),
    DeleteCommentResp DeleteComment(1:DeleteCommentReq req)
}
