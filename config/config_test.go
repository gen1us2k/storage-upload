package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
)

func TestParseConfig(t *testing.T) {
	os.Setenv("BIND_ADDR", ":8080")
	c, err := Parse()
	require.NoError(t, err)
	assert.Equal(t, ":8080", c.BindAddr)
	os.Setenv("BIND_ADDR", ":8081")
	c, err = Parse()
	require.NoError(t, err)
	assert.Equal(t, ":8081", c.BindAddr)

}
