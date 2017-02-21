package memory

import "generator/app/domain"

func (s *store) Get(campaignName string) (domain.Campaign, error) {
	s.storageLock.RLock()
	campaign, ok := s.campaigns[campaignName]
	s.storageLock.RUnlock()

	if !ok {
		return domain.Campaign{}, NoCampaignsError(campaignName)
	}

	return campaign, nil
}
