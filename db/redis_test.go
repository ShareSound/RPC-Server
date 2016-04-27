package db_test

import(
	"testing"
	"github.com/ShareSound/RPC-Server/db"
	"github.com/garyburd/redigo/redis"
)

//SET, GET, DEL
func TestRedisBasicCmds(t *testing.T) {
	c := db.GetNewAccountAuthConn()
	defer c.Connection.Close()

	if _, e := c.SetCmd("test_hello", "world"); e != nil {
		t.Fatalf("SET command failed: %s\n", e.Error())
		t.FailNow()
	}

	if str, e := redis.String(c.GetCmd("test_hello")); e != nil {
		t.Fatalf("GET command failed: %s\n", e.Error())
		t.FailNow()
	}else{
		if str != "world" {
			t.Fatalf("GET expect \"world\", get %s\n", str)
			t.FailNow()
		}
	}

	if _, e := c.DelCmd("test_hello"); e != nil {
		t.Fatalf("DEL command failed: %s\n", e.Error())
		t.FailNow()
	}
}
