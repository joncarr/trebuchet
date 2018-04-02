package hash

// TODO: This entire file is probably better suited outside of the app in general and just used as a package when creating the app
// * Leave the file inside of app directory for now until we see how this can be yanked out...

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// NewHMAC creates and returns a new HMAC object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// Hash will hash the provided input string using HMAC with
// the secret key provided when the HMAC object was created
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}

// HMAC is a wrapper around the crypto/hmac package making
// it a little easier to use in our code.
type HMAC struct {
	hmac hash.Hash
}
