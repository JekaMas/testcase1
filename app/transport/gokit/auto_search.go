package gokit

import (
	"fmt"
	"generator/app/inject"
	"generator/app/services"
	"net/http"

	"context"
	"generator/app/storage/repository"
	transport "github.com/go-kit/kit/transport/http"
)

//AutoSearch generates random user profile and search campaign by it
// @Title Search campaign by user(profile)
// @Accept  json
// @Success 200 domain.SearchResult
// @Router /search_auto  [GET]
func AutoSearch(ctx inject.FullCtx) http.Handler {
	endpoint := func(context context.Context, request interface{}) (interface{}, error) {
		reqCTX := ctx.WithRequest(context.Value(inject.Request).(*http.Request))

		db := reqCTX.GetDB()
		if db == nil {
			return nil, fmt.Errorf("DB is not exists")
		}

		getUserService := services.NewGetUserService()
		user, err := getUserService.GetUser()
		if err != nil {
			return nil, err
		}

		searchService := services.NewSearchService(repository.CampaignRepository(db))
		return searchService.Search(user)
	}

	return transport.NewServer(
		context.Background(),
		endpoint,
		emptyDecoder(),
		simpleEncoder(),
		transport.ServerBefore(AddRequest()),
	)
}
