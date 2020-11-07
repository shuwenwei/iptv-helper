package util

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"math/big"
)

var PwdEncoderInstance = PwdEncoder{PublicKey: ""}

type PwdEncoder struct {
	PublicKey string
}

func (pwdEncoder *PwdEncoder) EncodePassword(origData []byte) (string, error) {
	i := new(big.Int)
	i, ok := i.SetString(pwdEncoder.PublicKey, 16)

	if !ok {
		log.Fatal("set publicKey error")
	}
	pub := rsa.PublicKey{
		N: i,
		E: 0x10001,
	}
	bytePwd, err := rsa.EncryptPKCS1v15(rand.Reader, &pub, origData)
	if err != nil {
		log.Fatal("encode password error")
	}
	return fmt.Sprintf("%x", bytePwd), nil
}
