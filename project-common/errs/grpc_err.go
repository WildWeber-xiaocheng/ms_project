package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	common "test.com/project-common"
)

func GrpcError(err *BError) error {
	return status.Error(codes.Code(err.Code), err.Msg)
}

func ParseGrpcError(err error) (common.BusinessCode, string) {
	s, _ := status.FromError(err)
	return common.BusinessCode(s.Code()), s.Message()
}
