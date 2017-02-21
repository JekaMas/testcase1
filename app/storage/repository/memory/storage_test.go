package memory

import (
	"reflect"
	"testing"

	"generator/app/domain"
)

func Test_StoreCampaigns_Success(t *testing.T) {
	s := NewStorage()

	campaings := domain.CampaignCollection{
		{
			CampaignName: "campaign0",
			Price:        10,
			TargetList: domain.TargetCollection{
				domain.Target{
					AttributeName: "attr_A",
					Attributes:    domain.AttributeCollection{"A0"},
				},
			},
		},
	}

	if err := s.Store(campaings); err != nil {
		t.Fatal(err)
	}
}

func Test_GetCampaigns_Success(t *testing.T) {
	s := NewStorage()

	campaignName := "campaign1"
	campaign := domain.Campaign{
		CampaignName: campaignName,
		Price:        10.0,
		TargetList: domain.TargetCollection{
			domain.Target{
				AttributeName: "attr_B",
				Attributes:    domain.AttributeCollection{"B0"},
			},
		},
	}
	campaings := domain.CampaignCollection{campaign}

	if err := s.Store(campaings); err != nil {
		t.Fatal(err)
	}

	campaignStored, err := s.Get(campaignName)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(campaign, campaignStored) {
		t.Fatalf("Expected %#+v\nGiven %#+v\n", campaign, campaignStored)
	}
}

func Test_SearchCampaigns_Success(t *testing.T) {
	s := NewStorage()

	campaignName1 := "campaign2"
	campaign1 := domain.Campaign{
		CampaignName: campaignName1,
		Price:        10.0,
		TargetList: domain.TargetCollection{
			domain.Target{
				AttributeName: "attr_C",
				Attributes:    domain.AttributeCollection{"C0"},
			},
		},
	}

	campaignName2 := "campaign3"
	campaign2 := domain.Campaign{
		CampaignName: campaignName2,
		Price:        10.0,
		TargetList: domain.TargetCollection{
			domain.Target{
				AttributeName: "attr_D",
				Attributes:    domain.AttributeCollection{"D1"},
			},
			domain.Target{
				AttributeName: "attr_E",
				Attributes:    domain.AttributeCollection{"E1"},
			},
		},
	}

	campaings := domain.CampaignCollection{campaign1, campaign2}

	if err := s.Store(campaings); err != nil {
		t.Fatal(err)
	}

	user := &domain.User{
		User: "u1000",
		Profile: domain.ProfileCollection{
			"attr_C": "C0",
			"attr_E": "E1",
		},
	}

	campaignFind, err := s.Search(user.Profile)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(domain.CampaignCollection{campaign1}, campaignFind) {
		t.Fatalf("Expected %#+v\nGiven %#+v\n", domain.CampaignCollection{campaign1}, campaignFind)
	}
}
