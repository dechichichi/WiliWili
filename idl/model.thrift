namespace go model

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Video {
    1:required string video_id
    2:required string video_name
    3:required string video_url
    4:required i32 video_duration
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
