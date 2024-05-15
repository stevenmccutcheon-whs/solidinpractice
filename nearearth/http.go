package nearearth

import (
	"encoding/json"
	"net/http"
)

type ImpactHTTPService struct {
	url    string
	client *http.Client
}

func (s ImpactHTTPService) GetAll() (Impacts, error) {
	return s.readEventsFromHTTP()
}

func (s ImpactHTTPService) readEventsFromHTTP() (Impacts, error) {
	resp, err := s.client.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var res Impacts
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type HTTPServiceOption func(*ImpactHTTPService)

func WithTestURL(url string) HTTPServiceOption {
	return func(s *ImpactHTTPService) {
		s.url = url
	}
}
func NewImpactHTTPService(opts ...HTTPServiceOption) ImpactHTTPService {
	url := "https://data.nasa.gov/resource/2vr3-k9wn.json"
	client := http.DefaultClient
	defaultService := ImpactHTTPService{url: url, client: client}
	for _, opt := range opts {
		opt(&defaultService)
	}
	return defaultService
}
