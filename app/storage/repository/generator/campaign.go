package generator

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"generator/app/domain"
)

const (
	maxCampaigns  = 10000
	MaxTargets    = 26
	MaxAttributes = 100

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

	if numTargets > MaxTargets || numTargets <= 0 {
		return nil, fmt.Errorf("Number of targets should be positive less or equal %v. Given %v.", MaxTargets, numTargets)
	}

	if numAttributes > MaxAttributes || numAttributes <= 0 {
		return nil, fmt.Errorf("Number of attributes should be positive less or equal %v. Given %v.", MaxAttributes, numAttributes)
	}

	r := getRandomGenerator()

	campaigns := make(domain.CampaignCollection, 0, numCampaigns)
	for i := 0; i < numCampaigns; i++ {
		campaigns = append(campaigns, domain.Campaign{
			CampaignName: domain.CampaignPrefix + strconv.Itoa(int(campaignIDGenerator.Get())),
			Price:        float64(r.Intn(maxPrice)) / 100,
			TargetList:   newTargetCollection(numTargets, numAttributes, r),
		})
	}

	if !campaigns.Verify() {
		log.Printf("Incorrect Campaigns has been generated %#+v\n."+
			" numCampaigns %v,numTargets %v, numAttributes %v\n", campaigns, numCampaigns, numTargets, numAttributes)
		return nil, errors.New("Incorrect Campaigns has been generated")
	}

	return campaigns, nil
}
