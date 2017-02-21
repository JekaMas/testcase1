package services

import (
	"generator/app/domain"
	"generator/app/storage/repository"
)

//StoreCampaignsService service to store campaigns
type StoreCampaignsService struct {
	repo repository.CampaignRepository
}

//NewStoreCampaignsService construct service to store campaigns
func NewStoreCampaignsService(rep repository.CampaignRepository) StoreCampaignsService {
	return StoreCampaignsService{rep}
}

//StoreCampaigns in repository
func (this StoreCampaignsService) StoreCampaigns(campaigns domain.CampaignCollection) (bool, error) {
	err := this.repo.Store(campaigns)
	if err != nil {
		return false, err
	}

	return true, nil
}
