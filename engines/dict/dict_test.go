package dict_test

import (
	"strconv"
	"testing"

	"github.com/itpkg/web/engines/dict"
)

func TestString(t *testing.T) {
	for _, s := range []string{"ls /tmp", "| more", "@ more", "\n \t \b", "\" '' ``", "中文", "!test", ">/tmp/aaa", "</tmp/aaa"} {
		t.Log(strconv.QuoteToASCII(s))
	}
}

func tTestStarDict(t *testing.T) {
	test(t, &dict.StarDict{Dir: "/opt/dic"})
}

func test(t *testing.T, d dict.Provider) {
	if ds, er := d.List(); er == nil {
		t.Log(ds)
	} else {
		t.Fatal(er)
	}
	if rs, er := d.Query("one"); er == nil {
		t.Log(rs)
	} else {
		t.Fatal(er)
	}
}
