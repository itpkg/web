package web

import "crypto/rand"

//Random random bytes
func Random(n int) ([]byte, error) {
	b := make([]byte, n)
	_, e := rand.Read(b)
	return b, e
}
