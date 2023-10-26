package utils

import (
	"database/sql/driver"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func JSONScan(dbValue interface{}, value interface{}) error {
	switch v := dbValue.(type) {
	case []byte:
		bytes := v
		if len(bytes) > 0 {
			return Json.Unmarshal(bytes, value)
		}
		return nil
	case string:
		str := v
		if str == "" {
			return nil
		}
		return Json.Unmarshal([]byte(str), value)
	case nil:
		return nil
	default:
		return fmt.Errorf("cannot sql.Scan() from: %#v", value)
	}
}

func JSONValue(value interface{}) (driver.Value, error) {
	if zeroCheck, ok := value.(interface {
		IsZero() bool
	}); ok {
		if zeroCheck.IsZero() {
			return "", nil
		}
	}
	bytes, err := Json.Marshal(value)
	if err != nil {
		return "", err
	}
	str := string(bytes)
	if str == "null" {
		return "", nil
	}
	return str, nil
}
