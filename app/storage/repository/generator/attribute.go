package generator

import (
	"generator/app/domain"
	"strconv"
)

func newAttributeCollection(attributePrefix string, numAttributes int) domain.AttributeCollection {
	setLesserRandom(&numAttributes)

	attributes := make(domain.AttributeCollection, 0, numAttributes)
	for i := 0; i < numAttributes; i++ {
		attributes = append(attributes, domain.Attribute(attributePrefix+strconv.Itoa(i)))
	}

	return attributes
}
