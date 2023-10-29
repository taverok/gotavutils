package dbtypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Json[T any] struct {
	Type T
}

func (c *Json[T]) Value() (driver.Value, error) {
	return json.Marshal(c.Type)
}

func (c *Json[T]) Scan(value any) error {
	bb, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bb, &c.Type)
}
