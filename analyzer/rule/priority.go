package rule

import (
	"fmt"
	"math"
	"strconv"

	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

// Priority is the deserialized priority of rule
type Priority struct {
	val int64
}

// PriorityFromCtx gets the priority from the parsed ANTLR file.
func PriorityFromCtx(ctx cmpl.IPriorityContext) Priority {
	v := ctx.GetText()
	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		// In order not to break the whole policy/rule set, will use a default failure value here to make it as void
		return Priority{
			val: math.MaxInt64,
		}
	}
	return Priority{
		val: val,
	}
}

// String returns the serialized priority.
func (d *Priority) String() string {
	return fmt.Sprintf("%d", d.val)
}

// Val returns the priority integer value.
func (d *Priority) Val() int64 {
	return d.val
}
