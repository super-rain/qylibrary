package gcodex

import (
	"log"
	"os"
)

func GenBackendLayout() error {
	db := IDL.DB[0]
	// generate project layout
	for _, dir := range backendSpec {
		if dir == "" {
			break
		}
		dirN := "examples/backend/" + db.Name + "/" + dir
		err := os.MkdirAll(dirN, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			log.Printf("make target file dir fail error:%v\n", err)
			return err
		}
	}
	// generate model with entgo
	genErr := GenEntSchemaCode(db)
	if genErr != nil {
		return genErr
	}
	// generate services with gokit
	genErr = GenFuncCode(funcTmplList, db)
	return genErr
}

// TODO: Frontend APP Layout
// func GenFrontendLayout() error {
// 	db := IDL.DB[0]
// 	// generate project layout
// 	for _, dir := range backendSpec {
// 		if dir == "" {
// 			break
// 		}
// 		dirN := "examples/backend/" + db.Name + "/" + dir
// 		err := os.MkdirAll(dirN, os.ModePerm)
// 		if err != nil && !os.IsExist(err) {
// 			log.Printf("make target file dir fail error:%v\n", err)
// 			return err
// 		}
// 	}
// 	// generate model with entgo
// 	genErr := GenEntSchemaCode(db)
// 	if genErr != nil {
// 		return genErr
// 	}
// 	// generate services with gokit
// 	genErr = GenFuncCode(funcTmplList, db)
// 	return genErr
// }
