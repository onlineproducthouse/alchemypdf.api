package ioc

import (
	infrastructure "alchemypdf.api/lib/alchemypdf.api.infrastructure"
	model "alchemypdf.api/lib/alchemypdf.api.model"
	"alchemypdf.api/lib/alchemypdf.api.service/healthchecksvc"
	"alchemypdf.api/lib/alchemypdf.api.service/requestsvc"
)

type Ioc struct {
	infra infrastructure.IInfrastructure
	model model.IModel
}

func NewIoc(infra infrastructure.IInfrastructure, model model.IModel) Ioc {
	return Ioc{infra, model}
}

func (ioc Ioc) HealthCheck() healthchecksvc.HealthCheckService {
	return healthchecksvc.NewHealthCheckService(
		ioc.infra.Logger(),
		ioc.infra.DBConn(),
	)
}

func (ioc Ioc) RequestService() requestsvc.RequestService {
	return *requestsvc.NewRequestService(ioc.infra.Logger(), ioc.model.RequestModel())
}
