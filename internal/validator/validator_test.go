package validator

import (
	"reflect"
	"testing"
)

func TestValidator(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		v := New()
		if got := v.Valid(); !got {
			t.Errorf("Valid() got %v, want %v", got, true)
		}

		v.errors["age"] = "the age should be greater than 18"
		if got := v.Valid(); got {
			t.Errorf("Valid() got %v, want %v", got, false)
		}
	})
	t.Run("Check", func(t *testing.T) {
		v := New()
		key := "username"
		msg := "the username should be defined"

		v.Check(true, key, msg)
		if _, ok := v.errors[key]; ok {
			t.Error("the error message should not have been save")
		}

		v.Check(false, key, msg)
		if got, ok := v.errors[key]; !ok || got != msg {
			t.Error("the error message was not saved correctly")
		}
	})
	t.Run("addError", func(t *testing.T) {
		v := New()
		key := "password"
		msg := "the password should not be empty"
		v.addError(key, msg)
		if got, ok := v.errors[key]; !ok || got != msg {
			t.Error("the error message was not saved correctly")
		}

		msg = "the password should contains at least one special character"
		v.addError(key, msg)
		if got, ok := v.errors[key]; !ok || got == msg {
			t.Error("the original error message was replaced")
		}
	})
	t.Run("Errors", func(t *testing.T) {
		v := New()
		errors := map[string]string{
			"username": "the username should not be empty",
			"password": "the password should not be empty",
		}
		for key, msg := range errors {
			v.addError(key, msg)
		}
		if got := v.Errors(); !reflect.DeepEqual(got, errors) {
			t.Errorf("Errors() got %v, want %v", got, errors)
		}
	})
}

func TestIsUUID(t *testing.T) {
	s := "b27ba9b0-0eb6-4c68-a4ae-d64d582aa56e"
	if got := IsUUID(s); !got {
		t.Errorf("IsUUID(%v) got %v, want %v", s, got, true)
	}
	s = "123"
	if got := IsUUID(s); got {
		t.Errorf("IsUUID(%s) got %v, want %v", s, got, false)
	}
}

func TestIsURL(t *testing.T) {
	s := "https://example.com/image.png"
	if got := IsURL(s); !got {
		t.Errorf("IsURL(%v) got %v, want %v", s, got, true)
	}
	s = "example/image.png"
	if got := IsURL(s); got {
		t.Errorf("IsURL(%v) got %v, want %v", s, got, false)
	}
}
