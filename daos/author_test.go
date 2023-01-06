package daos

import (
	"context"
	"go-template/models"
	"reflect"
	"testing"
)

func TestCreateAuthor(t *testing.T) {
	type args struct {
		author models.Author
		ctx    context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Author
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAuthor(tt.args.author, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}
