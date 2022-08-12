package rule

import (
	"github.com/8treenet/iris/v12/context"
)

// Rule a superset of validators
type Rule interface {
	Claim(ctx context.Context) bool
	Valid(ctx context.Context) bool
}
