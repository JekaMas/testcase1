package gokit

import (
	"context"
	"fmt"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"

	"generator/app/domain"
	"generator/app/inject"
	"generator/app/services"
	"generator/app/storage/repository"
)

//AutoSearch search campaign by user(profile)
// @Title Search campaign by user(profile)
// @Accept  json
// @Success 200 domain.SearchResult
// @Router /search  [POST]
func Search(ctx inject.FullCtx) http.Handler {
	endpoint := func(context context.Context, request interface{}) (interface{}, error) {
		reqCTX := ctx.WithRequest(context.Value(inject.Request).(*http.Request))

		db := reqCTX.GetDB()
		if db == nil {
			return nil, fmt.Errorf("DB is not exists")
		}

		searchService := services.NewSearchService(repository.NewCampaignRepository(db))
		return searchService.Search(request.(domain.User))
	}

	return transport.NewServer(
		context.Background(),
		endpoint,
		userDecoder(),
		simpleEncoder(),
		transport.ServerBefore(AddRequest()),
	)
}
