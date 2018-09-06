package queries

import "testing"

func TestUser_GetQuery(t *testing.T) {
	q := (&User{}).GetQuery()
	if q !=  `
	type User {
		firstName: String!
	}` {
		t.Fatal("User query invalid")
	}

}

func TestUser_FirstName(t *testing.T) {
	u := &User{ firstName: "Rusith"}
	if u.FirstName() != "Rusith" {
		t.Fatalf("First name invalid")
	}
}
