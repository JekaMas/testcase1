package generator

import (
	"fmt"
	"generator/app/domain"
	"log"
	"strconv"
	"errors"
)

const (
	maxCampaigns  = 10000
	maxTargets    = 26
	maxAttributes = 100

	//maxPrice*100
	maxPrice = 1000000
)

var campaignIDGenerator Generate

func init() {
	campaignIDGenerator = Get(campaignsID)
}

func NewCampaignCollection(numCampaigns, numTargets, numAttributes int) (domain.CampaignCollection, error) {
	if numCampaigns > maxCampaigns || numCampaigns <= 0 {
		return nil, fmt.Errorf("Number of campaigns should be positive less or equal %v. Given %v.", maxCampaigns, numCampaigns)
	}

	if numTargets > maxTargets || numTargets <= 0 {
		return nil, fmt.Errorf("Number of targets should be positive less or equal %v. Given %v.", maxTargets, numTargets)
	}

	if numAttributes > maxAttributes || numAttributes <= 0 {
		return nil, fmt.Errorf("Number of attributes should be positive less or equal %v. Given %v.", maxAttributes, numAttributes)
	}

	r := getRandomGenerator()
	r.setLesserRandom(&numCampaigns)

	campaigns := make(domain.CampaignCollection, 0, numCampaigns)
	for i := 0; i < numCampaigns; i++ {
		campaigns = append(campaigns, domain.Campaign{
			CampaignName: domain.CampaignPrefix + strconv.Itoa(int(campaignIDGenerator.Get())),
			Price:        float64(r.Intn(maxPrice)) / 100,
			TargetList:   newTargetCollection(numTargets, numAttributes, r),
		})
	}

	if !campaigns.Verify() {
		log.Printf("Incorrect Campaigns has been generated %#+v\n." +
			" numCampaigns %v,numTargets %v, numAttributes %v\n", campaigns, numCampaigns, numTargets, numAttributes)
		return nil, errors.New("Incorrect Campaigns has been generated")
	}

	return campaigns, nil
}
