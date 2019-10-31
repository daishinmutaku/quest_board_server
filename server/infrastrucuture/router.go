package infrastrucuture

import (
	"github.com/daishinmutaku/quest_board_server/server/infrastrucuture/datastore"
	"github.com/daishinmutaku/quest_board_server/server/interface/controller"
	"github.com/daishinmutaku/quest_board_server/server/interface/database"
	"github.com/daishinmutaku/quest_board_server/server/usecase/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	r := gin.Default()
	r.Use(cors.Default())

	db := datastore.NewSqlHandler()

	//controller
	userController := controller.NewUserController(
		service.NewUserService(
			database.NewUserRepository(db),
		),
	)

	//user
	r.GET("/user/:id", userController.GetUser)
	r.POST("/user/create", userController.CreateUser)
	r.PUT("/user/update/:id", userController.UpdateUser)
	r.DELETE("/user/delete/:id", userController.GetUser)

	Router = r
}
