package helper

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

func IsEmptyUUID(id uuid.UUID) bool {
	for x := 0; x < 16; x++ {
		if id[x] != 0 {
			return false
		}
	}

	return true
}

func IsValidateCPF(cpf string) error {

	re := regexp.MustCompile(`[0-9]{3}\.?[0-9]{3}\.?[0-9]{3}\-?[0-9]{2}`)

	if !re.MatchString(cpf) || len(cpf) != 11 {
		return errors.New("error: Invalid document cpf")
	}

	return nil
}

func IsValidateCNPJ(cnpj string) error {

	re := regexp.MustCompile(`[0-9]{2}\.?[0-9]{3}\.?[0-9]{3}\/?[0-9]{4}\-?[0-9]{2}`)

	if !re.MatchString(cnpj) || len(cnpj) != 14 {
		return errors.New("error: Invalid document cnpj")
	}

	return nil
}
func IsValidateEMAIL(email string) error {

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(email) {
		return errors.New("error: Invalid adress of email")
	}

	return nil
}