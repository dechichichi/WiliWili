namespace go api.user
struct RegisterUserReq {
    1: required string username;
    2: required string password;
    3: required string email;
    4: required string gender;
    5: required string signature;
}

struct RegisterUserResp{
    1: required i64 id;
}
struct LoginRequest{
    1: required i64 id;
    2: required string password;
}

struct UserInfo{
    1: i64 userId,
    2: string name,
}
struct LoginResponse{
    2: required UserInfo user;
}

struct ProfileReq{
    1:required i64 userId,
}

struct ProfileResp{
    1: required string username;
    2: required string email;
    3: required string gender;
    4: required string signature;
}

struct  UserAvatarUploadReq {
    1: required i64 userId;
}

struct UserAvatarUploadResp {
    1: required bool success;
}

struct UserAvatarGetReq {
    1: required i64 userId;
}

struct UserAvatarGetResp {
    1: required string url;
}

service UserService {
    RegisterUserResp RegisterUser(1:RegisterUserReq req)(api.post="api/v1/user/register"),
    LoginResponse Login(1:LoginRequest req)(api.post="api/v1/user/login"),
    ProfileResp GetProfile(1:ProfileReq req)(api.get="api/v1/user/profile"),
    UserAvatarUploadResp UploadAvatar(1:UserAvatarUploadReq req)(api.post="api/v1/user/uploadavatar");
    UserAvatarGetResp GetAvatar(1:UserAvatarGetReq req)(api.get="api/v1/user/getavatar");
}
