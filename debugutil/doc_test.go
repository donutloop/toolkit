package debugutil_test

import (
	"fmt"
	"github.com/donutloop/toolkit/debugutil"
)

func ExamplePrettySprint() {

	str := debugutil.PrettySprint([]string{})
	fmt.Println(str)
	// Output: []string{
	//}
}
