package web_test

import (
	"crypto/aes"
	"testing"

	"github.com/itpkg/web"
)

const hello = "Hello, Husky!"

func TestAes(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	cip, e := aes.NewCipher(key)
	if e != nil {
		t.Fatal(e)
	}
	p := web.Encryptor{Cip: cip}

	if buf, err := p.Encode([]byte(hello)); err == nil {
		if s, err := p.Decode(buf); err == nil {
			t.Logf("%x, %s", buf, s)
			if string(s) != hello {
				t.Fatalf("Want %s, get %s", hello, s)
			}
		} else {
			t.Fatal(e)
		}
	} else {
		t.Fatal(e)
	}
}
