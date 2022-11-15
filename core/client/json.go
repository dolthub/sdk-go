package client

import (
	"reflect"
	"strings"
	"unicode/utf8"
)

func SnakeCaseEncode(i interface{}) map[string]interface{} {
	typ, value := reflect.TypeOf(i), reflect.ValueOf(i)

	if typ.Kind() == reflect.Ptr {
		i := reflect.Indirect(value).Interface()
		typ, value = reflect.TypeOf(i), reflect.ValueOf(i)
	}

	out := make(map[string]interface{}, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Tag.Get("json") == "-" {
			continue
		}

		if strings.Contains(typ.Field(i).Tag.Get("json"), "omitempty") &&
			value.Field(i).IsZero() {
			continue
		}

		// collapse embedded structs
		if typ.Field(i).Anonymous == true {
			vals := SnakeCaseEncode(value.Field(i).Interface())
			for k, v := range vals {
				out[k] = v
			}
			continue
		}

		k := snakeCase(typ.Field(i).Name)

		if value.Field(i).Kind() == reflect.Struct {
			out[k] = SnakeCaseEncode(value.Field(i).Interface())
		} else {
			out[k] = value.Field(i).Interface()
		}

	}

	return out
}

func snakeCase(s string) string {
	out := make([]rune, 0, utf8.RuneCountInString(s))

	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			r += 32

			if i > 0 {
				out = append(out, '_')
			}
		}

		out = append(out, r)
	}

	return string(out)
}
