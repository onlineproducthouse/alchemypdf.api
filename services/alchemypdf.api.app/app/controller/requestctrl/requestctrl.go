package requestctrl

import (
	"fmt"
	"time"

	constant "alchemypdf.api/lib/alchemypdf.api.constant"
	"alchemypdf.api/lib/alchemypdf.api.contract/requestcontract"
	"alchemypdf.api/lib/alchemypdf.api.infrastructure/config"
	"alchemypdf.api/lib/alchemypdf.api.service/requestsvc"
	"github.com/labstack/echo/v4"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponseutil"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpstatusutil"
	"github.com/onlineproducthouse/alchemypdf.api.logger/loggylog"
)

type (
	IRequestCtrl interface {
		HandleCreate(c echo.Context) error
		HandleGetByClientReference(c echo.Context) error
		HandleGetWithContentByClientReference(c echo.Context) error
		HandleGetPending(c echo.Context) error
		HandleComplete(c echo.Context) error
		HandleCallback(c echo.Context) error
	}

	RequestCtrl struct {
		config         config.IConfig
		logger         loggylog.ILoggyLog
		requestService requestsvc.IRequestService
	}

	CreateRequest struct {
		ClientReference string `json:"clientReference"`
		Content         string `json:"content"`
		CallbackURL     string `json:"callbackUrl"`
	}

	Request struct {
		RequestID         int    `json:"requestId"`
		RequestExternalID string `json:"requestExternalId"`
		RequestStateKey   string `json:"requestStateKey"`

		ClientReference string `json:"clientReference"`
		Content         string `json:"content"`
		CallbackURL     string `json:"callbackUrl"`
		AttemptCount    int    `json:"attemptCount"`

		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	CompleteRequest struct {
		Success   bool `json:"success"`
		RequestID int  `json:"requestId"`
	}

	CallbackRequest struct {
		Success   bool   `json:"success"`
		PDFString string `json:"pdfString"`
	}
)

func NewRequestCtrl(
	config config.IConfig,
	logger loggylog.ILoggyLog,
	requestService requestsvc.IRequestService,
) *RequestCtrl {
	return &RequestCtrl{
		config,
		logger,
		requestService,
	}
}

// Request/Create godoc
// @id Request.Create
// @tags Request
// @summary Create a request for converting HTML to PDF
// @router /v1/Request/Create [post]
// @accept json
// @produce json
// @success 200 {object} httpresponseutil.Response
// @Failure 400,500 {object} httpresponseutil.Response
// @param x-api-key header string true "API Key"
// @param payload body CreateRequest true "CreateRequest"
func (ctrl RequestCtrl) HandleCreate(c echo.Context) error {
	const op = "RequestCtrl.HandleCreate"

	ctrl.logger.Info(fmt.Sprintf("[%s]: starting", op))

	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		ctrl.logger.Debug(`"RequestCtrl.HandleCreate": failed to parse body`)
		ctrl.logger.AppError(httperrorutil.UnknownErr(err.Error(), op, err))

		statusCode, _ := httpstatusutil.InternalServerError()
		ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))
		return c.JSON(statusCode, httpresponseutil.Default(err.Error(), statusCode))
	}

	if err := ctrl.requestService.Create(requestcontract.CreateRequest{
		ClientReference: req.ClientReference,
		Content:         req.Content,
		CallbackURL:     req.CallbackURL,
	}); err != nil {
		ctrl.logger.AppError(httperrorutil.CatchErr(err, op))
		ctrl.logger.HTTPResponse(err.StatusCode(), c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))
		return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
	}

	statusCode, statusCodeText := httpstatusutil.Ok()

	ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))

	return c.JSON(statusCode, httpresponseutil.Default(statusCodeText, statusCode))
}

