//这个是完全使用 gokit 的方式实现日志收集中间件,没写完，与另一种方式差不多

package plugins

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
)

//type LoggerMiddware func(endpoint.Endpoint) endpoint.Endpoint

func MakeCreateUserMiddware(log kitlog.Logger, next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func(begin time.Time) {
			log.Log(
				"function", "CreateUser",
				"took", time.Since(begin),
				"request", request,
			)
		}(time.Now())
		req, err := next(ctx, request)
		return req, err
	}
}

func MakeGetUserListLoggingMidWare(log kitlog.Logger, next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		defer func(begin time.Time) {
			log.Log(
				"function", "GetUserList",
				"took", time.Since(begin),
				"request", request,
			)
		}(time.Now())
		res, err := next(ctx, request)
		return res, err
	}
}
