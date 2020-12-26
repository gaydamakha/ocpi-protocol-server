package emsp

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

var authRegexp = regexp.MustCompile("^Token ([!-~]+)$")

// baseRequest represents a form of all incoming requests to eMSP system
type baseRequest struct {
	Authorization string
	RawRequest    http.Request
}

// ExtractAuthorization extracts an authorization header
// Return an error if header could not be extracted or
// does not consist of printable, non-whitespace characters
func ExtractAuthorization(req *http.Request) (auth string, err error) {
	auth = req.Header.Get("Authorization")
	if auth == "" {
		return "", errors.New("Authorization header must be provided")
	}

	if !authRegexp.MatchString(auth) {
		return "", fmt.Errorf("Invalid token: must consist of printable, non-whitespace characters")
	}

	auth = authRegexp.FindStringSubmatch(auth)[1]
	return
}

func newBaseRequest(req *http.Request) (breq *baseRequest, err error) {
	auth, err := ExtractAuthorization(req)
	if err != nil {
		return nil, fmt.Errorf("Can't instantiate base request:" + err.Error())
	}

	return &baseRequest{
		Authorization: auth,
		RawRequest:    *req,
	}, nil
}
