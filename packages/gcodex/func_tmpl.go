package gcodex

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type dbT struct {
	PackageName string
	Columns     []string
	Indexes     []string
	TypeName    string
}

func genEntSchema() error {
	// New("名称")要和模版文件名，define 名称保持一致，如果多个模版和第一个模版文件名一致
	t, err := template.New("ent.tmpl").ParseFiles("./tmpl/model/ent.tmpl")
	if err != nil {
		log.Printf("parse tmpl error:%v\n", err)
		return err
	}
	// 	ss := `
	// 	{{define "ent"}}
	// Package {{.PackageName}}
	// name={{.TableName}}
	// {{end}}`
	// 	t, err := template.New("ent").Parse(ss)
	// 	if err != nil {
	// 		log.Printf("parse tmpl error:%v\n", err)
	// 		return err
	// 	}
	vv := dbT{"db", []string{"phoenix_house", "t2"}, []string{"222", "1"}, "User"}
	fname := strings.Trim(vv.TypeName, "") + ".go"

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create target file fail error:%v\n", err)
		return err
	}
	defer f.Close()

	ff := os.NewFile(f.Fd(), fname)
	err = t.Execute(ff, vv)
	if err != nil {
		log.Printf("tmpl exec fail,error:%v", err)
		return err
	}
	cmd := exec.Command("go", "fmt", fname)
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}
