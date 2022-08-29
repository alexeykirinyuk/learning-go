package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum is calculating right",
			args: args{x: 1, y: 2},
			want: 3,
		},
		{
			name: "sum is calculating right #2",
			args: args{x: 5, y: 10},
			want: 15,
		},
		{
			name: "sum is calculating right #3",
			args: args{x: 100500, y: 100},
			want: 100600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcSumFromStrings(t *testing.T) {
	t.Parallel()

	t.Run("successful case", func(t *testing.T) {
		t.Parallel()

		res, err := CalcSumFromStrings("1", "10")

		assert.Nil(t, err)
		assert.Equal(t, 11, res)
	})

	t.Run("error for broken string", func(t *testing.T) {
		t.Parallel()

		res, err := CalcSumFromStrings("test", "1")

		assert.NotNil(t, err)
		t.Log(err)
		assert.Equal(t, 0, res)
	})
}
