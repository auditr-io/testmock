package test

import (
	"testing"

	"github.com/honeycombio/libhoney-go"
	"github.com/honeycombio/libhoney-go/transmission"
	"github.com/stretchr/testify/assert"
)

// CreateHoneyMock creates a mock honeycomb client
func CreateHoneyMock(
	useMock bool,
	honeycombKey string,
	honeycombDataset string,
	t testing.TB,
) *libhoney.Client {
	var cfg libhoney.ClientConfig
	if useMock {
		cfg = libhoney.ClientConfig{
			APIKey:       "test-key",
			Dataset:      "test-dataset",
			APIHost:      "test-host",
			Transmission: &transmission.MockSender{},
			Logger:       &libhoney.DefaultLogger{},
		}
	} else {
		cfg = libhoney.ClientConfig{
			APIKey:  honeycombKey,
			Dataset: honeycombDataset,
		}
	}

	c, err := libhoney.NewClient(cfg)

	assert.Nil(t, err)
	return c
}
