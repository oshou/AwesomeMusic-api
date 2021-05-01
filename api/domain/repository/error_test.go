package repository

import "testing"

func TestNoRowsError_Error(t *testing.T) {
	tests := []struct {
		name string
		n    NoRowsError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NoRowsError{}
			if got := n.Error(); got != tt.want {
				t.Errorf("NoRowsError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
