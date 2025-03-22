namespace go video

include "model.thrift"


/*struct VideoSubmissionReq 视频提交请求
*@author userid 用户id
*@author video 视频信息
*/
struct VideoSubmissionReq{
1:required string video_name
2:required string video_duration
3:required binary video
}

/*struct VideoSubmissionResp 视频提交响应
*@author success 操作是否成功
*/
struct VideoSubmissionResp{
1:required model.BaseResp baseresp
2:required string video_id
3:required string video_url
}

/*struct VideoGetReq 视频请求
*@author userid 用户id
*@author page_num 页码
*@author page_size 页大小
*/
struct VideoGetReq{
required string video_id
}

/*struct VideoListResp 视频列表响应
*@author videos 视频列表
*/
struct VideoGetResp{
    1:required model.BaseResp baseresp
    2: required model.Video video,
}

/*struct VideoSearchReq 视频搜索请求
*@author keyword 关键字
*@author page_num 页码
*@author page_size 页大小
*/
struct VideoSearchReq{
    1:required string keyword
    2:required i64 page_num
    3:required i64 page_size
}
/*struct VideoSearchResp 视频搜索响应
*@author videos 视频列表
*/
struct VideoSearchResp{
    1:required model.BaseResp baseresp
    2:required list<model.Video> videos
}

/*struct VideoTrendingReq 视频热门榜请求
*@author page_num 页码
*@author page_size 页大小
*/
struct VideoTrendingReq{
    1:required i64 page_num
    2:required i64 page_size
}

/*struct VideoTrendingResp 视频热门榜响应
*@author video_ids 视频列表
*/
struct VideoTrendingResp{
    1:required model.BaseResp baseresp
    2:required list<model.Video> videos
}

service VideoService {    
    VideoSubmissionResp videoSubmission(1:VideoSubmissionReq req)
    VideoGetResp videoGet(1:VideoGetReq req)
    VideoSearchResp videoSearch(1:VideoSearchReq req)
    VideoTrendingResp videoTrending(1:VideoTrendingReq req)
}