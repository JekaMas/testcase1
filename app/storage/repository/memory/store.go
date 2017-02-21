package memory

import "generator/app/domain"

func (s *store) Store(campaignCollection domain.CampaignCollection) error {
	s.storageLock.Lock()
	for _, campaign := range campaignCollection {
		storageInstance.campaigns[campaign.CampaignName] = campaign
	}
	s.storageLock.Unlock()

	var (
		ok            bool
		attributeMap  attributeIndex
		campaignNames campaignAttributeCount
	)

	s.indexLock.Lock()
	for _, campaign := range campaignCollection {
		for _, target := range campaign.TargetList {
			attributeMap, ok = storageInstance.searchIndex[target.AttributeName]

			if !ok {
				attributeMap = make(attributeIndex)
			}

			for _, attribute := range target.Attributes {
				campaignNames, ok = attributeMap[attribute]
				if !ok {
					campaignNames = make(campaignAttributeCount)
				}


				campaignNames[campaign.CampaignName] = len(campaign.TargetList)

				attributeMap[attribute] = campaignNames
			}

			storageInstance.searchIndex[target.AttributeName] = attributeMap
		}
	}
	s.indexLock.Unlock()

	return nil
}