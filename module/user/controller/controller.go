package controller

import (
	"conceitoExato/module/user/repository"
	"conceitoExato/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	repo := repository.NewUserRepository()

	couldNotFindUser := repo.FindUserByLogin(ctx.Param("login"))

	if util.ContainsError(couldNotFindUser) {
		errorMessage := fmt.Sprintf("%v: %v", couldNotFindUser, ctx.Param("login"))

		ctx.JSON(http.StatusNotFound, errorMessage)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": repo.GetUser()})
}
