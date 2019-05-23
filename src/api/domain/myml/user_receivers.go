package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
)

const urlUsers = "https://api.mercadolibre.com/users/"

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"


	cb = gobreaker.NewCircuitBreaker(st)
}

func (user *User) Get() *apierrors.ApiError {
	final := fmt.Sprintf("%s%d", urlUsers, user.ID)

	_, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(final)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, user); err != nil {
			return nil, err
		}
		return body, nil
	})
	if err != nil {
		return &apierrors.ApiError{err.Error(),
			http.StatusInternalServerError}
	}

	return nil
}
