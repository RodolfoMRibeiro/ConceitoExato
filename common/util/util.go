package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ParseMapToJson(mp map[string]string) string {
	str, _ := json.Marshal(mp)
	return string(str)
}

func ContainsError(err error) bool {
	return err != nil
}

func HttpBadRequestMessage(ctx *gin.Context, errElement error) {
	errorMessage := fmt.Sprintf("%v", errElement)

	ctx.JSON(http.StatusBadRequest, errorMessage)
}

func HttpNotFoundMessage(ctx *gin.Context, errElement error) {
	errorMessage := fmt.Sprintf("%v", errElement)

	ctx.JSON(http.StatusNotFound, errorMessage)
}

func IsEqual(firstElement interface{}, secondElement interface{}) bool {
	firstElementType := reflect.ValueOf(firstElement).Kind()
	secondElementType := reflect.ValueOf(secondElement).Kind()

	if firstElementType != secondElementType {
		return false
	}

	return firstElement == secondElement
}
