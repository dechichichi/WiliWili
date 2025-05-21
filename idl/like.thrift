namespace go like
include "model.thrift"

/*LikeComment 给评论点赞
* @param commentId 评论ID
* @param userId 用户ID
* @param likeType 点赞类型 1:点赞 2:取消点赞
*/
struct LikeCommentReq {
    1:required string commentId // 评论ID
    2:required bool IsLike // 点赞类型 1:点赞 2:取消点赞
}

struct LikeCommentResp {
}

/*LikeVideo 给视频点赞
* @param postId 帖子ID
* @param userId 用户ID
* @param likeType 点赞类型 1:点赞 2:取消点赞
*/
struct LikeVideoReq {
    1:required string videoId // 视频ID
    2:required bool IsLike // 点赞类型 1:点赞 2:取消点赞
}
struct LikeVideoResp {
}

/*CommentLikeNum 获取评论点赞数目
*/
struct CommentLikeNumReq {
    1:required string commentId // 评论ID
}

struct CommentLikeNumResp {
    1:required i64 totalCount; // 总数量
}

/*VideoLikeNum 获取视频点赞数目
*/
struct VideoLikeNumReq {
    1:required string videoId // 视频ID
}

struct VideoLikeNumResp {
    1:required i64 totalCount; // 总数量
}

service LikeService {
     LikeCommentResp LikeComment(1:LikeCommentReq req);
     LikeVideoResp LikeVideo(1:LikeVideoReq req);
     CommentLikeNumResp CommentLikeNum(1:CommentLikeNumReq req);
     VideoLikeNumResp VideoLikeNum(1:VideoLikeNumReq req);
}
