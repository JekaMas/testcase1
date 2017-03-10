package domain

import "strings"

//map[AttributeName]Attribute
//easyjson:json
type ProfileCollection map[string]Attribute

func (p ProfileCollection) Verify() bool {
	return compose(
		p.IsNotEmptyProfile,
		p.IsCorrectAttributeNames,
		p.IsAttributesMatched,
		p.VerifyAttributes,
	)()
}

const (
	AttributePrefix         = "attr_"
	attributeNumberPosition = len(AttributePrefix)
)

func (p ProfileCollection) IsAttributesMatched() bool {
	for name, attr := range p {
		if name[attributeNumberPosition] != attr[0] {
			return false
		}
	}

	return true
}

func (p ProfileCollection) IsNotEmptyProfile() bool {
	return len(p) > 0
}

func (p ProfileCollection) IsCorrectAttributeNames() bool {
	var ok bool
	for name := range p {
		if ok = strings.HasPrefix(name, AttributePrefix); !ok {
			return false
		}

		if ok = IsLatinUpper(name, attributeNumberPosition); !ok {
			return false
		}
	}

	return true
}

func (p ProfileCollection) VerifyAttributes() bool {
	var ok bool
	for _, attr := range p {
		if ok = attr.Verify(); !ok {
			return false
		}
	}

	return true
}
