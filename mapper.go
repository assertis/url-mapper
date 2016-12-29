// Package mapping query string to struct using reflection.
// Mapper supports optional validation of input values through tags.
//
// Example:
// 	type Request struct {
// 		Origin string `query:"origin"`
//		Destination string `query:"destination"`
//		NumOfPassengers int `query:"adults"`
//		OutwardDate time.Time `query:"outward"`
//		ReturnDate time.Time `query:"inward,omitempty"`
//	}
//

package mapper

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"time"
	"github.com/yezooz/govalidator"
)

var (
	errWrongType = errors.New("Unmarshal only works with pointers")
)

func Unmarshal(path url.Values, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errWrongType
	}

	return mapToStruct(path, val.Elem())
}

func mapToStruct(values url.Values, v reflect.Value) error {
	mapToType := v.Type() // must be struct
	for i := 0; i < mapToType.NumField(); i++ {
		mapToField := mapToType.Field(i)

		// Ignore unexported fields
		if mapToField.PkgPath != "" && !mapToField.Anonymous {
			continue
		}

		mapToValue := v.Field(i)

		tag := mapToField.Tag.Get("query")
		if tag == "-" {
			continue
		}

		name, opts := TagOptionsFromString(tag)

		if opts.Contains("omitempty") && isEmptyValue(mapToValue) {
			continue
		}

		for mapToValue.Kind() == reflect.Ptr {
			if mapToValue.IsNil() {
				break
			}
			mapToValue = mapToValue.Elem()
		}

		if mapToValue.IsValid() && mapToValue.CanSet() {
			value := values.Get(name)

			// Time?
			if mapToValue.Type() == reflect.TypeOf(time.Time{}) {
				if values.Get(name) == "" {
					continue
				}

				if opts.Contains("rfc3339") && govalidator.IsRFC3339(value) {
					t, err := time.Parse(time.RFC3339, value)
					if err != nil {
						return err
					}
					mapToValue.Set(reflect.ValueOf(t))
				} else if opts.Contains("unix") && govalidator.IsInt(value) {
					i, _ := strconv.Atoi(value)

					t := time.Unix(int64(i), 0)
					mapToValue.Set(reflect.ValueOf(t))
				} else {
					return errors.New("Incorrect time value")
				}


				continue
			}

			switch mapToValue.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				i, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				mapToValue.SetInt(int64(i))

			case reflect.String:
				mapToValue.SetString(value)
			}

		}
	}

	return nil
}
