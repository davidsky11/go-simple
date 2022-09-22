package IOs

import (
	"fmt"
	"io"
	"strings"
)

func Reader() {
	r := strings.NewReader("Hello, Reader! Let's Go!")
	b := make([]byte, 8)

	// Read 用数据填充给定的字节切片并返回填充的字节数和错误值。
	// 在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
