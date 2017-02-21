package gokit

import (
	"context"
	"fmt"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"

	"generator/app/inject"
	"generator/app/services"
	"generator/app/storage/repository"
)

const (
	campaignsNumKey  = "z"
	targetsNumKey    = "y"
	attributesNumKey = "x"
)

// GetCampaigns - ganerate and get new campaigns
// @Title Generate new campaigns
// @Accept  json
// @Success 200 {array}  domain.CampaignsCollection
// @Router /campaigns [get]
func GetCampaigns(ctx inject.FullCtx) http.Handler {
	endpoint := func(context context.Context, request interface{}) (interface{}, error) {
		var (
			err           error
			campaignsNum  uint32
			targetsNum    uint32
			attributesNum uint32
		)

		reqCTX := ctx.WithRequest(context.Value(inject.Request).(*http.Request))

		campaignsNum, err = reqCTX.GetParamUint32(campaignsNumKey)
		if err != nil {
			return nil, fmt.Errorf("Wrong campaigns number")
		}

		targetsNum, err = reqCTX.GetParamUint32(targetsNumKey)
		if err != nil {
			return nil, fmt.Errorf("Wrong targets number")
		}

		attributesNum, err = reqCTX.GetParamUint32(attributesNumKey)
		if err != nil {
			return nil, fmt.Errorf("Wrong attributes number")
		}

		db := reqCTX.GetDB()
		if db == nil {
			return nil, fmt.Errorf("DB is not exists")
		}

		getService := services.NewGetService(repository.NewCampaignRepository(db))
		return getService.Get(int(campaignsNum), int(targetsNum), int(attributesNum))
	}

	return transport.NewServer(
		context.Background(),
		endpoint,
		emptyDecoder(),
		simpleEncoder(),
		transport.ServerBefore(AddRequest()),
	)
}
