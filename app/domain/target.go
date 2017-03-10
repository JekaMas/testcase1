package domain

import (
	"strings"
)

//easyjson:json
type Target struct {
	AttributeName string              `json:"target"`
	Attributes    AttributeCollection `json:"attr_list"`
}

func (target *Target) Verify() bool {
	return compose(
		target.IsCorrectFormat,
		target.IsAttributesValid,
	)()
}

func (target *Target) IsAttributesValid() bool {
	return target.Attributes.Verify()
}

func (target *Target) IsCorrectFormat() bool {
	return compose(
		target.IsFirstCapitalLetter,
		target.IsAttributeNameMatched,
	)()
}

func (target *Target) IsFirstCapitalLetter() bool {
	if len(target.AttributeName) <= 1 {
		return false
	}

	if ok := strings.HasPrefix(target.AttributeName, AttributePrefix); !ok {
		return false
	}

	if ok := IsLatinUpper(target.AttributeName, attributeNumberPosition); !ok {
		return false
	}

	return true
}

func (target *Target) IsAttributeNameMatched() bool {
	for _, attr := range target.Attributes {
		if attr[0] != target.AttributeName[attributeNumberPosition] {
			return false
		}
	}

	return true
}
