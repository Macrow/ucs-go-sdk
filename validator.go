package ucs

import "errors"

type Validator struct {
	publicKey []byte
}

func (v *Validator) ValidateJwt(tokenString string) (user *JwtUser, err error) {
	if len(v.publicKey) == 0 {
		err = errors.New("please provide rsa public key")
		return
	}
	return ValidateJwt(v.publicKey, tokenString)
}
