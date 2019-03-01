package gutil

import (
	"crypto/aes"
	"encoding/hex"
)

// AESEncryptData
func AESEncryptData(keystr, src string) (string, error) {
	key := []byte(keystr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	d := []byte(src)
	d = padData(d, bs)
	r := make([]byte, len(d))
	r1 := r
	for len(d) > 0 {
		block.Encrypt(r1, d)
		d = d[bs:]
		r1 = r1[bs:]
	}
	s := hex.EncodeToString(r)
	return s, err
}

// AESDecryptData
func AESDecryptData(keystr, hexSrc string) (string, error) {
	key := []byte(keystr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	src, err := hex.DecodeString(hexSrc)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	d := []byte(src)
	r := make([]byte, len(d))
	r1 := r
	for len(d) > 0 {
		block.Decrypt(r1, d)
		d = d[bs:]
		r1 = r1[bs:]
	}
	return string(removePad(r)), nil
}

// removePad
func removePad(r []byte) []byte {
	l := len(r)
	last := int(r[l-1])
	pad := r[l-last : l]
	isPad := true
	for _, v := range pad {
		if int(v) != last {
			isPad = false
			break
		}
	}
	if !isPad {
		return r
	}
	return r[:l-last]
}

// padData
func padData(d []byte, bs int) []byte {
	padedSize := ((len(d) / bs) + 1) * bs
	pad := padedSize - len(d)
	for i := len(d); i < padedSize; i++ {
		d = append(d, byte(pad))
	}
	return d
}
