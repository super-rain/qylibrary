package gcodex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_genEntSchema(t *testing.T) {
	err := genEntSchema()
	assert.Nil(t, err, "genEntSchema() success")
}
