package memory

import (
	"generator/app/domain"
	"generator/app/storage/repository/generator"
	"testing"
)

func Benchmark_SearchCampaigns(b *testing.B) {
	s := NewStorage()

	campaignName1 := "campaign4"
	campaign1 := domain.Campaign{
		CampaignName: campaignName1,
		Price:        10.0,
		TargetList: domain.TargetCollection{
			domain.Target{
				AttributeName: "attr_F",
				Attributes:    domain.AttributeCollection{"F0"},
			},
		},
	}

	campaignName2 := "campaign5"
	campaign2 := domain.Campaign{
		CampaignName: campaignName2,
		Price:        10.0,
		TargetList: domain.TargetCollection{
			domain.Target{
				AttributeName: "attr_G",
				Attributes:    domain.AttributeCollection{"G1"},
			},
			domain.Target{
				AttributeName: "attr_H",
				Attributes:    domain.AttributeCollection{"H1"},
			},
		},
	}

	campaings := domain.CampaignCollection{campaign1, campaign2}

	if err := s.Store(campaings); err != nil {
		b.Fatal(err)
	}

	user := &domain.User{
		User: "u1000",
		Profile: domain.ProfileCollection{
			"attr_C": "C0",
			"attr_E": "E1",
		},
	}

	var (
		err          error
		campaignFind domain.CampaignCollection
	)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		campaignFind, err = s.Search(user.Profile)
		if err != nil {
			b.Fatal(err)
		}
	}

	if len(campaignFind) != 1 {
		b.Fatal(campaignFind)
	}
}

func Benchmark_SearchCampaigns_Parallel(b *testing.B) {
	campaings, err := generator.NewCampaignCollection(10, 10, 50)
	if err != nil {
		b.Fatal(err)
	}

	s := NewStorage()
	if err := s.Store(campaings); err != nil {
		b.Fatal(err)
	}

	user := &domain.User{
		User: "u1000",
		Profile: domain.ProfileCollection{
			"attr_C": "C0",
			"attr_E": "E1",
		},
	}

	var campaignFind domain.CampaignCollection

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			campaignFind, err = s.Search(user.Profile)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	//FIXME make reasonable number
	if len(campaignFind) == 100000 {
		b.Fatal(campaignFind)
	}
}
