package wcl_crypto_test

import (
	"fmt"
	"testing"

	"github.com/wsva/common_lib/wcl_crypto"
)

func TestUUID(T *testing.T) {
	fmt.Println(wcl_crypto.NewDateUUID())
	fmt.Println(wcl_crypto.NewDateUUID())
	fmt.Println(wcl_crypto.NewDateUUID())
	fmt.Println(wcl_crypto.NewDateUUID())
	fmt.Println(wcl_crypto.NewDateUUID())
	fmt.Println(wcl_crypto.NewUUID())
	fmt.Println(wcl_crypto.NewUUID())
	fmt.Println(wcl_crypto.NewUUID())
	fmt.Println(wcl_crypto.NewUUID())
}
