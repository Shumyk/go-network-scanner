package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

const zeroID = 0

type (
	scanFunc      func(...any) error
	mapper[T any] func(scanFunc) (T, error)
)

func get[T any](
	db *sql.DB,
	query string,
	mapper mapper[T],
	args ...any,
) (T, error) {
	row := db.QueryRow(query, args...)
	return mapper(row.Scan)
}

func all[T any](
	db *sql.DB,
	query string,
	mapper mapper[T],
) ([]T, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve all rows")
	}
	defer rows.Close()

	var res []T
	for rows.Next() {
		scan, scanErr := mapper(rows.Scan)
		if scanErr != nil {
			return nil, scanErr
		}
		res = append(res, scan)
	}
	return res, nil
}

func upsert(
	db *sql.DB,
	query string,
	args ...any,
) (int, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return zeroID, errors.Wrap(err, "failed to upsert")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return zeroID, errors.Wrap(err, "failed to get last insert ID")
	}
	return int(id), nil
}
