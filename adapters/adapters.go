package adapters

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
}

type BaseAdapter struct {
	client *http.Client
}

func NewAdapter() *BaseAdapter {
	return &BaseAdapter{
		&http.Client{},
	}
}

func (a *BaseAdapter) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", "ldst_touchstone=1;ldst_is_support_browser=1;ldst_visit=1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1")
	
	return a.client.Do(req)
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
