package usecase

import "testing"

func TestNotFoundError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  NotFoundError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NotFoundError{}
			if got := err.Error(); got != tt.want {
				t.Errorf("NotFoundError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvalidParamError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  InvalidParamError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InvalidParamError{}
			if got := err.Error(); got != tt.want {
				t.Errorf("InvalidParamError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternalServerError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  InternalServerError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InternalServerError{}
			if got := err.Error(); got != tt.want {
				t.Errorf("InternalServerError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnauthorizedError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  UnauthorizedError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UnauthorizedError{}
			if got := err.Error(); got != tt.want {
				t.Errorf("UnauthorizedError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConflictError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  ConflictError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ConflictError{}
			if got := err.Error(); got != tt.want {
				t.Errorf("ConflictError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForbiddenError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  ForbiddenError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ForbiddenError{}
			if got := err.Error(); got != tt.want {
				t.Errorf("ForbiddenError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
