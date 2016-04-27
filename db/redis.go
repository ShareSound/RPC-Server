package db

import (
	"os"
	"fmt"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"time"
)

const(
	REDIS_MASTER_ADDRESS_ENV = "REDIS_MASTER_SERVICE_HOST"
	REDIS_MASTER_PORT_ENV = "REDIS_MASTER_SERVICE_PORT"

	ACCOUNT_AUTH_REDIS_DB = 0
	ACCOUNT_AUTH_KEY_EXPIRE_SEC = 3600 * 24 * 3 //3 days
)

var(
	accountAuthConnPool *redis.Pool
	redisHost string
	redisPort string
)

//Wrapper for future extension
type RedisConn struct {
	Connection redis.Conn
}
func (this *RedisConn) GetCmd(key string) (interface{}, error){
	return this.Connection.Do("GET", key)
}
func (this *RedisConn) SetCmd(key, value string) (interface{}, error) {
	//'Vanilla' SET command
	return this.Connection.Do("SET", key, value)
}
func (this *RedisConn) DelCmd(key string) (interface{}, error){
	return this.Connection.Do("DEL", key)
}

func initRedis(){

	//Get redis server info from env
	redisHost = os.Getenv(REDIS_MASTER_ADDRESS_ENV)
	if len(redisHost) == 0 {
		panic(fmt.Errorf("Empty redis host env: %s", REDIS_MASTER_ADDRESS_ENV))
	}

	port, e := strconv.Atoi(os.Getenv(REDIS_MASTER_PORT_ENV))
	if e != nil || port < 0 {
		panic(fmt.Errorf("Wrong redis port env: %s", REDIS_MASTER_PORT_ENV))
	}
	redisPort = os.Getenv(REDIS_MASTER_PORT_ENV)

	//Common arguments
	testBorrowFunc := func(c redis.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	}
	idle_timeout := 5 * time.Minute
	max_idle_count := 5

	//Account auth database
	accountAuthConnPool = &redis.Pool{
		MaxIdle: max_idle_count,
		IdleTimeout: idle_timeout,
		Dial: func() (redis.Conn, error){
			c, err := redis.Dial(
				"tcp", redisHost + ":" + redisPort,
				redis.DialDatabase(ACCOUNT_AUTH_REDIS_DB),
			)

			if err != nil { return nil, err }
			return c, nil
		},
		TestOnBorrow: testBorrowFunc,
	}
}

func GetNewAccountAuthConn() *RedisConn {
	return &RedisConn{
		Connection: accountAuthConnPool.Get(),
	}
}
