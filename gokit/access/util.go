package access

import "context"

var (
	defaultLimit = 30
	defaultPage  = 1
)

//Errorer interface
type Errorer interface {
	error() error
}

type newRequest struct {
	Req interface{}     `json:"req,omitempty"`
	Ctx context.Context `json:"ctx,omitempty"`
}

type newResponse struct {
	ID  string `json:"id,omitempty"`
	Err error  `json:"error,omitempty"`
}

func (r newResponse) error() error { return r.Err }

type messageResponse struct {
	Message interface{} `json:"message,omitempty"`
	Err     error       `json:"error,omitempty"`
}

func (r messageResponse) error() error { return r.Err }

type statusResponse struct {
	Status string `json:"status,omitempty"`
	Err    error  `json:"error,omitempty"`
}

func (r statusResponse) error() error { return r.Err }
