package example

import (
	"fmt"
	"testing"

	"github.com/jaybennett89/protoexperiment/protox"
	"github.com/stretchr/testify/assert"
)

func TestValidExample(t *testing.T) {
	example := &Example{
		Message:    "Hello World.",
		Percentage: 99,
		List:       []int32{1, 2, 3},
	}

	// do the Validate() test, expecting success
	validateTest(t, example, true)

	// do a serialization and deserialization and have them match
	data, err := protox.Marshal(example)
	assert.NoError(t, err)

	copy := &Example{}
	err = protox.Unmarshal(data, copy)
	assert.NoError(t, err)

	fmt.Printf("copy=%#v\n", copy)
	assert.Equal(t, true, assert.ObjectsAreEqualValues(example, copy))
}

func TestExampleMessagePattern(t *testing.T) {
	example := &Example{
		Message:    "-not an alphanumeric start to this message",
		Percentage: 99,
		List:       []int32{1, 2, 3},
	}
	err := validateTest(t, example, false)
	assert.Equal(t, "validation error: field 'message' did not match pattern ^[a-zA-Z].*", err.Error())
}

func TestExamplePercentageBelowMinValue(t *testing.T) {
	example := &Example{
		Message:    "Hello World.",
		Percentage: -1,
		List:       []int32{1, 2, 3},
	}
	err := validateTest(t, example, false)
	assert.Equal(t, "validation error: field 'percent' had value below minValue 1", err.Error())
}

func TestExamplePercentageAboveMaxValue(t *testing.T) {
	example := &Example{
		Message:    "Hello World.",
		Percentage: 9000,
		List:       []int32{1, 2, 3},
	}
	err := validateTest(t, example, false)
	assert.Equal(t, "validation error: field 'percent' had value above maxValue 1", err.Error())
}

func TestExampleListAboveMaxItems(t *testing.T) {
	example := &Example{
		Message:    "Hello World.",
		Percentage: 99,
		List:       []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	}
	err := validateTest(t, example, false)
	assert.Equal(t, "validation error: field 'list' had item count above maxItems 1", err.Error())
}

func TestExampleListBelowMinItems(t *testing.T) {
	example := &Example{
		Message:    "Hello World.",
		Percentage: 99,
		List:       []int32{},
	}
	err := validateTest(t, example, false)
	assert.Equal(t, "validation error: field 'list' had item count below minItems 1", err.Error())
}

// validateTest runs a Validate() on the message and asserts that success or failure to throw an error was as expected. It
// returns the error to the caller for further inspection of the error details.
func validateTest(t *testing.T, msg protox.Message, expectSuccess bool) error {
	err := msg.Validate()
	assert.Equal(t, expectSuccess, err == nil)
	return err
}
