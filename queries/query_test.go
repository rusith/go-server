package queries

import "testing"

func TestUserShouldReturnValue(t *testing.T) {
	u := (&Query{}).User()
	if u == nil {
		t.Fatal("User returned nil")
	}
}

func TestQuery_GetQuery(t *testing.T) {
	u := (&Query{}).GetQuery()
	if u != `
	type Query {
		user: User
	}
` {
		t.Fatal("User query is invalid")
	}
}