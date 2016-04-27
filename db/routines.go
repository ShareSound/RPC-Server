package db

import (
	"github.com/ShareSound/RPC-Server/rpc/shared"
	common "github.com/ShareSound/RPC-Server/shared"
	"github.com/garyburd/redigo/redis"
)

func GetAuthInfo(session *shared.Session) (string, *shared.AuthException) {
	conn := GetNewAccountAuthConn()
	defer conn.Connection.Close()

	token := session.GetAuthToken()
	if result_str, e := redis.String(conn.Connection.Do("GET", token)); e == nil {
		return result_str, nil
	}else{
		return "", &shared.AuthException{
			Message: "Auth Failed",
		}
	}
}

func UpdateAuthToken(session *shared.Session, value string) error {
	conn := GetNewAccountAuthConn()
	defer conn.Connection.Close()

	old_token := session.GetAuthToken()
	if len(old_token) > 0 {
		//Delete old token
		conn.Connection.Do("DEL", old_token)
	}

	session.AuthToken = shared.AuthToken(common.GetSecureHash(20))
	_, e := conn.Connection.Do("SETEX", session.GetAuthToken(), ACCOUNT_AUTH_KEY_EXPIRE_SEC, value)
	return e
}
