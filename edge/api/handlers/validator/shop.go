package validator

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/common"
	"github.com/erfansahebi/lamia_gateway/di"
)

type ShopResponse struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

// Create Shop

type CreateShopRequest struct {
	Name string `json:"name"`
}

func (csr *CreateShopRequest) Validate(ctx context.Context, container di.DIContainerInterface) error {
	switch {
	case csr.Name == "":
		return common.ErrEmptyFields
	}

	return nil
}

type CreateShopResponse struct {
	ShopResponse
}
