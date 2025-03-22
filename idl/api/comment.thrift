namespace go api.comment

struct CommentVideoReq{
    1: required i64 video_id,
    2: required string content,
    3: required i64 user_id,
}

struct CommentVideoResp{
    1: required bool success,
    2:required i64 comment_id,
}

struct ReplyCommentReq{
    1: required i64 comment_id,
    2: required string content,
    3: required i64 user_id,
}

struct ReplyCommentResp{
    1: required bool success,
    2:required i64 comment_id,
}

struct GetVideoCommentListReq{
    1: required i64 video_id,
    2: required i64 page_num,
    3: required i64 page_size,
}

struct Comment{
    1: required i64 comment_id
    2: required i64 user_id
    3: required string content
    4: required i64 create_time
}

struct GetVideoCommentListResp{
    1: required bool success,
    2: required list<Comment> comment_list,
}

struct GetCommentReplyListReq{
    1: required i64 comment_id,
    2: required i64 page_num,
    3: required i64 page_size,
}

struct GetCommentReplyListResp{
    1: required bool success,
    2: required list<Comment> comment_list,
}

struct DeleteCommentReq{
    1: required i64 comment_id,
    2: required i64 user_id,
}

struct DeleteCommentResp{
    1: required bool success,
}



service CommentService{
    CommentVideoResp CommentVideo(1:CommentVideoReq req)(api.post="api/v1/comment/commentvideo")
    ReplyCommentResp ReplyComment(1:ReplyCommentReq req)(api.post="api/v1/comment/commentreply")
    GetVideoCommentListResp GetVideoCommentList(1:GetVideoCommentListReq req)(api.get="api/v1/comment/videocommentlist")
    GetCommentReplyListResp GetCommentReplyList(1:GetCommentReplyListReq req)(api.get="api/v1/comment/commentreplylist")
    DeleteCommentResp DeleteComment(1:DeleteCommentReq req)(api.delete="api/v1/comment/deletecomment")
}