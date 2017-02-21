package inject

import (
	"sync"

	"generator/app/storage"
)

//AcquireFullCtx get from pool
func AcquireFullCtx(db storage.CampaignsDAO) FullCtx {
	v := ctxPool.Get()
	if v == nil {
		return NewTyped(db)
	}

	fullCtx := v.(FullCtx)
	return fullCtx
}

// AcquireCtx get clone from pool
func AcquireCtx(c *TypedCtx) *TypedCtx {
	v := ctxPool.Get()
	if v == nil {
		ctx := &TypedCtx{}
		clone(c, ctx)
		return ctx
	}

	ctx := v.(*TypedCtx)
	if ctx == nil {
		ctx = &TypedCtx{}
	}
	clone(c, ctx)

	return ctx
}

func clone(src, desc *TypedCtx) {
	desc.db = src.db

	//reset
	desc.Reset(false)
}

//Release context to pool
func Release(ctx FullCtx) {
	ctxPool.Put(ctx)
}

var ctxPool sync.Pool
