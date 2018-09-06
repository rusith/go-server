package queries

// The user type, represents a user
type User struct { firstName string }

func (u *User) FirstName() string {
	return u.firstName
}

func (_ *User) GetQuery() string {
	return `
	type User {
		firstName: String!
	}`
}
