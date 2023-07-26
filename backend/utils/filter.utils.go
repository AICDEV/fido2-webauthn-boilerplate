package utils

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/statics"
)

func FilterStorageQuery(i string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9@\.]+`).ReplaceAllString(i, "")
}

func FilterTokenStringFromAuthorizationHeader(req *http.Request) (string, error) {
	tokenHeader := req.Header.Get(statics.GLOBAL_AUTH_HEADER)

	if len(tokenHeader) == 0 {
		return "", errors.New("unable to parse token header")
	}

	token := strings.Split(strings.TrimSpace(tokenHeader), " ")

	if len(token) != 2 {
		return "", errors.New("unable to split token header")
	}

	return token[1], nil
}
