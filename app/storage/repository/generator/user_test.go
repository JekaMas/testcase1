package generator

import (
	"generator/app/domain"
	"strconv"
	"strings"
	"testing"
)

func Test_CorrectNewUser_Success(t *testing.T) {
	userIDGenerator.Reset()
	attributeIDGenerator.Reset()

	u1, err := NewUser()
	if err != nil {
		t.Fatal(err)
	}

	if !u1.Verify() {
		t.Fatalf("Correct User should be generated %#+v", u1)
	}
}

func Test_NewUserCountIncrease_Success(t *testing.T) {
	userIDGenerator.Reset()
	attributeIDGenerator.Reset()

	const count = 28*2 + 2

	var (
		u            *domain.User
		err          error
		userIDString string
		userID       int
		prevUserID   int = -1
	)

	for i := 0; i < count; i++ {
		u, err = NewUser()
		if err != nil {
			t.Fatal(err)
		}

		userIDString = strings.TrimLeft(u.User, domain.UserPrefix)

		userID, err = strconv.Atoi(userIDString)
		if err != nil {
			t.Fatalf("User ID should be number increased by cycle: 0..25. Got ID %v on %v user", u.User, i)
		}

		if userID <= prevUserID {
			t.Fatalf("User ID should be increasing number. Got ID %v on %v user", u.User, i)
		}
		prevUserID = userID

		if len(u.Profile) != i%26+1 {
			t.Fatalf("User should have an increasing number of profiles: 1..26. Got %v on %v user", len(u.Profile), i)
		}
	}
}
