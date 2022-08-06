// Output:./dist/ent/schema/*.go
// TODO: Extension API support HTTP Restfull API/grpc logging/metric/tracing
// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.
package gcodex

import (
	"log"
	"strings"
	"text/template"
)

//go generate
// TODO:根据 AST 生成 service 层源码
func GenFuncCode(tmplNames []string, idl *DBCfgdata) error {
	if len(idl.Tables) <= 0 {
		return ErrDBHasEmptyTable
	}
	for _, table := range idl.Tables {
		for _, tt := range tmplNames {
			if table.PackageName == "" {
				table.PackageName = table.Name
			}
			t, err := template.New(tt).Funcs(template.FuncMap{"ToCamel": ToCamel, "MaxRuneCount": MaxRuneCount, "MinRuneCount": MinRuneCount}).ParseFiles(idl.FilePath + "func/" + tt)
			if err != nil {
				log.Printf("parse tmpl error:%v\n", err)
				return err
			}
			fErr := createSrcFile(t, table, strings.Split(tt, ".")[0]+".go", idl.TarFilePath+"internal/"+table.Name+"/")
			if fErr != nil {
				return fErr
			}
		}
	}

	return nil
}
