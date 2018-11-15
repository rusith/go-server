package queries

import (
	"github.com/rusith/go-server/services"
	"testing"
)

func TestUserShouldReturnValue(t *testing.T) {
	u := ((*Query).New(nil, services.UserService{ })).User()
	if u == nil {
		t.Fatal("User returned nil")
	}
}