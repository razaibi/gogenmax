package logic

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

// ConvertToKebabCase converts a string to kebab-case
func ConvertToKebabCase(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	return strcase.ToKebab(str), nil
}

// ConvertToCamelCase converts a string to CamelCase
func ConvertToCamelCase(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	return strcase.ToCamel(str), nil
}

// ConvertToSnakeCase converts a string to snake_case
func ConvertToSnakeCase(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	return strcase.ToSnake(str), nil
}
