package requestsvc

import (
	"database/sql"
	"fmt"
	"time"

	constant "alchemypdf.api/lib/alchemypdf.api.constant"
	"alchemypdf.api/lib/alchemypdf.api.contract/requestcontract"
	"alchemypdf.api/lib/alchemypdf.api.infrastructure/loglocal"
	"alchemypdf.api/lib/alchemypdf.api.model/requestmodel"
	"github.com/google/uuid"
	alchemypdfapihttputils "github.com/onlineproducthouse/alchemypdf.api.httputils"
)

type RequestService struct {
	logger loglocal.ILogger
	model  requestmodel.IRequestModel
}

type IRequestService interface {
	Create(payload requestcontract.CreateRequest) *alchemypdfapihttputils.AppError
	GetByClientReference(clientReference string) ([]*requestcontract.Request, *alchemypdfapihttputils.AppError)
	GetWithContentByClientReference(clientReference string) ([]*requestcontract.Request, *alchemypdfapihttputils.AppError)
	GetByState(stateKey string) (*requestcontract.Request, *alchemypdfapihttputils.AppError)
	StateUpdate(payload requestcontract.StateUpdateRequest) *alchemypdfapihttputils.AppError
	AttemptCountIncrement(requestId int) *alchemypdfapihttputils.AppError
}

func NewRequestService(
	logger loglocal.ILogger,
	model requestmodel.IRequestModel,
) *RequestService {
	return &RequestService{
		logger,
		model,
	}
}

func (h RequestService) Create(payload requestcontract.CreateRequest) *alchemypdfapihttputils.AppError {
	const op = "RequestService.Create"

	h.logger.Info(fmt.Sprintf("[%s]: starting", op))

	if payload.ClientReference == "" {
		return alchemypdfapihttputils.ValidationErr("request client reference is required", op, nil)
	}

	if payload.Content == "" {
		return alchemypdfapihttputils.ValidationErr("request content is required", op, nil)
	}

	if payload.CallbackURL == "" {
		return alchemypdfapihttputils.ValidationErr("request callback url is required", op, nil)
	}

	_create, _createErr := h.model.Insert(requestmodel.Schema{
		RequestExternalID: uuid.NewString(),
		RequestStateKey:   constant.RequestStatePending,
		ClientReference:   payload.ClientReference,
		Content:           payload.Content,
		CallbackURL:       payload.CallbackURL,
		AttemptCount:      0,
		CreatedAt:         time.Now(),
	})
	if _createErr != nil {
		return alchemypdfapihttputils.CatchErr(_createErr, op)
	}

	if len(_create) <= 0 {
		return alchemypdfapihttputils.UnknownErr("failed to create request", op, nil)
	}

	h.logger.Info(fmt.Sprintf("[%s]: done", op))

	return nil
}

func (h RequestService) GetByClientReference(clientReference string) ([]*requestcontract.Request, *alchemypdfapihttputils.AppError) {
	const op = "RequestService.GetByClientReference"

	h.logger.Info(fmt.Sprintf("[%s]: starting", op))

	if clientReference == "" {
		return nil, alchemypdfapihttputils.ValidationErr("request client reference is required", op, nil)
	}

	_select, _selectErr := h.model.SelectByClientReference(clientReference)
	if _selectErr != nil {
		return nil, alchemypdfapihttputils.CatchErr(_selectErr, op)
	}

	h.logger.Info(fmt.Sprintf("[%s]: done", op))

	return mapListToDomain(_select), nil
}

func (h RequestService) GetWithContentByClientReference(clientReference string) ([]*requestcontract.Request, *alchemypdfapihttputils.AppError) {
	const op = "RequestService.GetWithContentByClientReference"

	h.logger.Info(fmt.Sprintf("[%s]: starting", op))

	if clientReference == "" {
		return nil, alchemypdfapihttputils.ValidationErr("request client reference is required", op, nil)
	}

	_select, _selectErr := h.model.SelectWithContentByClientReference(clientReference)
	if _selectErr != nil {
		return nil, alchemypdfapihttputils.CatchErr(_selectErr, op)
	}

	h.logger.Info(fmt.Sprintf("[%s]: done", op))

	return mapListToDomain(_select), nil
}

