package main

import "fmt"

func main()  {
	var x = []int{3,44,38,5,47,15,36,26,27,44,46,38,19,2}
	x = testSlice(x)
	fmt.Println(x)
}

//冒泡排序，与相邻元素比较，如果前者小于后者,交换位置
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

//选择排序,当前元素与后面的所有元素进行比较，如果有比当前元素小的，把最小值与当前对象交换位置
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

//插入排序
func insertionSort(x []int) []int {
	length := len(x)
	for i := 0; i <length; i++ {
		preIndex := i-1
		current := x[i]
		for preIndex >= 0 && current <= x[preIndex] {
			x[preIndex+1] = x[preIndex]
			preIndex -= 1
		}
		x[preIndex+1] = current
	}

	return x
}

//快速排序,找一个基准，然后两个指针向中间移动，直至两个指针指向同一个值，然后递归执行
func quickSort(nums []int, start int, end int)  {
	if start > end {
		return
	}

	temp := nums[start]
	i := start
	j := end

	for i != j {
		for nums[j] >= temp && j > i  {
			j--
		}

		for nums[i] <= temp && j > i {
			i++
		}

		//交换位置
		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[start] = nums[i]
	nums[i] = temp

	quickSort(nums, start, i-1)
	quickSort(nums, i+1, end)
}

func testSlice(x []int) []int {
	x = append(x, 100)
	return x
}
