package example

import (
	"fmt"
	"regexp"
)

// message Example {
// string message = 1; // `pattern: $[a-zA-Z].*`
// int32 percentage = 2; // `minValue: 0, maxValue: 100`
// repeated int32 list = 3; // `minItems: 1, maxItems: 10`
// }

const (
	ExampleMessagePattern  = "^[a-zA-Z].*"
	ExamplePercentMinValue = 1
	ExamplePercentMaxValue = 100
	ExampleListMinItems    = 1
	ExampleListMaxItems    = 10
)

func (m *Example) Validate() error {
	if !regexp.MustCompile(ExampleMessagePattern).MatchString(m.Message) {
		return fmt.Errorf("validation error: field 'message' did not match pattern %s", ExampleMessagePattern)
	}
	if m.Percentage < ExamplePercentMinValue {
		return fmt.Errorf("validation error: field 'percent' had value below minValue %d", ExamplePercentMinValue)
	}
	if m.Percentage > ExamplePercentMaxValue {
		return fmt.Errorf("validation error: field 'percent' had value above maxValue %d", ExamplePercentMinValue)
	}
	if len(m.List) < ExampleListMinItems {
		return fmt.Errorf("validation error: field 'list' had item count below minItems %d", ExamplePercentMinValue)
	}
	if len(m.List) > ExampleListMaxItems {
		return fmt.Errorf("validation error: field 'list' had item count above maxItems %d", ExamplePercentMinValue)
	}
	return nil
}
