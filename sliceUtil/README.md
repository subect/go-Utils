### InSlice 判断value是否在slice中

```bazaar
func InSlice(value T, slices []T) bool
```

### IsSlice 判断指定值i是否是slice类型

```bazaar
func IsSlice(slice any) bool
```

### MergeSlice 将多个slice合并成一个slice

```bazaar
func MergeSlice(slices ...[]T) []T
```

### SumSlice 对slice中的元素求和

```bazaar
func SumSlice(slice []T) T
```

### UniqueSlice 移除slice中重复的值

```bazaar
func UniqueSlice(slice []T) []T
```

### DeleteSlice 删除 index 处的元素
```bazaar
func DeleteSlice(src []T, index int) ([]T, T, error)
```

### AddIndex 在index处添加元素
```bazaar
func AddIndex(src []T, element T, index int) ([]T, error)
```