package memory

import (
	"sync"

	"generator/app/domain"
	"generator/app/storage"
	"generator/app/storage/repository/generator"
	"strconv"
)

// map[CampaignName]domain.Campaign
type campaigns map[string]domain.Campaign

// map[AttributeName]attributeIndex
type searchIndex [generator.MaxTargets]attributeIndex

// map[domain.Attribute]campaignAttributeCount
type attributeIndex [generator.MaxAttributeNum]campaignAttributeCount

// map[CampaignName]campaignAttributeCount
type campaignAttributeCount map[string]int

type store struct {
	campaigns
	*searchIndex

	indexLock   sync.RWMutex
	storageLock sync.RWMutex
}

const (
	firstLetter = 'A'
	firstNumber = '0'
)

func (index *searchIndex) getAttributes(name string) *attributeIndex {
	n := int(name[len(name)-1] - firstLetter)
	return &index[n]
}

func (attrIndex *attributeIndex) getCampaigns(name domain.Attribute) campaignAttributeCount {
	n, _ := strconv.Atoi(string(name[1:]))
	return attrIndex[n]
}

func (attrIndex *attributeIndex) setCampaigns(name domain.Attribute, campaignNames campaignAttributeCount) {
	n, _ := strconv.Atoi(string(name[1:]))
	attrIndex[n] = campaignNames
	return
}

var storageInstance = &store{
	make(campaigns),
	&searchIndex{},
	sync.RWMutex{},
	sync.RWMutex{},
}

func NewStorage() storage.CampaignsDAO {
	return storageInstance
}
