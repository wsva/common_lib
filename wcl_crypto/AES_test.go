package wcl_crypto_test

import (
	"fmt"
	"testing"

	"github.com/wsva/common_lib/wcl_crypto"
)

func TestAES128(T *testing.T) {
	ctext, err := wcl_crypto.AES128Encrypt("1", "2", "3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ctext)
}
