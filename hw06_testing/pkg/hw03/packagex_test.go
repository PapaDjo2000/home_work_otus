package packagex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	want := ` # # #
# # # 
 # # #
`
	assert.Equal(t, want, Chess(3, 6))
}
