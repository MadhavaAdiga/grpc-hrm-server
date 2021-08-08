package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	_ "github.com/lib/pq"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:docker_postgres@localhost:5432/hrm_db?sslmode=disable"
)

// Global varibles for db testing
var (
	testSQLStore *db.SQLStore
	testDB       *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	// testCache := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "test_password",
	// 	DB:       0,
	// })

	testSQLStore = db.NewSQlStore(testDB)

	os.Exit(m.Run())
}
