package db

import (
	"database/sql"

	"github.com/samber/lo"
)

func ScanResultAppend[T any](r *sql.Rows, c *[]T) error {
	var v T

	if err := r.Scan(&v); err != nil {
		return err
	}

	*c = append(lo.FromPtr(c), v)

	return nil
}
