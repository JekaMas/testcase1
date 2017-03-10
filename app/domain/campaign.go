package domain

import (
	"strings"
)

//easyjson:json
type Campaign struct {
	CampaignName string           `json:"compaign_name"`
	Price        float64          `json:"price"`
	TargetList   TargetCollection `json:"target_list"`
}

func (c *Campaign) Verify() bool {
	return compose(
		c.IsPricePositive,
		c.IsCorrectName,
		c.VerifyTargetList,
	)()
}

const (
	//FIXME: should use either math.Bigint or env dependant params
	PriceThreshold         = 0.001
	CampaignPrefix         = "campaign"
	campaignNumberPosition = len(CampaignPrefix)
)

func (c *Campaign) IsPricePositive() bool {
	return c.Price > PriceThreshold
}

func (c *Campaign) IsCorrectName() bool {
	var ok bool

	if ok = strings.HasPrefix(c.CampaignName, CampaignPrefix); !ok {
		return false
	}

	if ok = IsLatinNumber(c.CampaignName, campaignNumberPosition); !ok {
		return false
	}

	return true
}

func (c *Campaign) VerifyTargetList() bool {
	return c.TargetList.Verify()
}

