package db_test

import (
	"context"
	"os"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

// Global varibles for db testing
var (
	testCacheStore *db.CacheStore
)

func TestMain(m *testing.M) {
	var err error

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	err = client.Ping(context.Background()).Err()
	if err != nil {
		os.Exit(m.Run())
	}

	testCacheStore = db.NewCacheStore(client)

	os.Exit(m.Run())
}
