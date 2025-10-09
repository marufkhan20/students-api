package types

type Stundent struct {
	Id    int64
	Name  string `validate:"required"`
	Age   int    `validate:"required"`
	Email string `validate:"required,email"`
}