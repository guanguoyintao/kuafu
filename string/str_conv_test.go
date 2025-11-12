package estrings

import (
	"fmt"
	"github.com/guanguoyintao/kuafu/math/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvBase10To62(t *testing.T) {
	type args struct {
		num uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{num: 0}, want: "t"},
		{name: "61", args: args{num: 61}, want: "P"},
		{name: "UID", args: args{num: 806600011265120}, want: "i0RY2XW8Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvBase10To62(tt.args.num); got != tt.want {
				t.Errorf("ConvBase10To62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvBase62To10(t *testing.T) {
	var i uint64
	for i = 0; i < 10; i++ {
		t.Run(fmt.Sprintf("round %d", i), func(t *testing.T) {
			num := uint64(erand.RandomRange(0, 806600011265120))
			base62 := ConvBase10To62(num)
			base10, err := ConvBase62To10(base62)
			assert.Nil(t, err)
			assert.Equal(t, num, base10)
			fmt.Println(num, base62)
		})
	}
}

func TestInvalidCode(t *testing.T) {
	invalid := "iads 9ena"
	_, err := ConvBase62To10(invalid)
	assert.NotNil(t, err)
}
