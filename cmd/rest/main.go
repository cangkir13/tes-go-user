package main

import (
	"log"

	"github.com/cangkir13/tes-go-user/docs"
	usercontroller "github.com/cangkir13/tes-go-user/lib/controllerRest/user"
	"github.com/cangkir13/tes-go-user/lib/core/config"
	userrepo "github.com/cangkir13/tes-go-user/lib/core/repository/user"
	userusecase "github.com/cangkir13/tes-go-user/lib/core/usecase/user"
	"github.com/cangkir13/tes-go-user/lib/middlewares"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@cangkir13
//	@title			Swagger users go
//	@version		1.0
//	@description	This RESTFULL API users

//	@contact.name	API Support
//	@contact.email	umarbahabasi@gmail.com

//	@host						localhost:8080
//	@BasePath					/api
//	@query.collection.format	multi
func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// variables repositories usecases and handler/controllers
	var (
		// repositories
		userrepo = userrepo.NewUserRepository(db)

		// usecases
		userusecase = userusecase.NewUserUsecase(userrepo)

		// controller
		usercontroller = usercontroller.NewUserController(userusecase)
	)

	// Init Routers and Middlewares
	r := gin.Default()
	r.Use(middlewares.CORS())

	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	apiuser := api.Group("/users")
	{
		apiuser.GET("", usercontroller.GetListUser)
		apiuser.GET("/:id", usercontroller.GetOneID)
		apiuser.POST("", usercontroller.Create)
		apiuser.PUT("/:id", usercontroller.Update)
		apiuser.DELETE("/:id", usercontroller.Delete)
	}

	r.Run()
}
