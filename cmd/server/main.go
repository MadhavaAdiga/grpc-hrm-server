package main

import (
	"database/sql"
	"net"
	"os"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/organization"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/user"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:docker_postgres@localhost:5432/hrm_db?sslmode=disable"
	address  = "0.0.0.0:3000"
)

func main() {
	// create a loger with option
	loggerOption := &hclog.LoggerOptions{
		Name:  "grpcServer",
		Level: hclog.DefaultLevel,
	}
	log := hclog.New(loggerOption)

	// connect to database
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Error("unable to connect to database", "error", err)
		os.Exit(1)
	}
	log.Info("successfully connected to database", "driver", DBDriver)

	// create store
	store := db.NewSQlStore(conn)

	// create servers
	userServer := user.NewUserServer(store, log)
	organizationServer := organization.NewOrganizationServer(store, log)

	// create tcp connection
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("unable to start server", "address", address, "error", err)
		os.Exit(1)
	}
	// create a new grpc server
	grpcServer := grpc.NewServer()
	// register servers
	hrm.RegisterUserServiceServer(grpcServer, userServer)
	hrm.RegisterOrganizationServiceServer(grpcServer, organizationServer)

	reflection.Register(grpcServer)

	log.Info("server started successfully", "serving at port", address)

	// server
	grpcServer.Serve(listener)

}
