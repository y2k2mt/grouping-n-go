package model

import (
	"encoding/json"
	"fmt"
	e "github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/y2k2mt/grouping-n-go/errors"
	"github.com/y2k2mt/grouping-n-go/infra"
	"math/rand"
	"time"
	"unsafe"
)

type GroupId struct {
	Id string
}

type Candidates struct {
	N       int      `json:"n"`
	Members []string `json:"members"`
}

type Group struct {
	Members []string `json:"members"`
}

type IdentifiedGroups struct {
	Id     string  `json:"id"`
	Groups []Group `json:"groups"`
}

func GetGroup(id GroupId, datas infra.GroupingDatasource) (*IdentifiedGroups, error) {
	loaded, err := datas.GetGroup(id.Id)
	if err != nil {
		return nil, err
	}
	var groups []Group
	json.Unmarshal([]byte(loaded.Value), &groups)
	return &IdentifiedGroups{Id: loaded.Id, Groups: groups}, nil
}

func Grouping(candidates Candidates, datas infra.GroupingDatasource) (*IdentifiedGroups, error) {
	if candidates.N < 2 {
		return nil, e.WithMessage(errors.InsufficientGroupingNumber, fmt.Sprintf("%d", candidates.N))
	} else if candidates.N > len(candidates.Members) {
		return nil, e.WithMessage(errors.InsufficientGroupingMember, fmt.Sprintf("n=%d members=%d", candidates.N, len(candidates.Members)))
	}
	shuffled := candidates.Members
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
	groupedMembers := lo.Chunk(shuffled, (len(candidates.Members)/candidates.N)+(len(candidates.Members)%candidates.N))
	groups := lo.Map(groupedMembers, func(x []string, index int) Group {
		return Group{Members: x}
	})
	uid, _ := uuid.NewUUID()
	serialized, _ := json.Marshal(groups)
	err := datas.AddGroup(uid.String(), *(*string)(unsafe.Pointer(&serialized)))
	if err != nil {
		return nil, e.Wrap(err, "Failed to add group")
	}
	return &IdentifiedGroups{
		Id:     uid.String(),
		Groups: groups,
	}, nil
}
