package value

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Email struct {
	LiteralBase[string]
}

func (e Email) String() string {
	return e.v
}

func (e Email) Value() (driver.Value, error) {
	return e.v, nil
}

func (e Email) Validate() error {
	return validation.Validate(e.v,
		validation.Required,
		validation.Length(6, 319),
		is.Email,
	)
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.v)
}

func (e *Email) UnmarshalJSON(data []byte) error {
	var value string

	err := json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	e.v = value

	return nil
}

func (e *Email) UnmarshalParam(src string) error {
	var value string

	if _, err := fmt.Sscan(src, &value); err != nil {
		return err
	}

	e.v = value

	return nil
}

func (e *Email) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		e.v = v
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return nil
}

func NewEmail(email string) (Email, error) {
	v := Email{
		LiteralBase: LiteralBase[string]{
			v: email,
		},
	}

	if err := v.Validate(); err != nil {
		return Email{}, err
	}

	return v, nil
}
