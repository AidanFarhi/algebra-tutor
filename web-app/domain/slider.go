package domain

import (
	"errors"
)

type SliderElement struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Step float64 `json:"step"`
}

func (s SliderElement) Type() string {
	return "slider"
}

func (s SliderElement) Validate() error {
	if s.Min >= s.Max {
		return errors.New("min must be less than max")
	}
	if s.Step <= 0 {
		return errors.New("step must be positive")
	}
	return nil
}
