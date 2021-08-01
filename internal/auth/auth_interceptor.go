package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const authorizeHeader string = "authorization"

type AuthInterceptor struct {
	tokenManager    TokenManager
	accessableRoles map[string][]int32 // map of methods and permission-set
}

func NewAuthInterceptor(manager TokenManager, methods map[string][]int32) *AuthInterceptor {
	return &AuthInterceptor{
		tokenManager:    manager,
		accessableRoles: methods,
	}
}

/*
  Higher order method that returns a grpc.UnaryServerInterceptor
  vlidates for allowed method with security token
*/
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		log.Println("--------> Unary interceptor: " + info.FullMethod)

		err = interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}

}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	// get accesable method and roles
	accessableRoles, ok := interceptor.accessableRoles[method]
	if !ok {
		// public access method
		return nil
	}

	// get token from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "user authentication is required")
	}
	values := md[authorizeHeader]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token not provided")
	}

	token := values[0]
	payload, err := interceptor.tokenManager.VerifyToken(token)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "token is invalid: %v", err)
	}

	exists := make(map[int32]struct{})
	for _, role := range accessableRoles {
		exists[role] = struct{}{}
	}

	for _, permission := range payload.Permissions {
		if _, ok := exists[int32(permission)]; ok {
			return nil
		}
	}

	return status.Errorf(codes.PermissionDenied, "permission not granted to access rpc")
}
