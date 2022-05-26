package gcodex

import "time"

// type GenKind uint8
// const(
// 	GenKindSignal GenKind =itoa
// 	GenKindMult GenKind
// )

var uppercaseAcronym = map[string]string{
	"ID": "id",
}

// idl type ,toml结构须和 struct 层级一致
// struct 映射字段须导出
// 表名须一致，大小写敏感，字段大小写不敏感不支持特殊字符
// 建议表名和字段保持一致
type IDLConfig struct {
	DB []*DBCfgdata `json:"db_metadata"`
}

type DBCfgdata struct {
	tmplName    string // 带后缀
	filePath    string // 相对路径或绝对路径
	tarFilePath string // 相对路径或绝对路径
	Name    string
	Comment string
	Tables  []*TableMetaData
}

// type TableCfgData struct {
// 	Name      string     `json:"name"`
// 	Comment   string     `json:"comment"`
// 	ColumnSet [][]string `json:"columns"`
// }

// type TableCfgData struct {
// 	tmplName    string // 带后缀
// 	filePath    string // 相对路径或绝对路径
// 	tarFilePath string // 相对路径或绝对路径
// 	Name        string
// 	Comment     string
// 	Tables      []*TableMetaData
// }

type TableMetaData struct {
	PackageName string
	Name        string
	Comment     string
	Columns     []*ColumnType 
	Edges       []*ModelRelationShip
}

type ModelRelationShip struct {
	From     string
	FromName string
	To       string
	ToName   string
	RefName  string
}
type ColumnType struct {
	Name          string
	StructTag     string
	StorageKey    string
	Comment       string
	FieldType     string
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
