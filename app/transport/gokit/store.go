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

//StoreCampaigns import campaigns into service
// @Title import campaigns into service
// @Accept  json
// @Success 200
// @Router /import_camp [POST]
func StoreCampaigns(ctx inject.FullCtx) http.Handler {
	endpoint := func(context context.Context, request interface{}) (interface{}, error) {
		reqCTX := ctx.WithRequest(context.Value(inject.Request).(*http.Request))

		db := reqCTX.GetDB()
		if db == nil {
			return nil, fmt.Errorf("DB is not exists")
		}

		storeService := services.NewStoreCampaignsService(repository.NewCampaignRepository(db))
		return storeService.StoreCampaigns(request.(domain.CampaignCollection))
	}

	return transport.NewServer(
		context.Background(),
		endpoint,
		campaignCollectionDecoder(),
		simpleEncoder(),
		transport.ServerBefore(AddRequest()),
	)
}
