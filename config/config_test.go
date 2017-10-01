package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConfigFile(t *testing.T) {
	assert.Equal(t,true ,fileExists(), "The configuration file can not be found.")
}