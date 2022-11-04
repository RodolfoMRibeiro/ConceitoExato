package middleware

import (
	"conceitoExato/adapter/env"
	"conceitoExato/common/library"
	"conceitoExato/common/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(ctx *gin.Context) {
	signedString := []byte(env.Secret.SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().UTC().Add(10 * time.Minute).Unix()

	tokenString, unableToGetSignedToken := token.SignedString(signedString)

	if util.ContainsError(unableToGetSignedToken) {
		ctx.JSON(http.StatusOK, gin.H{"Error": unableToGetSignedToken, "Token": library.EMPTY_STRING})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Token": tokenString})
}
