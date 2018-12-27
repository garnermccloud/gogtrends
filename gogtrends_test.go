package gogtrends

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const locUS = "US"

func TestDailyTrending(t *testing.T) {
	resp, err := Daily(context.Background(), locUS)
	assert.NoError(t, err)
	assert.True(t, resp.StatusCode == http.StatusOK)
}

func TestRealtimeTrending(t *testing.T) {
	resp, err := Realtime(context.Background(), locUS)
	assert.NoError(t, err)
	assert.True(t, resp.StatusCode == http.StatusOK)
}