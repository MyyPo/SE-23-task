package router

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/myypo/btcinform/internal/config"
	"github.com/myypo/btcinform/internal/handlers"
	"github.com/myypo/btcinform/internal/repositories"
	"github.com/myypo/btcinform/internal/services"
	simpdb "github.com/myypo/btcinform/pkg/simpDB"
)

type RouterImpl struct{}

func NewRouterImpl() *RouterImpl {
	return &RouterImpl{}
}

func (r *RouterImpl) Serve() {
	router := gin.Default()
	config, err := config.NewConfigImpl()
	if err != nil {
		log.Fatalln(err)
	}

	db := simpdb.NewSimpDBProvider(*config.DBPath)
	repo := repositories.NewSubscriptionRepositoryImpl(db)
	exchangeRateService := services.NewExchangeRateServiceImpl(config, repo)

	handler := handlers.NewHandlerImpl(exchangeRateService)

	router.GET("/rate", handler.HandleGetRate())
	router.POST("/subscribe", handler.HandleSubscribe())
	router.POST("/sendEmails", handler.HandleSendEmails())

	router.Run(":" + *config.GetAppPort())
}
