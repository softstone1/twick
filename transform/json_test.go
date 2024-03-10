package transform

import (
	"reflect"
	"testing"
	"time"
)

func TestParseInput(t *testing.T) {
	// Test case 1: Valid input
	input := map[string]any{
		"key1": map[string]any{
			"BOOL": "1",
		},
		"key2": map[string]any{
			"subkey3": "value3",
			"subkey4": "value4",
		},
	}
	expected := map[string]any{
		"key1": true,
	}

	result, err := ParseInput(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 2: Invalid input
	invalidInput := "invalid"
	_, err = ParseInput(invalidInput)
	if err != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err)
	}

	// Test case 3: Empty input
	emptyInput := map[string]any{}
	_, err = ParseInput(emptyInput)
	if err != errEmptyData {
		t.Errorf("Expected errEmptyData, got: %v", err)
	}
}

func TestToString(t *testing.T) {
	// Test case 1: Valid string input
	input := "2022-01-01T12:00:00Z"
	expected := time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC).Unix()
	result, err := toString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 2: Empty string input
	emptyInput := ""
	_, err = toString(emptyInput)
	if err != errEmptyData {
		t.Errorf("Expected errEmptyData, got: %v", err)
	}

	// Test case 3: Invalid input
	invalidInput := 123
	_, err = toString(invalidInput)
	if err != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err)
	}
}

func TestToNumber(t *testing.T) {
	// Test case 1: Valid number input
	input := "123.45"
	expected := 123.45
	result, err := toNumber(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 2: Empty string input
	emptyInput := ""
	_, err = toNumber(emptyInput)
	if err != errEmptyData {
		t.Errorf("Expected errEmptyData, got: %v", err)
	}

	// Test case 3: Invalid input
	invalidInput := 123
	_, err = toNumber(invalidInput)
	if err != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err)
	}
}

func TestToBoolean(t *testing.T) {
	// Test case 1: Valid true input
	input := "true"
	expected := true
	result, err := toBoolean(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 2: Valid false input
	input = "false"
	expected = false
	result, err = toBoolean(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 3: Valid true input with leading/trailing spaces
	input = "  true  "
	expected = true
	result, err = toBoolean(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 4: Valid false input with leading/trailing spaces
	input = "  false  "
	expected = false
	result, err = toBoolean(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 5: Invalid input
	input = "invalid"
	_, err = toBoolean(input)
	if err != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err)
	}
}
func TestToNull(t *testing.T) {
	// Test case 1: Valid true input
	input := "true"
	expected := any(nil)
	result, err := toNull(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 2: Invalid false input
	input = "false"
	_, err = toNull(input)
	if err != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err)
	}

	// Test case 3: Valid true input with leading/trailing spaces
	input = "  true  "
	expected = any(nil)
	result, err = toNull(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result, expected)
	}

	// Test case 4: Invalid false input with leading/trailing spaces
	input = "  false  "
	_, err = toNull(input)
	if err != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err)
	}

	// Test case 5: Invalid input
	input = "invalid"
	_, err = toNull(input)
	if err != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err)
	}
}

