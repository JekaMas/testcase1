package memory

type NoCampaignsError string

func (e NoCampaignsError) Error() string {
	return "Not found campaign " + string(e)
}

func IsNoCampaignsError(err error) bool {
	_, ok := err.(NoCampaignsError)
	return ok
}
