package settings_test

import (
	"crypto/aes"
	"testing"
	"time"

	"github.com/itpkg/web"
	"github.com/itpkg/web/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestGorm(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	cip, e := aes.NewCipher(key)
	if e != nil {
		t.Fatal(e)
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&settings.Model{})

	test(t, &settings.DatabaseProvider{Db: db, Enc: &web.Encryptor{Cip: cip}})
}

type S struct {
	I int
	S string
	T time.Time
}

func test(t *testing.T, p settings.Provider) {
	s := S{I: 123, S: "hello", T: time.Now()}
	run(t, p, "k1", &s, false)
	run(t, p, "k2", &s, true)
}

func run(t *testing.T, p settings.Provider, k string, s *S, f bool) {

	if e := p.Set(k, s, f); e == nil {
		var s1 S
		if e := p.Get(k, &s1); e == nil {
			t.Logf("S1 %+v", s1)
			if s1.I != s.I || s1.S != s.S {
				t.Fatalf("wang %+v, get %+v", s, s1)
			}
		} else {
			t.Fatal(e)
		}
	} else {
		t.Fatal(e)
	}
}
