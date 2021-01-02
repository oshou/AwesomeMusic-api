package usecase

type NotFoundError struct{}

func (err NotFoundError) Error() string {
	return "Not Found"
}

type InvalidParamError struct{}

func (err InvalidParamError) Error() string {
	return "Invalid Parameter"
}

type InternalServerError struct{}

func (err InternalServerError) Error() string {
	return "Internal Server Error"
}

type UnauthorizedError struct{}

func (err UnauthorizedError) Error() string {
	return "Unauthorized Error"
}

type ConflictError struct{}

func (err ConflictError) Error() string {
	return "Conflict"
}

type ForbiddenError struct{}

func (err ForbiddenError) Error() string {
	return "Forbidden"
}
