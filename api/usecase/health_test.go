//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

func TestNewHealthUsecase(t *testing.T) {
	type args struct {
		repo repository.IHealthRepository
	}
	tests := []struct {
		name string
		args args
		want IHealthUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_healthUsecase_GetHealth(t *testing.T) {
	type fields struct {
		repo repository.IHealthRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hu := &healthUsecase{
				repo: tt.fields.repo,
			}
			if err := hu.GetHealth(); (err != nil) != tt.wantErr {
				t.Errorf("healthUsecase.GetHealth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
