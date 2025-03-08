namespace go api.networking

struct FollowOperationReq{
    1: required i64 userId,
    2: required i64 followerId
}

struct FollowOperationResp{
    1:required bool success,
}

struct FollowListReq{
    1: required i64 userId,
}

struct FollowListResp{
    1:required list<i64> followList,
}

struct FollowerListReq{
    1: required i64 userId,
}

struct FollowerListResp{
    1:required list<i64> followerList,
}

struct FriendListReq {
    1: required i64 userId
}

struct FriendListResp {
    1: required list<i64> friendList
}

service NetworkingService{
    FollowOperationResp followUser(1:FollowOperationReq req)(api.post="app/v1/networking/followUser"),
    FollowListResp getFollowList(1:FollowListReq req)(api.get="app/v1/networking/followList"),
    FollowerListResp getFollowerList(1:FollowerListReq req)(api.get="app/v1/networking/followerList"),
    FriendListResp getFriendList(1:FriendListReq req)(api.get="app/v1/networking/friendList"),
}