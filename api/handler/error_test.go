package handler

import (
	"net/http"
	"testing"
)

func Test_errorResponse(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		statusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorResponse(tt.args.w, tt.args.statusCode)
		})
	}
}

func Test_unauthorizedError(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unauthorizedError(tt.args.w)
		})
	}
}

func Test_internalServerError(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			internalServerError(tt.args.w)
		})
	}
}

func Test_badRequestError(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badRequestError(tt.args.w)
		})
	}
}

func Test_notFoundError(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notFoundError(tt.args.w)
		})
	}
}
