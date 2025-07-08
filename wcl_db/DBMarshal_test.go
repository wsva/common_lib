package wcl_db_test

import (
	"fmt"
	"testing"

	"github.com/wsva/common_lib/wcl_db"
)

func TestMarshal(T *testing.T) {
	db := wcl_db.DB{
		Type: wcl_db.DBTypeOracle,
		Oracle: wcl_db.Oracle{
			Username: "",
		},
	}
	_, jsonString, err := wcl_db.MarshalDB(db, true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(jsonString)
	}
}
