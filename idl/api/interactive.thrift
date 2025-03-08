namespace go api.interactive

struct LikeOperationReq {
    1: required i64 userId,
    2: required i64 postId,
}

struct LikeOperationResp {
    1: required bool success,
}

struct LikeListReq {
    1: required i64 userId,
}

struct LikeListResp {
    1: required list<i64> postIdList,
}

struct CommentReq {
    1: required i64 userId,
    2: required i64 postId,
    3: required string content,
}

struct CommentResp {
    1: required bool success,
}

struct CommentListReq {
    1: required i64 postId,
}

struct CommentListResp {
    1: required list<string> contentList,
}

struct DeleteCommentReq {
    1: required i64 userId,
    2: required i64 commentId,
}

struct DeleteCommentResp {
    1: required bool success,
}

service Interactive {
    LikeOperationResp likePost(1: LikeOperationReq req)(api.post="api/v1/interactive/likePost")
    LikeListResp getLikedPostList(1: LikeListReq req)(api.get="api/v1/interactive/getLikedPostList")
    CommentResp comment(1: CommentReq req)(api.post="api/v1/interactive/comment")
    CommentListResp getCommentList(1: CommentListReq req)(api.get="api/v1/interactive/getCommentList")
    DeleteCommentResp deleteComment(1: DeleteCommentReq req)(api.delete="api/v1/interactive/deleteComment")
}