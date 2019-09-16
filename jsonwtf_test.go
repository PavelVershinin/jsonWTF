package jsonWTF

import (
	"encoding/json"
	"testing"
)

func TestJsonWTF_UnmarshalJSON(t *testing.T) {
	var testStruct struct {
		BoolAsString  JsonWTF `json:"bool_as_string"`
		FloatAsString JsonWTF `json:"float_as_string"`
		IntAsString   JsonWTF `json:"int_as_string"`
		BoolValid     JsonWTF `json:"bool_valid"`
		IntValid      JsonWTF `json:"int_valid"`
		FloatValid    JsonWTF `json:"float_valid"`
		StringValid   JsonWTF `json:"string_valid"`
	}
	var testJson = []byte(`
		{
			"bool_as_string":"true",
			"float_as_string": "3,1415926535897932384626433832795",
			"int_as_string": "68465",
			"bool_valid": true,
			"int_valid": 68465,
			"float_valid": 3.1415926535897932384626433832795,
			"string_valid": "Мама мыла раму"
		}
	`)

	if err := json.Unmarshal(testJson, &testStruct); err != nil {
		t.Fatal(err)
	}

	if testStruct.BoolAsString.Bool() != true {
		t.Errorf("Wrong value of BoolAsString. Expected: true, real: %s", testStruct.BoolAsString)
	}

	if testStruct.FloatAsString.Float() != 3.1415926535897932384626433832795 {
		t.Errorf("Wrong value of FloatAsString. Expected: 3.1415926535897932384626433832795, real: %s", testStruct.FloatAsString)
	}

	if testStruct.IntAsString.Int64() != 68465 {
		t.Errorf("Wrong value of IntAsString. Expected: 68465, real: %s", testStruct.IntAsString)
	}

	if testStruct.BoolValid.Bool() != true {
		t.Errorf("Wrong value of BoolValid. Expected: true, real: %s", testStruct.BoolValid)
	}

	if testStruct.FloatValid.Float() != 3.1415926535897932384626433832795 {
		t.Errorf("Wrong value of FloatValid. Expected: 3.1415926535897932384626433832795, real: %s", testStruct.FloatValid)
	}

	if testStruct.IntValid.Int64() != 68465 {
		t.Errorf("Wrong value of IntValid. Expected: 68465, real: %s", testStruct.IntValid)
	}

	if testStruct.StringValid.String() != "Мама мыла раму" {
		t.Errorf("Wrong value of StringValid. Expected: \"Мама мыла раму\", real: %s", testStruct.StringValid)
	}
}
