package eid

import (
	"testing"
)

func TestHashID_Encode(t *testing.T) {
	type args struct {
		id uint64
	}
	hashID, err := NewHashID("test", 12)
	if err != nil {
		return
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "big int",
			args: args{
				id: 4560337509371871805,
			},
			want:    "DOMrQokEprE1V",
			wantErr: false,
		},
		{
			name: "small int",
			args: args{
				id: 3,
			},
			want:    "8k9mWoMdAMDl",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashID.Encode(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashID_Decode(t *testing.T) {
	type args struct {
		hashid string
	}
	hashID, err := NewHashID("test", 12)
	if err != nil {
		return
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "big int",
			args: args{
				hashid: "DOMrQokEprE1V",
			},
			want:    4560337509371871805,
			wantErr: false,
		},
		{
			name: "small int",
			args: args{
				hashid: "8k9mWoMdAMDl",
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashID.Decode(tt.args.hashid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
