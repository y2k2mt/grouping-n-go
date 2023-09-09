package errors

import (
	"github.com/cockroachdb/errors"
)

var (
	NoGroup                    = errors.New("No group for id")
	InsufficientGroupingNumber = errors.New("Grouping number must be over 2")
	InsufficientGroupingMember = errors.New("Grouping member must be over grouping number")
)
