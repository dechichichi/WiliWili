package comment

import (
	"wiliwili/app/comment/controllers/rpc"
	"wiliwili/app/comment/domain/service"
	"wiliwili/app/comment/infrastructure/mysql"
	"wiliwili/app/comment/infrastructure/redis"
	"wiliwili/app/comment/usecase"
	"wiliwili/kitex_gen/comment"
	"wiliwili/pkg/base/client"
	"wiliwili/pkg/constants"
)

func InjectCommentHandler() comment.CommentService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	redisCache, err := client.NewRedisClient(constants.RedisDBUSer)
	if err != nil {
		panic(err)
	}
	db := mysql.NewCommentDB(gormDB)
	re := redis.NewCommentCache(redisCache)
	svc := service.NewCommentService(db, re)
	uc := usecase.NewcommentUseCase(db, svc, re)

	return rpc.NewCommentHandler(uc)
}
