package validator

import (
	"errors"
	"math"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

/*
	x nonzero : string(len), int, array(len), float
	x max: int, float
	x min: int, float
	x length: string(len), int, array(len)
	x min length: string(len), int(len), array(len)
	x max length: string(len), int(len), array(len)
	x ends with: string (add '~' in front of the word for ignore case) (add '|' for compare more than one condition)
	x starts with: string (add '~' in front of the word for ignore case) (add '|' for compare more than one condition)
	x equal: string, int
	x dateformat: string (dd-mm-yyyy)
	x timeformat: string (hh:mm:ss or hh:mm)
	x contain: string (add '~' in front of the word for ignore case) (add '|' for compare more than one condition) (include 1 of them = true)
	x constains: string (add '~' in front of the word for ignore case) (add '|' for compare more than one condition) (include all of them = true)
	x email: string
	x must digit: string
	x must numeric: string
	x must alphabet: string
	x must lower case: string
	x must upper case: string
*/

func nonzero(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		valid = len(v.String()) != 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = v.Int() != 0

	case reflect.Float32, reflect.Float64:
		valid = v.Float() != 0

	case reflect.Struct:
		valid = true

	case reflect.Slice:
		valid = v.Len() != 0

	default:
		return "", errors.New(ERR_UNSUPPORTED)
	}

	if !valid {
		return errorMsg(ERR_NONZERO, param, attr), nil
	}

	return "", nil
}

func max(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Int() <= p

	case reflect.Float32, reflect.Float64:
		p, err := toFloat(param)

		if err != nil {
			return "", err
		}

		valid = v.Float() <= p

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_MAX, param, attr), nil
	}

	return "", nil
}

func min(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Int() >= p

	case reflect.Float32, reflect.Float64:
		p, err := toFloat(param)

		if err != nil {
			return "", err
		}

		valid = v.Float() >= p

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_MIN, param, attr), nil
	}

	return "", nil
}

func length(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = int64(len(v.String())) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Int()/int64(math.Pow10(int(p)-1)) > 0 && v.Int()/int64(math.Pow10(int(p)-1)) < 10

	case reflect.Slice:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Len() == int(p)

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_LENGTH, param, attr), nil
	}

	return "", nil
}

func minlength(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = int64(len(v.String())) >= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Int()/int64(math.Pow10(int(p)-1)) >= 1

	case reflect.Slice:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Len() >= int(p)

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_MIN_LENGTH, param, attr), nil
	}

	return "", nil
}

func maxlength(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = int64(len(v.String())) <= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Int()/int64(math.Pow10(int(p)-1)) <= 1

	case reflect.Slice:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Len() <= int(p)

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_MAX_LENGTH, param, attr), nil
	}

	return "", nil
}

func equal(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		//valid = int64(len(v.String())) == p
		valid = v.String() == param

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := toInt(param)

		if err != nil {
			return "", err
		}

		valid = v.Int() == p

	case reflect.Float32, reflect.Float64:
		p, err := toFloat(param)

		if err != nil {
			return "", err
		}

		valid = v.Float() == p

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_EQUAL, param, attr), nil
	}

	return "", nil
}

func dateformat(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		re, err := regexp.Compile(`\d{4}-\d{2}-\d{2}`)

		if err != nil {
			return "", err
		}

		valid = re.MatchString(v.String()) && len(v.String()) == 10

		t, err := time.Parse(`2006-01-02`, v.String())

		if err != nil {
			valid = false
		} else if t.Format(`2006-01-02`) == `0001-01-01` {
			valid = false
		}

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_INVALID_FORMAT, param, attr), nil
	}

	return "", nil
}

func timeformat(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		re1, err := regexp.Compile(`\d{2}:\d{2}`)
		re2, err := regexp.Compile(`\d{2}:\d{2}:\d{2}`)

		if err != nil {
			return "", err
		}

		valid = (re1.MatchString(v.String()) && len(v.String()) == 5) || (re2.MatchString(v.String()) && len(v.String()) == 8)

		t, err := time.Parse(`15:04:05`, v.String())

		if t.Format(`15:04:05`) == `00:00:00` {
			valid = false
		}

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_INVALID_FORMAT, param, attr), nil
	}

	return "", nil
}

