package gorut

import (
	"errors"
	"testing"
)

func TestIsValid(t *testing.T) {
	// Testing valid RUT
	rut := Rut{"14696787", "6"}
	actual, err := rut.IsValid()
	expected := true

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err != nil {
		t.Errorf("Got %v, expected %v", err, nil)
	}

	// Testing invalid RUT
	rut.Digit = "4"
	actual, err = rut.IsValid()
	expected = false

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err == nil {
		t.Errorf("Got %v, expected %v", nil, err)
	}

	// Testing RUT with len(Numbers) < 7.
	rut.Numbers = "1123256"
	actual, err = rut.IsValid()
	expected = false

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err == nil {
		t.Errorf("Got %v, expected %v", nil, errors.New("RUT numbers length is invalid"))
	}

	// Testing RUT with Numbers = ""
	rut.Numbers = ""
	actual, err = rut.IsValid()
	expected = false

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err == nil {
		t.Errorf("Got %v, expected %v", nil, errors.New("RUT numbers are required"))
	}

	// Testing RUT with Digit = ""
	rut.Numbers = "14696787"
	rut.Digit = ""
	actual, err = rut.IsValid()
	expected = false

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err == nil {
		t.Errorf("Got %v, expected %v", nil, errors.New("RUT digit is required"))
	}
}

func TestFormat(t *testing.T) {
	// Test RUT with len(Numbers) == 8
	rut := Rut{"14696787", "5"}
	actual := rut.Format()
	expected := "14.696.787-5"

	if actual != expected {
		t.Errorf("Got %s, expected %s", actual, expected)
	}

	// Test RUT with len(Numbers) == 7
	rut.Numbers = "8822842"
	rut.Digit = "1"
	actual = rut.Format()
	expected = "8.822.842-1"

	if actual != expected {
		t.Errorf("Got %s, expected %s", actual, expected)
	}

	// Test RUT with invalid Numbers
	rut.Numbers = "1"
	actual = rut.Format()
	expected = "1-1"

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}

func TestValidateRut(t *testing.T) {
	// Test valid RUT with dots and dash
	rut := "14.696.787-6"
	actual, err := ValidateRut(rut)
	expected := true

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err != nil {
		t.Errorf("Got %v, expected %v", err, nil)
	}

	// Test valid RUT with dash and no dots
	rut = "14696787-6"
	actual, err = ValidateRut(rut)
	expected = true

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err != nil {
		t.Errorf("Got %v, expected %v", err, nil)
	}

	// Test valid RUT with dots and no dash
	rut = "14.696.7876"
	actual, err = ValidateRut(rut)
	expected = true

	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

	if err != nil {
		t.Errorf("Got %v, expected %v", err, nil)
	}
}
