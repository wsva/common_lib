package wcl_crypto

import (
	"fmt"
	mathrand "math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func NewUUID() string {
	return fmt.Sprint(uuid.NewV4())
}

func NewDateUUID() string {
	t := time.Now()
	timestr := t.Format("20060102150405")
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	randstr := fmt.Sprintf("%v%v%v",
		strings.Repeat("0", 32), r.Int63n(1000000), r.Int63n(1000000))
	return fmt.Sprintf("%v%v",
		timestr, randstr[len(randstr)-(32-len(timestr)):])
}
