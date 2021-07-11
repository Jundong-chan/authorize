package main

import (
	"CommoditySpike/server/admin/config"
	"CommoditySpike/server/admin/endpoint"
	plugins "CommoditySpike/server/admin/plugins/log"
	"CommoditySpike/server/admin/service"
	"CommoditySpike/server/admin/transport"
	"CommoditySpike/server/pkg/mysql"
	"context"
	"flag"
	"net/http"
)

func main() {
	//预先处理
	var (
		//参数说明：命令行参数名称，默认值，参数说明（如果使用-help会给出这个说明提示）
		//使用变量时要加 *
		serverhost    = flag.String("host", "", "server host")
		serverport    = flag.String("port", "8809", "server port")
		mysqluser     = flag.String("mysqluser", "root", "mysql user")
		mysqlpassword = flag.String("mysqlpassword", "gzhuchan", "mysql password")
		mysqlhost     = flag.String("mysqlhost", "116.62.214.44", "mysql host")
		mysqlport     = flag.String("mysqlport", "3306", "mysql port")
		mysqlschema   = flag.String("mysqlschema", "Seckill", "mysql schema")
	)
	flag.Parse() //解析命令行参数
	ctx := context.Background()
	config.InitLog() //初始化log
	//定义服务及组件
	var svc1 service.ProductService = service.ProductServiceImpl{}
	var svc2 service.ActivityService = service.ActivityServiceImpl{}
	var svc3 service.UserService = service.UserServiceImpl{}
	var svc4 service.AdminService = service.AdminserviceImpl{}
	var svc5 service.OrderService = service.OrderServiceImpl{}

	//加入日志收中间件
	svc1 = plugins.ProductServiceLoggingMidWare(config.KitLogger)(svc1)
	svc2 = plugins.ActivityServiceLoggingMidware(config.KitLogger)(svc2)
	svc3 = plugins.UserServiceLoggingMidWare(config.KitLogger)(svc3)
	svc4 = plugins.HealthCheckLoggerMidWare(config.KitLogger)(svc4)
	svc5 = plugins.OrderServiceLoggingMidware(config.KitLogger)(svc5)

	CreateProductserv := endpoint.MakeCraeteProductEp(svc1)
	GetProductserv := endpoint.MakeGetProductListEp(svc1)
	ModifyProductserv := endpoint.MakeModifyProductEp(svc1)
	DeleteProductserv := endpoint.MakeDeleteProductEp(svc1)
	QureyProductserv := endpoint.MakeQureyProductEp(svc1)
	ProductCountsserv := endpoint.MakeProductCountEp(svc1)
	PageProductserv := endpoint.MakePageProductEp(svc1)
	CreateActivityserv := endpoint.MakeCreateActivityEp(svc2)
	GetActivserv := endpoint.MakeGetActivityListEp(svc2)
	CreateUserserv := endpoint.MakeCreateUserEp(svc3)
	GetUserserv := endpoint.MakeGetUserListEp(svc3)
	QureyUserserv := endpoint.MakeQueryUserEp(svc3)
	LoginUserServ := endpoint.MakeLoginEp(svc3)
	Healthserv := endpoint.MakeHealthCheckEp(svc4)
	QureyOrderserv := endpoint.MakeQureyOrderEp(svc5)
	ModifyOrderserv := endpoint.MakeModifyOrderEp(svc5)

	adminserv := endpoint.AdminEndpoints{
		CreateProductserv,
		GetProductserv,
		ModifyProductserv,
		DeleteProductserv,
		CreateActivityserv,
		GetActivserv,
		CreateUserserv,
		GetUserserv,
		QureyUserserv,
		QureyProductserv,
		LoginUserServ,
		ModifyOrderserv,
		QureyOrderserv,
		ProductCountsserv,
		PageProductserv,
		Healthserv,
	}

	handle := transport.MakeHttpHandler(ctx, adminserv, config.KitLogger)
	errchan := make(chan error, 1)
	//进行服务注册，开启监听
	go func() {
		mysql.Init(*mysqluser, *mysqlpassword, *mysqlhost, *mysqlport, *mysqlschema)
		config.Logger.Println("Http Server start at port:" + (*serverport))
		errchan <- http.ListenAndServe(*serverhost+":"+*serverport, handle)
	}()
	<-errchan
}
