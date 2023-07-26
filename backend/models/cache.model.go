package models

import "github.com/go-webauthn/webauthn/webauthn"

type CacheModel struct {
	User        User
	SessionData *webauthn.SessionData
}
