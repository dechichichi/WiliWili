namespace go api.video


struct Video {
1:required string video_id
2:required string video_name
3:required string video_url
4:required i32 video_duration
}

struct VideoSubmissionReq {
    1:required Video video
}

struct VideoSubmissionResp {
    1:required bool success
}

struct VideoListReq{
    1:required i32 page_num
    2:required i32 page_size
}

struct VideoListResp{
    1:required list<Video> videos
}

struct VideoSearchReq{
    1:required string keyword
}

struct VideoSearchResp{
    1:required list<Video> videos
}

struct VideoTrendingReq{
    1:required i32 page_num
    2:required i32 page_size
}

struct VideoTrendingResp{
    1:required list<Video> videos
}
service VideoService {
    VideoSubmissionResp submit_video(1:VideoSubmissionReq req)(api.post="api/v1/video/submit")
    VideoListResp get_video_list(1:VideoListReq req)(api.get="api/v1/video/list")
    VideoSearchResp search_video(1:VideoSearchReq req)(api.get="api/v1/video/search")
    VideoTrendingResp get_video_trending(1:VideoTrendingReq req)(api.get="api/v1/video/trending")
}