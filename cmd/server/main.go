package main

import (
	"database/sql"
	"net"
	"os"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/auth"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/employee"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/organization"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/payroll"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/role"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/user"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/go-redis/redis/v8"
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
	// check if the connection was established
	err = conn.Ping()
	if err != nil {
		log.Error("unable to connect to database", "error", err)
		os.Exit(1)
	}
	log.Info("successfully connected to database", "driver", DBDriver)

	// connect to NoSql server
	cacheClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// if status := cacheClient.Ping(context.Background()); err != nil {
	// 	log.Error("unable to connect to cache", "error", err)
	// 	os.Exit(1)
	// }

	// create store
	store := db.NewSQlStore(conn)
	cache := db.NewCacheStore(cacheClient)

	// create a tokoen manager
	manager, err := auth.NewPasetoManager()
	if err != nil {
		log.Error("Unable to create token", err)
	}

	// create servers
	authServer := auth.NewAuthServer(store, log, manager)
	userServer := user.NewUserServer(store, log)
	organizationServer := organization.NewOrganizationServer(store, log)
	employeeServer := employee.NewEmployeeServer(store, log)
	roleServer := role.NewRoleServer(store, log)
	payrollServer := payroll.NewPayrollServe(store, log)

	// permission set
	accessableMethods := map[string][]int32{
		"/OrganizationService/FindOrganization": {int32(hrm.Permission_ADMIN)},
	}

	// create a auth interceptor
	authInterceptor := auth.NewAuthInterceptor(manager, cache, accessableMethods)

	// create tcp connection
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("unable to start server", "address", address, "error", err)
		os.Exit(1)
	}
	// create a new grpc server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)
	// register servers
	hrm.RegisterAuthServiceServer(grpcServer, authServer)
	hrm.RegisterUserServiceServer(grpcServer, userServer)
	hrm.RegisterOrganizationServiceServer(grpcServer, organizationServer)
	hrm.RegisterEmployeeServiceServer(grpcServer, employeeServer)
	hrm.RegisterRoleServiceServer(grpcServer, roleServer)
	hrm.RegisterPayrollServiceServer(grpcServer, payrollServer)

	reflection.Register(grpcServer)

	log.Info("server started successfully", "serving at port", address)

	// server
	grpcServer.Serve(listener)

}
