package memory

import (
	"generator/app/domain"
)

func (s *store) Search(profile domain.ProfileCollection) (domain.CampaignCollection, error) {
	var (
		campaign           domain.Campaign
		campaigns          domain.CampaignCollection
		campaignNames      = make(campaignAttributeCount)
		foundCampaignNames = make(map[string]int)

		ok             bool
		attrIndex      *attributeIndex
		campaignsFound campaignAttributeCount
	)

	// TODO уменьшить сложность
	for attributeName, attr := range profile {
		attrIndex = s.searchIndex.getAttributes(attributeName)
		campaignsFound = attrIndex.getCampaigns(attr)

		s.indexLock.RLock()
		for foundCampaignName, count := range campaignsFound {
			foundCampaignNames[foundCampaignName]++

			//store targets count by campaigns
			campaignNames[foundCampaignName] = count
		}
		s.indexLock.RUnlock()
	}

	if len(foundCampaignNames) == 0 {
		return nil, nil
	}

	// Filter campaigns that doesnt fit user profile
	for foundCampaignName, count := range foundCampaignNames {
		if count != campaignNames[foundCampaignName] {

			delete(campaignNames, foundCampaignName)
		}
	}

	for campaignName := range campaignNames {

		s.storageLock.RLock()
		if campaign, ok = s.campaigns[campaignName]; ok {

			campaigns = append(campaigns, campaign)
		}
		s.storageLock.RUnlock()
	}

	return campaigns, nil
}
