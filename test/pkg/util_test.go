package pkg

import (
	"github.com/stretchr/testify/assert"
	"playground/internal/pkg/utils"
	"testing"
)

func TestHash(t *testing.T) {
	assert.EqualValues(t, 96354, utils.HashCode("abc"))
}
