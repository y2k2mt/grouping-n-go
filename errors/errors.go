package errors

import (
	"github.com/cockroachdb/errors"
)

var (
	NoGroup = errors.New("No group for id")
)
