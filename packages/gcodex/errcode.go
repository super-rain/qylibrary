package gcodex

import (
	"errors"
)

var (
	ErrInvalidParams           = errors.New("invalid params ")
	ErrInvalidSchemaOption     = errors.New("invalid schema option ")
	ErrIDLFileNotExists        = errors.New("IDL file not exists ")
	ErrInitIDLMetadata         = errors.New("loading IDL tmpl fail ")
	ErrInvalidIDLMetadata      = errors.New("invalid IDL metadata ")
	ErrInvalidEntitiesMetadata = errors.New("invalid entities metadata ")
	ErrDBHasEmptyTable         = errors.New("DB has empty table ")
)

// // 0～100 common code,>=1000 biz code
// type ErrCode int

// //go:generate stringer -type ErrCode -linecomment -output code_string.go
// const (
// 	StatusTimeout          ErrCode = -1 // 操作超时
// 	StatusFail             ErrCode = 0  // 操作失败
// 	StatusSuccess          ErrCode = 1  // 操作成功
// 	StatusUnknown          ErrCode = 2  // 未知错误
// 	StatusBadRoute         ErrCode = 3  // 无效或错误路由
// 	StatusInvalidParams    ErrCode = 4  // 无效参数
// 	StatusInvalidToken     ErrCode = 5  // 无效Token
// 	StatusIllegalSignature ErrCode = 6  // 非法签名
// 	StatusNoRecordFound    ErrCode = 7  // 未查询到受影响记录
// 	// biz ErrCode >=1000
// 	ErrInvalidParams ErrCode = 1000 // 业务无效参数
// 	// ErrXXXX ErrCode = 1001 // Xxxxx
// )

// func (e ErrCode) String() string {
// 	return GetDescription(e)
// }