func email(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		re, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		if err != nil {
			return "", err
		}

		valid = re.MatchString(v.String())

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_INVALID_FORMAT, param, attr), nil
	}

	return "", nil
}

func startsWith(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		val := strings.Trim(v.String(), " ")
		param_check := param

		if strings.HasPrefix(param, `~`) {
			param_check = strings.ToLower(param_check[1:])
			param = param[1:]
			val = strings.ToLower(val)
		}

		params := strings.Split(param_check, "|")

		temp_valid := false

		for _, param := range params {
			if strings.HasPrefix(val, param) {
				temp_valid = true
			}
		}

		valid = temp_valid

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_STARTS_WITH, param, attr), nil
	}

	return "", nil
}

func endsWith(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		val := strings.Trim(v.String(), " ")
		param_check := param

		if strings.HasPrefix(param, `~`) {
			param_check = strings.ToLower(param_check[1:])
			param = param[1:]
			val = strings.ToLower(val)
		}

		params := strings.Split(param_check, "|")

		temp_valid := false

		for _, param := range params {
			if strings.HasSuffix(val, param) {
				temp_valid = true
			}
		}

		valid = temp_valid

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_ENDS_WITH, param, attr), nil
	}

	return "", nil
}

func upperCase(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		val := strings.Trim(v.String(), " ")
		valid = val == strings.ToUpper(val)

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorList[ERR_UPPER], nil
	}

	return "", nil
}

func lowerCase(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		val := strings.Trim(v.String(), " ")
		valid = val == strings.ToLower(val)

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorList[ERR_LOWER], nil
	}

	return "", nil
}

func contain(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		val := strings.Trim(v.String(), " ")
		param_check := param

		if strings.HasPrefix(param, `~`) {
			param_check = strings.ToLower(param_check[1:])
			param = param[1:]
			val = strings.ToLower(val)
		}

		params := strings.Split(param_check, "|")

		temp_valid := false

		for _, param := range params {
			if strings.Contains(val, param) {
				temp_valid = true
			}
		}

		valid = temp_valid

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_CONTAIN, param, attr), nil
	}

	return "", nil
}

func contains(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		val := strings.Trim(v.String(), " ")
		param_check := param

		if strings.HasPrefix(param, `~`) {
			param_check = strings.ToLower(param_check[1:])
			param = param[1:]
			val = strings.ToLower(val)
		}

		params := strings.Split(param_check, "|")

		for _, param := range params {
			if !strings.Contains(val, param) {
				valid = false
				break
			}
		}

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorMsg(ERR_CONTAINS, param, attr), nil
	}

	return "", nil
}

func isDigit(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		for _, c := range v.String() {
			if !unicode.IsDigit(c) {
				valid = false
				break
			}
		}

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorList[ERR_DIGIT], nil
	}

	return "", nil
}

func isLetter(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		for _, c := range strings.Trim(v.String(), " ") {
			if !unicode.IsLetter(c) {
				valid = false
				break
			}
		}

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorList[ERR_LETTER], nil
	}

	return "", nil
}

func isNumeric(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		for _, c := range v.String() {
			if !unicode.IsNumber(c) {
				valid = false
				break
			}
		}

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorList[ERR_NUMERIC], nil
	}

	return "", nil
}

func isAlphaNumeric(v reflect.Value, param, attr string) (string, error) {
	valid := true

	switch v.Kind() {
	case reflect.String:
		isAlphaNum := true

		for _, c := range v.String() {
			if !unicode.IsNumber(c) && !unicode.IsLetter(c) {
				isAlphaNum = false
				break
			}
		}

		valid = isAlphaNum

	default:
		return "", errors.New(errorList[ERR_UNSUPPORTED] + v.Kind().String())
	}

	if !valid {
		return errorList[ERR_ALPHA_NUMERIC], nil
	}

	return "", nil
}
