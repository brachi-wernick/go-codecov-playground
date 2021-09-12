package dummy1

import (
	"github.com/stretchr/testify/assert"
	dummy1 "playground/internal/dummy1"
	"testing"
)

func TestFoo(t *testing.T) {
	assert.EqualValues(t, 2, dummy1.Foo2(1))
}
