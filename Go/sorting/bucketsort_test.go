/*
Bucket sort is mainly useful when input is uniformly distributed over a range. (Floating point nos)
*/

package sorting

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "test-01",
			args: args{
				arr: []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434},
			},
			want: []float64{0.1234, 0.3434, 0.565, 0.656, 0.665, 0.897},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BucketSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("Got:%v but want:%v", tt.args.arr, tt.want)
			}
		})
	}
}
