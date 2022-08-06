package gcodex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIDLConfig(t *testing.T) {
	LoadIDLConfig("./tmpl/model/", "phoenix_house")
	assert.NotNil(t, ViperCfg, "GetIDLConfig() successed.")
	assert.Empty(t, ViperCfg.GetString("db"), "GetIDLConfig() successed.")
}
