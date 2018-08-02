package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"os"
	"errors"
	"fmt"
)

func errPanic(writer http.ResponseWriter,
request *http.Request) error {
	panic("123")
}

type testUserError string

func (u testUserError)Error() string {
	return u.Message()
}

func (u testUserError)Message() string {
	return string(u)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnKnow(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unKnow error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnKnow, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _,tt := range tests {
		f := errWapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com",
			nil)

		f(response, request)

		resp := response.Result()
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _,tt := range tests {
		f := errWapper(tt.h)

		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")

	if resp.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("expect (%d,%s);" +
			"got (%d,%s)",
			expectedCode, expectedMessage,
			resp.StatusCode, body)
	}
}
