namespace go model

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Video {
    1:required string video_id
    2:required string video_name
    3:required string video_url
    4:required string video_duration
    5:required i64 likes_count
}

struct UserInfo {
    1: required string username;
    2: required i64 uid;
}

struct UserProfile {
    1: required string username;
    2: required string email;
    3: required string gender;
    4: required string signature;
}

struct Image{
    1: required i64 image_id
    2: required string image_url
}

struct Comment{
    1: required i64 comment_id
    2: required i64 user_id
    3: required string content
    4: required i64 create_time
}