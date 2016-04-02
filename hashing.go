package web

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

//SaltHashing Hashing with salt
type SaltHashing struct {
	H Hashing
}

//Sum append salt
func (p *SaltHashing) Sum(b []byte, l int) (string, error) {
	s := make([]byte, l)
	if _, e := rand.Read(s); e != nil {
		return "", e
	}
	return hex.EncodeToString(s) + p.H.Sum(append(b, s...)), nil
}

//Check check
func (p *SaltHashing) Check(h string, b []byte) bool {
	buf, e := hex.DecodeString(h)
	if e == nil {
		l := len(buf) - p.H.Size()
		if l > 0 {
			s := buf[0:l]
			return hex.EncodeToString(s)+p.H.Sum(append(b, s...)) == h
		}
	}
	return false
}

//Hashing hashing
type Hashing interface {
	Sum([]byte) string
	Size() int
}

//Sha512 sha512
type Sha512 struct {
}

//Sum sum
func (p *Sha512) Sum(bs []byte) string {
	buf := sha512.Sum512(bs)
	return hex.EncodeToString(buf[:])
}

//Size size
func (p *Sha512) Size() int {
	return sha512.Size
}

//Md5 md5
type Md5 struct {
}

//Sum sum
func (p *Md5) Sum(bs []byte) string {
	buf := md5.Sum(bs)
	return hex.EncodeToString(buf[:])
}

//Size size
func (p *Md5) Size() int {
	return md5.Size
}
