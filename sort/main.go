package main

func main()  {

}

//冒泡排序，与相邻元素比较，如果前者小于后者
func bubbleSort(x []int) []int {
	l := len(x)
	for i := 0; i < l - 1; i++ {
		for j := 0; j < l-i-1; j++ {
			if x[j] < x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}

	return x
}

//选择排序
func selectionSort(x []int) []int {
	length := len(x)
	for i := 0; i < length; i++ {
		min := i
		for j := i+1; j < length; j ++ {
			if x[min] > x[j] {
				min = j
			}
		}
		x[min], x[i] = x[i], x[min]
	}
	return x
}
