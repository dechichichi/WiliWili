namespace go networking

/*struct FollowOperationReq 关注操作请求
*   @param userId 用户id
*   @param followerId 关注者id
*/
struct FollowOperationReq {
    1: required i64 userId,
    2: required i64 followerId
}

/*struct FollowOperationResp 关注操作响应
*   @param success 操作是否成功
*/
struct FollowOperationResp {
1: required bool success
}

/*struct FollowListReq 获取关注列表请求
*   @param userId 用户id
*/
struct FollowListReq {
    1: required i64 userId
}

/*struct FollowListResp 获取关注列表响应
*   @param followList 关注列表
*/
struct FollowListResp {
    1: required list<i64> followList
}

/*struct FollowerListReq 获取粉丝列表请求
*   @param userId 用户id
*/
struct FollowerListReq {
    1: required i64 userId

}

/*struct FollowerListResp 获取粉丝列表响应
*   @param followerList 粉丝列表
*/
struct FollowerListResp {
    1: required list<i64> followerList
}

/*struct FriendListReq 获取好友列表请求
*   @param userId 用户id
*/
struct FriendListReq {
    1: required i64 userId
}

/*struct FriendListResp 获取好友列表响应
*   @param friendList 好友列表
*/
struct FriendListResp {
    1: required list<i64> friendList
}

service FollowService {
    FollowOperationResp follow(1: FollowOperationReq req)
    FollowListResp getFollowList(1: FollowListReq req)
    FollowerListResp getFollowerList(1: FollowerListReq req)
    FriendListResp getFriendList(1: FriendListReq req)
}