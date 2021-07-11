package endpoint

import (
	"github.com/Jundong-chan/authorize/service"
	"github.com/Jundong-chan/seckill/model"
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/gohouse/gorose"
)

type AdminEndpoints struct {
	CraeteProductEp   endpoint.Endpoint
	GetProductListEp  endpoint.Endpoint
	ModifyProductEp   endpoint.Endpoint
	DeleteProductEp   endpoint.Endpoint
	CreateActivityEp  endpoint.Endpoint
	GetActivityListEp endpoint.Endpoint
	CreateUserEp      endpoint.Endpoint
	GetUserListEp     endpoint.Endpoint
	QureyUserEp       endpoint.Endpoint
	QureyProductEp    endpoint.Endpoint
	LoginEp           endpoint.Endpoint
	ModifyOrderEp     endpoint.Endpoint
	QureyOrderEp      endpoint.Endpoint
	ProductCountsEp   endpoint.Endpoint
	PageProductEp     endpoint.Endpoint
	HealthCheckEp     endpoint.Endpoint
}

type GetListRequest struct {
}

type GetListResponse struct {
	Result []gorose.Data `json:"result"`
	Error  error         `json:"error"`
}
type CreateRequest struct {
}
type CreateResponse struct {
	Error error `json:"error"`
}
type DeleteRequest struct {
	Id string `json:"id"`
}

type ErrorResponse struct {
	Error error `json:"error"`
}

type QureyRequest struct {
	Token     string `json:"token"`
	Feild     string `json:"feild"`
	Condition string `json:"condition"`
	Value     string `json:"value"`
}

type QureyResponse struct {
	Result []gorose.Data `json:"result"`
	Token  string        `json:"token"`
	Total  int           `json:"total"`
	Code   int           `json:"code"`
	Error  error         `json:"error"`
}

type ModifyOrderRequest struct {
	Data  map[string]interface{} `json:"data"`
	Feild string                 `json:"feild"`
	Value string                 `json:"value"`
}

type ModifyOrderResponse struct {
	Error string `json:"feild"`
}

