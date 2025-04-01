package like

import (
	"wiliwili/kitex_gen/like"
	"wiliwili/app/like/controllers/rpc"
	"wiliwili/app/like/domain/service"
	"wiliwili/app/like/infrastructure/mysql"
	"wiliwili/app/like/infrastructure/redis"
	"wiliwili/app/like/usecase"
	"wiliwili/pkg/base/client"
	"wiliwili/pkg/constants"
)

func InjectLikeHandler() like.LikeService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	redisCache, err := client.NewRedisClient(constants.RedisDBLike)
	if err != nil {
		panic(err)
	}
	db := mysql.NewLikeDB(gormDB)
	re := redis.NewLikeCache(redisCache)
	svc := service.NewLikeService(db, re)
	uc := usecase.NewLikeUseCase(db, svc, re)

	return rpc.NewLikeHandler(uc)
}
