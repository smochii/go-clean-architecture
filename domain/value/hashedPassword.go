package value

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/crypto/bcrypt"
)

type HashedPassword struct {
	LiteralBase[string]
}

func (p HashedPassword) Validate() error {
	return validation.Validate(p.v,
		validation.Required,
	)
}

func (p HashedPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal("**********")
}

func NewHashedPassword(password Password) (HashedPassword, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.String()), bcrypt.DefaultCost)
	if err != nil {
		return HashedPassword{}, err
	}

	v := HashedPassword{
		LiteralBase: LiteralBase[string]{
			v: string(hashedPassword),
		},
	}
	if err := v.Validate(); err != nil {
		return HashedPassword{}, err
	}

	return v, nil
}
