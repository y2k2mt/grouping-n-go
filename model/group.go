package model

import (
	"fmt"
	e "github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"github.com/y2k2mt/grouping-n-go/errors"
	"github.com/y2k2mt/grouping-n-go/infra"
	"math/rand"
	"time"
)

type GroupId struct {
	Id string
}

type GroupResult struct {
	group *infra.Group
}

type Candidates struct {
	N       int      `json:"n"`
	Members []string `json:"members"`
}

type Group struct {
	Members []string
}

type Groups struct {
	Groups []Group
}

type IdentifiedGroups struct {
	Id     string
	Groups Groups
}

func GetGroup(id GroupId, repository infra.GroupRepository) (GroupResult, error) {
	group, err := repository.GetGroup(id.Id)
	if err != nil {
		return GroupResult{}, err
	}
	return GroupResult{group: group}, nil
}

func Grouping(candidates Candidates) (IdentifiedGroups, error) {
	if candidates.N < 2 {
		return IdentifiedGroups{}, e.WithMessage(errors.InsufficientGroupingNumber, fmt.Sprintf("%d", candidates.N))
	} else if candidates.N > len(candidates.Members) {
		return IdentifiedGroups{}, e.WithMessage(errors.InsufficientGroupingMember, fmt.Sprintf("n=%d members=%d", candidates.N, len(candidates.Members)))
	}
	shuffled := candidates.Members
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
	groupedMembers := lo.Chunk(shuffled, (len(candidates.Members)/candidates.N)+(len(candidates.Members)%candidates.N))
	groups := Groups{
		Groups: lo.Map(groupedMembers, func(x []string, index int) Group {
			return Group{Members: x}
		}),
	}
	return IdentifiedGroups{
		Id:     "FIXME",
		Groups: groups,
	}, nil
}
