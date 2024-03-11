package cue

import (
	"fmt"
)

type CUE string

// ParseToTemplateValue parse the cue script to cue.Value. It must include a valid template.
func ParseToTemplateValue(c []byte) (*Value, error) {
	// the cue script must be first, it could include the imports
	template := string(c)
	v, err := NewValue(template, "")
	if err != nil {
		return nil, fmt.Errorf("fail to parse the template:%w", err)
	}
	_, err = v.LookupValue("template")
	if err != nil {
		if v.Error() != nil {
			return nil, fmt.Errorf("the template cue is invalid:%w", v.Error())
		}
		return nil, fmt.Errorf("the template cue must include the template field:%w", err)
	}
	_, err = v.LookupValue("template", "parameter")
	if err != nil {
		return nil, fmt.Errorf("the template cue must include the template.parameter field")
	}
	return v, nil
}

// ParseToValue parse the cue script to cue.Value
func ParseToValue(c []byte) (*Value, error) {
	// the cue script must be first, it could include the imports
	template := string(c)
	v, err := NewValue(template, "")
	if err != nil {
		return nil, fmt.Errorf("fail to parse the template:%w", err)
	}
	return v, nil
}
