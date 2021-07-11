package encrypt

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

//加密解密密码，验证密码

type Encryption struct {
	Message   string
	Saltbytes int
}

func NewRandomSalt(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Printf("Gnerate Salt error:%v\n", err)
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (e *Encryption) Encryption() (result string, salt string, err error) {
	sha1hash1 := sha1.New()
	salt, err = NewRandomSalt(e.Saltbytes)
	if err != nil {
		return "", "", err
	}
	//先对信息加密
	sha1hash1.Write([]byte(e.Message))
	encmeesage := hex.EncodeToString(sha1hash1.Sum(nil))

	//对信息和盐加密
	mixedmessage := encmeesage + salt
	sha1hash2 := sha1.New()
	sha1hash2.Write([]byte(mixedmessage))
	ecryped := hex.EncodeToString(sha1hash2.Sum(nil))
	return ecryped, salt, nil
}

func TestifyEncrypt(Plaintext string, salt interface{}, ciphertext interface{}) bool {
	salts := salt.(string)
	sha1hash1 := sha1.New()
	sha1hash1.Write([]byte(Plaintext))
	encmeesage := hex.EncodeToString(sha1hash1.Sum(nil))

	mixedmessage := encmeesage + salts
	sha1hash2 := sha1.New()
	sha1hash2.Write([]byte(mixedmessage))
	ecryped := hex.EncodeToString(sha1hash2.Sum(nil))
	if ecryped != ciphertext {
		return false
	}
	return true
}
