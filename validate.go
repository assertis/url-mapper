package mapper

import (
	"fmt"
	"reflect"
)

func isValid(v reflect.Value, opts TagOptions) bool {
	fmt.Println(v.Kind())
	fmt.Println(opts)
	fmt.Println(opts.Contains("regexp"))

	if v.Kind() == reflect.String && opts.Contains("regexp") {
		fmt.Println(v.String())
		fmt.Println(opts)
	}

	return false
}
