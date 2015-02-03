package gorut

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Rut struct {
	Numbers string
	Digit   string
}

func (r *Rut) IsValid() (bool, error) {
	if r.Numbers == "" {
		return false, errors.New("RUT numbers are required")
	}

	if r.Digit == "" {
		return false, errors.New("RUT digit is required")
	}

	numbersLength := len(r.Numbers)

	if numbersLength > 8 || numbersLength <= 7 {
		return false, errors.New("RUT numbers length is invalid")
	}

	var sum int64
	multipliers := [8]int64{3, 2, 7, 6, 5, 4, 3, 2}

	for i, number := range r.Numbers {
		integer, err := strconv.ParseInt(string(number), 10, 64)

		if err != nil {
			return false, err
		}

		sum += integer * multipliers[i]
	}

	sum = 11 - (sum % 11)
	digit := strconv.FormatInt(sum, 10)

	if sum == 10 {
		digit = "K"
	}

	if digit != r.Digit {
		return false, errors.New("RUT is invalid")
	}

	return true, nil
}

func (r *Rut) Format() string {
	rx := regexp.MustCompile("([0-9]+)([0-9]{3})([0-9]{3})$")

	chunks := rx.FindStringSubmatch(r.Numbers)

	if len(chunks) == 0 {
		return fmt.Sprintf("%s-%s", r.Numbers, r.Digit)
	}

	return fmt.Sprintf("%s-%s", strings.Join(chunks[1:], "."), r.Digit)
}

func ValidateRut(rut string) (bool, error) {
	r := strings.NewReplacer(".", "", "-", "")
	cleanRut := r.Replace(strings.Trim(rut, " "))
	numbers := string(cleanRut[0 : len(cleanRut)-1])
	digit := string(cleanRut[len(cleanRut)-1])
	newRut := Rut{Numbers: numbers, Digit: digit}

	return newRut.IsValid()
}
