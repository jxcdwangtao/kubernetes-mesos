package config

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func is_default(c *Config, t *testing.T) {
	assert := assert.New(t)

	assert.Equal(DefaultOfferTTL,                           c.OfferTTL.Duration)
	assert.Equal(DefaultOfferLingerTTL,                     c.OfferLingerTTL.Duration)
	assert.Equal(DefaultListenerDelay,                      c.ListenerDelay.Duration)
	assert.Equal(DefaultUpdatesBacklog,                     c.UpdatesBacklog)
	assert.Equal(DefaultFrameworkIdRefreshInterval,         c.FrameworkIdRefreshInterval.Duration)
	assert.Equal(DefaultInitialImplicitReconciliationDelay, c.InitialImplicitReconciliationDelay.Duration)
	assert.Equal(DefaultExplicitReconciliationMaxBackoff,   c.ExplicitReconciliationMaxBackoff.Duration)
	assert.Equal(DefaultExplicitReconciliationAbortTimeout, c.ExplicitReconciliationAbortTimeout.Duration)
	assert.Equal(DefaultInitialPodBackoff,                  c.InitialPodBackoff.Duration)
	assert.Equal(DefaultMaxPodBackoff,                      c.MaxPodBackoff.Duration)
	assert.Equal(DefaultHttpHandlerTimeout,                 c.HttpHandlerTimeout.Duration)
}

// Check that SetDefaults sets the default values
func TestConfig_SetDefaults(t *testing.T) {
	c := &Config{}
	c.SetDefaults()
	is_default(c, t)
}

// Check that CreateDefaultConfig returns a default config
func TestConfig_CreateDefaultConfig(t *testing.T) {
	c := CreateDefaultConfig()
	is_default(c, t)
}

// Check that a config string can be parsed
func TestConfig_Read(t *testing.T) {
	assert := assert.New(t)

	c := CreateDefaultConfig()
	reader := strings.NewReader(`
	[scheduler]
	offer-ttl=42s
	offer-linger-ttl=42s
	listener-delay=42s
	updates-backlog=42
	framework-id-refresh-interval=42s
	initial-implicit-reconciliation-delay=42s
	explicit-reconciliantion-max-backoff=42s
	explicit-reconciliantion-abort-timeout=42s
	initial-pod-backoff=42s
	max-pod-backoff=42s
	http-handler-timeout=42s
	`)
	err := c.Read(reader)
	if err != nil {
		t.Fatal("Cannot parse scheduler config: " + err.Error())
	}

	assert.Equal(42 * time.Second, c.OfferTTL.Duration)
	assert.Equal(42 * time.Second, c.OfferLingerTTL.Duration)
	assert.Equal(42 * time.Second, c.ListenerDelay.Duration)
	assert.Equal(42,               c.UpdatesBacklog)
	assert.Equal(42 * time.Second, c.FrameworkIdRefreshInterval.Duration)
	assert.Equal(42 * time.Second, c.InitialImplicitReconciliationDelay.Duration)
	assert.Equal(42 * time.Second, c.ExplicitReconciliationMaxBackoff.Duration)
	assert.Equal(42 * time.Second, c.ExplicitReconciliationAbortTimeout.Duration)
	assert.Equal(42 * time.Second, c.InitialPodBackoff.Duration)
	assert.Equal(42 * time.Second, c.MaxPodBackoff.Duration)
	assert.Equal(42 * time.Second, c.HttpHandlerTimeout.Duration)
}