package ports

type ValidationService interface {
	Struct(any) error
}