// Request/GetByClientReference godoc
// @id Request.GetByClientReference
// @tags Request
// @summary Gets a list of request created by the client without content
// @router /v1/Request/GetByClientReference/{ClientReference} [get]
// @accept json
// @produce json
// @success 200 {object} []Request
// @Failure 400,500 {object} httpresponseutil.Response
// @param x-api-key header string true "API Key"
// @param ClientReference path string true "ClientReference"
func (ctrl RequestCtrl) HandleGetByClientReference(c echo.Context) error {
	const op = "RequestCtrl.HandleGetByClientReference"

	ctrl.logger.Info(fmt.Sprintf("[%s]: starting", op))

	res, err := ctrl.requestService.GetByClientReference(c.Param("ClientReference"))
	if err != nil {
		ctrl.logger.AppError(httperrorutil.CatchErr(err, op))
		return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
	}

	statusCode, _ := httpstatusutil.Ok()

	ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))

	return c.JSON(statusCode, mapListToDomain(res))
}

// Request/GetWithContentByClientReference godoc
// @id Request.GetWithContentByClientReference
// @tags Request
// @summary Gets a list of request created by the client with content
// @router /v1/Request/GetWithContentByClientReference/{ClientReference} [get]
// @accept json
// @produce json
// @success 200 {object} []Request
// @Failure 400,500 {object} httpresponseutil.Response
// @param x-api-key header string true "API Key"
// @param ClientReference path string true "ClientReference"
func (ctrl RequestCtrl) HandleGetWithContentByClientReference(c echo.Context) error {
	const op = "RequestCtrl.HandleGetWithContentByClientReference"

	ctrl.logger.Info(fmt.Sprintf("[%s]: starting", op))

	res, err := ctrl.requestService.GetWithContentByClientReference(c.Param("ClientReference"))
	if err != nil {
		ctrl.logger.AppError(httperrorutil.CatchErr(err, op))
		return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
	}

	statusCode, _ := httpstatusutil.Ok()

	ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))

	return c.JSON(statusCode, mapListToDomain(res))
}

// Request/GetPending godoc
// @id Request.GetPending
// @tags Request
// @summary Gets a list of request created by the client with content
// @router /v1/Request/GetPending [get]
// @accept json
// @produce json
// @success 200 {object} Request
// @Failure 400,500 {object} httpresponseutil.Response
// @param x-api-key header string true "API Key"
func (ctrl RequestCtrl) HandleGetPending(c echo.Context) error {
	const op = "RequestCtrl.HandleGetPending"

	ctrl.logger.Info(fmt.Sprintf("[%s]: starting", op))

	res, err := ctrl.getPending()
	if err != nil {
		ctrl.logger.AppError(httperrorutil.CatchErr(err, op))
		return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
	}

	if err := ctrl.requestService.StateUpdate(requestcontract.StateUpdateRequest{
		RequestID:       res.RequestID,
		RequestStateKey: constant.RequestStateInProgress,
	}); err != nil {
		ctrl.logger.AppError(httperrorutil.CatchErr(err, op))
		return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
	}

	if err := ctrl.requestService.AttemptCountIncrement(res.RequestID); err != nil {
		ctrl.logger.AppError(httperrorutil.CatchErr(err, op))
		return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
	}

	statusCode, _ := httpstatusutil.Ok()

	ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))

	return c.JSON(statusCode, res)
}

func mapListToDomain(l []*requestcontract.Request) []*Request {
	r := []*Request{}

	for _, v := range l {
		r = append(r, mapToDomain(v))
	}

	return r
}

func mapToDomain(v *requestcontract.Request) *Request {
	return &Request{
		RequestID:         v.RequestID,
		RequestExternalID: v.RequestExternalID,
		RequestStateKey:   v.RequestStateKey,
		ClientReference:   v.ClientReference,
		Content:           v.Content,
		CallbackURL:       v.CallbackURL,
		AttemptCount:      v.AttemptCount,
		CreatedAt:         v.CreatedAt,
		UpdatedAt:         v.UpdatedAt,
	}
}

