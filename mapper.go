// Package mapping query string to struct using reflection.
// Mapper supports optional validation of input values through tags.
//
// Example:
// 	type Request struct {
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
	"net/url"
	"reflect"
)

var (
	errWrongType        = errors.New("Unmarshal only works with pointers")
	errValidationFailed = errors.New("Field invalid")
)

func Unmarshal(path string, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errWrongType
	}

	values, err := url.ParseQuery(path)
	if err != nil {
		return err
	}

	return mapToStruct(values, val.Elem())
}

func mapToStruct(values url.Values, v reflect.Value) error {
	mapToType := v.Type() // must be struct
	for i := 0; i < mapToType.NumField(); i++ {
		mapToField := mapToType.Field(i)
		if mapToField.PkgPath != "" && !mapToField.Anonymous {
			// unexported
			continue
		}

		mapToValue := v.Field(i)

		tag := mapToField.Tag.Get("query")
		if tag == "-" {
			continue
		}

		name, opts := ParseTagsIntoMap(tag)

		if opts.Contains("omitempty") && isEmptyValue(mapToValue) {
			continue
		}

		//if mapToValue.Type() == timeType {
		//	values.Add(name, valueString(values.Get(name), opts))
		//	continue
		//}

		for mapToValue.Kind() == reflect.Ptr {
			if mapToValue.IsNil() {
				break
			}
			mapToValue = mapToValue.Elem()
		}

		if mapToValue.Kind() == reflect.Struct {
			mapToStruct(values, mapToValue)
			continue
		}

		// TODO: validation
		//if !isValid(sv, opts) {
		//	return errValidationFailed
		//}

		if mapToValue.IsValid() && mapToValue.CanSet() {
			if mapToValue.Kind() == reflect.String {
				mapToValue.SetString(values.Get(name))
			}
		}
	}

	return nil
}
