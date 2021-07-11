package endpoint

//用户授权与鉴权的endpoint
import (
	"github.com/Jundong-chan/authorize/authorize/model"
	"github.com/Jundong-chan/authorize/authorize/service"
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

var (
	ConvertRequestError = errors.New("Failed to convert CheckUserRequests")
	UserDataInvalid     = errors.New("The user's data is invalid")
	CreateTokenFaild    = errors.New("Create Token Faild")
)

type TokenAuthEndpoint struct {
	CreateTokenEp          endpoint.Endpoint //登陆服务，携带用户验证信息来请求token以及返回用户信息
	ParseAndRefreshTokenEp endpoint.Endpoint //检验并刷新token
	ParseTokenEp           endpoint.Endpoint //检验token的合法性
	HealthCheckEp          endpoint.Endpoint //健康检查
}

type CreateTokenRequest struct {
	UserId   string `json:"userid"`
	UserName string `josn:"username"`
}

type CreateTokenResponse struct {
	Token string `json:"token,omitempty"`
	Error error  `json:"error"`
}

type ParseAndRefreshTokenRequest struct {
	Token string `json:"token"`
}

type ParseAndRefreshTokenResponse struct {
	UserId string `json:"userid"`
	Token  string `json:"token"`
	Error  error  `json:"error"`
}

type ParseTokenRequest struct {
	Token string `json:"token"`
}

type ParseTokenResponse struct {
	IsValid bool  `json:"isvalid"`
	Error   error `json:"error"`
}
type HealthCheckRequest struct{}
type HealthCheckResponse struct {
	Status bool `json:"status"`
}

//验证用户信息，以及token(如果存在)成功则返回用户的基本信息以及两个token
func MakeCreateTokenEp(svc service.TokenService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateTokenRequest)
		if !ok {
			return nil, ConvertRequestError
		}
		fmt.Println("userid", req.UserId, "username", req.UserName)
		//创建和用户相关的payload
		user := model.UserDetail{
			UserId:   req.UserId,
			UserName: req.UserName,
		}
		claims := model.NewClaims(&user, model.InitTokenConfig())
		token, err := svc.CreateToken(claims)
		if err != nil {
			return CreateTokenResponse{
				Error: err,
			}, err
		}

		return CreateTokenResponse{
			Token: token,
			Error: nil,
		}, err
	}
}

//检验令牌，合法则返回true
func MakeParseTokenEp(svc service.TokenService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, ok := request.(ParseTokenRequest)
		if !ok {
			return nil, ConvertRequestError
		}
		_, err = svc.ParseToken(res.Token)
		if err != nil {
			return ParseTokenResponse{
				IsValid: false,
				Error:   nil,
			}, err
		}
		//fmt.Println(token)
		return ParseTokenResponse{
			IsValid: true,
			Error:   nil,
		}, nil
	}
}

//验证token是否合法，合法则返回新的token，以及userid
func MakeParseAndRefreshTokenEp(svc service.TokenService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, ok := request.(ParseAndRefreshTokenRequest)
		if !ok {
			return nil, ConvertRequestError
		}
		newtoken, claims, err := svc.ParseTokenAndReToken(res.Token)
		//fmt.Println("claims", claims)
		if err != nil {
			return nil, err
		}
		//fmt.Println("userid", (*claims).UserId)
		return ParseAndRefreshTokenResponse{
			UserId: claims.UserId,
			Error:  err,
			Token:  newtoken,
		}, nil
	}
}

func MakeHealthCheckEp(svc service.HealthCheckService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		status := svc.HealthCheck()
		return HealthCheckResponse{
			Status: status,
		}, nil
	}
}
