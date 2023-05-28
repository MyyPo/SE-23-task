package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/myypo/btcinform/internal/dto/requests"
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
		var request requests.GetRateRequest

		publicAPIResponse, err := h.ExchangeRateService.GetExchangeRate(request)
		if err != nil {
			log.Println("service HandleGetRate InternalError")
			sendResponse(ctx, http.StatusInternalServerError, nil, err)
			return
		}

		response := responses.NewGetRateResponse(publicAPIResponse.Rate)

		sendResponse(ctx, http.StatusOK, response, nil)
	}
}

func (h *HandlerImpl) HandleSubscribe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request requests.SubscribeRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			sendResponse(ctx, http.StatusBadRequest, nil, err)
			return
		}

		resp, err := h.ExchangeRateService.Subscribe(request)
		if err != nil {
			switch err {
			case services.DuplicateError{}:
				log.Println("service HandleSubscribe DuplicateError")
				sendResponse(ctx, http.StatusConflict, nil, err)
				return
			default:
				log.Println("service HandleSubscribe InternalError")
				sendResponse(ctx, http.StatusInternalServerError, nil, err)
				return
			}
		}
		sendResponse(ctx, http.StatusOK, resp, nil)
	}
}

func (h *HandlerImpl) HandleSendEmails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request requests.SendEmailsRequest

		resp, err := h.ExchangeRateService.SendEmails(request)
		if err != nil {
			switch err {
			default:
				log.Println("service HandleSendEmails InternalError")
				sendResponse(ctx, http.StatusInternalServerError, nil, err)
				return
			}
		}
		sendResponse(ctx, http.StatusOK, resp, nil)
	}
}

func sendResponse(ctx *gin.Context, statusCode int, data any, err error) {
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}
	resp := struct {
		ErrorMessage string `json:"error"`
		Data         any    `json:"data"`
	}{
		ErrorMessage: errorMessage,
		Data:         data,
	}
	ctx.JSON(statusCode, resp)
}
