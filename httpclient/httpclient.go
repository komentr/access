package httpclient

type HTTPClient struct {
	APIKey string
	URL    string
	Method string
	Data   interface{}
}

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type DomainHTTPClient interface {
	HandlerClient(HTTPClient) error
}

type WebhookMessageRequest struct {
	Name        string `json:"name"`
	Method      string `json:"method"`
	Description string `json:"description"`
	TID       string `json:"tid"`
	URL         string `json:"url"`
	APIKey      string `json:"api_key"`
	IsDisabled  bool   `json:"is_disabled"`
	IsDeleted   bool   `json:"is_deleted"`
}

type NewResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}
