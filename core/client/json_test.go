package client

import (
	"fmt"
	"testing"
)

func Test_decodeJsonMap(t *testing.T) {
	type args struct {
		jsn map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			"1",
			args{map[string]interface{}{
				"test":        "abcd",
				"license_key": 2,
				"embed": map[string]interface{}{
					"embedded_key": "efgh",
					"nul":          "afvaf",
				},
			}},
			map[string]interface{}{
				"Test":       "abcd",
				"LicenseKey": 2,
				"Embed": map[string]interface{}{
					"EmbeddedKey": "efgh",
					"Nul":         "afvaf",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeJsonMap(tt.args.jsn)
			if fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("decodeJsonMap() = %v, want %v", tt.args.jsn, tt.want)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"license_key"},
			"LicenseKey",
		},
		{
			"2",
			args{"license"},
			"License",
		},
		{
			"3",
			args{"license_key_test"},
			"LicenseKeyTest",
		},
		{
			"4",
			args{"licen123se"},
			"Licen123se",
		},
		{
			"5",
			args{"EmbeddedKey"},
			"EmbeddedKey",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.s); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
