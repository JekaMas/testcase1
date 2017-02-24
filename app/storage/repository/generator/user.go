package generator

import (
	"strconv"

	"generator/app/domain"
)

var (
	userIDGenerator      Generate
	attributeIDGenerator Generate
)

const MaxAttributeNum = 200

func init() {
	userIDGenerator = Get(userID)
	attributeIDGenerator = Get(attributeID)
}

func NewUser() (*domain.User, error) {
	u := &domain.User{
		User:    domain.UserPrefix + strconv.Itoa(int(userIDGenerator.Get())),
		Profile: make(domain.ProfileCollection),
	}

	var (
		attribute       string
		attributeNumber = getAttributeNumber()
	)

	r := getRandomGenerator()
	for i := 0; i <= attributeNumber; i++ {
		attribute = string('A' + i)
		u.Profile[domain.AttributePrefix+attribute] = domain.Attribute(attribute + strconv.Itoa(r.Intn(MaxAttributeNum)))
	}

	return u, nil
}

func getAttributeNumber() int {
	return int(attributeIDGenerator.Get()) % 26
}
