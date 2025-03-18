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
    2:required i64 commentId
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
    2:required i64 commentId
}

/*获取视频评论列表
*/
struct GetCommentListReq {
    1: required i64 videoId,
    2: required i64 page,
    3: required i64 pageSize
    4: required i64 CommentTpye//1:视频评论 2:回复评论
}

struct GetCommentListResp {
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

service CommentService {
    CommentVideoResp CommentVideo(1:CommentVideoReq req),
    ReplyCommentResp ReplyComment(1:ReplyCommentReq req),
    GetCommentListResp GetCommentList(1:GetCommentListReq req),
    DeleteCommentResp DeleteComment(1:DeleteCommentReq req)
}