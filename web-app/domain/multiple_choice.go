package domain

import "errors"

type MultipleChoiceElement struct {
	Options []string `json:"options"`
}

func (m MultipleChoiceElement) Type() string {
	return "multiple_choice"
}

func (m MultipleChoiceElement) Validate() error {
	if len(m.Options) == 0 {
		return errors.New("must have at least one option")
	}
	return nil
}
