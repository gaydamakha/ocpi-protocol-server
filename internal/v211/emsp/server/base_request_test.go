package server

import (
	"net/http"
	"testing"
)

func TestExtractAuthorization(t *testing.T) {
	token := `abcd123456789`
	req, _ := http.NewRequest(`GET`, `url`, http.NoBody)
	req.Header.Add(`Authorization`, "Token "+token)

	extracted, err := ExtractAuthorization(req)
	if extracted != token || err != nil {
		t.Fatalf(`Expected: %q, got %q ; err is \"%v\"`, token, extracted, err)
	}
}

func TestNoHeader(t *testing.T) {
	req, _ := http.NewRequest(`GET`, `url`, http.NoBody)

	extracted, err := ExtractAuthorization(req)
	if extracted != "" {
		t.Fatalf(`Expected "", got %q; error is "%v"`, extracted, err)
	}
	if err == nil {
		t.Fatalf(`Expected not nil error`)
	}
}

func TestEmptyAuthorization(t *testing.T) {
	req, _ := http.NewRequest(`GET`, `url`, http.NoBody)
	req.Header.Add(`Authorization`, "")

	extracted, err := ExtractAuthorization(req)
	if extracted != "" {
		t.Fatalf(`Expected "", got %q; error is "%v"`, extracted, err)
	}
	if err == nil {
		t.Fatalf(`Expected not nil error`)
	}
}

func TestInvalidAuthorization(t *testing.T) {
	var token string
	req, _ := http.NewRequest(`GET`, `url`, http.NoBody)

	t.Run(`Authorization is empty`, func(t *testing.T) {
		token = ""
	})
	t.Run(`Authorization contains spaces`, func(t *testing.T) {
		token = `abcd 123abcd`
	})
	t.Run(`Authorization contains newlines`, func(t *testing.T) {
		token = "abcd\n123abcd"
	})
	t.Run(`Authorization contains tabs`, func(t *testing.T) {
		token = "\tabcd123abcd"
	})

	req.Header.Add(`Authorization`, "Token "+token)

	extracted, err := ExtractAuthorization(req)
	if extracted != "" {
		t.Fatalf(`Expected "", got %q; error is "%v"`, extracted, err)
	}
	if err == nil {
		t.Fatalf(`Expected not nil error`)
	}
}

func TestNewRequest(t *testing.T) {
	token := `abcd123456789`
	req, _ := http.NewRequest(`GET`, `url`, http.NoBody)
	req.Header.Add(`Authorization`, "Token "+token)

	var breq *baseRequest
	var err error
	breq, err = newBaseRequest(req)

	if breq == nil {
		t.Fatal(`Expected baseRequest object, got nil`)
	}
	if err != nil {
		t.Fatalf(`Expected error to be nil, got "%v"`, err)
	}
	if breq.Authorization != token {
		t.Fatalf(`Expetected baseRequest.Authorization to be "%q", got "%q"`, token, breq.Authorization)
	}
}

func TestNewRequestEmptyAuthorization(t *testing.T) {
	req, _ := http.NewRequest(`GET`, `url`, http.NoBody)
	req.Header.Add(`Authorization`, "")

	var breq *baseRequest
	var err error
	breq, err = newBaseRequest(req)

	if breq != nil {
		t.Fatalf(`Expected baseRequest to be nil, got "%v"`, breq)
	}
	if err == nil {
		t.Fatalf(`Expected not nil error`)
	}
}
