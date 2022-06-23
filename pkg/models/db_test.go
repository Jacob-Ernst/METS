package models

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestOpenDB(t *testing.T) {
	type args struct {
		dbName string
	}
	tests := []struct {
		name    string
		args    args
		wantDb  *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDb, err := OpenDB(tt.args.dbName)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDb, tt.wantDb) {
				t.Errorf("OpenDB() = %v, want %v", gotDb, tt.wantDb)
			}
		})
	}
}
