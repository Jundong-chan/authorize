package transport

import (
	"CommoditySpike/server/admin/authorize/endpoint"
	"CommoditySpike/server/admin/authorize/pkg/pb"
	"context"
	"errors"

	"github.com/go-kit/kit/transport/grpc"
)

type grpcserver struct {
	createtoken          grpc.Handler
	parseandrefreshtoken grpc.Handler
	parsetoken           grpc.Handler
}

//远程调用服务配置
func (g *grpcserver) CreateToken(ctx context.Context, r *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	_, resp, err := g.createtoken.ServeGRPC(ctx, r) //返回context，response和error
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateTokenResponse), nil
}

func (g *grpcserver) ParseAndRefreshToken(ctx context.Context, r *pb.ParseAndRefreshTokenRequest) (*pb.ParseAndRefreshTokenResponse, error) {
	_, resp, err := g.parseandrefreshtoken.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ParseAndRefreshTokenResponse), nil
}

func (g *grpcserver) ParseToken(ctx context.Context, r *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	_, resp, err := g.parsetoken.ServeGRPC(ctx, r) //返回context，response和error
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ParseTokenResponse), nil
}

func DecodeCreateTokenRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.CreateTokenRequest)
	if !ok {
		return nil, errors.New("DecodeCreateTokenRequest error:Type error")
	}
	return endpoint.CreateTokenRequest{
		UserId:   req.Userid,
		UserName: req.Username,
	}, nil
}

func EncodeCreateTokenResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res, ok := r.(endpoint.CreateTokenResponse)
	if !ok {
		return nil, errors.New("EncodeCreateTokenResponse error:Type error")
	}
	if res.Error != nil {
		return &pb.CreateTokenResponse{
			Token: res.Token,
			Error: res.Error.Error(),
		}, nil
	}
	return &pb.CreateTokenResponse{
		Token: res.Token,
		Error: "",
	}, nil
}

func DecodeParseAndRefreshTokenRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.ParseAndRefreshTokenRequest)
	if !ok {
		return nil, errors.New("DecodeParseAndRefreshTokenRequest error:")
	}
	return endpoint.ParseAndRefreshTokenRequest{
		Token: req.Token,
	}, nil
}

func EncodeParseAndRefreshTokenResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res, ok := r.(endpoint.ParseAndRefreshTokenResponse)
	if !ok {
		return nil, errors.New("EncodeParseAndRefreshTokenResponse error:Type error")
	}
	if res.Error != nil {
		return &pb.ParseAndRefreshTokenResponse{
			UserId: "",
			Token:  res.Token,
			Error:  res.Error.Error(),
		}, nil
	}
	return &pb.ParseAndRefreshTokenResponse{
		UserId: res.UserId,
		Token:  res.Token,
		Error:  "",
	}, nil
}

func DecodeParseTokenRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.ParseTokenRequest)
	if !ok {
		return nil, errors.New("DecodeParseTokenRequest error:Type error")
	}
	return endpoint.ParseTokenRequest{
		Token: req.Token,
	}, nil
}

func EncodeParseTokenResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res, ok := r.(endpoint.ParseTokenResponse)
	if !ok {
		return nil, errors.New("EncodeParseTokenResponse error:Type error")
	}
	if res.Error != nil {
		return &pb.ParseTokenResponse{
			Isvalid: false,
			Error:   res.Error.Error(),
		}, nil
	}

	return &pb.ParseTokenResponse{
		Isvalid: true,
		Error:   "",
	}, nil
}

//将定义的服务enpoint注册到 自定义的grpcserver中
func MakeGrpcTokenServer(ctx context.Context, endpoints endpoint.TokenAuthEndpoint) pb.AuthorizeserviceServer {
	return &grpcserver{
		createtoken: grpc.NewServer(
			endpoints.CreateTokenEp,
			DecodeCreateTokenRequest,
			EncodeCreateTokenResponse,
		),
		parseandrefreshtoken: grpc.NewServer(
			endpoints.ParseAndRefreshTokenEp,
			DecodeParseAndRefreshTokenRequest,
			EncodeParseAndRefreshTokenResponse,
		),
		parsetoken: grpc.NewServer(
			endpoints.ParseTokenEp,
			DecodeParseTokenRequest,
			EncodeParseTokenResponse,
		),
	}
}
