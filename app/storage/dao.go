package storage

import "generator/app/domain"

type CampaignsDAO interface {
	Store(campaignCollection domain.CampaignCollection) error
	Get(campaignName string) (domain.Campaign, error)
	Search(profile domain.ProfileCollection) (domain.CampaignCollection, error)
}
