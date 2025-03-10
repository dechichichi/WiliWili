// Code generated by hertz generator. DO NOT EDIT.

package networking

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	networking "wiliwili/app/gateway/handler/api/networking"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_app := root.Group("/app", _appMw()...)
		{
			_v1 := _app.Group("/v1", _v1Mw()...)
			{
				_networking := _v1.Group("/networking", _networkingMw()...)
				_networking.GET("/followList", append(_getfollowlistMw(), networking.GetFollowList)...)
				_networking.POST("/followUser", append(_followuserMw(), networking.FollowUser)...)
				_networking.GET("/followerList", append(_getfollowerlistMw(), networking.GetFollowerList)...)
				_networking.GET("/friendList", append(_getfriendlistMw(), networking.GetFriendList)...)
			}
		}
	}
}
