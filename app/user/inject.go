package user

import (
	"wiliwili/app/user/controllers/rpc"
	"wiliwili/app/user/domain/service"
	"wiliwili/app/user/infrastructure/mysql"
	"wiliwili/app/user/infrastructure/redis"
	"wiliwili/app/user/usecase"
	"wiliwili/kitex_gen/user"
	"wiliwili/pkg/base/client"

	"wiliwili/pkg/constants"
)

func InjectUserHandler() user.UserService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	redisCache, err := client.NewRedisClient(constants.RedisDBUSer)
	if err != nil {
		panic(err)
	}
	db := mysql.NewUserDB(gormDB)
	re := redis.NewUserCache(redisCache)
	svc := service.NewUserService(db, re)
	uc := usecase.NewUserUseCase(db, svc, re)

	return rpc.NewUserHandler(uc)
}
