package queries

import "testing"

func TestUserShouldReturnValue(t *testing.T) {
	u := (&Query{}).User()
	if u == nil {
		t.Fatal("User returned nil")
	}
}
