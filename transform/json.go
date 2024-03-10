package transform

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"
)
var (
	errInvalidData     = errors.New("invalid data")
	errInvalidType    = errors.New("invalid data type")
	errEmptyData      = errors.New("empty data")
)

// ParseInput recursively parses the schema-less JSON input data and returns a desired json output.
func ParseInput(input any) (any, error) {
	data, ok := input.(map[string]any)
	if !ok {
		return nil, errInvalidData
	}
	unsortedResult := make(map[string]any, len(data))
	for k, v := range data {
		if len(k) == 0 {
			continue
		}
		val, ok := v.(map[string]any)
		if !ok {
			continue
		}
		_, newVal, err := toType(val)
		if err != nil {
			continue
		}
		unsortedResult[k] = newVal
	}
	if len(unsortedResult) == 0 {
		return nil, errEmptyData
	}
	var sortedKeys []string
	for k := range unsortedResult {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	// Create a new result map with sorted keys
	sortedResult := make(map[string]any, len(unsortedResult))
	for _, k := range sortedKeys {
		sortedResult[k] = unsortedResult[k]
	}

	return sortedResult, nil
}
	
// toString validates and transforms to a string.
func toString(value any) (any, error) {
	str, ok := value.(string)
	if !ok {
		return "", errInvalidData
	}
	v := strings.TrimSpace(str)
	if len(v) == 0 {
		return "", errEmptyData
	}
	if t, err := time.Parse(time.RFC3339, v); err == nil {
		return t.Unix(), nil
	}
	return v, nil
}

// toNumber validates and transforms to a number.
func toNumber(value any) (float64, error) {
	str, ok := value.(string)
	if !ok {
		return 0, errInvalidData
	}
	v := strings.TrimLeft(strings.TrimSpace(str), "0")
	if len(v) == 0 {
		return 0, errEmptyData
	}
	num, err := strconv.ParseFloat(string(v), 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// toBoolean validates and transforms to a boolean.
func toBoolean(value any) (bool, error) {
	str, ok := value.(string)
	if !ok {
		return false, errInvalidData
	}

	v := strings.TrimSpace(str)

	switch v {
	case "1", "t", "T", "TRUE", "true", "True":
		return true, nil
	case "0", "f", "F", "FALSE", "false", "False":
		return false, nil
	}
	return false, errInvalidData
}

// toNull validates and transforms to a null.
func toNull(value any) (any, error) {
	if b, err := toBoolean(value); !b || err != nil {
		return nil, errInvalidData
	}

	return nil, nil
}

// toType validates and transforms to a desired data type.
func toType(value map[string]any) (string, any, error) {
	for k, v := range value {
		k := strings.TrimSpace(k)
		var newVal any
		var err error
		switch k {
		case "S":
			newVal, err = toString(v)
		case "N":
			newVal, err = toNumber(v)
		case "BOOL":
			newVal, err = toBoolean(v)
		case "NULL":
			newVal, err = toNull(v)
		case "L":
			newVal, err = toList(v)
		case "M":
			newVal, err = ParseInput(v)
		default:
			newVal, err = nil, errInvalidType
		}
		return k, newVal, err
	}
	return "", nil, errInvalidData
}

// toList validates and transforms to a list of desired data types.
func toList(value any) (any, error) {
	data, ok := value.([]any)
	if !ok {
		return nil, errInvalidData
	}
	var result []any
	for _, v := range data {
		val, ok := v.(map[string]any)
		if !ok {
			continue
		}
		t, newVal, err := toType(val)
		if err != nil {
			continue
		}
		if t == "NULL" || t == "L" || t == "M" {
			continue
		}

		result = append(result, newVal)
	}
	if len(result) == 0 {
		return nil, errEmptyData
	}
	return result, nil
}


