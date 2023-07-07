package handlers

import (
	"encoding/json"
	"github.com/erfansahebi/lamia_gateway/common"
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers/validator"
	shopProto "github.com/erfansahebi/lamia_shared/go/proto/shop"
	"io"
	"net/http"
)

func (h *Handler) CreateShop(r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrFailedToReadRequestBody
	}

	userID, err := common.GetUserIDFromContext(r.Context())
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var pendData validator.CreateShopRequest
	if err = json.Unmarshal(body, &pendData); err != nil {
		return nil, http.StatusBadRequest, common.ErrFailedToUnmarshalRequestBody
	}

	if err = pendData.Validate(r.Context(), h.Di); err != nil {
		return nil, http.StatusBadRequest, err
	}

	createdShopData, err := h.Di.Services().Shop().Client().Create(r.Context(), &shopProto.CreateRequest{
		Shop: &shopProto.ShopStruct{
			Id:     "",
			UserId: userID.String(),
			Name:   pendData.Name,
		},
	})
	if err != nil {
		return nil, http.StatusBadRequest, HandleErrorFromGrpc(err)
	}

	return validator.CreateShopResponse{
		ShopResponse: validator.ShopResponse{
			ID:     createdShopData.Shop.Id,
			UserID: createdShopData.Shop.UserId,
			Name:   createdShopData.Shop.Name,
		},
	}, http.StatusOK, nil
}

func (h *Handler) GetShop(r *http.Request) (interface{}, int, error) {
	var pendData validator.GetShopRequest
	if err := pendData.Validate(r.Context(), h.Di); err != nil {
		return nil, http.StatusBadRequest, err
	}

	fetchedShopData, err := h.Di.Services().Shop().Client().Get(r.Context(), &shopProto.GetRequest{
		Id: pendData.ShopID.String(),
	})
	if err != nil {
		return nil, http.StatusBadRequest, HandleErrorFromGrpc(err)
	}

	return validator.GetShopResponse{
		ShopResponse: validator.ShopResponse{
			ID:     fetchedShopData.Shop.Id,
			UserID: fetchedShopData.Shop.UserId,
			Name:   fetchedShopData.Shop.Name,
		},
	}, http.StatusOK, nil
}
