package generator

import (
	"generator/app/domain"
	"testing"
	"reflect"
	"fmt"
)

func Test_IncorrectCampaignNumber_Fail(t *testing.T) {
	cases := []int{-1, 0, maxCampaigns + 1, maxCampaigns * 10}

	var err error
	for _, testCase := range cases {
		_, err = NewCampaignCollection(testCase, 1, 1)
		if err == nil {
			t.Fatal("Error expected ad value", testCase)
		}
	}
}

func Test_IncorrectTargetsNumber_Fail(t *testing.T) {
	cases := []int{-1, 0, maxTargets + 1, maxTargets * 10}

	var err error
	for _, testCase := range cases {
		_, err = NewCampaignCollection(1, testCase, 1)
		if err == nil {
			t.Fatal("Error expected ad value", testCase)
		}
	}
}

func Test_IncorrectAttributesNumber_Fail(t *testing.T) {
	cases := []int{-1, 0, maxAttributes + 1, maxAttributes * 10}

	var err error
	for _, testCase := range cases {
		_, err = NewCampaignCollection(1, 1, testCase)
		if err == nil {
			t.Fatal("Error expected ad value", testCase)
		}
	}
}

func Test_GenerateCampaigns_Success(t *testing.T) {
	cases := []struct {
		campaignNum  int
		targetNum    int
		attributeNum int
		expected     domain.CampaignCollection
	}{
		{
			campaignNum:  1,
			targetNum:    1,
			attributeNum: 1,
			expected: domain.CampaignCollection{
				{
					CampaignName: "campaign0",
					TargetList: domain.TargetCollection{
						domain.Target{
							AttributeName: "attr_A",
							Attributes:    domain.AttributeCollection{"A0"},
						},
					},
				},
			},
		},
	}

	var (
		err error
		res domain.CampaignCollection
	)
	for _, testCase := range cases {
		res, err = NewCampaignCollection(testCase.campaignNum, testCase.targetNum, testCase.attributeNum)
		if err != nil {
			t.Fatal("No errors expected. Got %#+v", err)
		}

		if err = priceCheck(&res); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(testCase.expected, res) {
			t.Fatalf("Expected %#+v\nGiven %#+v\n", testCase.expected, res)
		}
	}
}

// Checks campaigns prices and set them to 0
func priceCheck(campaigns *domain.CampaignCollection) error {
	updatedCampaigns := make(domain.CampaignCollection, 0, len(*campaigns))

	for _, campaign := range *campaigns {
		if campaign.Price <= domain.PriceThreshold {
			return fmt.Errorf("Expect positive prices. Got %v", campaign.Price)
		}

		campaign.Price = 0

		updatedCampaigns = append(updatedCampaigns, campaign)
	}

	*campaigns = updatedCampaigns

	return nil
}