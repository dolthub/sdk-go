package core_models

import (
	"bytes"
	"time"
)

type NilTime struct {
	*time.Time
}

type InstallationFileResponse struct {
	InstallationFile string  `json:"installation_file,omitempty"`
	Version          string  `json:"version,omitempty"`
	RequiresVersion  string  `json:"requires_version,omitempty"`
	HashMd5          string  `json:"hash_md_5,omitempty"`
	ReleaseDate      NilTime `json:"release_date,omitempty"`
}

func (t *NilTime) UnmarshalJSON(data []byte) (err error) {

	// by convention, unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
	if bytes.Equal(data, []byte("null")) || bytes.Equal(data, []byte("\"\"")) {
		return nil
	}

	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse("\""+time.RFC3339+"\"", string(data))
	*t = NilTime{&tt}
	return
}
