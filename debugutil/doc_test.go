package debugutil

import (
	"fmt"

	"github.com/donutloop/gdp/util/debugutil"
)

func ExamplePrettySprint() {

	str := debugutil.PrettySprint([]string{})
	fmt.Println(str)
	// Output: []string{
	//}
}
