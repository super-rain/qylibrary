package arrow

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/apache/arrow/go/arrow"
// 	"github.com/apache/arrow/go/arrow/array"
// 	"github.com/apache/arrow/go/arrow/memory"
// )

// func main(){

// 	pool :=memory.NewGoAllocator()

// 	dtype:=arrow.StructOf([]arrow.Field{
// 		{Name:"f1",Type:arrow.ListOf(arrow.PrimitiveTypes.Uint8)},
// 		{Name:"f2",Type:arrow.PrimitiveTypes.Int32},
// 	}...)
// 	sb:=array.NewStructBuilder(pool,dtype)
// 	defer sb.Release()

// 	f1b:=sb.FiledBuilder(0).(*array.ListBuilder)
// 	f1vb:=f1b.ValueBuilder().(*array.Uint8Builder)
// 	f2b:=sb.FiledBuilder(1).(*array.Int32Builder)

// 	sb.Reserve(4)
// 	f1vb.Reserve(7)
// 	f2b.Reserve(3)

// 	sb.Append(true)
// 	f1b.Append(true)
// 	f1vb.AppendValues([]byte("joe"),nil)
// 	f2b.Append(1)

// 	for i:=0;i<arr.Len();i++{
// 		if !arr.isValid(i){
// 			fmt.Printf("(),i",i)
// 			continue
// 		}
// 		fmt.Printf("sfd",i)
// 		pos:=int(offsets[i])
// 		switch{
// 		case list.IsValid(pos):
// 			fmt.Printf("[")
// 			for j:=offsets[i];i<offsets[i+1];j++{
// 				if j!=offsets[i]{
// 					fmt.Printf(",")
// 				}
// 				fmt.Printf("fds",string(varr.Value(int(j))))
// 			}
// 			fmt.Printf("]")
// 		default:
// 			fmt.Printf("(null),")
// 		}
// 		fmt.Printf("fdsa",ints.Value(i))
// 	}
// }
