package car

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRank(t *testing.T) {
	p := Car{"honda", "青", 500, ""}
	result := Rank(p)
	assert.Equal(t, "middle", result)
}
