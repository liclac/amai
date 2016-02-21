package base

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

// Error raised when an HTTP request returns the wrong status code.
type WrongCodeError struct {
	Status string
	StatusCode int
}

// Return a string representation of the error.
func (e WrongCodeError) Error() string {
	return e.Status
}

// Interface for game-specific adapters.
type Adapter interface {
	Get(url string) (*http.Response, error)
	GetDocument(url string) (*goquery.Document, error)
	
	GetCharacter(id string) (interface{}, error)
}

// A basic adapter.
type BaseAdapter struct {
	Client *http.Client
	Headers map[string]string
}

// Creates a new adapter.
func NewAdapter(headers map[string]string) *BaseAdapter {
	return &BaseAdapter{
		&http.Client{},
		headers,
	}
}

// Get the raw response at the given URL, giving the adapter a chance to modify
// the request in various ways (such as appending headers).
func (a *BaseAdapter) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	for key, val := range a.Headers {
		req.Header.Add(key, val)
	}
	
	return a.Client.Do(req)
}

// Makes a goquery document out of a response.
func (a *BaseAdapter) MakeDocument(res *http.Response) (doc *goquery.Document, err error) {
	if res.StatusCode != 200 {
		return nil, WrongCodeError { res.Status, res.StatusCode }
	}
	
	return goquery.NewDocumentFromResponse(res)
}

// Gets a goquery document from the given URL, giving the adapter a chance to
// modify the request as needed.
func (a *BaseAdapter) GetDocument(url string) (doc *goquery.Document, err error) {
	res, err := a.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	return a.MakeDocument(res)
}
