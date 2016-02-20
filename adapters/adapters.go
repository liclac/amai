package adapters

import (
	"net/http"
)

type Adapter interface {
	Get(url string) (*http.Response, error)
}
