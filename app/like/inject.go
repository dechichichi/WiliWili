package like

import (
	"wiliwili/app/like/controllers/rpc"
	"wiliwili/app/like/domain/service"
	"wiliwili/app/like/infrastructure/mysql"
	"wiliwili/app/like/usecase"
	"wiliwili/kitex_gen/like"
	"wiliwili/pkg/base/client"
)

func InjectLikeHandler() like.LikeService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}

	db := mysql.NewLikeDB(gormDB)
	svc := service.NewLikeService(db)
	uc := usecase.NewLikeUseCase(db, svc)

	return rpc.NewLikeHandler(uc)
}
