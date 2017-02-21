package services

import (
	"generator/app/domain"
	"generator/app/storage/repository"
	"generator/app/storage/repository/generator"
)

//GetService - service to get campaigns
type GetService struct {
	repo repository.CampaignRepository
}

//NewGetService - construct service to get campaigns
func NewGetService(repo repository.CampaignRepository) GetService {
	return GetService{repo}
}

func (this GetService) Get(attributesNum int, targetsNum int, campaignsNum int) (domain.CampaignCollection, error) {
	campaigns, err := generator.NewCampaignCollection(attributesNum, targetsNum, campaignsNum)
	if err != nil {
		return nil, err
	}

	err = this.repo.Store(campaigns)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}
