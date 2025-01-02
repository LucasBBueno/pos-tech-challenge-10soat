package handler

import (
	"errors"
	"net/http"
	"post-tech-challenge-10soat/internal/core/domain"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var errorStatusMap = map[error]int{
	domain.ErrInternal:        http.StatusInternalServerError,
	domain.ErrDataNotFound:    http.StatusNotFound,
	domain.ErrConflictingData: http.StatusConflict,
	domain.ErrForbidden:       http.StatusForbidden,
}

func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}

func validationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

type errorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

type clientResponse struct {
	ID    uuid.UUID `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Name  string    `json:"name" example:"John Doe"`
	Email string    `json:"email" example:"john-doe@email.com"`
}

func newClientReponse(client *domain.Client) clientResponse {
	return clientResponse{
		ID:    client.Id,
		Name:  client.Name,
		Email: client.Email,
	}
}

type categoryResponse struct {
	ID   uuid.UUID `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Name string    `json:"name" example:"Lanche"`
}

func newCategoryResponse(category *domain.Category) categoryResponse {
	return categoryResponse{
		ID:   category.Id,
		Name: category.Name,
	}
}

type productResponse struct {
	ID          uuid.UUID        `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Name        string           `json:"name" example:"Lanche 1"`
	Description string           `json:"description" example:"Lanche com bacon"`
	Image       string           `json:"image" example:"https://"`
	Value       float32          `json:"value" example:"10.90"`
	Category    categoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt   time.Time        `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

func newProductResponse(product *domain.Product) productResponse {
	return productResponse{
		ID:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Value:       product.Value,
		Category:    newCategoryResponse(product.Category),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
