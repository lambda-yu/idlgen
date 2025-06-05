package anchor

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
	"os"
	"testing"
)

func Test_parse(t *testing.T) {
	open, err := os.Open("../swap.json")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	content, err := io.ReadAll(open)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	var idl Idl

	err = json.Unmarshal(content, &idl)
	if err != nil {
		fmt.Println("json错误:", err)
		return
	}
	spew.Dump(idl)
}
