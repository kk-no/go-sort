package selection

import (
	"go_sort/assets"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		data []int
	}
	type wants struct {
		data []int
	}
	tests := []struct {
		args args
		wants wants
	}{
		{
			args: args{
				data: []int{5, 3, 4, 1, 2},
			},
			wants: wants {
				data: []int{1, 2, 3, 4, 5},
			},
		},
		{
			args: args{
				data: []int{5, 3, 2, 1, 4},
			},
			wants: wants{
				data: []int{1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sort(tt.args.data); !reflect.DeepEqual(got, tt.wants.data) {
				t.Errorf("selection.Sort() = %v, want %v", got, tt.wants.data)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(assets.CopyInts())
	}
}