package db

import "database/sql"

type Mapper[T any] func(func(...any) error) (T, error)

func Get[T any](
	db *sql.DB,
	query string,
	mapper Mapper[T],
	args ...any,
) (T, error) {
	row := db.QueryRow(query, args...)
	return mapper(row.Scan)
}

func All[T any](
	db *sql.DB,
	query string,
	mapper Mapper[T],
) ([]T, error) {
	rows, err := db.Query(SCAN_SELECT_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []T
	for rows.Next() {
		scan, err := mapper(rows.Scan)
		if err != nil {
			return nil, err
		}
		res = append(res, scan)
	}
	return res, nil
}

func Upsert(
	db *sql.DB,
	query string,
	args ...any,
) (int, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
