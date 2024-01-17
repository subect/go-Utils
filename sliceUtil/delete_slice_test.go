package sliceUtil

import (
	"fmt"
	"testing"
)

func TestDeleteInt(t *testing.T) {
	newSlice, val, err := DeleteSlice([]int{1, 2, 3, 4}, 2)
	fmt.Printf("新切片：%v,删掉的元素值:%v,err:%v\n", newSlice, val, err)
}

func TestDeleteStr(t *testing.T) {
	newSlice, val, err := DeleteSlice([]string{"中", "国", "人"}, 2)
	fmt.Printf("新切片：%v,删掉的元素值:%v,err:%v\n", newSlice, val, err)
}
