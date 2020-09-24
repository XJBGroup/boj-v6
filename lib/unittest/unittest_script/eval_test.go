package unittest_script

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEval(t *testing.T) {
	v, err := Eval(ResultEvalContext{nil}, "date.now()")
	assert.NoError(t, err)
	x := v.(time.Time).Sub(time.Now())
	assert.True(t, Near(x, time.Second), NearMsg(x, time.Second))
}

func NearMsg(x time.Duration, abs time.Duration) string {
	return fmt.Sprintf("does not near, want |%v| < %v", x, abs)
}

func Near(x time.Duration, abs time.Duration) bool {
	return x < abs
}
