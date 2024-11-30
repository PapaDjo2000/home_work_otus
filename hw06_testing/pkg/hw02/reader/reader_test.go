package reader

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw02/types"
	"github.com/stretchr/testify/assert"
)

func TestReader_ReadJSON(t *testing.T) {
	filepath := "../../../json/data.json"
	var data []types.Employee

	f, _ := os.Open(filepath)
	bytes, _ := io.ReadAll(f)
	_ = json.Unmarshal(bytes, &data)

	testCases := []struct {
		name   string
		result []types.Employee
	}{
		{
			name:   "1",
			result: data,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			test, err := ReadJSON(filepath)
			assert.NoError(t, err)
			assert.Equal(t, tC.result, test)
		})
	}
}
