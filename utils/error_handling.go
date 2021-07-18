package utils

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
  Utility function to handele context error
*/
func ContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is cancled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline has exceeded")
	default:
		return nil
	}
}
