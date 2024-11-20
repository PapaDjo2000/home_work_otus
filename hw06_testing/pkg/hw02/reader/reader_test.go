package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader_ReadJSON(t *testing.T) {
	filepath := "../../../json/data.json"

	emp, err := ReadJSON(filepath)
	assert.NoError(t, err)
	assert.NotNil(t, emp)

}
