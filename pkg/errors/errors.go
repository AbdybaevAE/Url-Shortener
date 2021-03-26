package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var InvalidLink = status.Error(codes.InvalidArgument, "invalid link provided")
var InvalidLinkKey = status.Error(codes.InvalidArgument, "invalid link key")
