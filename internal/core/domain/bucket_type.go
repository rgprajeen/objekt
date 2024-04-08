package domain

import (
	"bytes"
	"encoding/json"
)

type BucketType int

const (
	AWS BucketType = iota
	Azure
	OCI
)

func (bt BucketType) String() string {
	return toString[bt]
}

var toString = map[BucketType]string{
	AWS:   "AWS",
	Azure: "Azure",
	OCI:   "OCI",
}

var toID = map[string]BucketType{
	"AWS":   AWS,
	"Azure": Azure,
	"OCI":   OCI,
}

// MarshalJSON marshals the enum as a quoted json string
func (bt BucketType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[bt])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (bt *BucketType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'AWS' in this case.
	*bt = toID[j]
	return nil
}
