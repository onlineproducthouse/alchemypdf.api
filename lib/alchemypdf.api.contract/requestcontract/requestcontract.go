package requestcontract

import "time"

type (
	CreateRequest struct {
		ClientReference string
		Content         string
		CallbackURL     string
	}

	Request struct {
		RequestID         int
		RequestExternalID string
		RequestStateKey   string

		ClientReference string

		Content      string
		CallbackURL  string
		AttemptCount int

		CreatedAt time.Time
		UpdatedAt time.Time
	}

	StateUpdateRequest struct {
		RequestID       int
		RequestStateKey string
	}
)
