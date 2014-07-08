package cmd

import (
	"errors"
	"flag"
	"reflect"
	"strconv"
)

func GetArguments(variable interface{}) error {
	//Check is this a pointer to struct
	variableType := reflect.TypeOf(variable)
	if variableType.Kind() != reflect.Ptr {
		return errors.New(ERR_PASS_BY_VALUE)
	} else if variableValue := variableType.Elem(); variableValue.Kind() != reflect.Struct {
		return errors.New(ERR_NOT_A_STRUCT)
	}

	// Parse parameters

	variableValue := variableType.Elem()

	for i := 0; i < variableValue.NumField(); i++ {
		structField := variableValue.Field(i)
		fieldName := structField.Name
		fieldType := structField.Type
		fieldTag := structField.Tag

		// Check if field have description
		if fieldTag.Get("description") == nil {
			return errors.New(ERR_NO_DESCRIPTION)
		}

		// Check type of field and set flag variable

		switch fieldType.Kind() {
		case reflect.Bool:
			defaultValue, err := strconv.ParseBool(fieldTag.Get("default"))
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.BoolVar(&structField, fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Int:
			defaultValue, err := strconv.Atoi(fieldTag.Get("default"))
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.IntVar(&structField, fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Int64:
			defaultValue, err := strconv.ParseInt(fieldTag.Get("default"), 10, 64)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.Int64Var(&structField, fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.String:
			defaultValue := fieldTag.Get("default")
			flag.StringVar(&structField, fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Float64:
			defaultValue, err := strconv.ParseFloat(fieldTag.Get("default"), 64)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.Float64Var(&structField, fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Uint64:
			defaultValue, err := strconv.ParseUint(fieldTag.Get("default"), 10, 64)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.Uint64Var(&structField, fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Uint:
			defaultValue, err := strconv.ParseUint(fieldTag.Get("default"), 10, 0)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.UintVar(&structField, fieldName, defaultValue, fieldTag.Get("description"))

		default:
			return errors.New(ERR_TYPE_ERROR)
		}

	}
	return nil
}
