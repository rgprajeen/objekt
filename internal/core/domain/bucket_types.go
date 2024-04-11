package domain

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (bt BucketType) Value() (string, error) {
	if bt == InvalidType {
		return "", errors.New("invalid bucket type")
	}
	return typeToString[bt], nil
}

func (bt *BucketType) Scan(value any) error {
	switch value := value.(type) {
	case string:
		*bt = typeToID[value]
	case []byte:
		*bt = typeToID[string(value)]
	case int:
		*bt = BucketType(value)
	default:
		return errors.New("incompatible type for BucketType")
	}
	return nil
}
