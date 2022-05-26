package gcodex

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func InitEntSchema(tmplName, filePath, tarFilePath string,idl *IDLConfig) (*DBCfgdata, error) {
	// idl, err := LoadIDLConfig("./tmpl/model/", "phoenix_house")
	// if err != nil {
	// 	return nil, err
	// }
	var sch DBCfgdata
	sch.filePath = filePath
	sch.tmplName = tmplName
	sch.tarFilePath = tarFilePath
	if len(idl.DB) <= 0 {
		return nil, ErrInvalidIDLMetadata
	}
	
	return &sch, nil
}

func GenSchemaCode(opt *DBCfgdata) error {
	if opt == nil {
		return ErrInvalidSchemaOption
	}

	// New("名称")要和模版文件名，define 名称保持一致，如果多个模版和第一个模版文件名一致
	t, err := template.New(opt.tmplName).Funcs(template.FuncMap{"ToCamel": ToCamel, "MaxRuneCount": MaxRuneCount, "MinRuneCount": MinRuneCount}).ParseFiles(opt.filePath + opt.tmplName)
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
		tmpTarget := opt.tarFilePath + "schema/"
		if opt.Tables[i].PackageName == "" {
			opt.Tables[i].PackageName = "schema"
		}

		fname := strings.Trim(opt.Tables[i].Name, "") + ".go"
		err := os.MkdirAll(tmpTarget, 0777)
		if err != nil && !os.IsExist(err) {
			log.Printf("make target file dir fail error:%v\n", err)
			return err
		}
		// f, err := os.OpenFile(tmpTarget+fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		f, err := os.OpenFile(tmpTarget+fname, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("create target file fail error:%v\n", err)
			return err
		}
		defer f.Close()

		ff := os.NewFile(f.Fd(), fname)
		err = t.Execute(ff, opt.Tables[i])
		if err != nil {
			log.Printf("tmpl exec fail,error:%v", err)
			return err
		}
		// 格式化
		runGoCMD("go", "fmt", tmpTarget+fname)
	}

	genT, err := template.New("generate.tmpl").ParseFiles(opt.filePath + "generate.tmpl")
	if err != nil {
		log.Printf("parse tmpl error:%v\n", err)
		return err
	}
	gf, err := os.OpenFile(opt.tarFilePath+"generate.go", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create target file fail error:%v\n", err)
		return err
	}
	defer gf.Close()

	genf := os.NewFile(gf.Fd(), "generate.go")
	err = genT.Execute(genf, nil)
	if err != nil {
		log.Printf("tmpl exec fail,error:%v", err)
		return err
	}
	runGoCMD("go", "fmt", opt.tarFilePath+"generate.go")

	return nil
}

func runGoCMD(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
