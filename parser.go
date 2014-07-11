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

	variableValue := reflect.Indirect(reflect.ValueOf(variable))

	for i := 0; i < variableValue.NumField(); i++ {
		structField := variableValue.Field(i)
		fieldType := variableValue.Type().Field(i)
		fieldName := fieldType.Name
		fieldTag := fieldType.Tag

		// Check if field have description
		if fieldTag.Get("description") == "" {
			return errors.New(ERR_NO_DESCRIPTION)
		}

		// Check type of field and set flag variable

		switch structField.Kind() {
		case reflect.Bool:
			defaultValue, err := strconv.ParseBool(fieldTag.Get("default"))
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.BoolVar(structField.Addr().Interface().(*bool), fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Int:
			defaultValue, err := strconv.Atoi(fieldTag.Get("default"))
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.IntVar(structField.Addr().Interface().(*int), fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Int64:
			defaultValue, err := strconv.ParseInt(fieldTag.Get("default"), 10, 64)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.Int64Var(structField.Addr().Interface().(*int64), fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.String:
			defaultValue := fieldTag.Get("default")
			flag.StringVar(structField.Addr().Interface().(*string), fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Float64:
			defaultValue, err := strconv.ParseFloat(fieldTag.Get("default"), 64)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.Float64Var(structField.Addr().Interface().(*float64), fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Uint64:
			defaultValue, err := strconv.ParseUint(fieldTag.Get("default"), 10, 64)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.Uint64Var(structField.Addr().Interface().(*uint64), fieldName, defaultValue, fieldTag.Get("description"))

		case reflect.Uint:
			defaultValue, err := strconv.ParseUint(fieldTag.Get("default"), 10, 0)
			if err != nil {
				return errors.New(ERR_PARSE_ERROR + err.Error())
			}
			flag.UintVar(structField.Addr().Interface().(*uint), fieldName, uint(defaultValue), fieldTag.Get("description"))

		default:
			return errors.New(ERR_TYPE_ERROR)
		}

	}

	flag.Parse()
	return nil
}
