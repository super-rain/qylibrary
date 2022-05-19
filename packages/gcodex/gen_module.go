package gcodex

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

type SchemaOption struct {
	tmplName    string // 带后缀
	idlFilePath string // 相对路径或绝对路径
	tables      []*TableMetaData
	tarFilePath string // 相对路径或绝对路径
}

type TableMetaData struct {
	PackageName string
	Name        string
	Comment     string
	Columns     []*ColumnType // []columns
	Edges       []*ModelRelationShip
}

type ModelRelationShip struct {
	From string
	To   string
	Ref  string
}
type ColumnType struct {
	Name          string
	StructTag     string
	StorageKey    string
	Comment       string
	FieldType     DBFieldType
	DefaultVal    any
	MaxLen        int
	MinLen        int
	IsNonNegative bool
	// scale: 2, precision: 10, default: '0.00'
	Scale        int
	Precision    int
	AutoInc      bool
	IsSensitive  bool
	IsImmutable  bool
	IsOptional   bool
	Enumerable   bool
	EnumVals     any
	IsPrimary    bool
	PkName       string
	IsForeignKey bool
	FkName       string
	IsUnique     bool
	UidxName     string
	IsIndex      bool
	IdxName      string
}

type DBFieldType string

const (
	TypeString    DBFieldType = "string"
	TypeInt       DBFieldType = "int"
	TypeUInt      DBFieldType = "uint"
	TypeFloat     DBFieldType = "float"
	TypeFloat64   DBFieldType = "float64"
	TypeDecimal   DBFieldType = "decimal"
	TypeBool      DBFieldType = "bool"
	TypeDatetime  DBFieldType = "datetime"
	TypeTimestamp DBFieldType = "timestamp"
)

var DefaultValue map[DBFieldType]any = map[DBFieldType]any{
	TypeString:    "",
	TypeInt:       0,
	TypeUInt:      0,
	TypeFloat:     0.00,
	TypeFloat64:   0.00,
	TypeDecimal:   0.00,
	TypeBool:      false,
	TypeDatetime:  time.Now(),
	TypeTimestamp: time.Now().Unix(),
}

func genSchemaCode(opt *SchemaOption) error {
	if opt == nil {
		return ErrInvalidSchemaOption
	}

	// New("名称")要和模版文件名，define 名称保持一致，如果多个模版和第一个模版文件名一致
	t, err := template.New(opt.tmplName).ParseFiles(opt.idlFilePath + opt.tmplName)
	if err != nil {
		log.Printf("parse tmpl error:%v\n", err)
		return err
	}

	//生成
	if len(opt.tables) <= 0 {
		return ErrInvalidEntitiesMetadata
	}
	// TODO: 并行生成 entity schema
	for i := 0; i < len(opt.tables); i++ {
		if opt.tables[i].PackageName == "" {
			opt.tables[i].PackageName = "schema"
		}

		fname := strings.Trim(opt.tables[i].Name, "") + ".go"
		err := os.MkdirAll(opt.tarFilePath, 0777)
		if err != nil && !os.IsExist(err) {
			log.Printf("make target file dir fail error:%v\n", err)
			return err
		}
		// f, err := os.OpenFile(opt.tarFilePath+fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		f, err := os.OpenFile(opt.tarFilePath+fname, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("create target file fail error:%v\n", err)
			return err
		}
		defer f.Close()

		ff := os.NewFile(f.Fd(), fname)
		err = t.Execute(ff, opt.tables[i])
		if err != nil {
			log.Printf("tmpl exec fail,error:%v", err)
			return err
		}
		// 格式化
		runGoCMD("go", "fmt", opt.tarFilePath+fname)
	}

	return nil
}

func runGoCMD(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// import(

// )

// type XXer interface{

// }

// type XX struct{

// }

// func (x XX) NewXX()*XX{
// 	return &XX{

// 	}
// }

// func (x XX) Add()error{

// }
