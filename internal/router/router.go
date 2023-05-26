package router

import (
	"github.com/gin-gonic/gin"

	"github.com/myypo/btcinform/internal/handlers"
)

type RouterImpl struct{}

func NewRouterImpl() *RouterImpl {
	return &RouterImpl{}
}

func (r *RouterImpl) Serve() {
	router := gin.Default()

	handler := handlers.NewHandlerImpl()

	router.GET("/rate", handler.HandleGetRate())
	router.POST("/subscribe", handler.HandleSubscribe())
	router.POST("/sendEmails", handler.HandleSendEmails())

	router.Run(":8080")
}
