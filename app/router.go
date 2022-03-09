package app

import (
	"github.com/gin-gonic/gin"
	"github.com/iammunir/clean-architecture-template/delivery"
	"github.com/iammunir/clean-architecture-template/repository"
	"github.com/iammunir/clean-architecture-template/usecase"
	"gorm.io/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {

	repo := repository.NewRepository(mysqlConn)
	use := usecase.NewUseCase(repo)
	deliver := delivery.NewDelivery(use)

	router := gin.Default()
	router.GET("/ping", deliver.Ping)

	return router
}
