package handlers

import (
	"github.com/erfansahebi/lamia_gateway/common"
	authProto "github.com/erfansahebi/lamia_shared/services/auth"
	"net/http"
)

func (h *Handler) UserDetail(r *http.Request) (interface{}, int, error) {
	userID, err := common.GetUserIDFromContext(r.Context())
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	userData, err := h.Di.Services().Auth().Client().GetUser(r.Context(), &authProto.GetUserRequest{
		UserId: userID.String(),
	})
	if err != nil {
		return nil, http.StatusBadRequest, HandleErrorFromGrpc(err)
	}

	return UserDetailResponse{AuthorizationUserResponse{
		ID:        userData.User.Id,
		FirstName: userData.User.FirstName,
		LastName:  userData.User.LastName,
		Email:     userData.User.Email,
	}}, http.StatusOK, nil
}
