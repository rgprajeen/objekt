package domain

import (
	"bytes"
	"encoding/json"
)

type BucketType int

const (
	InvalidType BucketType = iota
	AWS
	Azure
	OCI
)

func (bt BucketType) String() string {
	return typeToString[bt]
}

var typeToString = map[BucketType]string{
	AWS:   "aws",
	Azure: "azure",
	OCI:   "oci",
}

var typeToID = map[string]BucketType{
	"aws":   AWS,
	"azure": Azure,
	"oci":   OCI,
}

// MarshalJSON marshals the enum as a quoted json string
func (bt BucketType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(typeToString[bt])
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
	*bt = typeToID[j]
	return nil
}
