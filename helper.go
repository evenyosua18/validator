package validator

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func splitComma(str string) []string {
	pattern := regexp.MustCompile(`((?:^|[^\\])(?:\\\\)*),`)

	var result []string
	index := pattern.FindAllStringIndex(str, -1)
	last := 0

	for _, is := range index {
		result = append(result, str[last:is[1]-1])
		last = is[1]
	}

	result = append(result, str[last:])

	return result
}

func (mv *Validator) parseTags(tagList []string) ([]tag, error) {
	tags := make([]tag, 0, len(tagList))

	for _, i := range tagList {
		i = strings.Replace(i, `\,`, ",", -1)

		tg := tag{}

		v := strings.SplitN(i, "=", 2)

		tg.Name = strings.Trim(v[0], " ")

		if tg.Name == "" {
			return []tag{}, errors.New(errorList[ERR_UNKNOWN_TAG])
		}

		if len(v) > 1 {
			tg.Param = strings.Trim(v[1], " ")
		}

		var found bool
		if tg.Fn, found = mv.validationFuncs[tg.Name]; !found {
			return []tag{}, errors.New(errorList[ERR_UNKNOWN_TAG])
		}
		tags = append(tags, tg)
	}

	return tags, nil
}

func errorMsg(errMsg, param, attr string) string {
	var buffer bytes.Buffer

	buffer.WriteString(attr)
	buffer.WriteString(errorList[errMsg])
	buffer.WriteString(param)

	return strings.Trim(buffer.String(), " ")
}

func toInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 64)

	if err != nil {
		return 0, errors.New(errorList[ERR_BAD_PARAM])
	}

	return i, nil
}

func toFloat(param string) (float64, error) {
	i, err := strconv.ParseFloat(param, 32)

	if err != nil {
		return 0.0, errors.New(errorList[ERR_BAD_PARAM])
	}

	return i, nil
}
