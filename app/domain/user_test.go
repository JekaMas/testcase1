package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUserMarshal(t *testing.T) {
	u := &User{
		User: "u3",
		Profile: ProfileCollection{
			"attr_C": Attribute("C45"),
			"attr_A": Attribute("A23"),
			"attr_B": Attribute("B132"),
		},
	}
	expected := []byte(`{"user":"u3","profile":{"attr_A":"A23","attr_B":"B132","attr_C":"C45"}}`)

	res, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Got %q. Expect %q", res, expected)
	}

	u1 := &User{}
	err = json.Unmarshal(res, u1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(u, u1) {
		t.Fatal("Marshal and unmarshal should be consistant")
	}
}
