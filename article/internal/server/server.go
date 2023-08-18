package server

import (
	"context"
	v1 "github.com/2pgcn/auth/api/auth/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"strconv"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

func authCtx() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			meta, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.Unauthorized(v1.ErrorReason_USER_AUTH_ERROR.String(), "auth not uid")
			}
			uid := meta.RequestHeader().Get("uid")
			var intUid int
			intUid, err = strconv.Atoi(uid)
			if err != nil {
				return nil, errors.Unauthorized(v1.ErrorReason_USER_AUTH_ERROR.String(), "auth not uid")
			}
			ctx = context.WithValue(ctx, "uid", int64(intUid))
			return handler(ctx, req)
		}
	}
}
