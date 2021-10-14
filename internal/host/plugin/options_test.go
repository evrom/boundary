package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetOpts(t *testing.T) {
	t.Parallel()
	t.Run("WithName", func(t *testing.T) {
		opts := getOpts(WithName("test"))
		testOpts := getDefaultOptions()
		testOpts.withName = "test"
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithDescription", func(t *testing.T) {
		opts := getOpts(WithDescription("test desc"))
		testOpts := getDefaultOptions()
		testOpts.withDescription = "test desc"
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithPreferredEndpoints", func(t *testing.T) {
		opts := getOpts(WithPreferredEndpoints([]string{"foo"}))
		testOpts := getDefaultOptions()
		testOpts.withPreferredEndpoints = []string{"foo"}
		assert.EqualValues(t, opts, testOpts)
	})
	t.Run("withDnsNames", func(t *testing.T) {
		opts := getOpts(withDnsNames([]string{"foo"}))
		testOpts := getDefaultOptions()
		testOpts.withDnsNames = []string{"foo"}
		assert.EqualValues(t, opts, testOpts)
	})
	t.Run("withIpAddresses", func(t *testing.T) {
		opts := getOpts(withIpAddresses([]string{"foo"}))
		testOpts := getDefaultOptions()
		testOpts.withIpAddresses = []string{"foo"}
		assert.EqualValues(t, opts, testOpts)
	})
}
