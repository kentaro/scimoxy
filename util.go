package scimoxy

import (
	"reflect"
)

func typeOf(s interface{}) string {
	return reflect.TypeOf(s).String()
}
