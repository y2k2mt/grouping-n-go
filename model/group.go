package model

import (
	"github.com/y2k2mt/grouping-n-go/infra"
)

type GroupId struct {
	Id string
}

type GroupResult struct {
	group *infra.Group
}

func GetGroup(id GroupId, repository infra.GroupRepository) (GroupResult, error) {
	group, err := repository.GetGroup(id.Id)
	if err != nil {
		return GroupResult{}, err
	}
	return GroupResult{group: group}, nil
}
