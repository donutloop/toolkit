// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package debugutil

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	bracketOpen  string = "{\n"
	bracketClose string = "}"
	pointerSign  string = "&"
	nilSign      string = "nil"
)

// PrettyPrint generates a human readable representation of the value v.
func PrettySprint(v interface{}) string {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Struct:
		str := fullName(value.Type()) + bracketOpen
		for i := 0; i < value.NumField(); i++ {
			l := string(value.Type().Field(i).Name[0])
			if strings.ToUpper(l) == l {
				str += fmt.Sprintf("%s: %s,\n", value.Type().Field(i).Name, PrettySprint(value.Field(i).Interface()))
			}
		}
		str += bracketClose
		return str
	case reflect.Map:
		str := "map[" + fullName(value.Type().Key()) + "]" + fullName(value.Type().Elem()) + bracketOpen
		for _, k := range value.MapKeys() {
			str += fmt.Sprintf(`"%s":%s,\n`, k.String(), PrettySprint(value.MapIndex(k).Interface()))
		}
		str += bracketClose
		return str
	case reflect.Ptr:
		if e := value.Elem(); e.IsValid() {
			return fmt.Sprintf("%s%s", pointerSign, PrettySprint(e.Interface()))
		}
		return nilSign
	case reflect.Slice:
		str := "[]" + fullName(value.Type().Elem()) + bracketOpen
		for i := 0; i < value.Len(); i++ {
			str += fmt.Sprintf("%s,\n", PrettySprint(value.Index(i).Interface()))
		}
		str += bracketClose
		return str
	default:
		return fmt.Sprintf("%#v", v)
	}
}

func pkgName(t reflect.Type) string {
	pkg := t.PkgPath()
	c := strings.Split(pkg, "/")
	return c[len(c)-1]
}

func fullName(t reflect.Type) string {
	if pkg := pkgName(t); pkg != "" {
		return pkg + "." + t.Name()
	}
	return t.Name()
}
