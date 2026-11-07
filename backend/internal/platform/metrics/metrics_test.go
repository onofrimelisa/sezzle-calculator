package metrics_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/platform/metrics"
	"github.com/stretchr/testify/assert"
)

func TestNew_CreatesMetrics(t *testing.T) {
	m := metrics.New()
	assert.NotNil(t, m.RequestsTotal)
	assert.NotNil(t, m.RequestDuration)
	assert.NotNil(t, m.ErrorsTotal)
}
