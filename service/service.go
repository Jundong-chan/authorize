package service

//健康检查接口
type AdminService interface {
	HealthCheck() bool
}

type AdminserviceImpl struct {
}

func (s AdminserviceImpl) HealthCheck() bool {
	return true
}

//定义服务中间件形式
type ServiceMidWare func(AdminService) AdminService
