package web

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"
)

//Serial serial
type Serial interface {
	From([]byte, interface{}) error
	To(interface{}) ([]byte, error)
}

//JSONSerial json serial
type JSONSerial struct {
}

//From bytes-to-object
func (p *JSONSerial) From(j []byte, o interface{}) error {
	return json.Unmarshal(j, o)
}

//To object-to-bytes
func (p *JSONSerial) To(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

//BytesSerial serial by gob
type BytesSerial struct {
}

//From bytes-to-object
func (p *BytesSerial) From(d []byte, o interface{}) error {
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(d)
	err := dec.Decode(o)
	if err != nil {
		return err
	}
	return nil
}

//To object-to-bytes
func (p *BytesSerial) To(o interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(o)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}
func init() {
	gob.Register(time.Time{})
}
