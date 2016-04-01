package web

import "encoding/base64"

//FromBase64 string-to-bytes
func FromBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

//ToBase64 bytes-to-string
func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
