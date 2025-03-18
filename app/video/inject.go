package video

import (
	"wiliwili/app/video/controllers/rpc"
	"wiliwili/app/video/domain/service"
	"wiliwili/app/video/infrastructure/mysql"
	"wiliwili/app/video/infrastructure/redis"
	"wiliwili/app/video/usecase"
	"wiliwili/kitex_gen/video"
	"wiliwili/pkg/base/client"
	"wiliwili/pkg/constants"
)

func InjectUserHandler() video.VideoService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	redisCache, err := client.NewRedisClient(constants.RedisDBUSer)
	if err != nil {
		panic(err)
	}
	db := mysql.NewVideoDB(gormDB)
	re := redis.NewVideoCache(redisCache)
	svc := service.NewVideoService(db, re)
	uc := usecase.NewVideoUsecase(db, svc, re)

	return rpc.NewVideoHandler(uc)
}
