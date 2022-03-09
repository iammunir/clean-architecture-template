package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iammunir/clean-architecture-template/usecase"
)

type Delivery interface {
	Ping(*gin.Context)
}

type delivery struct {
	usecase usecase.UseCase
}

func NewDelivery(usecase usecase.UseCase) Delivery {
	return &delivery{
		usecase: usecase,
	}
}

func (deliver *delivery) Ping(c *gin.Context) {
	err := deliver.usecase.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success"})
}
