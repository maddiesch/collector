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

func ScanRowMap(r *sql.Rows) (map[string]any, error) {
	columns, err := r.Columns()
	if err != nil {
		return nil, err
	}

	valuePointers := lo.Map(make([]any, len(columns)), func(v any, _ int) any {
		return lo.ToPtr(v)
	})

	if err := r.Scan(valuePointers...); err != nil {
		return nil, err
	}

	results := make(map[string]any, len(columns))

	for i, column := range columns {
		results[column] = lo.FromPtr(valuePointers[i].(*any))
	}

	return results, nil
}
