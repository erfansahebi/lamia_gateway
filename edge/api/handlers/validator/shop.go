package validator

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/common"
	"github.com/erfansahebi/lamia_gateway/di"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

// Get Shop

type GetShopRequest struct {
	ShopID uuid.UUID `json:"shop_id"`
}

func (gsr *GetShopRequest) Validate(ctx context.Context, container di.DIContainerInterface) error {
	shopID, err := uuid.Parse(chi.URLParamFromCtx(ctx, "shop_id"))
	if err != nil {
		return common.ErrEmptyFields
	}

	gsr.ShopID = shopID

	return nil
}

type GetShopResponse struct {
	ShopResponse
}
