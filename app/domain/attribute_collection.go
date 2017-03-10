package domain

import (
	"log"
	"strconv"
)

//easyjson:json
type AttributeCollection []Attribute

func (c AttributeCollection) Verify() bool {
	return compose(
		c.IsSorted,
		c.VerifyAttributes,
	)()
}

func (c AttributeCollection) IsSorted() bool {
	var (
		prev int
		next int
		err  error
	)

	for i := 0; i < len(c)-1; i++ {
		prev, err = strconv.Atoi(string(c[i][1:]))
		if err != nil {
			log.Printf("Error in attributes: %#+v. On value: %q", c, c[i])
		}

		next, err = strconv.Atoi(string(c[i+1][1:]))
		if err != nil {
			log.Printf("Error in attributes: %#+v. On value: %q", c, c[i+1])
		}

		if prev > next {
			return false
		}
	}

	return true
}

func (c AttributeCollection) VerifyAttributes() bool {
	var ok bool
	for _, attr := range c {
		if ok = attr.Verify(); !ok {
			return false
		}
	}

	return true
}
