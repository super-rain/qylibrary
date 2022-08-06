package gcodex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenBackendLayout(t *testing.T) {
	LoadIDLConfig("./tmpl/model/", "phoenix_house")
	IDL.DB[0].TmplName = "ent.tmpl"
	IDL.DB[0].FilePath = tmplFilePath
	IDL.DB[0].TarFilePath = "examples/backend/phoenix_house/"
	err := GenBackendLayout()
	assert.Nil(t, err, "GenBackendLayout() success!")
}
