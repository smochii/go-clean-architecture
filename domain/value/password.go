package value

import (
	"encoding/json"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Password struct {
	LiteralBase[string]
}

func (p Password) Validate() error {
	return validation.Validate(p.v,
		validation.Required,
		validation.RuneLength(8, 64),
		validation.Match(regexp.MustCompile("[a-z]")),
		validation.Match(regexp.MustCompile("[0-9]")),
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9\\^$*.[\\]{}()?\"!@#%&/\\\\,<>':;|+_~`=+-]+$")),
	)
}

func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal("**********")
}

func NewPassword(password string) (Password, error) {
	v := Password{
		LiteralBase: LiteralBase[string]{
			v: password,
		},
	}
	if err := v.Validate(); err != nil {
		return Password{}, err
	}
	return v, nil
}
