package bubble

import (
	"go_sort/assets"
	"reflect"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
				data: assets.Ints,
			},
			wants: wants {
				data: func() []int {
					c := assets.CopyInts()
					sort.Ints(c)
					return c
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			c := make([]int, len(tt.args.data))
			copy(c, tt.args.data)
			if got := Sort(c); !reflect.DeepEqual(got, tt.wants.data) {
				t.Errorf("bubble.Sort() = %v, want %v", got, tt.wants.data)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(assets.CopyInts())
	}
}