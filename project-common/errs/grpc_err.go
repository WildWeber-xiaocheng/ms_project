package errs

import (
	common "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcError(err *BError) error {
	return status.Error(codes.Code(err.Code), err.Msg)
}

func ParseGrpcError(err error) (common.BusinessCode, string) {
	s, _ := status.FromError(err)
	return common.BusinessCode(s.Code()), s.Message()
}

func ToBError(err error) *BError {
	s, _ := status.FromError(err)
	return NewError(ErrorCode(s.Code()), s.Message())
}