func (ctrl RequestCtrl) getPending() (*Request, *httperrorutil.AppError) {
	const op = "RequestCtrl.getPending"

	ctrl.logger.Info(fmt.Sprintf("[%s]: starting", op))

	res, err := ctrl.requestService.GetByState(constant.RequestStatePending)
	if err != nil {
		return nil, httperrorutil.CatchErr(err, op)
	}

	if res.AttemptCount >= 3 {
		if err := ctrl.requestService.StateUpdate(requestcontract.StateUpdateRequest{
			RequestID:       res.RequestID,
			RequestStateKey: constant.RequestStateAttemptLimitReached,
		}); err != nil {
			return nil, httperrorutil.CatchErr(err, op)
		}

		return ctrl.getPending()
	}

	return mapToDomain(res), nil
}

// Request/Complete godoc
// @id Request.Complete
// @tags Request
// @summary Completes a request, sets it to complete if successful, pending if failed
// @router /v1/Request/Complete [post]
// @accept json
// @produce json
// @success 200 {object} httpresponseutil.Response
// @Failure 400,500 {object} httpresponseutil.Response
// @param x-api-key header string true "API Key"
// @param payload body CompleteRequest true "CompleteRequest"
func (ctrl RequestCtrl) HandleComplete(c echo.Context) error {
	const op = "RequestCtrl.HandleComplete"

	ctrl.logger.Info(fmt.Sprintf("[%s]: starting", op))

	var completeReq CompleteRequest
	if err := c.Bind(&completeReq); err != nil {
		ctrl.logger.Debug(fmt.Sprintf("[%s]: failed to parse body", op))
		ctrl.logger.AppError(httperrorutil.UnknownErr(err.Error(), op, err))

		statusCode, _ := httpstatusutil.InternalServerError()
		ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))
		return c.JSON(statusCode, httpresponseutil.Default(err.Error(), statusCode))
	}

	getRequest, getRequestErr := ctrl.requestService.GetByID(completeReq.RequestID)
	if getRequestErr != nil {
		ctrl.logger.AppError(httperrorutil.CatchErr(getRequestErr, op))
		return c.JSON(getRequestErr.StatusCode(), httpresponseutil.Default(getRequestErr.Error(), getRequestErr.StatusCode()))
	}

	if getRequest.RequestStateKey == constant.RequestStateInProgress {
		requestStateKey := constant.RequestStatePending
		if completeReq.Success {
			requestStateKey = constant.RequestStateCompleted
		}

		if err := ctrl.requestService.StateUpdate(requestcontract.StateUpdateRequest{
			RequestID:       completeReq.RequestID,
			RequestStateKey: requestStateKey,
		}); err != nil {
			ctrl.logger.AppError(httperrorutil.CatchErr(err, op))
			return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
		}
	}

	statusCode, statusCodeText := httpstatusutil.Ok()
	return c.JSON(statusCode, httpresponseutil.Default(statusCodeText, statusCode))
}

// Request/Callback godoc
// @id Request.Callback
// @tags Request
// @summary Callback for integration tests
// @router /v1/Request/Callback [post]
// @accept json
// @produce json
// @success 200 {object} httpresponseutil.Response
// @Failure 400,500 {object} httpresponseutil.Response
// @param payload body CallbackRequest true "CallbackRequest"
func (ctrl RequestCtrl) HandleCallback(c echo.Context) error {
	const op = "RequestCtrl.HandleCallback"

	ctrl.logger.Info(fmt.Sprintf("[%s]: starting", op))

	var callbackReq CallbackRequest
	if err := c.Bind(&callbackReq); err != nil {
		ctrl.logger.Debug(fmt.Sprintf("[%s]: failed to parse body", op))
		ctrl.logger.AppError(httperrorutil.UnknownErr(err.Error(), op, err))

		statusCode, _ := httpstatusutil.InternalServerError()
		ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))
		return c.JSON(statusCode, httpresponseutil.Default(err.Error(), statusCode))
	}

	ctrl.logger.Info(fmt.Sprintf("callbackReq: %t", callbackReq.Success))
	ctrl.logger.Info(fmt.Sprintf("callbackReq: %d", len(callbackReq.PDFString)))

	statusCode, statusCodeText := httpstatusutil.Ok()
	return c.JSON(statusCode, httpresponseutil.Default(statusCodeText, statusCode))
}
