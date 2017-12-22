// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package debugutil

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

const (
	bracketOpen  string = "{\n"
	bracketClose string = "}"
	pointerSign  string = "&"
	nilSign      string = "<nil>"
)

// PrettySprint generates a human readable representation of the value v.
func PrettySprint(v interface{}) string {

	value := reflect.ValueOf(v)

	// nil
	switch value.Kind() {
	case reflect.Interface, reflect.Map, reflect.Slice, reflect.Ptr:
		if value.IsNil() {
			return nilSign
		}
	}

	buff := bytes.Buffer{}
	switch value.Kind() {
	case reflect.Struct:
		buff.WriteString(fullName(value.Type()) + bracketOpen)
		for i := 0; i < value.NumField(); i++ {
			l := string(value.Type().Field(i).Name[0])
			if strings.ToUpper(l) == l {
				buff.WriteString(fmt.Sprintf("%s: %s,\n", value.Type().Field(i).Name, PrettySprint(value.Field(i).Interface())))
			}
		}
		buff.WriteString(bracketClose)
		return buff.String()
	case reflect.Map:
		buff.WriteString("map[" + fullName(value.Type().Key()) + "]" + fullName(value.Type().Elem()) + bracketOpen)
		for _, k := range value.MapKeys() {
			buff.WriteString(fmt.Sprintf(`"%s":%s,\n`, k.String(), PrettySprint(value.MapIndex(k).Interface())))
		}
		buff.WriteString(bracketClose)
		return buff.String()
	case reflect.Ptr:
		if e := value.Elem(); e.IsValid() {
			return fmt.Sprintf("%s%s", pointerSign, PrettySprint(e.Interface()))
		}
		return nilSign
	case reflect.Slice:
		buff.WriteString("[]" + fullName(value.Type().Elem()) + bracketOpen)
		for i := 0; i < value.Len(); i++ {
			buff.WriteString(fmt.Sprintf("%s,\n", PrettySprint(value.Index(i).Interface())))
		}
		buff.WriteString(bracketClose)
		return buff.String()
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
