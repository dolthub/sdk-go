package client

import (
	"encoding/json"
	"log"
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

		tag := typ.Field(i).Tag.Get("json")

		if strings.Contains(tag, "omitempty") &&
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

		tagValue := tagValue(tag)
		var key string

		if tagValue != "" {
			key = tagValue
		} else {
			key = snakeCase(typ.Field(i).Name)
		}

		if value.Field(i).Kind() == reflect.Struct {
			out[key] = SnakeCaseEncode(value.Field(i).Interface())
		} else {
			out[key] = value.Field(i).Interface()
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

func SnakeCaseDecode(jsn []byte) ([]byte, error) {
	m := map[string]interface{}{}

	err := json.Unmarshal(jsn, &m)
	if err != nil {
		log.Panic(err)
	}

	decoded := decodeJsonMap(m)
	res, _ := json.Marshal(decoded)

	return res, nil
}

func tagValue(tag string) string {
	if tag == "" {
		return ""
	}
	if !strings.Contains(tag, ",") {
		return tag
	}
	return tag[:strings.IndexByte(tag, ',')]
}

func decodeJsonMap(jsn map[string]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for key, val := range jsn {
		switch val.(type) {
		case map[string]interface{}:
			val = decodeJsonMap(val.(map[string]interface{}))
		}

		ccKey := CamelCase(key)
		res[ccKey] = val
	}

	return res
}

func CamelCase(s string) string {
	out := make([]rune, 0, utf8.RuneCountInString(s))

	capital := true

	for _, r := range s {
		if r == '_' {
			capital = true
			continue
		}
		if r >= 'a' && r <= 'z' && capital {
			r -= 32
			capital = false
		}
		if r >= 'A' && r <= 'Z' && capital {
			capital = false
		}

		out = append(out, r)
	}

	return string(out)
}
