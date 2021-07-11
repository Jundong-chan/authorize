package main

import (
	"github.com/Jundong-chan/authorize/authorize/endpoint"
	"github.com/Jundong-chan/authorize/authorize/pkg/mysql"
	"github.com/Jundong-chan/authorize/authorize/pkg/pb"
	"github.com/Jundong-chan/authorize/authorize/service"
	"github.com/Jundong-chan/authorize/authorize/transport"
	"github.com/Jundong-chan/authorize/config"
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func main() {
	var (
		//参数说明：命令行参数名称，默认值，参数说明（如果使用-help会给出这个说明提示）
		//使用变量时要加 *
		serverhost    = flag.String("host", "", "server host")
		serverport    = flag.String("port", "8808", "server port")
		mysqluser     = flag.String("mysqluser", "root", "mysql user")
		mysqlpassword = flag.String("mysqlpassword", "gzhuchan", "mysql password")
		mysqlhost     = flag.String("mysqlhost", "127.0.0.1", "mysql host")
		mysqlport     = flag.String("mysqlport", "3306", "mysql port")
		mysqlschema   = flag.String("mysqlschema", "Seckill", "mysql schema")
	)
	flag.Parse() //解析命令行参数
	ctx := context.Background()
	config.InitLog()                                                             //初始化log
	mysql.Init(*mysqluser, *mysqlpassword, *mysqlhost, *mysqlport, *mysqlschema) //初始化mysql

	//定义服务
	//var authsvc service.AutorizeService = service.AutorizeServiceImpl{}
	var tokensvc service.TokenService = service.TokenServiceImpl{}
	var heathsvc service.HealthCheckService = service.AdminserviceImpl{}

	//生成endpoint
	CreatetokenServ := endpoint.MakeCreateTokenEp(tokensvc)
	RefreshServ := endpoint.MakeParseAndRefreshTokenEp(tokensvc)
	ParseServ := endpoint.MakeParseTokenEp(tokensvc)
	HealthServ := endpoint.MakeHealthCheckEp(heathsvc)
	TokenAuthServ := endpoint.TokenAuthEndpoint{
		CreatetokenServ,
		RefreshServ,
		ParseServ,
		HealthServ,
	}

	//将服务注册进路由
	mux := transport.MakeHttpHandler(ctx, TokenAuthServ, config.KitLogger)
	Httpsev := &http.Server{
		Handler: mux,
	}
	//定义主监听
	mainlisten, err := net.Listen("tcp", *serverhost+":"+*serverport)
	if err != nil {
		log.Fatal(err)
	}
	// Create a cmux.
	m := cmux.New(mainlisten)
	grpclistener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httplistener := m.Match(cmux.HTTP1Fast())

	//将服务注册进grpc路由,并开启服务监听
	gmux := transport.MakeGrpcTokenServer(ctx, TokenAuthServ)
	grpcserver := grpc.NewServer()
	pb.RegisterAuthorizeserviceServer(grpcserver, gmux)
	errchan := make(chan error, 1)
	go func() {
		config.Logger.Println("grpc Server start at port:" + (*serverport))
		errchan <- grpcserver.Serve(grpclistener)
	}()

	//开始http服务监听

	go func() {
		config.Logger.Println("Http Server start at port:" + (*serverport))
		errchan <- Httpsev.Serve(httplistener)
		// errchan <- http.ListenAndServe(*serverhost+":"+*serverport, mux)
	}()
	m.Serve()
	<-errchan
}
