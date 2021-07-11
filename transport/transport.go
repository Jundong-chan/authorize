package transport

import (
	endpoint "CommoditySpike/server/admin/endpoint"
	"CommoditySpike/server/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	gokithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	decodeErr = errors.New("decode Request error")
	encodeErr = errors.New("encode Response error")
)

//获取列表
func decodeGetListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoint.GetListRequest{}, nil
}
func encodeGetListResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.GetListResponse)
	if !ok {
		return errors.New("failed to encodeGetListResponse")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

//创建
func decodeCreateProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return nil, err
	}
	return product, nil

}
func decodeCreateActivityRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var activity model.Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		return nil, err
	}
	return activity, nil
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}

func encodecreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.CreateResponse)
	if !ok {
		return errors.New("failed to encodecreateResponse")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	return json.NewEncoder(w).Encode(res)
}

//查询信息
func decodeQureyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var qreq endpoint.QureyRequest

	if err := json.NewDecoder(r.Body).Decode(&qreq); err != nil {
		return nil, errors.New("failed to decodeQureyRequest")
	}
	fmt.Println("解析到数据：", qreq)
	return qreq, nil
}
func encodeQureyResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.QureyResponse)

	if !ok {
		return errors.New("failed to encodeQureyResponse")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

//修改商品
func decodeModifyProRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		return nil, errors.New("failed to decodeModifyRequest")
	}
	return product, nil
}
func encodeModifyProResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.ErrorResponse)
	if !ok {
		return errors.New("failed to encodeModifyResponse")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

//登录
//解析 cookie中是否存在 token
func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var CKRequest endpoint.LoginRequest
	var token string
	cookie := r.Header.Get("Cookie")
	if cookie != "" {
		fmt.Println("解析到cookie：", cookie)
		//去掉空格
		cookie = strings.Replace(cookie, " ", "", -1)
		cookies := strings.Split(cookie, ";")
		if strings.Contains(cookies[0], "token=") {
			token = strings.TrimPrefix(cookies[0], "token=")
		} else {
			token = strings.TrimPrefix(cookies[1], "token=")
		}
		CKRequest.Token = token
	}

	err := json.NewDecoder(r.Body).Decode(&CKRequest)
	if err != nil && token == "" {
		return nil, errors.New("failed to decodeLoginRequest")
	}
	fmt.Println("CKRequest:", CKRequest)
	//如果没有明确登陆的账号类型，就默认为邮箱账号登陆
	if CKRequest.Feild == "" {
		CKRequest.Feild = "email"
	}

	return CKRequest, nil

}

//登陆成功的话，将token写进cookie
func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.LoginResponse)
	if !ok {
		return encodeErr
	}
	if res.Error != nil {
		return json.NewEncoder(w).Encode(res)
	}
	cookie1 := http.Cookie{
		Name:     "token",
		Value:    res.Token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 10),
	}
	http.SetCookie(w, &cookie1)
	cookie2 := http.Cookie{
		Name:     "userid",
		Value:    res.User[0]["user_id"].(string),
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 10),
	}
	http.SetCookie(w, &cookie2)

	w.Header().Set("Content-Type", "application/json;charset=utf-8") //该函数要在 WriteHeader函数之前调用
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

//删除商品
func decodeDeleteProRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.New("failed to decodeDeleteProductRequest")
	}
	return req, nil
}
func encodeDeleteProResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.ErrorResponse)
	if !ok {
		return encodeErr
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

func decodeProductCountsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoint.ProductCountRequest{}, nil
}

//查询商品总数
func encodeProductCountsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.ProductCountResponse)
	if !ok {
		return errors.New("failed to encodeProductCountsResponse")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

//分页查询商品
func decodePageProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.PageProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.New("failed to decodePageProductRequest")
	}
	return req, nil
}

func encodePageProductResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.PageProductResponse)
	if !ok {
		return errors.New("failed to encodePageProductResponse")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

//操作订单

func decodeModifyOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.ModifyOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.New("failed to decodeModifyOrderRequest")
	}
	return req, nil
}

func encodeModifyOrderResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.ModifyOrderResponse)
	if !ok {
		return encodeErr
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

func decodeQureyOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.QureyRequest
	var userid, token string
	cookie := r.Header.Get("Cookie")
	if cookie == "" {
		return nil, errors.New("decodeQureyOrderRequest: no token")
	}
	fmt.Println("解析到cookie：", cookie)
	//去掉空格
	cookie = strings.Replace(cookie, " ", "", -1)
	cookies := strings.Split(cookie, ";")
	if strings.Contains(cookies[0], "token=") {
		userid = strings.TrimPrefix(cookies[1], "userid=")
		token = strings.TrimPrefix(cookies[0], "token=")
	} else {
		userid = strings.TrimPrefix(cookies[0], "userid=")
		token = strings.TrimPrefix(cookies[1], "token=")
	}
	fmt.Println("解析到userid：", userid)
	if token == "" {
		return nil, errors.New("decodeQureyOrderRequest: no token")
	}
	fmt.Println("解析到token：", token)
	req.Token = token
	return req, nil
}

func encodeQureyOrderResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.QureyResponse)
	if !ok {
		return encodeErr
	}
	if res.Error != nil {
		return json.NewEncoder(w).Encode(res)
	}
	cookie1 := http.Cookie{
		Name:     "token",
		Value:    res.Token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 10),
	}
	http.SetCookie(w, &cookie1)

	cookie2 := http.Cookie{
		Name:     "userid",
		Value:    res.Result[0]["user_id"].(string),
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 10),
	}
	http.SetCookie(w, &cookie2)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

func decodeHealthCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoint.HealthCheckRequest{}, nil
}

func encodeHealthCheckResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.HealthCheckResponse)
	if !ok {
		return errors.New("failed to encodecreateResponse")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

func encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch err {

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func MakeHttpHandler(ctx context.Context, enp endpoint.AdminEndpoints, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	options := []gokithttp.ServerOption{
		//加上自定义的出错日志收集
		gokithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		//加上自定义的错误代码处理函数，自定义出错时的responsewriter
		gokithttp.ServerErrorEncoder(encodeErrorResponse),
	}
	r.Methods("GET").Path("/product/list").Handler(gokithttp.NewServer(
		enp.GetProductListEp,
		decodeGetListRequest,
		encodeGetListResponse,
		options...,
	))

	r.Methods("POST").Path("/product/register").Handler(gokithttp.NewServer(
		enp.CraeteProductEp,
		decodeCreateProductRequest,
		encodecreateResponse,
		options...,
	))

	r.Methods("POST").Path("/product/modify").Handler(gokithttp.NewServer(
		enp.ModifyProductEp,
		decodeModifyProRequest,
		encodeModifyProResponse,
		options...,
	))

	r.Methods("POST").Path("/product/delete").Handler(gokithttp.NewServer(
		enp.DeleteProductEp,
		decodeDeleteProRequest,
		encodeDeleteProResponse,
		options...,
	))
	r.Methods("POST").Path("/product/qurey").Handler(gokithttp.NewServer(
		enp.QureyProductEp,
		decodeQureyRequest,
		encodeQureyResponse,
		options...,
	))
	r.Methods("GET", "POST").Path("/product/qurey/count").Handler(gokithttp.NewServer(
		enp.ProductCountsEp,
		decodeProductCountsRequest,
		encodeProductCountsResponse,
		options...,
	))
	r.Methods("POST").Path("/product/qurey/page").Handler(gokithttp.NewServer(
		enp.PageProductEp,
		decodePageProductRequest,
		encodePageProductResponse,
		options...,
	))

	r.Methods("GET").Path("/activity/list").Handler(gokithttp.NewServer(
		enp.GetActivityListEp,
		decodeGetListRequest,
		encodeGetListResponse,
		options...,
	))
	r.Methods("POST").Path("/activity/register").Handler(gokithttp.NewServer(
		enp.CreateActivityEp,
		decodeCreateActivityRequest,
		encodecreateResponse,
		options...,
	))
	r.Methods("GET").Path("/user/list").Handler(gokithttp.NewServer(
		enp.GetUserListEp,
		decodeGetListRequest,
		encodeGetListResponse,
		options...,
	))
	r.Methods("GET").Path("/user/qurey").Handler(gokithttp.NewServer(
		enp.QureyUserEp,
		decodeQureyRequest,
		encodeQureyResponse,
		options...,
	))

	r.Methods("POST").Path("/user/register").Handler(gokithttp.NewServer(
		enp.CreateUserEp,
		decodeCreateUserRequest,
		encodecreateResponse,
		options...,
	))
	r.Methods("POST").Path("/login").Handler(gokithttp.NewServer(
		enp.LoginEp,
		decodeLoginRequest,
		encodeLoginResponse,
		options...,
	))

	r.Methods("GET").Path("/order/qurey").Handler(gokithttp.NewServer(
		enp.QureyOrderEp,
		decodeQureyOrderRequest,
		encodeQureyOrderResponse,
		options...,
	))

	r.Methods("POST").Path("/order/modify").Handler(gokithttp.NewServer(
		enp.ModifyOrderEp,
		decodeModifyOrderRequest,
		encodeModifyOrderResponse,
		options...,
	))

	r.Methods("GET").Path("/healthcheck").Handler(gokithttp.NewServer(
		enp.HealthCheckEp,
		decodeHealthCheckRequest,
		encodeHealthCheckResponse,
		options...,
	))
	r.Methods("GET", "POST").PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	return r
}
