package gcodex

import (
	"log"
	"strings"
	"text/template"
)

func GenEntSchemaCode(opt *DBCfgdata) error {
	if opt == nil {
		return ErrInvalidSchemaOption
	}
	// New("名称")要和模版文件名，define 名称保持一致，如果多个模版和第一个模版文件名一致
	t, err := template.New(opt.TmplName).Funcs(template.FuncMap{"ToCamel": ToCamel, "MaxRuneCount": MaxRuneCount, "MinRuneCount": MinRuneCount}).ParseFiles(opt.FilePath + "model/" + opt.TmplName)
	if err != nil {
		log.Printf("parse tmpl error:%v\n", err)
		return err
	}
	//生成
	if len(opt.Tables) <= 0 {
		return ErrInvalidEntitiesMetadata
	}

	// TODO: 并行生成 entity schema
	for i := 0; i < len(opt.Tables); i++ {
		tmpTarget := opt.TarFilePath + "internal/ent/schema/"
		if opt.Tables[i].PackageName == "" {
			opt.Tables[i].PackageName = "schema"
		}

		fname := strings.Trim(opt.Tables[i].Name, "") + ".go"
		fErr := createSrcFile(t, opt.Tables[i], fname, tmpTarget)
		if fErr != nil {
			return fErr
		}
	}

	genT, err := template.New("generate.tmpl").ParseFiles(opt.FilePath + "model/" + "generate.tmpl")
	if err != nil {
		log.Printf("parse tmpl error:%v\n", err)
		return err
	}
	fErr := createSrcFile(genT, nil, "generate.go", opt.TarFilePath+"internal/ent/")
	if fErr != nil {
		return fErr
	}

	return nil
}
