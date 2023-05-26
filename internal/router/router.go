package router

import (
	"github.com/gin-gonic/gin"

	"github.com/myypo/btcinform/internal/handlers"
	"github.com/myypo/btcinform/internal/services"
)

type RouterImpl struct{}

func NewRouterImpl() *RouterImpl {
	return &RouterImpl{}
}

func (r *RouterImpl) Serve() {
	router := gin.Default()

	exchangeRateService := services.NewExchangeRateServiceImpl()

	handler := handlers.NewHandlerImpl(exchangeRateService)

	router.GET("/rate", handler.HandleGetRate())
	router.POST("/subscribe", handler.HandleSubscribe())
	router.POST("/sendEmails", handler.HandleSendEmails())

	router.Run(":8080")
}
