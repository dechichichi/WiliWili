namespace go api.video

struct Video {
    1:required string video_id
    2:required string video_name
    3:required string video_url
    4:required string video_duration
}

struct VideoSubmissionReq {
    1:required string video_name
    2:required string video_duration
}

struct VideoSubmissionResp {
    1:required bool success
    2:required string video_id
}

struct VideoGetReq{
    1:required string video_id
}

struct VideoGetResp{
    1:required Video video
}

struct VideoSearchReq{
    1:required string keyword
    2:required i64 page_size
    
}

struct VideoSearchResp{
    1:required list<Video> videos
}

struct VideoTrendingReq{
    1:required i64 page_num
    2:required i64 page_size
}

struct VideoTrendingResp{
    1:required list<Video> videos
}
service VideoService {
    VideoSubmissionResp submit_video(1:VideoSubmissionReq req)(api.post="api/v1/video/submit")
    VideoGetResp get_video(1:VideoGetReq req)(api.get="api/v1/video/get")
    VideoSearchResp search_video(1:VideoSearchReq req)(api.get="api/v1/video/search")
    VideoTrendingResp get_video_trending(1:VideoTrendingReq req)(api.get="api/v1/video/trending")
}