package queries

// This is the root queries of the graph
type Query struct { }

func (q *Query) User() *User {
	return &User { firstName: "John" }
}

func (_ *Query) GetQuery() string {
	return `
	type Query {
		user: User
	}
`
}