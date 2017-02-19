package domain

type Attribute string

func (attr Attribute) Verify() bool {
	return compose(
		attr.IsCorrectFormat,
	)()
}

// First latin letter + uint number
func (attr Attribute) IsCorrectFormat() bool {
	var ok bool

	if len(attr) < 2 {
		return false
	}

	if ok = IsLatinUpper(string(attr), 0); !ok {
		return false
	}

	if ok = IsLatinNumber(string(attr), 1); !ok {
		return false
	}

	if ok = IsLeadingZero(string(attr), 1); ok {
		return false
	}

	return true
}