package repository

import (
	"generator/app/domain"
	"generator/app/storage"
)

type CampaignRepository interface {
	Store(domain.CampaignCollection) error
	Get(string) (domain.Campaign, error)
	Search(domain.ProfileCollection) (domain.CampaignCollection, error)
}

//campaignRepository
type campaignRepository struct {
	dao storage.CampaignsDAO
}

//NewCampaignRepository - constructor
func NewCampaignRepository(dao storage.CampaignsDAO) CampaignRepository {
	return &campaignRepository{
		dao,
	}
}

func (repo *campaignRepository) Store(campaignCollection domain.CampaignCollection) error {
	return repo.dao.Store(campaignCollection)
}

func (repo *campaignRepository) Get(campaignName string) (domain.Campaign, error) {
	return repo.dao.Get(campaignName)
}

func (repo *campaignRepository) Search(profile domain.ProfileCollection) (domain.CampaignCollection, error) {
	return repo.dao.Search(profile)
}
