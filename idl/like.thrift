namespace go like
include "model.thrift"
/*LikeComment 给评论点赞
* @param commentId 评论ID
* @param userId 用户ID
* @param likeType 点赞类型 1:点赞 2:取消点赞
*/
struct LikeCommentReq {
    1:required i64 commentId // 评论ID
    2:required i64 userId // 用户ID
    3:required i32 likeType // 点赞类型 1:点赞 2:取消点赞
}

struct LikeCommentResp {
    1:required model.BaseResp baseResp;
}


/*LikeVideo 给视频点赞
* @param postId 帖子ID
* @param userId 用户ID
* @param likeType 点赞类型 1:点赞 2:取消点赞
*/
struct LikeVideoReq {
    1:required i64 videoId // 视频ID
    2:required i64 userId // 用户ID
    3:required i32 likeType // 点赞类型 1:点赞 2:取消点赞
}
struct LikeVideoResp {
    1:required model.BaseResp baseResp;
}

/*CommentLikeNum 获取评论点赞数目
*/
struct CommentLikeNumReq {
    1:required i64 commentId // 评论ID
}

struct CommentLikeNumResp {
    1:required i32 totalCount; // 总数量
}

/*VideoLikeNum 获取视频点赞数目
*/
struct VideoLikeNumReq {
    1:required i64 videoId // 视频ID
}

struct VideoLikeNumResp {
    1:required i32 totalCount; // 总数量
}

service LikeService {
     LikeCommentResp LikeComment(1:LikeCommentReq req);
     LikeVideoResp LikeVideo(1:LikeVideoReq req);
     CommentLikeNumResp CommentLikeList(1:CommentLikeNumReq req);
     VideoLikeNumResp VideoLikeList(1:VideoLikeNumReq req);
}