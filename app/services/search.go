package services

import (
	"generator/app/storage/repository"
	"generator/app/storage/repository/generator"
	"generator/app/domain"
)

type SearchService struct {
	repo repository.CampaignRepository
}

func NewSearchService(repo repository.CampaignRepository) SearchService {
	return SearchService{repo}
}

const noneWinner = "none"

func (this SearchService) Search(user domain.User) (domain.SearchResult, error) {
	campaigns, err := this.repo.Search(user.Profile)
	if err != nil {
		return domain.SearchResult{}, err
	}

	result := domain.SearchResult{
		Counter: generator.SearchResultID(),
	}

	maxPriceCampaign, ok := this.getMaxPriceCampaign(campaigns)
	if !ok {
		result.Winner = noneWinner
		return result, nil
	}

	result.Winner = maxPriceCampaign.CampaignName

	return result, nil
}

func (this SearchService) getMaxPriceCampaign(collection domain.CampaignCollection) (domain.Campaign, bool) {
	if len(collection) == 0 {
		return domain.Campaign{}, false
	}

	maxPriceCampaign := collection[0]

	for _, campaign := range collection {
		if (maxPriceCampaign.Price - campaign.Price) < domain.PriceThreshold {
			maxPriceCampaign = campaign
		}
	}

	return maxPriceCampaign, true
}
