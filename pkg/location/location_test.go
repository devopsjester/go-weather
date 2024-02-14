package location

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetLocation(t *testing.T) {
	ls := &LocationService{
		Client: &MockHttpClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				json := `{"country":"US","places":[{"place name":"Beverly Hills","state":"California"}]}`
				r := io.NopCloser(bytes.NewReader([]byte(json)))
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			},
		},
	}

	location, err := ls.GetLocation("90210")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if location.Places[0].City != "Beverly Hills" {
		t.Errorf("expected city to be 'Beverly Hills', got '%s'", location.Places[0].City)
	}

	if location.Places[0].State != "California" {
		t.Errorf("expected state to be 'California', got '%s'", location.Places[0].State)
	}
}