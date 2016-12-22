// Package mapping query string to struct using reflection.
// Mapper supports optional validation of input values through tags.
//
// Example:
// 	type Request {
// 		Origin string `query:"origin,regexp=^[A-Z]{3}$"`
//		Destination string `query:"destination,regexp=^[A-Z]{3}$"`
//		Adults int `query:"adults,default=1,max=9"`
// 		Children int `query:"children,optional,default=0,max=9"`
//		Outward time.Time `query:"outward,dateFormat=RFC_3339"`
//		Return *time.Time `query:"return,optional,dateFormat=RFC_3339"`
//	}
//

package mapper

import (
	"errors"
	"reflect"
)

var sliceOfBytes = reflect.TypeOf([]byte(nil))

var errWrongType = errors.New("Unmarshal can work with string or slice of bytes")

func Unmarshal(v interface{}) error {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.String && val.Type() != sliceOfBytes {
		return errWrongType
	}

	return nil
}
