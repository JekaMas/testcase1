package generator

import (
	"strconv"

	"generator/app/domain"
)

func newAttributeCollection(attributePrefix string, numAttributes int, r *randomGenerator) domain.AttributeCollection {
	r.setLesserRandom(&numAttributes)

	attributes := make(domain.AttributeCollection, 0, numAttributes)
	for i := 0; i < numAttributes; i++ {
		attributes = append(attributes, domain.Attribute(attributePrefix+strconv.Itoa(i)))
	}

	return attributes
}
