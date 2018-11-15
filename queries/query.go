package queries

import "github.com/rusith/go-server/services"

// This is the root queries of the graph
type Query struct { userService services.UserService  }

func (*Query) New(userService services.UserService) *Query {
	return &Query{ userService }
}

func (q *Query) User() *UserType {
	return &UserType{ firstName: "John" }
}

func (_ *Query) GetQuery() string {
	return `
	type Query {
		user: User
	}
`
}