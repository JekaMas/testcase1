package generator

import (
	"generator/app/domain"
)

const firstLatinLetter = 'A'

func newTargetCollection(numTargets, numAttributes int) domain.TargetCollection {
	setLesserRandom(&numTargets)

	targets := make(domain.TargetCollection, 0, numTargets)
	var attributePrefix string

	for i := 0; i < numTargets; i++ {
		attributePrefix = string(firstLatinLetter + i)

		targets = append(targets, domain.Target{
			AttributeName: domain.AttributePrefix + attributePrefix,
			Attributes:    newAttributeCollection(attributePrefix, numAttributes),
		})
	}

	return targets
}
