package models

import (
	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	RegistrationID []byte
	ID             string `bson:"_id,omitempty"`
	Email          string
	Firstname      string
	Lastname       string
	DisplayName    string
	Credentials    []webauthn.Credential `bson:"credentials"`
}

func (u *User) WebAuthnID() []byte {

	if len(u.RegistrationID) < 1 {
		return []byte(u.ID)
	}

	return u.RegistrationID
}

func (u *User) WebAuthnName() string {
	return u.Email
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u *User) WebAuthnIcon() string {
	// FOR CONVENIENCE: EMPTY
	return ""
}
