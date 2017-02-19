package generator

import (
	"generator/app/domain"
	"strconv"
	"log"
	"errors"
)

var (
	userIDGenerator      IGenerate
	attributeIDGenerator IGenerate
)

const maxAttributeNum = 200

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
		attribute string
		attributeNumber = getAttributeNumber()
	)
	for i := 0; i <= attributeNumber; i++ {
		attribute = string('A' + i)
		u.Profile[domain.AttributePrefix+attribute] = domain.Attribute(attribute + strconv.Itoa(r.Intn(maxAttributeNum)))
	}

	if !u.Verify() {
		log.Printf("Incorrect uset has been generated %#+v\n.", u)
		return nil, errors.New("Incorrect user has been generated")
	}

	return u, nil
}

func getAttributeNumber() int {
	return int(attributeIDGenerator.Get()) % 26
}
