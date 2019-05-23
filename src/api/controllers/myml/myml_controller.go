package myml

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml/src/api/services/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
)

func Myml(context *gin.Context) {
	userID := context.Param("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		context.JSON(apiErr.Status, apiErr)
		return
	}
	user, apiErr := myml.GetUser(id)
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}
	context.JSON(http.StatusOK, user)

}
