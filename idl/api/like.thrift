namespace go api.like

struct LikeCommentReq{
    1: required i64 comment_id,
    2: required i64 user_id,
    3: required bool is_like,
}

struct LikeCommentResp{
    1: required bool success,
}

struct LikeVideoReq{
    1: required i64 video_id,
    2: required i64 user_id,
    3: required bool is_like,
}

struct LikeVideoResp{
    1: required bool success,
}

struct CommentLikeNumReq{
    1: required i64 comment_id,
}

struct CommentLikeNumResp{
    1: required i32 like_num,
}

struct VideoLikeNumReq{
    1: required i64 video_id,
}

struct VideoLikeNumResp{
    1: required i32 like_num,
}



service LikeService{
    LikeCommentResp LikeComment(1: LikeCommentReq req)(api.post="api/v1/like/likecomment")
    LikeVideoResp LikeVideo(1: LikeVideoReq req)(api.post="api/v1/like/likevideo")
    CommentLikeNumResp GetCommentLikeNum(1: CommentLikeNumReq req)(api.get="api/v1/like/commentlike")
    VideoLikeNumResp GetVideoLikeNum(1: VideoLikeNumReq req)(api.get="api/v1/like/videolike")
}