func TestToType(t *testing.T) {
	// Test case 1: Valid string input
	value1 := map[string]any{
		"S": "Hello, World!",
	}
	expectedKey1 := "S"
	expectedVal1 := "Hello, World!"
	resultKey1, resultVal1, err1 := toType(value1)
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
	if resultKey1 != expectedKey1 || resultVal1 != expectedVal1 {
		t.Errorf("Unexpected result. Got: %v, %v, Expected: %v, %v", resultKey1, resultVal1, expectedKey1, expectedVal1)
	}

	// Test case 2: Valid number input
	value2 := map[string]any{
		"N": "123.45",
	}
	expectedKey2 := "N"
	expectedVal2 := 123.45
	resultKey2, resultVal2, err2 := toType(value2)
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err2)
	}
	if resultKey2 != expectedKey2 || resultVal2 != expectedVal2 {
		t.Errorf("Unexpected result. Got: %v, %v, Expected: %v, %v", resultKey2, resultVal2, expectedKey2, expectedVal2)
	}

	// Test case 3: Valid boolean input
	value3 := map[string]any{
		"BOOL": "true",
	}
	expectedKey3 := "BOOL"
	expectedVal3 := true
	resultKey3, resultVal3, err3 := toType(value3)
	if err3 != nil {
		t.Errorf("Unexpected error: %v", err3)
	}
	if resultKey3 != expectedKey3 || resultVal3 != expectedVal3 {
		t.Errorf("Unexpected result. Got: %v, %v, Expected: %v, %v", resultKey3, resultVal3, expectedKey3, expectedVal3)
	}

	// Test case 4: Invalid null input
	value4 := map[string]any{
		"NULL": "null",
	}
	expectedVal4 := any(nil)
	_, resultVal4, err4 := toType(value4)
	if err4 != errInvalidData {
		t.Errorf("Expected errInvalidType, got: %v", err4)
	}
	if resultVal4 != expectedVal4 {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", resultVal4, expectedVal4)
	}

	// Test case 5: invalid list input
	value5 := map[string]any{
		"L": []any{"item1", "item2", "item3"},
	}
	expectedKey5 := "L"
	expectedVal5 := any(nil)
	resultKey5, resultVal5, err5 := toType(value5)
	if err5 != errEmptyData {
		t.Errorf("Unexpected error: %v", err5)
	}
	if resultKey5 != expectedKey5 || !reflect.DeepEqual(resultVal5, expectedVal5) {
		t.Errorf("Unexpected result. Got: %v, %v, Expected: %v, %v", resultKey5, resultVal5, expectedKey5, expectedVal5)
	}

	// Test case 6: Invalid map input
	value6 := map[string]any{
		"M": map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}
	expectedKey6 := "M"
	expectedVal6 := any(nil)
	resultKey6, resultVal6, err6 := toType(value6)
	if err6 != errEmptyData {
		t.Errorf("Unexpected error: %v", err6)
	}
	if resultKey6 != expectedKey6 || !reflect.DeepEqual(resultVal6, expectedVal6) {
		t.Errorf("Unexpected result. Got: %v, %v, Expected: %v, %v", resultKey6, resultVal6, expectedKey6, expectedVal6)
	}

	// Test case 7: Invalid input
	value7 := map[string]any{
		"INVALID": "value",
	}
	expectedVal7 := any(nil)
	_, resultVal7, err7 := toType(value7)
	if err7 != errInvalidType {
		t.Errorf("Expected errInvalidType, got: %v", err7)
	}
	if resultVal7 != expectedVal7 {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", resultVal7, expectedVal7)
	}

	// Test case 8: Empty input
	value8 := map[string]any{}
	expectedVal8 := any(nil)
	_, resultVal8, err8 := toType(value8)
	if err8 != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err8)
	}
	if resultVal8 != expectedVal8 {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", resultVal8, expectedVal8)
	}
}

func TestToList(t *testing.T) {
	// Test case 1: Valid list input
	input1 := []any{
		map[string]any{
			"S": "item1",
		},
		map[string]any{
			"N": "123.45",
		},
		map[string]any{
			"BOOL": "true",
		},
	}
	expected1 := []any{"item1", 123.45, true}
	result1, err1 := toList(input1)
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result1, expected1)
	}

	// Test case 2: Valid list input with invalid values
	input2 := []any{
		map[string]any{
			"S": "item1",
		},
		map[string]any{
			"N": "123.45",
		},
		map[string]any{
			"BOOL": "true",
		},
		map[string]any{
			"NULL": "null",
		},
		map[string]any{
			"L": []any{"item2", "item3"},
		},
		map[string]any{
			"M": map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
		map[string]any{
			"INVALID": "value",
		},
	}
	expected2 := []any{"item1", 123.45, true}
	result2, err2 := toList(input2)
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err2)
	}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", result2, expected2)
	}

	// Test case 3: Empty list input
	input3 := []any{}
	_, err3 := toList(input3)
	if err3 != errEmptyData {
		t.Errorf("Expected errEmptyData, got: %v", err3)
	}

	// Test case 4: Invalid list input
	input4 := "invalid"
	_, err4 := toList(input4)
	if err4 != errInvalidData {
		t.Errorf("Expected errInvalidData, got: %v", err4)
	}
}
