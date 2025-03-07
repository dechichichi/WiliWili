namespace go video

/*struct VideoSubmissionReq 视频提交请求
*@author userid 用户id
*@author video_id 视频id
*@author video_name 视频名称
*@author video_url 视频地址
*@author video_cover_url 视频封面地址
*@author video_desc 视频描述
*@author video_tags 视频标签
*@author video_duration 视频时长
*@author video_size 视频大小
*/
struct VideoSubmissionReq{
required string userid
required string video_id
required string video_name
required string video_url
required string video_cover_url
required string video_desc
required string video_tags
required i32 video_duration
required i32 video_size
}

/*struct VideoSubmissionResp 视频提交响应
*@author success 操作是否成功
*@author error 错误信息
*/
struct VideoSubmissionResp{
1: required bool success,
2: optional string error
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
*@author video_ids 视频id列表
*@author video_names 视频名称列表
*@author video_urls 视频地址列表
*@author video_cover_urls 视频封面地址列表
*@author video_descs 视频描述列表
*@author video_tags 视频标签列表
*/
struct VideoListResp{
1: required list<string> video_ids,
2: required list<string> video_names,
3: required list<string> video_urls,
4: required list<string> video_cover_urls,
5: required list<string> video_descs,
6: required list<string> video_tags,
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
*@author video_ids 视频id列表
*@author video_names 视频名称列表
*/
struct VideoSearchResp{
    1:required list<string> video_ids
    2:required list<string> video_names
}

/*struct VideoTrendingLearderBoardReq 视频热门榜请求
*@author page_num 页码
*@author page_size 页大小
*/
struct VideoTrendingLearderBoardReq{
    1:required i32 page_num
    2:required i32 page_size
}

/*struct VideoTrendingLearderBoardResp 视频热门榜响应
*@author video_ids 视频id列表
*@author video_names 视频名称列表
*@author video_urls 视频地址列表
*@author video_cover_urls 视频封面地址列表
*@author video_descs 视频描述列表
*@author video_tags 视频标签列表
*/
struct VideoTrendingLearderBoardResp{
    1:required list<string> video_ids
    2:required list<string> video_names
    3:required list<string> video_urls
    4:required list<string> video_cover_urls
    5:required list<string> video_descs
    6:required list<string> video_tags
}

service VideoService {    
    VideoSubmissionResp videoSubmission(1:VideoSubmissionReq req)
    VideoListResp videoList(1:VideoListReq req)
    VideoSearchResp videoSearch(1:VideoSearchReq req)
    VideoTrendingLearderBoardResp videoTrendingLearderBoard(1:VideoTrendingLearderBoardReq req)
}