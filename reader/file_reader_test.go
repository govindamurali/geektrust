package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStrings(t *testing.T) {
	inputs, err := GetStrings("./test_files/test_input_file.txt")
	assert.Nil(t, err)
	assert.Equal(t, inputs, []string{"somecommand 1", "somecommand 2"})
}

func TestGetStrings_Invalid(t *testing.T) {
	inputs, err := GetStrings("randompath")
	assert.Error(t, err)
	assert.Nil(t, inputs)
}
