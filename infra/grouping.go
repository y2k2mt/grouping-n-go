package infra

import (
	"context"
	"database/sql"
	e "github.com/cockroachdb/errors"
	pgx "github.com/jackc/pgx/v4/pgxpool"
	"github.com/y2k2mt/grouping-n-go/errors"
)

type PersistedGroup struct {
	Id    string
	Value string
}

type GroupingDatasource struct {
	DB *pgx.Pool
}

func (g *GroupingDatasource) GetGroup(id string) (*PersistedGroup, error) {
	ctx := context.Background()
	group := PersistedGroup{}
	err := g.DB.QueryRow(ctx, "SELECT id, value FROM groupings WHERE id = $1", id).Scan(&group.Id, &group.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.WithMessage(errors.NoGroup, id)
			//FIXME: Not ErrNoRows?
		} else if err.Error() == "no rows in result set" {
			return nil, e.WithMessage(errors.NoGroup, id)
		} else {
			return nil, e.Wrap(err, "quering one group in database")
		}
	}
	return &group, nil
}

func (g *GroupingDatasource) AddGroup(id string, value string) error {
	ctx := context.Background()
	_, err := g.DB.Exec(ctx, "INSERT INTO groupings (id,value) VALUES ($1,$2)", id, value)
	if err != nil {
		if err == sql.ErrNoRows {
			return e.WithMessage(errors.NoGroup, id)
			//FIXME: Not ErrNoRows?
		} else if err.Error() == "no rows in result set" {
			return e.WithMessage(errors.NoGroup, id)
		} else {
			return e.Wrap(err, "quering one group in database")
		}
	}
	return nil
}
