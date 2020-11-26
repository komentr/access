package auth

//Errorer interface declaration
type Errorer interface {
	error() error
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
