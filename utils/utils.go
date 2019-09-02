package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	schema "github.com/gorilla/Schema"
	"golang.org/x/crypto/bcrypt"
)

// ParseForm parses request body to given target
func ParseForm(dest interface{}, req *http.Request) {
	err := req.ParseForm()
	Must(err)
	decoder := schema.NewDecoder()
	err = decoder.Decode(dest, req.PostForm)
	Must(err)
}

// CompareHashAndPassword returns error if password and hash doesn't match
func CompareHashAndPassword(token []byte, pwd []byte) error {
	if err := bcrypt.CompareHashAndPassword(token, pwd); err != nil {
		return err
	}
	return nil
}

// Hash generates hash for given input with secret key of hmac object
func Hash(input string) string {
	h := hmac.New(sha256.New, []byte("somekey")) // this can be taken from env variable too
	h.Reset()
	h.Write([]byte(input))
	b := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}

// GenerateRemeberToken returns a 32 bytes random token string using crypto/rand packages
func GenerateRemeberToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	Must(err)
	return base64.URLEncoding.EncodeToString(b)
}

// Must panic if error is present
func Must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

// MustAndSendError exits programs if error is present and also write it to response
func MustAndSendError(w http.ResponseWriter, err error) bool {
	if err != nil {
		fmt.Println("Error ", err.Error())
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error()+"<br><a class=\"btn btn-primary\" href=\"/\">GOTO HOMEPAGE</a>")
		return true
	}
	return false
}
