package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/PapaDjo2000/home_work_otus/hw06_testing/pkg/hw02/types"
	"github.com/stretchr/testify/assert"
)

func TestReader_ReadJSON(t *testing.T) {
	filepath := "../../../json/data.json"
	nonexistent := "../../data.json"
	invalidfile := "../../../json/invalid.json"
	var data []types.Employee

	_, want := os.Open(nonexistent)
	_, result := ReadJSON(nonexistent)
	assert.Equal(t, want, result)

	f, err := os.Open(invalidfile)
	if err != nil {
		fmt.Print(err)
	}
	_, want = io.ReadAll(f)
	_, result = ReadJSON(invalidfile)
	assert.Nil(t, want, result)

	f, err = os.Open(filepath)
	if err != nil {
		fmt.Print(err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Print(err)
	}

	testCases := []struct {
		name string
		want []types.Employee
	}{
		{
			name: "1",
			want: data,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			test, erro := ReadJSON(filepath)
			if erro != nil {
				fmt.Print(erro)
			}
			assert.NoError(t, err)
			assert.Equal(t, tC.want, test)
		})
	}
}
