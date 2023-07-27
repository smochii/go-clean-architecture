package value

import (
	"database/sql/driver"

	"github.com/gofrs/uuid/v5"
)

type PrimaryIdBase struct {
	uuid.UUID
	value string
}

func (e PrimaryIdBase) Value() (driver.Value, error) {
	return e.value, nil
}

func (e *PrimaryIdBase) Scan(src interface{}) error {
	switch src := src.(type) {
	case uuid.UUID:
		e.UUID = src
		e.value = src.String()
		return nil

	case []byte:
		u := uuid.UUID{}
		if len(src) == uuid.Size {
			if err := u.UnmarshalBinary(src); err != nil {
				return err
			}
		}
		if err := u.UnmarshalText(src); err != nil {
			return err
		}
		e.UUID = u
		e.value = u.String()
		return nil
	case string:
		uu, err := uuid.FromString(src)
		e.UUID = uu
		e.value = uu.String()
		return err
	}
	return nil
}

func newPrimaryIdBase() PrimaryIdBase {
	u := uuid.Must(uuid.NewV7())
	return PrimaryIdBase{
		UUID:  u,
		value: u.String(),
	}
}

func newPrimaryIdBaseFromString(id string) PrimaryIdBase {
	u := uuid.Must(uuid.FromString(id))
	return PrimaryIdBase{
		UUID:  u,
		value: u.String(),
	}
}
