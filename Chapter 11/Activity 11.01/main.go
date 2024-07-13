// Building a program to validate Social Security Numbers
package main

import (
	"errors"
	"log"
	"strconv"
)

var (
	ErrInvalidSSNLength  = errors.New("invalid SSN length")
	ErrInvalidSSNNumbers = errors.New("invalid SSN numbers")
	ErrInvalidSSNPrefix  = errors.New("invalid SSN prefix")
	ErrInvalidDigitPlace = errors.New("invalid digit place")
)

func checkSSNLength(ssn string) error {
	if len(ssn) != 9 {
		return ErrInvalidSSNLength
	}
	return nil
}

func checkSSNNumbers(ssn string) error {
	for i, v := range ssn {
		switch i {
		case 3, 6:
			continue
		}
		if _, err := strconv.Atoi(string(v)); err != nil {
			return ErrInvalidSSNNumbers
		}
	}
	return nil
}

func checkSSNPrefix(ssn string) error {
	if ssn[0:3] == "000" {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func checkDigitPlace(ssn string) error {
	if ssn[0] == '9' {
		if ssn[4] == '7' || ssn[4] == '9' {
			return ErrInvalidDigitPlace
		}
	}
	return nil
}

func main() {
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	for _, ssn := range validateSSN {
		if err := checkSSNLength(ssn); err != nil {
			log.Printf("Validating value %s caused error %s\n", ssn, err)
		}
		if err := checkSSNNumbers(ssn); err != nil {
			log.Printf("Validating value %s caused error %s\n", ssn, err)
		}
		if err := checkSSNPrefix(ssn); err != nil {
			log.Printf("Validating value %s caused error %s\n", ssn, err)
		}
		if err := checkDigitPlace(ssn); err != nil {
			log.Printf("Validating value %s caused error %s\n", ssn, err)
		}
	}
}
