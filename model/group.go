package model

import (
	"github.com/y2k2mt/grouping-n-go/infra"
)

type GroupResult struct {
	group *infra.Group
}

func GetGroup(id string, repository infra.GroupRepository) (GroupResult, error) {
	group, err := repository.GetGroup(id)
	if err != nil {
		return GroupResult{}, err
	}
	return GroupResult{group: group}, nil
}
