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
	//var embedded []reflect.Value

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
		//if opts == emptyTags {
		//	if mapToField.Anonymous && mapToValue.Kind() == reflect.Struct {
		//		// save embedded struct for later processing
		//		embedded = append(embedded, mapToValue)
		//		continue
		//	}
		//
		//	name = sf.Name
		//}

		if opts.Contains("omitempty") && isEmptyValue(mapToValue) {
			continue
		}

		//if sv.Type().Implements(encoderType) {
		//	if !reflect.Indirect(sv).IsValid() {
		//		sv = reflect.New(sv.Type().Elem())
		//	}
		//
		//	m := sv.Interface().(Encoder)
		//	if err := m.EncodeValues(name, &values); err != nil {
		//		return err
		//	}
		//	continue
		//}

		//if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
		//	var del byte
		//	if opts.Contains("comma") {
		//		del = ','
		//	} else if opts.Contains("space") {
		//		del = ' '
		//	} else if opts.Contains("semicolon") {
		//		del = ';'
		//	} else if opts.Contains("brackets") {
		//		name = name + "[]"
		//	}
		//
		//	if del != 0 {
		//		s := new(bytes.Buffer)
		//		first := true
		//		for i := 0; i < sv.Len(); i++ {
		//			if first {
		//				first = false
		//			} else {
		//				s.WriteByte(del)
		//			}
		//			s.WriteString(valueString(sv.Index(i), opts))
		//		}
		//		values.Add(name, s.String())
		//	} else {
		//		for i := 0; i < sv.Len(); i++ {
		//			k := name
		//			if opts.Contains("numbered") {
		//				k = fmt.Sprintf("%s%d", name, i)
		//			}
		//			values.Add(k, valueString(sv.Index(i), opts))
		//		}
		//	}
		//	continue
		//}

		//if sv.Type() == timeType {
		//	values.Add(name, valueString(sv, opts))
		//	continue
		//}

		//for sv.Kind() == reflect.Ptr {
		//	if sv.IsNil() {
		//		break
		//	}
		//	sv = sv.Elem()
		//}

		//if sv.Kind() == reflect.Struct {
		//	mapToStruct(values, sv)
		//	continue
		//}

		//fmt.Println(fmt.Sprintf("%v=", sv.String()))

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

	//for _, f := range embedded {
	//	if err := reflectValue(values, f, scope); err != nil {
	//		return err
	//	}
	//}

	return nil
}
