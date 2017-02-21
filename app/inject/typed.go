package inject

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"generator/app/storage"
)

//NewTyped - constructor for FullCtx
func NewTyped(db storage.CampaignsDAO) *TypedCtx {
	return &TypedCtx{
		db: db,
	}
}

//TypedCtx -TypedCtx
type TypedCtx struct {
	queryParams url.Values
	pathParams  map[string]string
	db          storage.CampaignsDAO
	context.Context
}

// Release - release into pool
func (this *TypedCtx) Release() {
	if this == nil {
		return
	}

	Release(this)
}

//Reset - reset incoming params in context
func (this *TypedCtx) Reset(full bool) {
	this.Context = nil
	this.pathParams = nil
	this.queryParams = nil

	if full {
		this.db = nil
		this.pathParams = nil
		this.queryParams = nil
	}
}

//WithRequest - WithRequest
func (this *TypedCtx) WithRequest(req *http.Request) FullCtx {
	this.queryParams = req.URL.Query()
	this.pathParams = mux.Vars(req)
	return this
}

//GetDB function to get db manager
func (this *TypedCtx) GetDB() storage.CampaignsDAO {
	return this.db
}

//GetParamString - GetParamString
func (this *TypedCtx) GetParamString(name string) (string, error) {
	paramQ, ok := this.queryParams[name]
	if ok == true && len(paramQ) > 0 {
		return paramQ[0], nil
	}

	paramP, ok := this.pathParams[name]
	if ok == true && len(paramP) > 0 {
		return paramP, nil
	}

	return "", fmt.Errorf("Incorrect get param %q", name)
}

//GetParamBytes - GetParamBytes
func (this *TypedCtx) GetParamBytes(name string) ([]byte, error) {
	paramQ, ok := this.queryParams[name]
	if ok == true && len(paramQ) > 0 {
		return []byte(paramQ[0]), nil
	}

	paramP, ok := this.pathParams[name]
	if ok == true && len(paramP) > 0 {
		return []byte(paramP), nil
	}

	return nil, fmt.Errorf("Incorrect get param %q", name)
}

//GetParamUint32 parse uint32 type from context param
func (this *TypedCtx) GetParamUint32(name string) (uint32, error) {
	if v, err := this.GetParamBytes(name); err == nil {
		v = bytes.TrimSpace(v)

		intval, err := toUint32(v)
		if err != nil {
			return 0, fmt.Errorf("wrong input parameter type - '%s'. Expected type - '%s'", name, "uint32")
		}

		return intval, nil
	}
	return 0, fmt.Errorf("Incorrect get param %q", name)
}

//GetParamUint64 parse uint64 type from context param
func (this *TypedCtx) GetParamUint64(name string) (uint64, error) {
	if v, err := this.GetParamBytes(name); err == nil {
		v = bytes.TrimSpace(v)

		intval, err := toUint64(v)
		if err != nil {
			return 0, fmt.Errorf("wrong input parameter type - '%s'. Expected type - '%s'", name, "uint64")
		}

		return intval, nil
	}
	return 0, fmt.Errorf("Incorrect get param %q", name)
}
