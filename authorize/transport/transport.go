package transport

import (
	"CommoditySpike/server/admin/authorize/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	gokithttp "github.com/go-kit/kit/transport/http"
)

var (
	decodeErr = errors.New("Decode Request error")
	encodeErr = errors.New("Encode Response error")
)

//解析，如果存在token则解析出来
func decodeCreateTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var CKRequest endpoint.CreateTokenRequest

	err := json.NewDecoder(r.Body).Decode(&CKRequest)
	if err != nil {
		return nil, decodeErr
	}
	return CKRequest, nil
}

func encodeCreateTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.CreateTokenResponse)
	if !ok {
		return encodeErr
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	cookie := http.Cookie{
		Name:     "token",
		Value:    res.Token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 10),
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

func decodeParseAndRefresTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var RfreshTRequest endpoint.ParseAndRefreshTokenRequest
	err := json.NewDecoder(r.Body).Decode(&RfreshTRequest)
	if err != nil {
		return nil, decodeErr
	}
	return RfreshTRequest, nil
}

func encodeParseAndRefreshTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.ParseAndRefreshTokenResponse)
	if !ok {
		return encodeErr
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(res)
}

func decodeParseTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var ParseTRequest endpoint.ParseTokenRequest
	err := json.NewDecoder(r.Body).Decode(&ParseTRequest)
	if err != nil {
		return nil, decodeErr
	}
	return ParseTRequest, nil
}

func encodeParseTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(endpoint.ParseTokenResponse)
	if !ok {
		return encodeErr
	}
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
		return encodeErr
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

func MakeHttpHandler(ctx context.Context, enp endpoint.TokenAuthEndpoint, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()
	options := []gokithttp.ServerOption{
		//加上自定义的出错日志收集
		gokithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		//加上自定义的错误代码处理函数，自定义出错时的responsewriter
		gokithttp.ServerErrorEncoder(encodeErrorResponse),
	}
	r.Methods("POST").Path("/oauth/createtoken").Handler(gokithttp.NewServer(
		enp.CreateTokenEp,
		decodeCreateTokenRequest,
		encodeCreateTokenResponse,
		options...,
	))
	r.Methods("POST").Path("/oauth/refreshtoken").Handler(gokithttp.NewServer(
		enp.ParseAndRefreshTokenEp,
		decodeParseAndRefresTokenRequest,
		encodeParseAndRefreshTokenResponse,
		options...,
	))

	r.Methods("POST").Path("/oauth/parsetoken").Handler(gokithttp.NewServer(
		enp.ParseTokenEp,
		decodeParseTokenRequest,
		encodeParseTokenResponse,
		options...,
	))

	r.Methods("GET").Path("/oauth/healthcheck").Handler(gokithttp.NewServer(
		enp.HealthCheckEp,
		decodeHealthCheckRequest,
		encodeHealthCheckResponse,
		options...,
	))
	return r
}
