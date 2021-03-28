package migrations

import (
	"github.com/jmoiron/sqlx"
)

type migrations struct {
	db *sqlx.DB
}
type migrationFunc = func(m *migrations) error

func Run(db *sqlx.DB) error {
	m := &migrations{db: db}
	all := []migrationFunc{
		ensureDefaultAlgorithm,
	}
	for _, f := range all {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil

}

func ensureDefaultAlgorithm(m *migrations) error {
	return nil
}
