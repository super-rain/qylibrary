package gcodex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	err := Run("backend", "top_starter", "tmpl/model/", "examples/backend/")
	assert.Nil(t, err, "gcodex.Run() success")
}
