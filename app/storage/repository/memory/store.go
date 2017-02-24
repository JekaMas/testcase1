package memory

import "generator/app/domain"

func (s *store) Store(campaignCollection domain.CampaignCollection) error {
	s.storageLock.Lock()
	for _, campaign := range campaignCollection {
		storageInstance.campaigns[campaign.CampaignName] = campaign
	}
	s.storageLock.Unlock()

	var (
		attributeMap  *attributeIndex
		campaignNames campaignAttributeCount
	)

	s.indexLock.Lock()
	for _, campaign := range campaignCollection {
		for _, target := range campaign.TargetList {
			attributeMap = storageInstance.searchIndex.getAttributes(target.AttributeName)

			for _, attribute := range target.Attributes {
				campaignNames = attributeMap.getCampaigns(attribute)
				if campaignNames == nil {
					campaignNames = make(campaignAttributeCount)
				}
				campaignNames[campaign.CampaignName] = len(campaign.TargetList)

				attributeMap.setCampaigns(attribute, campaignNames)
			}

			//storageInstance.searchIndex[target.AttributeName] = attributeMap
		}
	}
	s.indexLock.Unlock()

	return nil
}
