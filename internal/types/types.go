package types

type Student struct {
	Id    string `validate:"required"`
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}

type Response struct {
	Success bool
	Message string
}
