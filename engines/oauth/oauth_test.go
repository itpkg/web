package oauth_test

import (
	"testing"

	"github.com/itpkg/web/engines/oauth"
)

func TestGoogle(t *testing.T) {
	if g, e := oauth.ReadGoogle("google.json"); e == nil {
		t.Logf("%+v", g)
	} else {
		t.Fatal(e)
	}
}
