package service

import (
	"github.com/DeniesKresna/brigunaduty/config"
	"github.com/DeniesKresna/brigunaduty/service/middlewares"
	benghandler "github.com/DeniesKresna/brigunaduty/service/modules/bengkel/handler"
	bengrepo "github.com/DeniesKresna/brigunaduty/service/modules/bengkel/repository"
	bengcase "github.com/DeniesKresna/brigunaduty/service/modules/bengkel/usecase"
	userhandler "github.com/DeniesKresna/brigunaduty/service/modules/user/handler"
	userrepo "github.com/DeniesKresna/brigunaduty/service/modules/user/repository"
	"github.com/DeniesKresna/brigunaduty/service/modules/user/usecase"
	usercase "github.com/DeniesKresna/brigunaduty/service/modules/user/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setRoutes(cfg *config.Config) (r *gin.Engine, err error) {
	const (
		USER     = "user"
		EMPLOYEE = "employee"
		ADMIN    = "administrator"
	)
	r = gin.New()

	// user modules definition
	userRepo := userrepo.UserCreateRepository(cfg.DB)
	userCase := usercase.UserCreateUsecase(userRepo)
	userHandler := userhandler.UserCreateHandler(userCase)

	// bengkel modules definitions
	bengkelRepo := bengrepo.BengkelCreateRepository(cfg.DB)
	bengkelCase := bengcase.BengkelCreateUsecase(bengkelRepo, userCase)
	bengkelHandler := benghandler.BengkelCreateHandler(bengkelCase)

	r.Use(corsConfig())
	r.Use(middlewares.ActivityLogger())

	r.StaticFile("/", "./storage/dist/index.html")
	r.StaticFile("/caricon.png", "./storage/dist/caricon.png")
	r.Static("/static", "./storage/dist/static")
	r.Static("/assets", "./storage/dist/assets")
	r.NoRoute(func(c *gin.Context) {
		c.File("./storage/dist/index.html")
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/auth/login", userHandler.AuthLogin)

		// admin
		adminRoute := v1.Use(roleCheck(userCase, ADMIN))
		{
			adminRoute.POST("/user", userHandler.UserCreate)
			adminRoute.PUT("/user", userHandler.UserUpdate)
			adminRoute.GET("/user/:id", userHandler.UserGetByID)
			adminRoute.POST("/user/search", userHandler.UserSearch)
		}

		adminNemployeeRoute := v1.Use(roleCheck(userCase, ADMIN, EMPLOYEE))
		// admin & employee
		{
			adminNemployeeRoute.POST("/user/byemail", userHandler.UserGetByEmail)
			adminNemployeeRoute.POST("/user/employee-list", userHandler.UserGetAllEmployee)

			adminNemployeeRoute.GET("/car/:id", bengkelHandler.CarGetByID)
			adminNemployeeRoute.POST("/car/search", bengkelHandler.CarSearch)
			adminNemployeeRoute.POST("/car/plat-list", bengkelHandler.CarListGetByPlat)
			adminNemployeeRoute.POST("/car", bengkelHandler.CarCreate)
			adminNemployeeRoute.PUT("/car", bengkelHandler.CarUpdate)

			adminNemployeeRoute.GET("/customer/:id", bengkelHandler.CustomerGetByID)
			adminNemployeeRoute.POST("/customer/search", bengkelHandler.CustomerSearch)
			adminNemployeeRoute.POST("/customer/byphone", bengkelHandler.CustomerGetByPhone)
			adminNemployeeRoute.POST("/customer/name-list", bengkelHandler.CustomerListGetByName)
			adminNemployeeRoute.POST("/customer", bengkelHandler.CustomerCreate)
			adminNemployeeRoute.PUT("/customer", bengkelHandler.CustomerUpdate)

			adminNemployeeRoute.GET("/execution/:id", bengkelHandler.ExecutionGetByID)
			adminNemployeeRoute.POST("/execution/search", bengkelHandler.ExecutionSearch)
			adminNemployeeRoute.POST("/execution", bengkelHandler.ExecutionCreate)
			adminNemployeeRoute.PUT("/execution", bengkelHandler.ExecutionUpdate)
			adminNemployeeRoute.POST("/execution/download", bengkelHandler.ExecutionDownload)
		}
	}

	return
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Include "Content-Type" in the list of allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}

func roleCheck(userCase usecase.UserUsecase, roles ...string) gin.HandlerFunc {
	return middlewares.CheckRole(userCase, roles)
}

func Start(cfg *config.Config) (err error) {
	eng, err := setRoutes(cfg)

	eng.Run(cfg.App.Host + ":" + cfg.App.Port)
	return
}