func (h RequestService) GetByState(stateKey string) (*requestcontract.Request, *alchemypdfapihttputils.AppError) {
	const op = "RequestService.GetByState"

	h.logger.Info(fmt.Sprintf("[%s]: starting", op))

	if stateKey == "" {
		return nil, alchemypdfapihttputils.ValidationErr("request state key is required", op, nil)
	}

	_select, _selectErr := h.model.SelectByState(stateKey)
	if _selectErr != nil {
		return nil, alchemypdfapihttputils.CatchErr(_selectErr, op)
	}

	if len(_select) <= 0 {
		return nil, alchemypdfapihttputils.NotFoundErr("request not found", op, nil)
	}

	h.logger.Info(fmt.Sprintf("[%s]: done", op))

	return mapToDomain(_select[0]), nil
}

func mapListToDomain(l []*requestmodel.Schema) []*requestcontract.Request {
	r := []*requestcontract.Request{}

	for _, v := range l {
		r = append(r, mapToDomain(v))
	}

	return r
}

func mapToDomain(v *requestmodel.Schema) *requestcontract.Request {
	return &requestcontract.Request{
		RequestID:         v.RequestID,
		RequestExternalID: v.RequestExternalID,
		RequestStateKey:   v.RequestStateKey,
		ClientReference:   v.ClientReference,
		Content:           v.Content,
		CallbackURL:       v.CallbackURL,
		AttemptCount:      v.AttemptCount,
		CreatedAt:         v.CreatedAt,
		UpdatedAt:         v.UpdatedAt.Time,
	}
}

func (h RequestService) StateUpdate(payload requestcontract.StateUpdateRequest) *alchemypdfapihttputils.AppError {
	const op = "RequestService.StateUpdate"

	h.logger.Info(fmt.Sprintf("[%s]: starting", op))

	if payload.RequestID <= 0 {
		return alchemypdfapihttputils.ValidationErr("request id must be greater than zero", op, nil)
	}

	if payload.RequestStateKey == "" {
		return alchemypdfapihttputils.ValidationErr("request state key is required", op, nil)
	}

	updateState, updateStateErr := h.model.StateUpdate(requestmodel.Schema{
		RequestID:       payload.RequestID,
		RequestStateKey: payload.RequestStateKey,
		UpdatedAt:       sql.NullTime{Time: time.Now()},
	})
	if updateStateErr != nil {
		h.logger.AppError(updateStateErr)
		return alchemypdfapihttputils.CatchErr(updateStateErr, op)
	}

	if len(updateState) <= 0 {
		return alchemypdfapihttputils.UnknownErr("failed to update request state", op, nil)
	}

	h.logger.Info(fmt.Sprintf("[%s]: done", op))

	return nil
}

func (h RequestService) AttemptCountIncrement(requestId int) *alchemypdfapihttputils.AppError {
	const op = "RequestService.AttemptCountIncrement"

	h.logger.Info(fmt.Sprintf("[%s]: starting", op))

	if requestId <= 0 {
		return alchemypdfapihttputils.ValidationErr("request id must be greater than zero", op, nil)
	}

	attemptCountIncrement, attemptCountIncrementErr := h.model.AttemptCountIncrement(requestmodel.Schema{
		RequestID: requestId,
		UpdatedAt: sql.NullTime{Time: time.Now()},
	})
	if attemptCountIncrementErr != nil {
		h.logger.AppError(attemptCountIncrementErr)
		return alchemypdfapihttputils.CatchErr(attemptCountIncrementErr, op)
	}

	if len(attemptCountIncrement) <= 0 {
		return alchemypdfapihttputils.UnknownErr("failed to update request attempt count", op, nil)
	}

	h.logger.Info(fmt.Sprintf("[%s]: done", op))

	return nil
}
