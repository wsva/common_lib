package wcl_compress_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wsva/common_lib/wcl_compress"
)

func TestGzip(T *testing.T) {
	fmt.Println("TestGzip")
	text := strings.Repeat("12345678901234567890", 300)
	ctext, err := wcl_compress.GzipCompress(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ctext)
	dtext, err := wcl_compress.GzipDecompress(ctext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dtext)
}
