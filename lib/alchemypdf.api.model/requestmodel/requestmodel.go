package requestmodel

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"alchemypdf.api/lib/alchemypdf.api.util/errorlocal"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Schema struct {
		RequestID         int    `db:"request_id"`
		RequestExternalID string `db:"request_external_id"`
		RequestStateKey   string `db:"request_state_key"`

		ClientReference string `db:"client_reference"`

		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`

		Content      string `db:"content"`
		CallbackURL  string `db:"callback_url"`
		AttemptCount int    `db:"attempt_count"`
	}

	IRequestModel interface {
		Insert(payload Schema) ([]*Schema, *errorlocal.AppError)
		SelectByClientReference(clientReference string) ([]*Schema, *errorlocal.AppError)
		SelectWithContentByClientReference(clientReference string) ([]*Schema, *errorlocal.AppError)
		SelectByState(stateKey string) ([]*Schema, *errorlocal.AppError)
		StateUpdate(payload Schema) ([]*Schema, *errorlocal.AppError)
		AttemptCountIncrement(payload Schema) ([]*Schema, *errorlocal.AppError)
	}

	RequestModel struct {
		conn *pgxpool.Pool
	}
)

func NewRequestModel(conn *pgxpool.Pool) RequestModel {
	return RequestModel{conn}
}

func (m RequestModel) Insert(payload Schema) ([]*Schema, *errorlocal.AppError) {
	const op = "RequestModel.Insert"

	const query string = `
		INSERT INTO request
				(
					request_external_id
					, request_state_id
					, client_reference
					, content
					, callback_url
					, attempt_count
					, created_at
				)

			VALUES
			(
				$1
				, (SELECT rs.request_state_id FROM request_state rs WHERE rs.request_state_key = $2)
				, $3
				, $4
				, $5
				, $6
				, $7
			)

			RETURNING request_id, request_external_id
	`

	var dst []*Schema

	if err := pgxscan.Select(
		context.Background(),
		m.conn,
		&dst,
		query,
		strings.TrimSpace(payload.RequestExternalID),
		strings.TrimSpace(payload.RequestStateKey),
		strings.TrimSpace(payload.ClientReference),
		strings.TrimSpace(payload.Content),
		strings.TrimSpace(payload.CallbackURL),
		payload.AttemptCount,
		payload.CreatedAt,
	); err != nil {
		return nil, errorlocal.UnknownErr(err.Error(), op, err)
	}

	return dst, nil
}

func (m RequestModel) SelectByClientReference(clientReference string) ([]*Schema, *errorlocal.AppError) {
	const op = "RequestModel.SelectByClientReference"

	const query string = `
		SELECT r.request_id
			, r.request_external_id
			, rs.request_state_key
			, r.client_reference
			, r.callback_url
			, r.attempt_count
			, r.created_at
			, r.updated_at

			FROM request r

			INNER JOIN request_state rs
				ON rs.request_state_id = r.request_state_id

			WHERE r.client_reference = $1
	`

	var dst []*Schema

	if err := pgxscan.Select(
		context.Background(),
		m.conn,
		&dst,
		query,
		clientReference,
	); err != nil {
		return nil, errorlocal.UnknownErr(err.Error(), op, err)
	}

	return dst, nil
}

func (m RequestModel) SelectWithContentByClientReference(clientReference string) ([]*Schema, *errorlocal.AppError) {
	const op = "RequestModel.SelectWithContentByClientReference"

	const query string = `
		SELECT r.request_id
			, r.request_external_id
			, rs.request_state_key
			, r.client_reference
			, r.content
			, r.callback_url
			, r.attempt_count
			, r.created_at
			, r.updated_at

			FROM request r

			INNER JOIN request_state rs
				ON rs.request_state_id = r.request_state_id

			WHERE r.client_reference = $1
	`

	var dst []*Schema

	if err := pgxscan.Select(
		context.Background(),
		m.conn,
		&dst,
		query,
		clientReference,
	); err != nil {
		return nil, errorlocal.UnknownErr(err.Error(), op, err)
	}

	return dst, nil
}

func (m RequestModel) SelectByState(stateKey string) ([]*Schema, *errorlocal.AppError) {
	const op = "RequestModel.SelectByState"

	const query string = `
		SELECT r.request_id
			, r.request_external_id
			, rs.request_state_key
			, r.client_reference
			, r.content
			, r.callback_url
			, r.attempt_count
			, r.created_at
			, r.updated_at

			FROM request r

			INNER JOIN request_state rs
				ON rs.request_state_id = r.request_state_id

			WHERE rs.request_state_key = $1
				AND r.attempt_count <= 3

			LIMIT 1
	`

	var dst []*Schema

	if err := pgxscan.Select(
		context.Background(),
		m.conn,
		&dst,
		query,
		stateKey,
	); err != nil {
		return nil, errorlocal.UnknownErr(err.Error(), op, err)
	}

	return dst, nil
}

func (m RequestModel) StateUpdate(payload Schema) ([]*Schema, *errorlocal.AppError) {
	const op = "RequestModel.StateUpdate"

	const query string = `
		UPDATE request

			SET updated_at = $2
				, request_state_id = (SELECT request_state_id FROM request_state WHERE request_state_key = $3)

				WHERE request_id = $1

			RETURNING request_id, request_external_id
	`

	var dst []*Schema

	if err := pgxscan.Select(
		context.Background(),
		m.conn,
		&dst,
		query,
		payload.RequestID,
		payload.UpdatedAt.Time,
		strings.TrimSpace(payload.RequestStateKey),
	); err != nil {
		return nil, errorlocal.UnknownErr(err.Error(), op, err)
	}

	return dst, nil
}

func (m RequestModel) AttemptCountIncrement(payload Schema) ([]*Schema, *errorlocal.AppError) {
	const op = "RequestModel.AttemptCountIncrement"

	const query string = `
		UPDATE request

			SET updated_at = $2
				, attempt_count = attempt_count + 1

				WHERE request_id = $1

			RETURNING request_id, request_external_id
	`

	var dst []*Schema

	if err := pgxscan.Select(
		context.Background(),
		m.conn,
		&dst,
		query,
		payload.RequestID,
		payload.UpdatedAt.Time,
	); err != nil {
		return nil, errorlocal.UnknownErr(err.Error(), op, err)
	}

	return dst, nil
}
