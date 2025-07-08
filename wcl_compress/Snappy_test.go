package wcl_compress_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wsva/common_lib/wcl_compress"
)

func TestSnappy(T *testing.T) {
	fmt.Println("TestSnappy")
	text := strings.Repeat("12345678901234567890", 300)
	ctext := wcl_compress.SnappyCompress(text)
	fmt.Println(ctext)
	dtext, _ := wcl_compress.SnappyDecompress(ctext)
	fmt.Println(dtext)
}
