package domain

import (
	"errors"
)

type GraphElement struct {
	XMin float64 `json:"x_min"`
	XMax float64 `json:"x_max"`
	YMin float64 `json:"y_min"`
	YMax float64 `json:"y_max"`
}

func (g GraphElement) Type() string {
	return "graph"
}

func (g GraphElement) Validate() error {
	if g.XMin >= g.XMax {
		return errors.New("invalid x range")
	}
	if g.YMin >= g.YMax {
		return errors.New("invalid y range")
	}
	return nil
}
