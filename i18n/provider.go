package i18n

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/language"
)

//Provider i18n provider
type Provider interface {
	Set(lang *language.Tag, code, message string)
	Get(lang *language.Tag, code string) string
	Del(lang *language.Tag, code string)
	Keys(lang *language.Tag) []string
}

//I18n i18n helper
type I18n struct {
	Provider Provider
	Locales  map[string]map[string]string
	Logger   *log.Logger
}

//Items list all items
func (p *I18n) Items(lng *language.Tag) map[string]interface{} {
	rt := make(map[string]interface{})
	if items, ok := p.Locales[lng.String()]; ok {
		for k, v := range items {
			if strings.HasPrefix(k, "web.") {
				k = k[4:]
				codes := strings.Split(k, ".")
				tmp := rt
				for i, c := range codes {
					if i+1 == len(codes) {
						tmp[c] = v
					} else {
						if tmp[c] == nil {
							tmp[c] = make(map[string]interface{})
						}
						tmp = tmp[c].(map[string]interface{})
					}
				}

			}
		}
	}
	return rt
}

//Exist is lang exist?
func (p *I18n) Exist(lang string) bool {
	_, ok := p.Locales[lang]
	return ok
}

//Load load locales from filesystem
func (p *I18n) Load(dir string) {
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		p.Logger.Printf("Find locale file %s", path)
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			ss := strings.Split(info.Name(), ".")
			if len(ss) != 3 {
				return fmt.Errorf("Ingnore locale file %s", info.Name())
			}
			code := ss[0]
			lang := language.Make(ss[1])

			fd, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fd.Close()
			sc := bufio.NewScanner(fd)
			for sc.Scan() {
				line := sc.Text()
				idx := strings.Index(line, "=")
				if idx <= 0 || line[0] == '#' {
					continue
				}
				p.set(&lang, strings.TrimSpace(code+"."+line[0:idx]), strings.TrimSpace(line[idx+1:len(line)]))
			}

		}
		return nil
	}); err != nil {
		p.Logger.Println(err)
	}

	for lang := range p.Locales {
		lng := language.Make(lang)
		for _, k := range p.Provider.Keys(&lng) {
			p.Locales[lang][k] = p.Provider.Get(&lng, k)
		}
		p.Logger.Printf("Find locale %s, %d items.", lang, len(p.Locales[lang]))
	}
}

func (p *I18n) set(lng *language.Tag, code, message string) {
	lang := lng.String()
	if _, ok := p.Locales[lang]; !ok {
		p.Locales[lang] = make(map[string]string)
	}
	p.Locales[lang][code] = message
}

//Ts translate by lang
func (p *I18n) Ts(lng string, code string, args ...interface{}) string {
	l := language.Make(lng)
	return p.T(&l, code, args...)
}

//T translate by lang tag
func (p *I18n) T(lng *language.Tag, code string, args ...interface{}) string {
	lang := lng.String()
	msg := p.Provider.Get(lng, code)
	if len(msg) == 0 {
		if items, ok := p.Locales[lang]; ok {
			msg = items[code]
		}
	}
	return fmt.Sprintf(msg, args...)
}
