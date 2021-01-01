/*
	Parent : (i-1)/2
	Left Child: (2*parent)+1
	Right Child: (2*parent)+2

	Note: The below code sorts the array in ascending order using max-heap! For descending, min-heap should be used
*/

package sorting

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name   string
		args   args
		passed []int
	}{
		{
			name: "test-01",
			args: args{
				arr: []int{3, 14, 10, 12, 24, 6, 2, 17},
			},
			passed: []int{2, 3, 6, 10, 12, 14, 17, 24},
		},
		{
			name: "test-02",
			args: args{
				arr: []int{10, 2, 1, 3, 12, 8, 20},
			},
			passed: []int{1, 2, 3, 8, 10, 12, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.passed) {
				t.Errorf("Got:%v but want:%v", tt.args.arr, tt.passed)
			}
		})
	}
}
