package generator

var (
	searchResultIDGenerator Generate
)

func init() {
	searchResultIDGenerator = Get(searchResultID)
}

func SearchResultID() int {
	return int(searchResultIDGenerator.Get())
}