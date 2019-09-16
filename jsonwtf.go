package jsonWTF

import (
	"regexp"
	"strconv"
	"strings"
)

type JsonWTF struct {
	stringValue string
}

var (
	regIsNumber = regexp.MustCompile(`^[0-9]*[.|,]?[0-9]*$`)
)

// MarshalJSON Правила для упаковки в Json
func (j JsonWTF) MarshalJSON() ([]byte, error) {
	if regIsNumber.MatchString(j.stringValue) {
		return []byte(j.stringValue), nil
	}
	if j.stringValue == "true" || j.stringValue == "false" {
		return []byte(j.stringValue), nil
	}
	return []byte(`"` + strings.Replace(j.stringValue, `"`, `\"`, -1)), nil
}

// UnmarshalJSON Правила для распаковки из Json
func (j *JsonWTF) UnmarshalJSON(b []byte) error {
	j.stringValue = strings.Trim(string(b), `"`)
	j.stringValue = strings.Replace(j.stringValue, `\"`, `"`, -1)
	return nil
}

// String Вернёт строку как есть
func (j JsonWTF) String() string {
	return j.stringValue
}

// Float Попытается привести значение к float64
// Если НЕ получится, вернёт .0
func (j JsonWTF) Float() float64 {
	f, _ := strconv.ParseFloat(strings.Replace(j.stringValue, ",", ".", 1), 64)
	return f
}

// Int Попытается привести значение к int64
// Если НЕ получится, вернёт 0
func (j JsonWTF) Int64() int64 {
	return int64(j.Float())
}

// Bool Попытается привести значение к bool
// Если НЕ получится, вернёт false
func (j JsonWTF) Bool() bool {
	b, _ := strconv.ParseBool(j.stringValue)
	return b
}
