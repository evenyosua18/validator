package validator

import "reflect"

const (
	ERR_UNKNOWN_TAG    = "ErrUnknownTag"
	ERR_UNSUPPORTED    = "ErrUnsupported"
	ERR_NONZERO        = "ErrNonzero"
	ERR_MIN            = "ErrMin"
	ERR_MAX            = "ErrMax"
	ERR_BAD_PARAM      = "ErrBadParamater"
	ERR_EQUAL          = "ErrEqual"
	ERR_INVALID_FORMAT = "ErrInvalidFormat"
	ERR_LENGTH         = "ErrLength"
	ERR_ENDS_WITH      = "ErrEndsWith"
	ERR_STARTS_WITH    = "ErrStartsWith"
	ERR_MIN_LENGTH     = "ErrMinLength"
	ERR_MAX_LENGTH     = "ErrMaxLength"
	ERR_CONTAINS       = "ErrContains"
	ERR_CONTAIN        = "ErrContain"
	ERR_DIGIT          = "ErrDigit"
	ERR_NUMERIC        = "ErrNumeric"
	ERR_LETTER         = "ErrLetter"
	ERR_LOWER          = "ErrLower"
	ERR_UPPER          = "ErrUpper"
	ERR_ALPHA_NUMERIC  = "ErrAlphaNumeric"
)

type ValidationFunc func(v reflect.Value, param, attr string) (string, error)

type Validator struct {
	tagName         string
	validationFuncs map[string]ValidationFunc
}

type tag struct {
	Name  string
	Fn    ValidationFunc
	Param string
}

type ErrorArray []string

var errorList = map[string]string{
	ERR_UNKNOWN_TAG:    "Unknown Tag",
	ERR_BAD_PARAM:      "Bad paramater",
	ERR_UNSUPPORTED:    "Unsupported validation",
	ERR_NONZERO:        " cannot be empty or zero value",
	ERR_MIN:            " cannot less than equal to ",
	ERR_MAX:            " cannot be more than equal to ",
	ERR_EQUAL:          " must be equal to ",
	ERR_INVALID_FORMAT: " format is invalid",
	ERR_LENGTH:         " length or digit must be equal to ",
	ERR_ENDS_WITH:      " ends with ",
	ERR_STARTS_WITH:    " starts with ",
	ERR_MIN_LENGTH:     " length or digit must be more than equal to ",
	ERR_MAX_LENGTH:     " length or digit must be less than equal to ",
	ERR_CONTAINS:       " must be contains all of ",
	ERR_CONTAIN:        " must be contains at least one of ",
	ERR_LETTER:         "String must be letter",
	ERR_DIGIT:          "String must be digit",
	ERR_NUMERIC:        "String must be numeric",
	ERR_LOWER:          "String must be lower case",
	ERR_UPPER:          "String must be upper case",
	ERR_ALPHA_NUMERIC:  "String must be alphanumeric",
}

func NewValidator() *Validator {
	return &Validator{
		tagName: "validate",
		validationFuncs: map[string]ValidationFunc{
			"nonzero":   nonzero,
			"required":  nonzero,
			"max":       max,
			"min":       min,
			"equal":     equal,
			"date":      dateformat,
			"email":     email,
			"len":       length,
			"minlen":    minlength,
			"maxlen":    maxlength,
			"starts":    startsWith,
			"ends":      endsWith,
			"uppercase": upperCase,
			"lowercase": lowerCase,
			"contain":   contain,
			"contains":  contains,
			"letter":    isLetter,
			"digit":     isDigit,
			"numeric":   isNumeric,
			"time":      timeformat,
			"alphanum":  isAlphaNumeric,
		},
	}
}
