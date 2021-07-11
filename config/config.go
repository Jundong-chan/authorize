package config

//这是配置包，用于配置工具的初始化
import (
	"log"
	"os"
	"sync"

	kitlog "github.com/go-kit/kit/log"
	"github.com/samuel/go-zookeeper/zk"
)

var Logger *log.Logger
var KitLogger kitlog.Logger
var ZKConn *zk.Conn
var ZKProductPath = "/product"

func InitLog() {
	//LstdFlags表示 log输出时默认加上日期+时间
	Logger = log.New(os.Stderr, "", log.LstdFlags)
	KitLogger = kitlog.NewLogfmtLogger(os.Stderr)
	//With函数加上log日志的记录字段
	KitLogger = kitlog.With(KitLogger, "calling time", kitlog.DefaultTimestampUTC)
	KitLogger = kitlog.With(KitLogger, "caller", kitlog.DefaultCaller)
}

//存储在zk的商品信息
type ProductInfo struct {
	ProductId    string  `json:"productid"` //数据库自增
	ProductName  string  `json:"productname"`
	Price        float64 `json:"price,string"`
	Total        int     `json:"total,string"`
	Status       int     `json:"status,string"`
	Detail       string  `json:"detail"`
	SecstartTime string  `json:"secstarttime"`
	SecendTime   string  `json:"secendtime"`
}

//由商品id与商品信息构成的map,存放从zookeeper中读取的商品数据
var Secproductinfo = make(map[string]*ProductInfo, 1024)
var ProductInfoMapRWLocker sync.Mutex
