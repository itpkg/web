package web_test

import (
	"crypto/aes"
	"testing"

	"github.com/itpkg/web"
)

const hello = "Hello, IT-PACKAGE."

func TestMd5(t *testing.T) {
	var md5 web.Hashing
	md5 = &web.Md5{}
	t.Logf("md5('%s') = %s", hello, md5.Sum([]byte(hello)))
}

func TestSha512(t *testing.T) {
	var sha web.Hashing
	sha = &web.Sha512{}
	t.Logf("sha512('%s') = %s", hello, sha.Sum([]byte(hello)))
}

func TestAes(t *testing.T) {
	key, _ := web.Random(32)
	cip, _ := aes.NewCipher(key)
	var en web.Encryptor
	en = &web.Aes{Cip: cip}
	if h, e := en.Encode([]byte(hello)); e == nil {
		t.Logf("aes('%s') = %s", hello, string(h))
		if d, e := en.Decode(h); e == nil {
			t.Logf("Want %s, get %s", hello, d)
		} else {
			t.Errorf("bad in decode %v", e)
		}
	} else {
		t.Errorf("bad in encode aes %v", e)
	}
}

func TestSalt(t *testing.T) {
	sh := web.SaltHashing{H: &web.Sha512{}}
	if h, e := sh.Sum([]byte(hello), 6); e == nil {
		t.Logf("salt sum '%s' = %s", hello, h)
		if !sh.Check(h, []byte(hello)) {
			t.Errorf("bad in check sum")
		}
	} else {
		t.Errorf("bad in salt %v", e)
	}

	t.Logf("UUID: %s", web.UUID())
}
