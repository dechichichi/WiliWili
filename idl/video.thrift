namespace go video

include "common.thrift"


/*struct VideoSubmissionReq 视频提交请求
*@author userid 用户id
*@author video 视频信息
*/
struct VideoSubmissionReq{
1:required string userid
2:required model.Video video
}

/*struct VideoSubmissionResp 视频提交响应
*@author success 操作是否成功
*/
struct VideoSubmissionResp{
1: required bool success,
}

/*struct VideoListReq 视频列表请求
*@author userid 用户id
*@author page_num 页码
*@author page_size 页大小
*/
struct VideoListReq{
required string userid
required i32 page_num
required i32 page_size

}

/*struct VideoListResp 视频列表响应
*@author videos 视频列表
*/
struct VideoListResp{
1: required list<model.Video> videos,
}

/*struct VideoSearchReq 视频搜索请求
*@author keyword 关键字
*@author page_num 页码
*@author page_size 页大小
*/
struct VideoSearchReq{
    1:required string keyword
    2:required i32 page_num
    3:required i32 page_size
}
/*struct VideoSearchResp 视频搜索响应
*@author videos 视频列表
*/
struct VideoSearchResp{
    1:required list<model.Video> videos
}

/*struct VideoTrendingReq 视频热门榜请求
*@author page_num 页码
*@author page_size 页大小
*/
struct VideoTrendingReq{
    1:required i32 page_num
    2:required i32 page_size
}

/*struct VideoTrendingResp 视频热门榜响应
*@author video_ids 视频列表
*/
struct VideoTrendingResp{
    1:required list<model.Video> videos
}

service VideoService {    
    VideoSubmissionResp videoSubmission(1:VideoSubmissionReq req)
    VideoListResp videoList(1:VideoListReq req)
    VideoSearchResp videoSearch(1:VideoSearchReq req)
    VideoTrendingResp videoTrending(1:VideoTrendingReq req)
}