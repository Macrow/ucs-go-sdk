package ucs

import "errors"

type Validator struct {
	publicKey []byte
}

func (v *Validator) ValidateJwt(tokenString string) (ok bool, user *JwtUser, err error) {
	if len(v.publicKey) == 0 {
		err = errors.New("please provide rsa public key")
		return
	}
	_, user, err = ValidateJwt(v.publicKey, tokenString)
	ok = err == nil
	return
}
