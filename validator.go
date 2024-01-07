package validator

import (
	"errors"
	"reflect"
)

var defaultValidator = NewValidator()

func Validate(v interface{}) (map[string]ErrorArray, error) {
	return defaultValidator.Validate(v)
}

func GetFirstError(msg map[string]ErrorArray) error {
	for _, value := range msg {
		return errors.New(value[0])
	}

	return nil
}

func (mv *Validator) Validate(v interface{}) (map[string]ErrorArray, error) {
	sv := reflect.ValueOf(v)
	st := reflect.TypeOf(v)

	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return nil, errors.New(errorList[ERR_UNSUPPORTED])
	}

	if sv.Kind() != reflect.Struct && sv.Kind() != reflect.Interface {
		return nil, errors.New(errorList[ERR_UNSUPPORTED])
	}

	mapErrors := map[string]ErrorArray{}

	for i := 0; i < sv.NumField(); i++ {
		if st.Field(i).Tag.Get(defaultValidator.tagName) == "" {
			continue
		}

		tags, err := mv.parseTags(splitComma(st.Field(i).Tag.Get(defaultValidator.tagName)))

		if err != nil {
			return nil, errors.New(errorList[ERR_UNSUPPORTED])
		}

		errs := make(ErrorArray, 0, len(tags))

		for _, tag := range tags {

			if errMsg, err := tag.Fn(sv.Field(i), tag.Param, st.Field(i).Name); err != nil {
				return nil, err
			} else if errMsg != "" {
				errs = append(errs, errMsg)
			}
		}

		if len(errs) > 0 {
			mapErrors[st.Field(i).Name] = errs
		}
	}

	return mapErrors, nil
}

func Validation(v interface{}) (string, error) {
	return defaultValidator.Validation(v)
}

func (mv *Validator) Validation(v interface{}) (string, error) {
	sv := reflect.ValueOf(v)
	//st := reflect.TypeOf(v)

	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return "", errors.New(errorList[ERR_UNSUPPORTED])
	}

	if sv.Kind() != reflect.Struct && sv.Kind() != reflect.Interface {
		return "", errors.New(errorList[ERR_UNSUPPORTED])
	}

	//mapErrors := map[string]ErrorArray{}

	for i := 0; i < sv.NumField(); i++ {
		//fmt.Println(st.Field(i).Tag.Get(defaultValidator.tagName), "err")

	}

	return "", nil
}
