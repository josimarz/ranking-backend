package validator

import (
	"net/url"

	"github.com/google/uuid"
)

type Validator struct {
	errors map[string]string
}

func New() *Validator {
	return &Validator{
		errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.errors) == 0
}

func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.addError(key, msg)
	}
}

func (v *Validator) addError(key, msg string) {
	if _, ok := v.errors[key]; !ok {
		v.errors[key] = msg
	}
}

func (v *Validator) Errors() map[string]string {
	return v.errors
}

func IsUUID(s string) bool {
	return uuid.Validate(s) == nil
}

func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}
