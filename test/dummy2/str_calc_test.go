package dummy2

import (
	"github.com/stretchr/testify/assert"
	"playground/internal/dummy2"
	"testing"
)

func TestFoo(t *testing.T) {
	assert.EqualValues(t, "hello a", dummy2.Foo("a"))
}
