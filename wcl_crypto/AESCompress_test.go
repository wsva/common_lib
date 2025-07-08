package wcl_crypto_test

import (
	"fmt"
	"testing"

	"github.com/wsva/common_lib/wcl_compress"
	"github.com/wsva/common_lib/wcl_crypto"
)

var text = `
11111111111111111111111111111111111111111111111111
22222222222222222222222222222222222222222222222222
33333333333333333333333333333333333333333333333333
44444444444444444444444444444444444444444444444444
55555555555555555555555555555555555555555555555555
66666666666666666666666666666666666666666666666666
77777777777777777777777777777777777777777777777777
88888888888888888888888888888888888888888888888888
99999999999999999999999999999999999999999999999999
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
cccccccccccccccccccccccccccccccccccccccccccccccccc
`

func TestAESCompress(T *testing.T) {
	fmt.Println("source", len(text))
	ctext := wcl_compress.SnappyCompress(text)
	fmt.Println("snappy", len(ctext))
	ctext, err := wcl_compress.GzipCompress(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("gzip", len(ctext))
	ctext, err = wcl_crypto.GA128Encode("1", "2", text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(ctext))
	fmt.Println(ctext)
}
