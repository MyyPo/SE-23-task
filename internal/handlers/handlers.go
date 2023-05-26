package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/myypo/btcinform/internal/dto/responses"
	"github.com/myypo/btcinform/internal/services"
)

type Handler interface {
	HandleGetRate() gin.HandlerFunc
	HandleSubscribe() gin.HandlerFunc
	HandleSendEmails() gin.HandlerFunc
}

type HandlerImpl struct {
	ExchangeRateService services.ExchangeRateService
}

func NewHandlerImpl(exchangeRateService services.ExchangeRateService) *HandlerImpl {
	return &HandlerImpl{
		ExchangeRateService: exchangeRateService,
	}
}

func (h *HandlerImpl) HandleGetRate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		publicAPIResponse, err := h.ExchangeRateService.GetExchangeRate()
		if err != nil {
			sendResponse(ctx, http.StatusInternalServerError, nil, err)
			return
		}

		response := responses.NewGetRateResponse(publicAPIResponse.Rate)

		sendResponse(ctx, http.StatusOK, response, nil)
	}
}

func (h *HandlerImpl) HandleSubscribe() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *HandlerImpl) HandleSendEmails() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func sendResponse(ctx *gin.Context, statusCode int, data any, err error) {
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}
	resp := struct {
		errorMessage string
		data         any
	}{
		errorMessage: errorMessage,
		data:         data,
	}
	ctx.JSON(statusCode, resp)
}
