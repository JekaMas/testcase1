package gokit

import (
	"generator/app/inject"
	"generator/app/services"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"context"
)

//GetUser - get new user
// @Title Get just generated user
// @Accept  json
// @Success 200 domain.User
// @Router /user/ [GET]
func GetUser(ctx inject.FullCtx) http.Handler {
	endpoint := func(context context.Context, request interface{}) (interface{}, error) {
		getUserService := services.NewGetUserService()
		return getUserService.GetUser()
	}

	return transport.NewServer(
		context.Background(),
		endpoint,
		emptyDecoder(),
		simpleEncoder(),
		transport.ServerBefore(AddRequest()),
	)
}
