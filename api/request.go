package api

import (
	"net/http"
)

// Every request to AlphaVantage should implement this interface
type AlphaVantageRequest interface {
	QueryFunction() QueryFunction
	URLParams() map[string]string
	ParseResponse(*http.Response) (interface{}, error)
}
