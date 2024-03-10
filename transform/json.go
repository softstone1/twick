package transform

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ToString(value any) (any, error) {
	str, ok := value.(string)
	if !ok {
		return "", errors.New("not string")
	}
	v := strings.TrimSpace(str)
	if len(v) == 0 {
		return "", errors.New("empty string")
	}
	if t, err := time.Parse(time.RFC3339, v); err == nil {
		return t.Unix(), nil
	}
	return v, nil
}

func ToNumber(value any) (float64, error) {
	str, ok := value.(string)
	if !ok {
		return 0, errors.New("not a string")
	}
	v := strings.TrimLeft(strings.TrimSpace(str), "0")
	if len(v) == 0 {
		return 0, errors.New("empty string")
	}
	num, err := strconv.ParseFloat(string(v), 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func ToBoolean(value any) (bool, error) {
	str, ok := value.(string)
	if !ok {
		return false, errors.New("not a string")
	}
	// Sanitize the string value by trimming leading and trailing whitespace.
	v := strings.TrimSpace(str)

	// Transform specified string representations to boolean values.
	switch v {
	case "1", "t", "T", "TRUE", "true", "True":
		return true, nil
	case "0", "f", "F", "FALSE", "false", "False":
		return false, nil
	}

	// Omit fields with invalid Boolean values by returning an error.
	return false, errors.New("invalid boolean value")
}

func ToNull(value any) (any, error) {
	if b, err := ToBoolean(value); !b || err != nil {
		return nil, errors.New("omit field")
	}

	return nil, nil
}

// ParseType applies transformation logic based on value type.
func ParseType(value map[string]any) (string, any, error) {
	for k, v := range value {
    k := strings.TrimSpace(k)
		var newVal any
		var err error
		switch k {
		case "S":
			newVal, err = ToString(v)
		case "N":
			newVal, err = ToNumber(v)
		case "BOOL":
			newVal, err = ToBoolean(v)
		case "NULL":
			newVal, err = ToNull(v)
		case "L":
			newVal, err = ParseList(v)
		case "M":
			newVal, err = ParseMap(v)
		default:
			newVal, err = nil, errors.New("invald type")
		}
		return k, newVal, err
	}
	return "", nil, errors.New("invaild value")
}

// ParseList updated to exclude Null, List, or Map types and omit unsupported or empty values.
func ParseList(value any) (any, error) {
  data, ok := value.([]map[string]any)
  if !ok {
    return nil, errors.New("invalid value")
  }
  var result []any
  for _, v := range data {
    t, newVal, err := ParseType(v)
    if err != nil {
      continue
    }
    if t == "NULL" || t == "L" || t == "M" {
      continue
    }

    result = append(result, newVal)
  }
  return result, nil
}

func ParseMap(value any) (any, error) {
  data, ok := value.(map[string]any)
  if !ok {
    return nil, errors.New("invalid value")
  }
  result := make(map[string]any, len(data))
  for k, v := range data {
    if len(k) == 0 {
      continue
    }
    val, ok := v.(map[string]any)
    if !ok {
      continue
    }
    _, newVal, err := ParseType(val)
    if err != nil {
      continue
    }
    result[k] = newVal
  }
  return result, nil
}