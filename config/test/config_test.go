package test

import (
	"testing"

	"DataWall/config"

	//"github.com/stretchr/testify/assert"
)

func TestGetIpAddressNotNil(t *testing.T) {
	//_assert := assert.New(t)
	conf := config.Get()
	if conf != nil {
		//_assert.True(t, true)
	} else {
		//assert.True(conf.IpAddress, false)
	}
}
