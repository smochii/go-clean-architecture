package value

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type LiteralOnly interface {
	int64 | float64 | string
}

type LiteralBase[T LiteralOnly] struct {
	v T
}

func (b LiteralBase[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.v)
}

func (b *LiteralBase[T]) UnmarshalJSON(data []byte) error {
	var value T
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	b.v = value
	return nil
}

func (b *LiteralBase[T]) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case T:
		b.v = v
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return nil
}

func (b LiteralBase[T]) Value() (driver.Value, error) {
	return b.v, nil
}

func (b LiteralBase[T]) LiteralValue() T {
	return b.v
}

func (b LiteralBase[T]) Validate() error {
	return errors.New("not implemented")
}

func (b LiteralBase[T]) String() string {
	return fmt.Sprintf("%v", b.v)
}

func (b LiteralBase[T]) Equal(target LiteralBase[T]) bool {
	return b.v == target.v
}
