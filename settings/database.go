package settings

import (
	"github.com/itpkg/web"
	"github.com/jinzhu/gorm"
	"gopkg.in/vmihailenco/msgpack.v2"
)

//DatabaseProvider provider of gorm
type DatabaseProvider struct {
	Db  *gorm.DB
	Enc *web.Encryptor
}

//Set set
func (p *DatabaseProvider) Set(k string, v interface{}, f bool) error {
	buf, err := msgpack.Marshal(v)
	if err != nil {
		return err
	}
	if f {
		buf, err = p.Enc.Encode(buf)
		if err != nil {
			return err
		}
	}
	var m Model
	null := p.Db.Where("key = ?", k).First(&m).RecordNotFound()
	m.Key = k
	m.Val = buf
	m.Flag = f
	if null {
		err = p.Db.Create(&m).Error
	} else {
		err = p.Db.Save(&m).Error
	}
	return err
}

//Get get
func (p *DatabaseProvider) Get(k string, v interface{}) error {
	var m Model
	err := p.Db.Where("key = ?", k).First(&m).Error
	if err != nil {
		return err
	}
	if m.Flag {
		if m.Val, err = p.Enc.Decode(m.Val); err != nil {
			return err
		}
	}
	return msgpack.Unmarshal(m.Val, v)
}
