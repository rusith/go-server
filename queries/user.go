package queries

// The user type, represents a user
type UserType struct { firstName string }

func (u *UserType) FirstName() string {
	return u.firstName
}

func (_ *UserType) GetQuery() string {
	return `
	type User {
		firstName: String!
	}`
}
