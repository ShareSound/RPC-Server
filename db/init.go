package db

const(
	DATASTORE_ACCOUNT_USER_PROFILE_KIND = "account-user-profile"
)

func init(){

	initDatastore()

	initRedis()
}