type LoginRequest struct {
	Feild    string `json:"feild"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Token    string
}

type LoginResponse struct {
	User  []gorose.Data `json:"user"`
	Token string        `json:"token"`
	Error error         `json:"error,omitempty"`
}

type ParseTokenRequest struct {
	Token string
}
type ParseTokenResponse struct {
	UserId string
	Token  string
	Error  error
}

type ProductCountRequest struct {
}

type ProductCountResponse struct {
	Count int   `json:"count"`
	Error error `json:"error"`
}

type PageProductRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
type PageProductResponse struct {
	Data  []gorose.Data `json:"data"`
	Error error         `json:"error"`
}

type HealthCheckRequest struct {
}

type HealthCheckResponse struct {
	Status bool `json:"status"`
}

func MakeCraeteProductEp(svc service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(model.Product)
		err := svc.CreateProduct(&req)
		if err != nil {
			return CreateResponse{
				Error: err,
			}, nil
		}
		return CreateResponse{
			Error: nil,
		}, nil
	}
}

func MakeGetProductListEp(svc service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		result, err := svc.GetProductList()
		if err != nil {
			return GetListResponse{
				Result: nil,
				Error:  err,
			}, nil
		}
		return GetListResponse{
			Result: result,
			Error:  nil,
		}, nil
	}
}

func MakeModifyProductEp(svc service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(model.Product)
		err := svc.ModifyProduct(&req)
		if err != nil {
			return ErrorResponse{
				Error: errors.New("modify product failed"),
			}, err
		}

		return ErrorResponse{
			Error: nil,
		}, nil
	}
}

func MakeDeleteProductEp(svc service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(DeleteRequest)
		err := svc.DeleteProduct(req.Id)
		if err != nil {
			return ErrorResponse{
				Error: errors.New("delete product failed"),
			}, err
		}
		return ErrorResponse{
			Error: nil,
		}, nil
	}
}

func MakeQureyProductEp(svc service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		fmt.Println("进入 QureyProduct处理函数")
		req := request.(QureyRequest)

		data, err := svc.QureyProductByUserId(req.Feild, req.Condition, req.Value)
		if err != nil {
			return ErrorResponse{
				Error: errors.New("qurey product failed"),
			}, err
		}
		return QureyResponse{
			Result: data,
			Error:  nil,
		}, nil
	}
}

//请求商品总数量
func MakeProductCountEp(svc service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		fmt.Println("进入 QureyProduct处理函数")
		_ = request.(ProductCountRequest)
		count, err := svc.GetProductCounts()
		if err != nil {
			return ProductCountResponse{
				Count: 0,
				Error: errors.New("qurey product failed"),
			}, err
		}
		return ProductCountResponse{
			Count: count,
			Error: nil,
		}, nil
	}

}

//请求商品分页查询
func MakePageProductEp(svc service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(PageProductRequest)
		result, err := svc.PageProducts(req.Page, req.Limit)
		if err != nil {
			return PageProductResponse{
				Data:  nil,
				Error: err,
			}, err
		}
		return PageProductResponse{
			Data:  result,
			Error: nil,
		}, nil
	}
}

func MakeCreateActivityEp(svc service.ActivityService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(model.Activity)
		err := svc.CreateActivity(&req)
		if err != nil {
			return CreateResponse{
				Error: err,
			}, nil
		}
		return CreateResponse{
			Error: nil,
		}, nil
	}
}

func MakeGetActivityListEp(svc service.ActivityService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		result, err := svc.GetActivityList()
		if err != nil {
			return GetListResponse{
				Result: nil,
				Error:  err,
			}, nil
		}
		return GetListResponse{
			Result: result,
			Error:  nil,
		}, nil
	}
}

func MakeCreateUserEp(svc service.UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(model.User)
		fmt.Println(req)
		err := svc.CreateUser(&req)
		if err != nil {
			return CreateResponse{
				Error: err,
			}, nil
		}
		return CreateResponse{
			Error: nil,
		}, nil
	}
}

func MakeGetUserListEp(svc service.UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		result, err := svc.GetUserList()
		if err != nil {
			return GetListResponse{
				Result: nil,
				Error:  err,
			}, nil
		}
		return GetListResponse{
			Result: result,
			Error:  nil,
		}, nil
	}
}

func MakeQueryUserEp(svc service.UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(QureyRequest)
		result, err := svc.QureyUser(req.Feild, req.Condition, req.Value)
		return QureyResponse{
			Result: result,
			Error:  err,
		}, err
	}
}

//订单操作相关函数
func MakeQureyOrderEp(svc service.OrderService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(QureyRequest)
		rsvc := service.RpcTokenSvcimpl{}
		//对token进行验证
		ntoken, userid, err := rsvc.RpcParseReToken(req.Token)
		fmt.Println("得到ntoken: ", ntoken)
		fmt.Println("得到userid: ", userid)
		if err != nil {
			return nil, errors.New("rpc faild: " + err.Error())
		} else {
			//token合法
			result, err := svc.QureyOrder(userid)
			fmt.Println("orders:", result)
			return QureyResponse{
				Result: result,
				Token:  ntoken,
				Code:   0,
				Total:  len(result),
			}, err
		}
	}
}

func MakeModifyOrderEp(svc service.OrderService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(ModifyOrderRequest)
		err := svc.ModifyOrder(req.Data, req.Feild, req.Value)
		if err != nil {
			return ModifyOrderResponse{
				Error: err.Error(),
			}, err
		}
		return ModifyOrderResponse{
			Error: "",
		}, err
	}
}

//调用rpc方法得到token 以及验证用户的登陆信息，只有token没有用户信息也可以登录,如果同时有token和用户信息，就使用用户信息登录
func MakeLoginEp(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, e error) {
		req := request.(LoginRequest)
		fmt.Println(req)

		rsvc := service.RpcTokenSvcimpl{}
		//对token进行验证
		if req.Token != "" && req.Account == "" && req.Password == "" {
			ntoken, userid, err := rsvc.RpcParseReToken(req.Token)
			fmt.Println("得到ntoken: ", ntoken)
			fmt.Println("得到userid: ", userid)
			if err != nil {
				//如果token不合法，就检查请求是否有用户的验证信息
				if req.Account == "" || req.Password == "" || req.Feild == "" {
					return LoginResponse{
						Error: errors.New("rpc faild: " + err.Error()),
					}, nil
				}
			} else {
				//token合法
				result, err := svc.QureyUser("user_id", "=", userid)
				fmt.Println("user:", result)
				result[0]["salt"] = ""
				return LoginResponse{
					User:  result,
					Token: ntoken,
					Error: err,
				}, nil
			}
		}
		//对用户信息进行验证
		result, err := svc.CheckUser(req.Feild, req.Account, req.Password)
		if err != nil {
			return nil, err
		}
		username := result[0]["user_name"].(string)
		userid := result[0]["user_id"].(string)
		result[0]["salt"] = ""
		ntoken, err := rsvc.RpcCreateToken(userid, username)
		if err != nil {
			return LoginResponse{
				User:  result,
				Error: errors.New("no token"),
			}, errors.New("rpc faild: " + err.Error())
		}
		return LoginResponse{
			User:  result,
			Token: ntoken,
		}, nil

	}
}

func MakeHealthCheckEp(svc service.AdminService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		status := svc.HealthCheck()
		return HealthCheckResponse{
			Status: status,
		}, nil
	}
}
