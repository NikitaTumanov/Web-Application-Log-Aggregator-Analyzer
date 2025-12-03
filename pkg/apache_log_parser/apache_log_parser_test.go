package apachelogparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	line := `192.168.1.100 - - [10/Oct/2023:14:32:13 +0300] "GET /api/v1/not-found-endpoint HTTP/1.1" 404 234 "https://example.com/dashboard" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"`
	apacheParser := NewApacheLogParser()

	logEntity, err := apacheParser.Parse(line)
	require.NoError(t, err)

	assert.Equal(t, logEntity.IP, "192.168.1.100")
	assert.Equal(t, logEntity.Timestamp, time.Time(time.Date(2023, time.October, 10, 14, 32, 13, 0, time.Local)))
	assert.Equal(t, logEntity.Method, "GET")
	assert.Equal(t, logEntity.Path, "/api/v1/not-found-endpoint")
	assert.Equal(t, logEntity.Status, 404)
	assert.Equal(t, logEntity.Size, 234)
	assert.Equal(t, logEntity.Referer, "https://example.com/dashboard")
	assert.Equal(t, logEntity.UserAgent, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
}
