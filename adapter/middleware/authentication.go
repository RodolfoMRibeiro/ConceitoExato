package middleware

import (
	"conceitoExato/adapter/env"
	"conceitoExato/common/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAuthorized(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request

		if request.Header["Token"] == nil {
			util.HttpNoTokenFound(ctx)
			return
		}

		token, tokenParseError := jwt.Parse(request.Header["Token"][0], keyFunc)

		if tokenParseError != nil {

			util.HttpTokenHasExpired(ctx)
			return
		}

		if _, isTokenOk := token.Claims.(jwt.MapClaims); isTokenOk && token.Valid {
			handler(ctx)
			return
		}

		util.HttpTokenNotAuthorized(ctx)
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	secretKeyInBytes := []byte(env.Secret.SECRET_KEY)

	if _, isTokenOk := token.Method.(*jwt.SigningMethodECDSA); isTokenOk {
		return nil, fmt.Errorf("there was an error in parsing")
	}

	return secretKeyInBytes, nil
}
