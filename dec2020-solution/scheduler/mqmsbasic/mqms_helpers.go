package mqmsbasic

import (
	"fmt"
	"strings"
)

func (q *multiQueueScheduler) queueLengths() string {
	var b strings.Builder
	for i, s := range q.state {
		b.WriteString(fmt.Sprintf("%d", s.Len()))
		if i != len(q.state)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}
