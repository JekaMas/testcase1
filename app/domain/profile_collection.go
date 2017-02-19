package domain

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"
)

//map[AttributeName]Attribute
type ProfileCollection map[string]Attribute

//FIXME: add sync.pool for buffer
func (p ProfileCollection) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	buffer.WriteByte('{')

	//add sort
	names := make([]string, 0, len(p))
	for name := range p {
		names = append(names, name)
	}
	sort.Strings(names)

	var (
		i             int
		err           error
		marshaledPart []byte
	)

	for _, name := range names {
		marshaledPart, err = json.Marshal(name)
		if err != nil {
			return nil, err
		}
		buffer.Write(marshaledPart)

		buffer.WriteByte(':')

		marshaledPart, err = json.Marshal(p[name])
		if err != nil {
			return nil, err
		}
		buffer.Write(marshaledPart)

		i++
		if i < len(p) {
			buffer.WriteByte(',')
		}
	}

	buffer.WriteByte('}')

	return buffer.Bytes(), nil
}

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
