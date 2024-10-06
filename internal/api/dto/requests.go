package dto

// IDWrapper is a wrapper for the ID parameter.
type IDWrapper struct {
	ID int `param:"id" validate:"required"`
}
