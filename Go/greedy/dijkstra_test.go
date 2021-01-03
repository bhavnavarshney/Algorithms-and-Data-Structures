package greedy

import (
	"testing"
)

func TestDijkstraGraphCreation(t *testing.T) {
	tests := []struct {
		name string
		want map[*Node]int
	}{
		{
			name: "test-01",
			want: map[*Node]int{&Node{"a"}: 0, &Node{"b"}: 2, &Node{"c"}: 4, &Node{"d"}: 7, &Node{"e"}: 5, &Node{"f"}: 5, &Node{"g"}: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DijkstraGraphCreation()
			for key, val := range got {
				for testkey, testval := range tt.want {
					if key == testkey && val != testval {
						t.Errorf("DijkstraGraphCreation: Got={%v:%v} Want={%v:%v}", key, val, testkey, testval)
					}
				}
			}
		})
	}
}
