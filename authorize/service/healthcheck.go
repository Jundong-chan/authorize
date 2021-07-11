package service

type HealthCheckService interface {
	HealthCheck() bool
}
type AdminserviceImpl struct {
}

func (s AdminserviceImpl) HealthCheck() bool {
	return true
}
