package rust_func

/*
#cgo LDFLAGS: -L../lib -lwcloud
#include "../lib/libwcloud.h"
*/
import "C"
import (
	"fmt"
	"github.com/jizizr/LanMei/server/service/wordcloud/biz/model"
	"unsafe"
)

func _Cadd() {
	result := C.add(5, 3)
	fmt.Println("Result from Rust:", result)
}

func Wcloud(words []model.Word) string {
	cWords := make([]C.Word, len(words))
	for i, w := range words {
		cWords[i] = C.Word{
			word: C.CString(w.Word),
			freq: C.size_t(w.Count),
		}
	}
	defer func() {
		for _, w := range cWords {
			C.free(unsafe.Pointer(w.word))
		}
	}()
	// 调用Rust函数
	result := C.wcloud(&cWords[0], C.int(len(cWords)))
	defer C.free_string(result)
	// 将结果转换为Go字符串
	return C.GoString(result)
}
