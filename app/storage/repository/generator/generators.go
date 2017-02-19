package generator

type (
	userIDGeneratorKey          struct{}
	userAttributeIDGeneratorKey struct{}
	campaignsIDGeneratorKey     struct{}
)

var (
	userID      = userIDGeneratorKey{}
	attributeID = userAttributeIDGeneratorKey{}
	campaignsID = campaignsIDGeneratorKey{}
)
