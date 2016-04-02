package web_test

import (
	"testing"

	"github.com/itpkg/web"
)

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
