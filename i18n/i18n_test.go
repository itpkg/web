package i18n_test

import (
	"log"
	"os"
	"testing"

	"github.com/itpkg/web/engines/base"
	"github.com/itpkg/web/i18n"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
)

var logger = log.New(os.Stdout, "[test]", 0)
var lang = &language.SimplifiedChinese

func TestDatabase(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&i18n.Locale{})
	testProvider(t, &i18n.DatabaseProvider{Db: db, Logger: logger})
}

func TestRedis(t *testing.T) {
	re := base.Redis{Host: "localhost", Port: 6379}
	testProvider(t, &i18n.RedisProvider{Redis: re.Open(), Logger: logger})

}

func testProvider(t *testing.T, p i18n.Provider) {
	key := "hello"
	val := "你好"
	p.Set(lang, key, val)
	if val1 := p.Get(lang, key); val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
	ks := p.Keys(lang)
	if len(ks) == 0 {
		t.Errorf("empty keys")
	} else {
		t.Log(ks)
	}
	p.Del(lang, key)
}
