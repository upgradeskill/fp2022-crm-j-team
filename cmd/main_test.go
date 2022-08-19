package main

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_setupDatabase(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setupDatabase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setupApp(t *testing.T) {
	tests := []struct {
		name string
		want *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setupApp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupApp() = %v, want %v", got, tt.want)
			}
		})
	}
}
