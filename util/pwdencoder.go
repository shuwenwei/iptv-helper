package util

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
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
		return "", errors.New("err")
	}
	pub := rsa.PublicKey{
		N: i,
		E: 0x10001,
	}
	bytePwd, err := rsa.EncryptPKCS1v15(rand.Reader, &pub, origData)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytePwd), nil
}
