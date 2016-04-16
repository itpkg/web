package cms

import "github.com/itpkg/web/engines/oauth"

//Article model
type Article struct {
	Author oauth.User
}
