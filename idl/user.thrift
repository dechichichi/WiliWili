namespace go user  

include "model.thrift"

/*
*struct UserRegisterReq 用户注册请求
* @param string username 用户名
* @param string password 密码
* @param string email 邮箱
* @param string gender 性别
* @param string signature 签名
*/
struct  UserRegisterReq {
    1: required string username;
    2: required string password;
    3: required string email;
    4: required string gender;
    5: required string signature;
}

struct  UserRegisterResp {
    1: required model.BaseResp baseResp;
    2: required i64 Uid;
}


/*
*struct UserLoginReq 用户登录请求
* @param string username 用户名
* @param string password 密码
*/
struct  UserLoginReq {
    1: required i64 Uid;
    2: required string password;
}


struct  UserLoginResp {
    1:model.BaseResp baseResp;
    2:model.UserInfo userInfo;
}

/*
*struct UserProfileReq 用户资料请求
* @param string token 登录令牌
*/
struct  UserProfileReq {
    1: required i64 Uid;
}


struct  UserProfileResp {
   1:model.BaseResp baseResp;
   2:model.UserProfile userProfile;
}

/*
*struct UserAvatarUploadReq 用户头像上传请求
* @param binary avatar 头像文件
*/
struct  UserAvatarUploadReq {
    1: required binary avatar;
}


struct  UserAvatarUploadResp {
    1:model.BaseResp baseResp;
    2:model.Image image;
}

service UserService {
    UserRegisterResp userRegister(1: UserRegisterReq req);
    UserLoginResp userLogin(1: UserLoginReq req);
    UserProfileResp userProfile(1: UserProfileReq req);
    UserAvatarUploadResp userAvatarUpload(1: UserAvatarUploadReq req);
}