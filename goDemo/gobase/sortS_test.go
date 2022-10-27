package main

import (
	"reflect"
	"testing"
)

func Test_maoPao(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{num: []int{1, 3, 7, 5, 4, 8, 2, 6, 14, 3}}, want: []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maoPao(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maoPao() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maoPao1(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maoPao1() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
