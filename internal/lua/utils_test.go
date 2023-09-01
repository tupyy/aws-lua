package lua

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestToLTable(t *testing.T) {
	RegisterTestingT(t)

	o := Object{
		"Vpc": map[string]interface{}{
			"Id": "id",
		},
	}

	res := toLTable(o)
	Expect(res).ToNot(BeNil())
	fmt.Printf("%+v\n", *res)
}

// func TestFromLTable(t *testing.T) {
// 	RegisterTestingT(t)

// 	lv := &lua.LTable{
// 		"tag"
// }
