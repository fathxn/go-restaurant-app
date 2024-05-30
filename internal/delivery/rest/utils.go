package rest

import (
	"errors"
	"go-restaurant-app/internal/model"
	"net/http"
	"strings"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, " ")
	if len(splitString) != 2 {
		return model.UserSession{}, errors.New("unauthorized")
	}
	accessToken := splitString[1]

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}
