package base

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type WrongCodeError struct {
	Status string
	StatusCode int
}

func (e WrongCodeError) Error() string {
	return e.Status
}

type Adapter interface {
	Get(url string) (*http.Response, error)
	GetDocument(url string) (*goquery.Document, error)
	
	GetCharacter(id string) (Character, error)
}

type BaseAdapter struct {
	Client *http.Client
	Headers map[string]string
}

func NewAdapter(headers map[string]string) *BaseAdapter {
	return &BaseAdapter{
		&http.Client{},
		headers,
	}
}

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

func (a *BaseAdapter) MakeDocument(res *http.Response) (doc *goquery.Document, err error) {
	if res.StatusCode != 200 {
		return nil, WrongCodeError { res.Status, res.StatusCode }
	}
	
	return goquery.NewDocumentFromResponse(res)
}

func (a *BaseAdapter) GetDocument(url string) (doc *goquery.Document, err error) {
	res, err := a.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	return a.MakeDocument(res)
}
