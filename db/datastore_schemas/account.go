package datastore_schemas

import "time"

type UserAuthInfo struct {
	PasswordBcryptHash	string
	PasswordBcryptCost	int
}
type User struct {
	Id			HashId

	Username		string
	Email			string
	CreatedTimeStamp	time.Time

	Auth			UserAuthInfo
}
