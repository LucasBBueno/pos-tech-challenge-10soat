package client

import (
	"github.com/google/uuid"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/utils"
)

type ClientResponse struct {
	ID    uuid.UUID `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Name  string    `json:"name" example:"John Doe"`
	Email string    `json:"email" example:"john-doe@email.com"`
}

func NewClientResponse(client *domain.Client) *ClientResponse {

	return &ClientResponse{
		ID:    utils.StringToUuid(client.Id),
		Name:  client.Name,
		Email: client.Email,
	}
}
