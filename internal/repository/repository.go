package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/currency-task/internal/model"
)

type Repository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func New(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		log: log,
		db:  db,
	}
}

func (r *Repository) Save(currencies []*model.Currency, date string) {
	l := r.log.With(slog.String("method", "repository:save"))

	ctx := context.Background()

	dateObj, err := time.Parse("02.01.2006", date)
	if err != nil {
		l.Error("time", err)
	}
	formattedDate := dateObj.Format("2006-01-02")

	for _, cur := range currencies {
		_, err := r.db.Exec(ctx, `
			INSERT INTO R_CURRENCY (TITLE, CODE, VALUE, A_DATE) VALUES ($1, $2, $3, $4)
		`, cur.Title, cur.Code, cur.Value, formattedDate)
		if err != nil {
			l.Error("exec", err)
		}
	}

}

func (r *Repository) Get(date string, code string) ([]*model.Currency, error) {
	l := r.log.With(slog.String("method", "repository:get"))
	ctx := context.Background()
	currencies := []*model.Currency{}

	dateObj, err := time.Parse("02.01.2006", date)
	if err != nil {
		l.Error("time", err)
		return nil, fmt.Errorf("cannot parse time")
	}
	formattedDate := dateObj.Format("2006-01-02")

	var rows pgx.Rows
	if code != "" {
		rows, err = r.db.Query(ctx, `
            SELECT * FROM R_CURRENCY WHERE A_DATE = $1 AND CODE = $2
        `, formattedDate, code)
	} else {
		rows, err = r.db.Query(ctx, `
            SELECT * FROM R_CURRENCY WHERE A_DATE = $1
        `, formattedDate)
	}

	if err != nil {
		l.Error("query", err)
		return nil, fmt.Errorf("cannot get currencies")
	}

	defer rows.Close()

	for rows.Next() {
		var cur model.Currency
		err := rows.Scan(&cur.Id, &cur.Title, &cur.Code, &cur.Value, &cur.Date)
		if err != nil {
			l.Error("scan", err)
			return nil, fmt.Errorf("cannot get currencies")
		}

		currencies = append(currencies, &cur)
	}

	return currencies, nil

}
