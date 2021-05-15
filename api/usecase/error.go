package usecase

type InvalidParamError struct{}

func (err InvalidParamError) Error() string {
	return "Invalid Parameter"
}

type UnauthorizedError struct{}

func (err UnauthorizedError) Error() string {
	return "Unauthorized Error"
}

type ForbiddenError struct{}

func (err ForbiddenError) Error() string {
	return "Forbidden"
}

type NotFoundError struct{}

func (err NotFoundError) Error() string {
	return "Not Found"
}

type ConflictError struct{}

func (err ConflictError) Error() string {
	return "Conflict"
}

type InternalServerError struct{}

func (err InternalServerError) Error() string {
	return "Internal Server Error"
}
