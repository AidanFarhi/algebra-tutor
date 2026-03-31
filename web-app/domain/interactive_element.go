package domain

import (
	"encoding/json"
	"fmt"
)

type InteractiveElement interface {
	Type() string
	Validate() error
}

type InteractiveElementRecord struct {
	Id              int
	UnitComponentId int
	Type            string
	InputConfig     json.RawMessage
	OrderIndex      int
}

func NewInteractiveElementFromRecord(record InteractiveElementRecord) (InteractiveElement, error) {
	switch record.Type {

	case "slider":
		var cfg SliderElement
		if err := json.Unmarshal(record.InputConfig, &cfg); err != nil {
			return nil, err
		}
		return cfg, cfg.Validate()

	case "graph":
		var cfg GraphElement
		if err := json.Unmarshal(record.InputConfig, &cfg); err != nil {
			return nil, err
		}
		return cfg, cfg.Validate()

	case "multiple_choice":
		var cfg MultipleChoiceElement
		if err := json.Unmarshal(record.InputConfig, &cfg); err != nil {
			return nil, err
		}
		return cfg, cfg.Validate()

	default:
		return nil, fmt.Errorf("unknown interactive element type: %s", record.Type)
	}
}
