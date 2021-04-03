package key

import (
	"database/sql"
	"fmt"
	"strings"

	typed_errors "github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) KeyRepo {
	return &repo{db: db}
}

const delete_one_by_algo_id = `
	delete from keys
	where key_id = (
		select key_id 
		from keys 
		where algo_id = $1
		limit 1
	)
	returning key_value`

func (r *repo) DeleteOne(algoId int) (string, error) {
	var value string
	if err := r.db.QueryRow(delete_one_by_algo_id, algoId).Scan(&value); err != nil {
		if err == sql.ErrNoRows {
			return "", typed_errors.NoKeys

		}
		return "", err
	}
	return value, nil
}

const bulkInsertNamedQuery = `
	insert into keys 
		(key_value, algo_id)
	values 
		(:key_value, :algo_id)
`

func (r *repo) InsertMany(keys []models.Key) error {
	var sb strings.Builder
	sb.WriteString("insert into keys (key_value, algo_id) values ")
	args := make([]interface{}, 0)
	for i := 0; i < len(keys); i++ {
		args = append(args, keys[i].Value, keys[i].AlgoId)
		sb.WriteString(fmt.Sprintf(" ($%v, $%v)", 2*i+1, 2*i+2))
		if i != len(keys)-1 {
			sb.WriteByte(',')
		}
	}
	row, err := r.db.Exec(sb.String(), args...)
	if err != nil {
		return err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return typed_errors.KeyInsertError
	}
	return nil
}
