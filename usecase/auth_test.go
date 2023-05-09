package usecase

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models/payload"
	"github.com/labstack/echo/v4"
)

func TestLoginUser(t *testing.T) {
	type args struct {
		user *payload.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoginUser(c, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		user *payload.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
