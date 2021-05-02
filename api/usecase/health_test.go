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
	tests := []struct {
		name    string
		hu      *healthUsecase
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.hu.GetHealth(); (err != nil) != tt.wantErr {
				t.Errorf("healthUsecase.GetHealth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
