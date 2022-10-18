package gcodex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenFuncCode(t *testing.T) {
	LoadIDLConfig("./tmpl/model/", "phoenix_house")
	opt := &DBCfgdata{
		TmplName:    "ent.tmpl",
		FilePath:    "./tmpl/func/",
		TarFilePath: "./dist/ent/",
		Tables:      IDL.DB[0].Tables,
	}
	err := GenFuncCode(funcTmplList, opt)
	assert.Nil(t, err, "genSchemaCode() success.")
}
