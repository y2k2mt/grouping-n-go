package infra

import (
	"context"
	"database/sql"
	e "github.com/cockroachdb/errors"
	"github.com/y2k2mt/grouping-n-go/errors"
)

type Group struct {
	id    string
	value string
}

type GroupRepository struct {
}

func (g *GroupRepository) GetGroup(id string) (*Group, error) {
	ctx := context.Background()
	db := GetDatabase()
	var group *Group
	err := db.QueryRow(ctx, "SELECT id, value FROM groupings WHERE id = $1", id).Scan(&group.id, &group.value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.WithMessage(errors.NoGroup, id)
		} else {
			return nil, e.Wrap(err, "quering one group in database")
		}
	}
	return group, nil
}