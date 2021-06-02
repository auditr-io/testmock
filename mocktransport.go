package testmock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type RoundTrip func(m *MockTransport, req *http.Request) (*http.Response, error)

// MockTransport is a mock Transport client
type MockTransport struct {
	mock.Mock
	http.Transport
	roundTrip RoundTrip
}

// RoundTrip invokes the mocked roundTrip function
func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTrip(m, req)
}
