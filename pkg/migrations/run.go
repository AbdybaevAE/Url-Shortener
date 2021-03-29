package migrations

import (
	"math/rand"
	"strings"
	"time"

	"github.com/abdybaevae/url-shortener/pkg/models"
	algo_srv "github.com/abdybaevae/url-shortener/pkg/services/algo"
	"github.com/jmoiron/sqlx"
)

type migrations struct {
	db *sqlx.DB
}
type migrationFunc = func(m *migrations) error

func Run(db *sqlx.DB) error {
	m := &migrations{db: db}
	all := []migrationFunc{
		ensureBase62Algorithm,
	}
	for _, f := range all {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil

}

const (
	base64AlgoName   = string(algo_srv.BASE_62)
	numberStartValue = 10000
	incrementValue   = 500
)

func ensureBase62Algorithm(m *migrations) error {
	var count int
	if err := m.db.QueryRow("select count(*) from algorithm where name = $1", base64AlgoName).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	var numId int
	if err := m.db.QueryRow("insert into number (number_value) values ($1) returning number_id", numberStartValue).Scan(&numId); err != nil {
		return err
	}
	algo := &models.Algo{
		Strategy:       base64AlgoName,
		IncrementValue: incrementValue,
		Dict:           generateBase62Dict(),
		NumberId:       numId,
	}
	if _, err := m.db.Exec("insert into algo (algo_strategy, number_id, increment_value) values ($1, $2, $3)", algo.Strategy, algo.NumberId, algo.IncrementValue); err != nil {
		return err
	}
	return nil
}
func generateBase62Dict() string {
	var sb strings.Builder
	for i := 0; i < 26; i++ {
		if i < 10 {
			sb.WriteRune(rune(int('a') + i))
		}
		sb.WriteRune(rune(int('a') + i))
		sb.WriteRune(rune(int('A') + i))
	}
	ans := []byte(sb.String())
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(ans), func(i, j int) {
		ans[i], ans[j] = ans[j], ans[i]
	})
	return string(ans)
}
