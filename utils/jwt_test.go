package utils

import (
	"os"
	"testing"
)

func TestCreateJwt(t *testing.T) {
	os.Setenv("JWT_SECRET", "STUBBLE")
	email := "myemail@email.com"

	returningValue, expiresAt := CreateJwt(email)

	if returningValue != "" && expiresAt > 0 {
		t.Log("Passed CreateJwt")
	} else {
		t.Error("Failed CreateJwt")
	}
}

func TestVerifyJwt(t *testing.T) {
	os.Setenv("JWT_SECRET", "STUBBLE")
	email := "myemail@email.com"

	returningValue, expiresAt := CreateJwt(email)

	if returningValue != "" && expiresAt > 0 {
		t.Error("Failed CreateJwt")
	}

	if returnedValue, err := VerifyJwt(returningValue); err != nil || returnedValue != email {
		t.Error("Failed VerifyJwt")
	} else {
		t.Log("Passed VerifyJwt")
	}
}
