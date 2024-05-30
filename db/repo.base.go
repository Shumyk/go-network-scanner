package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

type (
	scanFunc      func(...any) error
	mapper[T any] func(scanFunc) (T, error)
)

// RepoBasic is a repository with basic methods.
type RepoBasic[T any] interface {
	All() ([]*T, error)
	Upsert(entity *T) (int, error)
}

// Repo is a repository with all methods.
type Repo[T any] interface {
	Get(name string) (*T, error)
	RepoBasic[T]
}

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

	//nolint:prealloc // We don't know the number of rows.
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
		return 0, errors.Wrap(err, "failed to upsert")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get last insert ID")
	}
	return int(rowsAffected), nil
}
