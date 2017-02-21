package generator

type (
	userIDGeneratorKey          struct{}
	userAttributeIDGeneratorKey struct{}
	campaignsIDGeneratorKey     struct{}
	searchResultIDGeneratorKey     struct{}
)

var (
	userID      = userIDGeneratorKey{}
	attributeID = userAttributeIDGeneratorKey{}
	campaignsID = campaignsIDGeneratorKey{}
	searchResultID = searchResultIDGeneratorKey{}
)
