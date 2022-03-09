package pkg

import "testing"

func TestSum_Add(t *testing.T) {
	tests := []struct {
		name string
		s    *Sum
		want int
	}{
		{
			name: "test add func",
			s: &Sum{
				Num1: 1,
				Num2: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Add(); got != tt.want {
				t.Errorf("Sum.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
