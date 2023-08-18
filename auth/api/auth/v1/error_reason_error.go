package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserNotFound(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(http.StatusNotFound, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserAuthError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_AUTH_ERROR.String() && e.Code == http.StatusUnauthorized
}

func UserAuthError(format string, args ...interface{}) *errors.Error {
	return errors.New(http.StatusUnauthorized, ErrorReason_USER_AUTH_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsInternalError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Internal_Error.String() && e.Code == http.StatusUnauthorized
}

func InternalError(format string, args ...interface{}) *errors.Error {
	return errors.New(http.StatusUnauthorized, ErrorReason_Internal_Error.String(), fmt.Sprintf(format, args...))
}
