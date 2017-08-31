package trace

import (
	"fmt"
	"io"
)

// Tracer interface
type Tracer interface {
	Trace(...interface{})
}

// New Create new Tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
