package config

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mstgnz/microservice/dto"
)

func TestReadJSON(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		r    *http.Request
		data any
	}
	var loginDto dto.LoginDTO
	var rw http.ResponseWriter
	empty := httptest.NewRequest("GET", "/test", strings.NewReader(``))
	invalid := httptest.NewRequest("GET", "/test", strings.NewReader(`{"email":"", "pass":pass}`))
	valid := httptest.NewRequest("GET", "/test", strings.NewReader(`{"email":"m@g.com", "password":"pass"}`))
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid", args: args{rw, empty, &loginDto}, wantErr: true},
		{name: "invalid", args: args{rw, invalid, &loginDto}, wantErr: true},
		{name: "valid", args: args{rw, valid, &loginDto}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ReadJSON(tt.args.w, tt.args.r, tt.args.data)
			log.Printf("PRINT %v", loginDto)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteJSON(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		httpStatus int
		data       Response
		headers    []http.Header
	}
	var rw = httptest.NewRecorder()
	var loginDto = dto.LoginDTO{Email: "mes", Password: "gen"}
	var resp = Response{false, "successful", nil, loginDto}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid", args: args{w: rw, httpStatus: 200, data: resp}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteJSON(tt.args.w, tt.args.httpStatus, tt.args.data, tt.args.headers...); (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
