package web

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

//Encryptor encryptor
type Encryptor interface {
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

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

//Aes aes
type Aes struct {
	//16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法
	Cip cipher.Block
}

//Encode encrypt
func (p *Aes) Encode(pn []byte) ([]byte, error) {

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(p.Cip, iv)
	ct := make([]byte, len(pn))
	cfb.XORKeyStream(ct, pn)

	return append(ct, iv...), nil

}

//Decode decrypt
func (p *Aes) Decode(sr []byte) ([]byte, error) {
	bln := len(sr)
	cln := bln - aes.BlockSize
	ct := sr[0:cln]
	iv := sr[cln:bln]

	cfb := cipher.NewCFBDecrypter(p.Cip, iv)
	pt := make([]byte, cln)
	cfb.XORKeyStream(pt, ct)
	return pt, nil
}
