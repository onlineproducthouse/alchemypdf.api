package model

import (
	"alchemypdf.api/lib/alchemypdf.api.model/requestmodel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	IModel interface {
		RequestModel() requestmodel.RequestModel
	}

	Model struct {
		requestModel requestmodel.RequestModel
	}
)

func NewModel(conn *pgxpool.Pool) Model {
	return Model{
		requestModel: requestmodel.NewRequestModel(conn),
	}
}

func (m Model) RequestModel() requestmodel.RequestModel {
	return m.requestModel
}
