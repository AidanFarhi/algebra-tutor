package domain

type InteractiveElement interface {
	GetType() string
	Validate() error
}
