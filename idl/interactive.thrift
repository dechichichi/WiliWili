namespace go interactive


/* LikeOperationReq点赞操作请求
* @param userId 用户id
* @param postId 帖子id
*/
struct LikeOperationReq {
    1: required i64 userId,
    2: required i64 postId,
}

/* LikeOperationResp点赞操作响应
* @param success 操作是否成功
*/
struct LikeOperationResp {
    1: required bool success, 
}

/* LikeListReq获取点赞列表请求
* @param userId 用户id
*/  
struct LikeListReq {
    1: required i64 userId,
}
/* LikeListResp获取点赞列表响应
* @param postIdList 帖子id列表
*/
struct LikeListResp {
    1: required list<i64> postIdList,
}

/* CommentReq评论请求
* @param userId 用户id
* @param postId 帖子id
* @param content 评论内容
*/
struct CommentReq {
    1: required i64 userId,
    2: required i64 postId,
    3: required string content,
}

/* CommentResp评论响应
* @param success 操作是否成功
*/
struct CommentResp {
    1: required bool success,
}

/* CommentListReq获取评论列表请求
* @param postId 帖子id
*/
struct CommentListReq {
    1: required i64 postId,
}

/* CommentListResp获取评论列表响应
* @param commentList 评论列表
*/
struct CommentListResp {
    1: required list<string> commentList,
}

/* DeleteCommentReq删除评论请求
* @param userId 用户id
* @param commentId 评论id
*/
struct DeleteCommentReq {
    1: required i64 userId,
    2: required i64 commentId,
}

/* DeleteCommentResp删除评论响应
* @param success 操作是否成功
* @param error 错误信息
*/
struct DeleteCommentResp {
    1: required bool success,
    2: optional string error,
}



service Interactive {
    LikeOperationResp likeOperation(1: LikeOperationReq req),
    LikeListResp likeList(1: LikeListReq req),
    CommentResp comment(1: CommentReq req),
    CommentListResp commentList(1: CommentListReq req),
    DeleteCommentResp deleteComment(1: DeleteCommentReq req),
}