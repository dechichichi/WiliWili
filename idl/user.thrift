namespace go user  


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

/*
*struct UserRegisterResp 用户注册响应
* @param string token 登录令牌
*/
struct  UserRegisterResp {
    1: required string token;
}


/*
*struct UserLoginReq 用户登录请求
* @param string username 用户名
* @param string password 密码
*/
struct  UserLoginReq {
    1: required string username;
    2: required string password;
}

/*
*struct UserLoginResp 用户登录响应
* @param string token 登录令牌
*/
struct  UserLoginResp {
    1: required string token;
}

/*
*struct UserProfileReq 用户资料请求
* @param string token 登录令牌
*/
struct  UserProfileReq {
    1: required string token;
}

/*
*struct UserProfileResp 用户资料响应
* @param string username 用户名
* @param string email 邮箱
* @param string gender 性别
* @param string signature 签名
*/
struct  UserProfileResp {
    1: required string username;
    2: required string email;
    3: required string gender;
    4: required string signature;
}

/*
*struct UserAvatarUploadReq 用户头像上传请求
* @param string token 登录令牌
* @param binary avatar 头像文件
*/
struct  UserAvatarUploadReq {
    1: required string token;
    2: required binary avatar;
}

/*
*struct UserAvatarUploadResp 用户头像上传响应    
* @param string url 头像URL
*/
struct  UserAvatarUploadResp {
    1: required string url;
}

service UserService {
    UserRegisterResp userRegister(1: UserRegisterReq req);
    UserLoginResp userLogin(1: UserLoginReq req);
    UserProfileResp userProfile(1: UserProfileReq req);
    UserAvatarUploadResp userAvatarUpload(1: UserAvatarUploadReq req);
}