package memory

import (
	"generator/app/domain"
	"generator/app/storage"
	"sync"
)

// From AttributeName to map[domain.Attribute]campaignAttributeCount
type searchIndex map[string]attributeIndex

// map[domain.Attribute]campaignAttributeCount
type attributeIndex map[domain.Attribute]campaignAttributeCount

// map[CampaignName]int
type campaignAttributeCount map[string]int

// map[CampaignName]domain.Campaign
type campaigns map[string]domain.Campaign

type store struct {
	searchIndex
	campaigns

	indexLock   sync.RWMutex
	storageLock sync.RWMutex
}

var storageInstance = &store{
	make(searchIndex),
	make(campaigns),
	sync.RWMutex{},
	sync.RWMutex{},
}

func NewStorage() storage.CampaignsDAO {
	return storageInstance
}

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

func (s *store) Get(campaignName string) (domain.Campaign, error) {
	s.storageLock.RLock()
	campaign, ok := s.campaigns[campaignName]
	s.storageLock.RUnlock()

	if !ok {
		return domain.Campaign{}, NoCampaignsError(campaignName)
	}

	/*
	if !campaign.Verify() {
		log.Printf("Incorrect Campaign in base: %#+v", campaign)
		return domain.Campaign{}, NoCampaignsError(campaignName)
	}
	*/

	return campaign, nil
}

func (s *store) Search(profile domain.ProfileCollection) (domain.CampaignCollection, error) {
	var (
		campaign           domain.Campaign
		campaigns          domain.CampaignCollection
		campaignNames      = make(campaignAttributeCount)
		foundCampaignNames = make(map[string]int)
	)

	s.indexLock.RLock()
	for attributeName, attr := range profile {
		if attrIndex, ok := s.searchIndex[attributeName]; ok {
			if campaigns, ok := attrIndex[attr]; ok {
				for foundCampaignName, count := range campaigns {
					foundCampaignNames[foundCampaignName]++
					campaignNames[foundCampaignName] = count
				}
			}
		}
	}
	s.indexLock.RUnlock()

	//пройти по foundCampaignNames
	// посчитать сколько каждой campaingName встречается
	for foundCampaignName, count := range foundCampaignNames {
		// если это отлично от len(profile), то campaingName не берем
		// иначе берем в 1 экземпляре

		if count != campaignNames[foundCampaignName] {
			delete(campaignNames, foundCampaignName)
		}
	}

	var ok bool
	s.storageLock.RLock()
	for campaignName := range campaignNames {
		if campaign, ok = s.campaigns[campaignName]; ok {

			/*
			if !campaign.Verify() {
				log.Printf("Incorrect Campaign in base: %#+v", campaign)
				continue
			}
			*/

			campaigns = append(campaigns, campaign)
		}
	}
	s.storageLock.RUnlock()

	return campaigns, nil
}
