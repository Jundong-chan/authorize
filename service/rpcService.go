package service

//远程调用服务，并且使用断路器控制
import (
	pkghystrix "github.com/Jundong-chan/authorize/plugins/hystrix"
	"github.com/Jundong-chan/authorize/plugins/pb"
	"context"
	"errors"

	"github.com/afex/hystrix-go/hystrix"
	"google.golang.org/grpc"
)

// var gatewayurl = "49.234.130.100:8000"
var grpcServiceAddress = "47.113.111.58:8808"

type RpcTokenService interface {
	RpcParseReToken(token string) (ntoken string, userid string, err error)
	RpcCreateToken(userid string, username string) (string, error)
}
type RpcTokenSvcimpl struct {
}

func (RT RpcTokenSvcimpl) RpcParseReToken(token string) (ntoken string, userid string, err error) {
	//设置断路器
	breaker := pkghystrix.NewDefaultCommand()
	err = hystrix.Do(breaker, func() error {
		//开启rpc调用请求token服务
		conn, err := grpc.Dial(grpcServiceAddress, grpc.WithInsecure())
		if err != nil {
			return errors.New("connect rpc failed" + err.Error())
		}
		defer conn.Close()
		tokenSevClient := pb.NewAuthorizeserviceClient(conn)
		ctx := context.TODO()
		tokenRequest := &pb.ParseAndRefreshTokenRequest{Token: token}
		reply, err := tokenSevClient.ParseAndRefreshToken(ctx, tokenRequest)
		if err != nil {
			return errors.New("Token wrong:" + err.Error())
		}
		ntoken = (*reply).Token
		userid = (*reply).UserId
		return err
	}, func(e error) error {
		return errors.New("Hystrix is on , RPC Token service busy")
	})
	if err != nil {
		return "", "", err
	}
	return ntoken, userid, nil

}

func (RT RpcTokenSvcimpl) RpcCreateToken(userid string, username string) (string, error) {
	//开启rpc调用请求token服务
	conn, err := grpc.Dial(grpcServiceAddress, grpc.WithInsecure())
	if err != nil {
		return "", errors.New("connect rpc failed" + err.Error())
	}
	defer conn.Close()
	ctx := context.TODO()
	tokenServClient := pb.NewAuthorizeserviceClient(conn)
	tokenRequest := &pb.CreateTokenRequest{Userid: userid, Username: username}
	reply, err := tokenServClient.CreateToken(ctx, tokenRequest)
	if err != nil {
		return "", errors.New("Rpc createtoken failed: " + err.Error())
	}
	return (*reply).Token, nil
}
