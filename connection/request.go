package connection

import (
	"net/http"
)

// Every request to AlphaVantage should implement this interface
type AlphaVantageRequest interface {
	QueryFunction() QueryFunction
	QueryParams() map[string]string
	ParseResponse(response *http.Response) (interface{}, error)
}
