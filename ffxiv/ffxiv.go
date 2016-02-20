package ffxiv

import (
	"net/http"
)

type FFXIVAdapter struct {
	client *http.Client
}

func NewAdapter() *FFXIVAdapter {
	return &FFXIVAdapter{
		&http.Client{},
	}
}

func (a *FFXIVAdapter) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", "ldst_touchstone=1;ldst_is_support_browser=1;ldst_visit=1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1")
	
	return a.client.Do(req)
}
