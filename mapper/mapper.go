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
	sliceOfBytes = reflect.TypeOf([]byte(nil))
	errWrongType = errors.New("Unmarshal only works with pointers")
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
	var embedded []reflect.Value

	typ := v.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous {
			// unexported
			//fmt.Println("unexported")
			continue
		}

		sv := v.Field(i)

		//fmt.Println(sv)

		tag := sf.Tag.Get("query")
		if tag == "-" {
			continue
		}

		name, opts := parseTag(tag)
		if name == "" {
			if sf.Anonymous && sv.Kind() == reflect.Struct {
				// save embedded struct for later processing
				embedded = append(embedded, sv)
				continue
			}

			name = sf.Name
		}

		//fmt.Println(name)
		//fmt.Println(opts)

		if opts.Contains("omitempty") && isEmptyValue(sv) {
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

		if sv.IsValid() && sv.CanSet() {
			if sv.Kind() == reflect.String {
				//sv.SetString(valueString(sv, opts))

				//fmt.Println("newVal=" + values.Get(name))
				sv.SetString(values.Get(name))
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
