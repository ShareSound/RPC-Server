package handlers

import (
	"github.com/ShareSound/RPC-Server/rpc/shared"
	"github.com/ShareSound/RPC-Server/rpc/account"
	"github.com/ShareSound/RPC-Server/db"
	"github.com/ShareSound/RPC-Server/db/datastore_schemas"
	common "github.com/ShareSound/RPC-Server/shared"
	"google.golang.org/cloud/datastore"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const(
	USER_ID_HASH_LENGTH = 25
)

type AccountServiceHandler struct{}

func NewAccountServiceHandler() *AccountServiceHandler {
	return &AccountServiceHandler{}
}

func (this *AccountServiceHandler) RegisterAccount(email , username , password string) (r *shared.Session, err error){
	if len(email) == 0 || len(username) == 0 || len(password) == 0{
		//Format error
		return nil, &shared.AuthException{
			Message: "Format error",
		}
	}

	if client,err := db.GetNewDataStoreClient(); err == nil {

		//Ensure account and hash id not exist
		q := datastore.NewQuery(db.DATASTORE_ACCOUNT_USER_PROFILE_KIND).
				Filter("Email =",email)
		if k,_ := client.Run(q).Next(nil); k != nil {
			//User exist
			return nil,&shared.AuthException{
				Message: "User exist",
			}
		}

		//Ensure id not taken
		var id_hash string = ""
		var key *datastore.Key
		for{
			id_hash = common.GetSecureHash(USER_ID_HASH_LENGTH)
			key = datastore.NewKey(client.Ctx, db.DATASTORE_ACCOUNT_USER_PROFILE_KIND, id_hash, 0, nil)
			p := datastore_schemas.User{}
			if e := client.Client.Get(client.Ctx, key, &p); e == datastore.ErrNoSuchEntity{
				break;
			}
		}

		pwd_cost := bcrypt.DefaultCost
		pwd_hash,_ := bcrypt.GenerateFromPassword([]byte(password), pwd_cost)
		profile := datastore_schemas.User{
			Id: datastore_schemas.HashId(id_hash),

			Username: username,
			Email: email,
			CreatedTimeStamp: time.Now(),
			Auth: datastore_schemas.UserAuthInfo{
				PasswordBcryptHash: string(pwd_hash),
				PasswordBcryptCost: pwd_cost,
			},
		}

		if _, e := client.Client.Put(client.Ctx, key, &profile); e != nil {
			common.LogE.Println("Insert new user failed: " + e.Error())
			return nil, &shared.AuthException{
				Message: "Internal Error: " + e.Error(),
			}
		}

		session := shared.NewSession()
		if e := db.UpdateAuthToken(session, string(profile.Id)); e != nil {
			common.LogE.Println("Update auth token failed: " + e.Error())
		}

		return session,nil
	}else{
		return nil, &shared.AuthException{
			Message: "Internal Error: " + err.Error(),
		}
	}

	return nil,nil
}

func (this *AccountServiceHandler) Login(email string, password string) (r *shared.Session, err error) {

	if client, e := db.GetNewDataStoreClient(); e == nil {
		q := datastore.NewQuery(db.DATASTORE_ACCOUNT_USER_PROFILE_KIND).
				Filter("Email =", email)
		it := client.Run(q)
		var profile datastore_schemas.User
		//Only pick one
		if _, err := it.Next(&profile); err != datastore.Done && err != nil {
			return nil, &shared.AuthException{
				Message: "Email or Password Error",
			}
		}

		pwd_hash := profile.Auth.PasswordBcryptHash
		if e := bcrypt.CompareHashAndPassword([]byte(pwd_hash), []byte(password)); e != nil {
			return nil, &shared.AuthException{
				Message: "Email or Password Error",
			}
		}

		//Create new session
		session := shared.NewSession()
		if e := db.UpdateAuthToken(session, string(profile.Id)); e != nil {
			common.LogE.Println("Update auth token failed: " + e.Error())
		}

		return session,nil
	}else{
		return nil, &shared.AuthException{
			Message: "Internal Error: " + err.Error(),
		}
	}
}

func (this *AccountServiceHandler) Logout(ctx *shared.Session) (err error) {
	id, e := db.GetAuthInfo(ctx)
	if e != nil {
		return e
	}

	conn := db.GetNewAccountAuthConn()
	defer conn.Connection.Close()

	if _, e := conn.DelCmd(id); e != nil {
		return &shared.AuthException{
			Message: "Internal Error",
		}
	}

	return nil
}

func (this *AccountServiceHandler) GetProfile(ctx *shared.Session) (r *account.ProfileResult_, err error) {
	id, e := db.GetAuthInfo(ctx)
	if e != nil { return nil, e }

	if client,err := db.GetNewDataStoreClient(); err == nil {
		key := datastore.NewKey(client.Ctx, db.DATASTORE_ACCOUNT_USER_PROFILE_KIND, id, 0, nil)
		profile := datastore_schemas.User{}
		if e := client.Client.Get(client.Ctx, key, &profile); e != nil {
			return nil, &shared.AuthException{
				Message: "Retreive Profile Error: " + e.Error(),
			}
		}

		if e := db.UpdateAuthToken(ctx, string(profile.Id)); e != nil {
			common.LogE.Println("Update auth token failed: " + e.Error())
		}
		result := account.ProfileResult_{
			Session: ctx,

			Email: profile.Email,
			Username: profile.Username,
		}

		return &result, nil
	}else{
		return nil, &shared.AuthException{
			Message: "Internal Error: " + err.Error(),
		}
	}
}


