package db

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestGetDB(t *testing.T) {
	tests := []struct {
		name string
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
