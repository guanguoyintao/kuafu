package estrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnyStrNotEmpty(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pass",
			args: args{strs: []string{"", "", "a"}},
			want: true,
		},
		{
			name: "fail",
			args: args{strs: []string{"", "", ""}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnyStrNotEmpty(tt.args.strs...); got != tt.want {
				t.Errorf("IsAnyStrNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllNotEmpty(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// todo: Add test cases.
		{args: args{strs: []string{"a", "b"}}, want: true},
		{args: args{strs: []string{"", "b"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsAllNotEmpty(tt.args.strs...), "IsAllNotEmpty(%v)", tt.args.strs)
		})
	}
}

func TestGetDomainByUrl(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name       string
		args       args
		wantDomain string
		wantErr    bool
	}{
		{
			name:       "empty url",
			args:       args{urlStr: ""},
			wantDomain: "",
			wantErr:    false,
		},
		{
			name:       "invalid host",
			args:       args{urlStr: "https://mktsocom/api/uc?k=v"},
			wantDomain: "mktsocom",
			wantErr:    false,
		},
		{
			name:       "invalid host 2",
			args:       args{urlStr: "https://mktsocom.com/api/uc?k=v"},
			wantDomain: "mktsocom.com",
			wantErr:    false,
		},
		{
			name:       "invalid host 3",
			args:       args{urlStr: "http://N/A"},
			wantDomain: "mktsocom.com",
			wantErr:    false,
		},
		{
			name:       "invalid host 4",
			args:       args{urlStr: "http://n/a"},
			wantDomain: "",
			wantErr:    false,
		},
		{
			name:       "domain",
			args:       args{urlStr: "mktso.com"},
			wantDomain: "mktso.com",
			wantErr:    false,
		},
		{
			name:       "not a url",
			args:       args{urlStr: "mktso"},
			wantDomain: "mktso",
			wantErr:    false,
		},
		{
			name:       "domain",
			args:       args{urlStr: "https://mktso.com/api/uc?k=v"},
			wantDomain: "mktso.com",
			wantErr:    false,
		},
		{
			name:       "sub domain",
			args:       args{urlStr: "https://api.mktso.com/api/uc?k=v"},
			wantDomain: "mktso.com",
			wantErr:    false,
		},
		{
			name:       "port",
			args:       args{urlStr: "https://api.mktso.com:8080/api/uc?k=v"},
			wantDomain: "mktso.com",
			wantErr:    false,
		},
		{
			name:       "null",
			args:       args{urlStr: ""},
			wantDomain: "",
			wantErr:    false,
		},
		{
			name:       "not found www 1",
			args:       args{urlStr: "https://vodafone.com.au/"},
			wantDomain: "vodafone.com.au",
			wantErr:    false,
		},
		{
			name:       "not found www 2",
			args:       args{urlStr: "http://gvbconservancy.co.za"},
			wantDomain: "gvbconservancy.co.za",
			wantErr:    false,
		},
		{
			name:       "not found www 3",
			args:       args{urlStr: "https://jobs.nike.com/zh/"},
			wantDomain: "nike.com",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDomain, err := GetDomainByUrl(tt.args.urlStr)
			fmt.Println(gotDomain)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDomainByUrl() error = %v, wantErr %v", err, tt.wantErr)
				panic(err)
				return
			}
			if gotDomain != tt.wantDomain {
				t.Errorf("GetDomainByUrl() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
		})
	}
}
