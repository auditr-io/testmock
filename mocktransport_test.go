package testmock

import (
	"net/http"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	m := &MockTransport{
		RoundTripFn: func(m *MockTransport, req *http.Request) (*http.Response, error) {
			m.MethodCalled("RoundTrip")

			return &http.Response{
				StatusCode: 200,
			}, nil
		},
	}

	m.On("RoundTrip").Return().Once()

	c := &http.Client{
		Transport: m,
	}

	req, _ := http.NewRequest(http.MethodGet, "fake", nil)
	c.Do(req)
	m.AssertExpectations(t)
}
