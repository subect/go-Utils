package sliceUtil

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			addVal:    233,
			index:     1,
			wantSlice: []int{123, 233, 124, 125},
		},
		{
			name:    "index out of range",
			slice:   []int{123, 100},
			index:   12,
			wantErr: NewErrIndexOutOfRange(2, 12),
		},
		{
			name:    "index less than 0",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: NewErrIndexOutOfRange(2, -1),
		},
		{
			name:      "index last",
			slice:     []int{123, 100, 101, 102, 102, 102},
			addVal:    233,
			index:     5,
			wantSlice: []int{123, 100, 101, 102, 102, 233, 102},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := AddIndex(tc.slice, tc.addVal, tc.index)
			if err != nil {
				fmt.Printf("err:%v\n", err)
				return
			}
			fmt.Println(res)
		})
	}
}
