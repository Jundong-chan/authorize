package plugins

//这是工具包，相关依赖在此定义，比如中间件
import (
	"github.com/Jundong-chan/authorize/service"
	"github.com/Jundong-chan/seckill/model"
	"fmt"
	"time"

	kitlog "github.com/go-kit/kit/log"
	"github.com/gohouse/gorose"
)

//定义实现该服务接口的中间件，加入logger包裹service的服务实现
type serviceLoggingMidWare struct {
	next   service.AdminService
	logger kitlog.Logger
}

//使用闭包包装中间件
func HealthCheckLoggerMidWare(logger kitlog.Logger) service.ServiceMidWare {
	return func(next service.AdminService) service.AdminService {
		return serviceLoggingMidWare{next, logger}
	}
}

func (mw serviceLoggingMidWare) HealthCheck() bool {

	status := mw.next.HealthCheck()

	defer func(begin time.Time) {
		mw.logger.Log("function", "HealthCheck", "Status", status, "took", time.Since(begin))
	}(time.Now())
	return status
}

type userServiceLoggingMidWare struct {
	next   service.UserService
	logger kitlog.Logger
}

func UserServiceLoggingMidWare(logger kitlog.Logger) service.UserServiceMidWare {
	return func(next service.UserService) service.UserService {
		return userServiceLoggingMidWare{next, logger}
	}
}

func (mw userServiceLoggingMidWare) CreateUser(user *model.User) error {
	fmt.Print("CreateUser")
	fmt.Println(*user)
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "CreateUser",
			"took", time.Since(begin),
		)
	}(time.Now())
	err := mw.next.CreateUser(user)
	return err
}

func (mw userServiceLoggingMidWare) GetUserList() ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function: ", "GetUserList ", "took: ", time.Since(begin))
	}(time.Now())
	Result, err := mw.next.GetUserList()
	return Result, err
}

func (mw userServiceLoggingMidWare) QureyUser(field string, condition string, value interface{}) ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function: ", "QureyUser ", "took: ", time.Since(begin))
	}(time.Now())
	Result, err := mw.next.QureyUser(field, condition, value)
	return Result, err
}

func (mw userServiceLoggingMidWare) CheckUser(field string, username string, password string) ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function: ", "CheckUser ", "took: ", time.Since(begin))
	}(time.Now())
	Result, err := mw.next.CheckUser(field, username, password)
	return Result, err
}

// func (mw userServiceLoggingMidWare) Login(token string, field string, account string, password string) ([]gorose.Data, string, error) {
// 	defer func(begin time.Time) {
// 		mw.logger.Log("function: ", "Login ", "took: ", time.Since(begin))
// 	}(time.Now())
// 	Result, token, err := mw.next.Login(token, field, account, password)
// 	return Result, token, err
// }

type productServiceLoggingMidWare struct {
	next   service.ProductService
	logger kitlog.Logger
}

func ProductServiceLoggingMidWare(logger kitlog.Logger) service.ProductServiceMidWare {
	return func(next service.ProductService) service.ProductService {
		return productServiceLoggingMidWare{next, logger}
	}
}

func (mw productServiceLoggingMidWare) CreateProduct(product *model.Product) error {
	defer func(begin time.Time) {
		mw.logger.Log("function", "CreateProduct", "took", time.Since(begin))
	}(time.Now())
	err := mw.next.CreateProduct(product)
	return err
}

func (mw productServiceLoggingMidWare) GetProductList() ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function", "CreateProduct", "took", time.Since(begin))
	}(time.Now())
	Result, err := mw.next.GetProductList()
	return Result, err
}
func (mw productServiceLoggingMidWare) ModifyProduct(product *model.Product) error {
	defer func(begin time.Time) {
		mw.logger.Log("function", "ModifyProduct", "took ", time.Since(begin))
	}(time.Now())
	err := mw.next.ModifyProduct(product)
	return err
}
func (mw productServiceLoggingMidWare) DeleteProduct(id string) error {
	defer func(begin time.Time) {
		mw.logger.Log("function", "DeleteProduct", "took ", time.Since(begin))
	}(time.Now())
	err := mw.next.DeleteProduct(id)
	return err
}
func (mw productServiceLoggingMidWare) QureyProductByUserId(field string, condition string, value interface{}) ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function", "QureyProductByUserId", "took ", time.Since(begin))
	}(time.Now())
	result, err := mw.next.QureyProductByUserId(field, condition, value)
	return result, err
}
func (mw productServiceLoggingMidWare) GetProductCounts() (int, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function", "GetproductCounts", "took ", time.Since(begin))
	}(time.Now())
	result, err := mw.next.GetProductCounts()
	return result, err
}
func (mw productServiceLoggingMidWare) PageProducts(page int, limit int) ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function", "PageProducts", "took ", time.Since(begin))
	}(time.Now())
	result, err := mw.next.PageProducts(page, limit)
	return result, err
}

type activityServiceLoggingMidware struct {
	next   service.ActivityService
	logger kitlog.Logger
}

func ActivityServiceLoggingMidware(logger kitlog.Logger) service.ActivityServiceMidWare {
	return func(next service.ActivityService) service.ActivityService {
		return activityServiceLoggingMidware{next, logger}
	}
}
func (mw activityServiceLoggingMidware) CreateActivity(activity *model.Activity) error {
	defer func(begin time.Time) {
		mw.logger.Log("function", "CreateActivity", "took", time.Since(begin))
	}(time.Now())
	return mw.next.CreateActivity(activity)
}

func (mw activityServiceLoggingMidware) GetActivityList() ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function", "GetActivityList", "took", time.Since(begin))
	}(time.Now())
	return mw.next.GetActivityList()
}

type orderServiceLoggingMidware struct {
	next   service.OrderService
	logger kitlog.Logger
}

func OrderServiceLoggingMidware(logger kitlog.Logger) service.OrderServiceMidWare {
	return func(next service.OrderService) service.OrderService {
		return orderServiceLoggingMidware{next, logger}
	}
}

func (mw orderServiceLoggingMidware) QureyOrder(value interface{}) ([]gorose.Data, error) {
	defer func(begin time.Time) {
		mw.logger.Log("function", "QureyOrder", "took", time.Since(begin))
	}(time.Now())
	return mw.next.QureyOrder(value)
}

func (mw orderServiceLoggingMidware) ModifyOrder(data map[string]interface{}, field string, value interface{}) error {
	defer func(begin time.Time) {
		mw.logger.Log("function", "ModifyOrder", "took", time.Since(begin))
	}(time.Now())
	return mw.next.ModifyOrder(data, field, value)
}
