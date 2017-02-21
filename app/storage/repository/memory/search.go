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
		attrIndex      attributeIndex
		campaignsFound campaignAttributeCount
	)

	for attributeName, attr := range profile {
		s.indexLock.RLock()

		if attrIndex, ok = s.searchIndex[attributeName]; ok {
			if campaignsFound, ok = attrIndex[attr]; ok {
				for foundCampaignName, count := range campaignsFound {
					foundCampaignNames[foundCampaignName]++

					//store targets count by campaigns
					campaignNames[foundCampaignName] = count
				}
			}
		}

		s.indexLock.RUnlock()
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
