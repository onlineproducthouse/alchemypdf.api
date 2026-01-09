package ioc

import (
	"alchemypdf.api/lib/alchemypdf.api.service/healthchecksvc"
	"alchemypdf.api/lib/alchemypdf.api.service/requestsvc"
)

type IIoc interface {
	HealthCheck() healthchecksvc.HealthCheckService
	RequestService() requestsvc.RequestService
}
