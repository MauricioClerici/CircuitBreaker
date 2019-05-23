package myml

import (
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
)


func GetUser(id int64) (*myml.User, *apierrors.ApiError) {
	if id == 0 {
		return nil, &apierrors.ApiError{
			"userId invalido",
			http.StatusBadRequest,
		}
	}
	user := &myml.User{ID:id}
	if apiErr := user.Get();
	apiErr != nil {
		return nil, apiErr
	}
	return user, nil


}
