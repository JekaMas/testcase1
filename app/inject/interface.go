package inject

import (
	"context"
	"net/http"

	"generator/app/storage"
)

type (
	request struct{}
	dbV     struct{}
)

var (
	//Request - request key value
	Request = request{}
	//DB - db manager key value
	DB = dbV{}
)

// FullCtx full dependencies context
type FullCtx interface {
	context.Context
	DBInjector
	GetParam
	RequestInjector
}

//GetParam get query params
type GetParam interface {
	GetParamUint32(name string) (uint32, error)
	GetParamUint64(name string) (uint64, error)
}

//DBInjector db injection
type DBInjector interface {
	GetDB() storage.CampaignsDAO
}

type RequestInjector interface {
	WithRequest(req *http.Request) FullCtx
}
