package ecrypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestVerifyPassword(t *testing.T) {
	type args struct {
		hashed string
		passwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pass",
			args: args{hashed: "$2a$10$LA6JD8TqZZKwuNtD4TgQru/EqIgc6B86nADn69wOqW6eGhd7xXaGm", passwd: "123456"},
			want: true,
		},
		{
			name: "success",
			args: args{hashed: "$2a$10$h6DGq1lT9BjuxTXpqrsSV.G8PZgM/ppezoJgtoEhYc/BPK6Iz/wDG", passwd: "a123456"},
			want: true,
		},
		{
			name: "fail",
			args: args{hashed: "$2a$10$LA6JD8TqZZKwuNtD4TgQru/EqIgc6B86nADn69wOqW6eGhd7xXaGm", passwd: "123456"},
			want: false,
		},
		{
			name: "success",
			args: args{hashed: "$2a$10$bU79Z.1sZhvNAJ1tlcJ8U.O9mh2G6aYqFX9UaVamFo824W0JzJTHu", passwd: "test123"},
			want: false,
		},
		{
			name: "andydesai108@gamil.com",
			args: args{hashed: "$2a$10$EXjSzRafpK2mRjmdzLkOa.t0Q0eQ7UuR0EjZ.z3iyRtxIRKKHJRmS", passwd: "Kiran108$"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VerifyPassword(tt.args.hashed, tt.args.passwd)
			fmt.Println(got)
			if got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPasswordWithCost(t *testing.T) {
	type args struct {
		passwd string
		cost   int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				passwd: "test123",
				cost:   bcrypt.DefaultCost,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				passwd: "000139abc",
				cost:   bcrypt.DefaultCost,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				passwd: "Kiran108$",
				cost:   bcrypt.DefaultCost,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPasswordWithCost(tt.args.passwd, tt.args.cost)
			assert.NoError(t, err)
			fmt.Println(got)
		})
	}
}
