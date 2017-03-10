package domain

import (
	"strings"
)

//easyjson:json
type User struct {
	User    string            `json:"user"`
	Profile ProfileCollection `json:"profile"`
}

func (u *User) Verify() bool {
	return compose(
		u.IsCorrectName,
		u.VerifyProfile,
	)()
}

func (u *User) VerifyProfile() bool {
	return u.Profile.Verify()
}

const (
	UserPrefix         = "u"
	userNumberPosition = len(UserPrefix)
)

func (u *User) IsCorrectName() bool {
	var ok bool

	if ok = strings.HasPrefix(u.User, UserPrefix); !ok {
		return false
	}

	if ok = IsLatinNumber(u.User, userNumberPosition); !ok {
		return false
	}

	return true
}
