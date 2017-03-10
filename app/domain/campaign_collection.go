package domain

//easyjson:json
type CampaignCollection []Campaign

func (c CampaignCollection) Verify() bool {
	var ok bool
	for _, campaign := range c {
		if ok = campaign.Verify(); !ok {
			return false
		}
	}

	return true

}
