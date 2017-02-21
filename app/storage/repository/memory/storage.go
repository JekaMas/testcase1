package memory

import (
	"generator/app/domain"
	"generator/app/storage"
	"sync"
)

// map[CampaignName]domain.Campaign
type campaigns map[string]domain.Campaign

// map[AttributeName]attributeIndex
type searchIndex map[string]attributeIndex

// map[domain.Attribute]campaignAttributeCount
type attributeIndex map[domain.Attribute]campaignAttributeCount

// map[CampaignName]campaignAttributeCount
type campaignAttributeCount map[string]int

type store struct {
	campaigns
	searchIndex

	indexLock   sync.RWMutex
	storageLock sync.RWMutex
}

var storageInstance = &store{
	make(campaigns),
	make(searchIndex),
	sync.RWMutex{},
	sync.RWMutex{},
}

func NewStorage() storage.CampaignsDAO {
	return storageInstance
}
