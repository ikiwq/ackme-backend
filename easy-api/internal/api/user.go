package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ikiwq/ackme/easy-api/internal/types"
	"github.com/ikiwq/ackme/utils"
)

func (a *api) login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var loginRequest types.UserLoginRequest
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		http.Error(w, "Error while parsing JSON", http.StatusBadRequest)
		return
	}

	user, err := a.easyUserRepository.GetByUsernameAndPassword(ctx, loginRequest.Username, loginRequest.Password)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		utils.ErrorResponse(w, r, http.StatusUnauthorized, "Wrong username or password")
	case err != nil:
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "An unknown error has occurred")
	default:
		utils.WriteJSON(w, http.StatusOK, user)
	}
}