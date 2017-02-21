package gokit

import (
	"context"
	"errors"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"

	"generator/app/inject"
)

//HandleNotFound - 400 error
// @Title HandleNotFound
// @Description HandleNotFound
// @Accept  json
// @Success 400 {bool}
func HandleNotFound(ctx inject.FullCtx) http.Handler {
	endpoint := func(context context.Context, request interface{}) (interface{}, error) {
		return nil, errors.New("Incorrect handler")
	}

	return transport.NewServer(
		context.Background(),
		endpoint,
		emptyDecoder(),
		simpleEncoder(),
		transport.ServerBefore(AddRequest()),
	)
}
