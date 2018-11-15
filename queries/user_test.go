package queries

import "testing"

//func TestUser_GetQuery(t *testing.T) {
//	q := (&UserType{}).GetQuery()
//	if q !=  `
//	type UserType {
//		firstName: String!
//	}` {
//		t.Fatal("UserType query invalid")
//	}
//
//}

func TestUser_FirstName(t *testing.T) {
	u := &UserType{ firstName: "Rusith"}
	if u.FirstName() != "Rusith" {
		t.Fatalf("First name invalid")
	}
}
