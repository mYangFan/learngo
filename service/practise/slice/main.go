package main

import "fmt"

func main() {
	//data := []int{0,1,2,3,4,5}
	//
	//s := data[2:4]
	//s[0] += 100
	//s[1] += 200
	//
	//s = append(s, 300, 600)
	//fmt.Println(s)
	//fmt.Println(data)
	stack := []int{1}
	stack = stack[:len(stack) - 1]
	fmt.Println(stack)
}

func sliceTest()  {
	data := []int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(s)    //[102 203]
	fmt.Println(data) //[0 1 102 203 4 5]

	//需要注意的现象
	data1 := make([]int, 6, 7)
	for i := 0; i < 6; i++ {
		data1[i] = i
	}
	s1 := data1[2:4]
	s1 = append(s1, 10, 20, 30, 40,50) // 这个操作会影响 s1,data1 值， 因为都看同一个值
	//
	fmt.Printf("%v\n", s1)    //[0 1 10]
	fmt.Printf("%v\n", data1) //[0 1 10 3 4 5]

	data2 := make([]int, 6, 30)
	for i := 0; i < 6; i++ {
		data2[i] = i
	}
	s2 := data2[3:] //s2的值是 指向s2[0] 所以s2和data2的指针地址不一样，但是s2[1],和data2[4] 指向的都是同一个内存空间。
	s2[0] = 10
	// s2 = append(s2, 10)

	fmt.Printf("%v, %v, %v, %p\n", s2, len(s2), cap(s2), &s2[1])             //[10 4 5], 3, 27, 0xc000108110
	fmt.Printf("%v, %v, %v, %p\n", data2, len(data2), cap(data2), &data2[4]) //[0 1 2 10 4 5], 6, 30, 0xc000108110
}

func sliceTest1() {
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1) //0xc000004078

	s1 = append(s1, 1)
	fmt.Println(s1)
	fmt.Printf("%p\n", &s1) //0xc000004078

	s2 := s1
	s2 = append(s1, 3)
	fmt.Printf("%p, %p\n", &s1, &s2) //0xc000004078, 0xc000004090
	fmt.Println(s1, s2)              //[1] [1 3]

	s2 = append(s2, 4)
	fmt.Printf("%p, %p\n", &s1, &s2) //0xc000004078, 0xc000004090

	fmt.Println(s1, s2) //[1] [1 3 4]
}

func sliceTest2() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	fmt.Println(data)
	s := data[:2:3]

	fmt.Println(s)

	s = append(s, 100) // 第一次 append 不超出 s.cap 限制。

	fmt.Println(s, data)         // 不会从新分配底层数组。         [0 1 100] [0 1 100 3 4 0 0 0 0 0 0]
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。 0xc000086060 0xc000086060

	s = append(s, 200) // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。    [0 1 100 200] [0 1 100 3 4 0 0 0 0 0 0]
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。             0xc0000c0060 0xc000086060
}


//移除有序数组中的重复元素
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	n := 1
	for i := 1; i <= length; i++ {
		//有序数组，相邻元素如果一致就说明有重复元素
		if nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}

	return n

}

func isValid(s string) bool {
	n := len(s)
	pairsMap := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}

	stack := make([]byte, 0)
	//遍历字符串
	for i := 0; i < n; i++ {
		//说明有符合条件的又括号出现
		if pairsMap[s[i]] > 0 {
			//开头是又括号肯定不符合条件，又括号和左括号不匹配也不符合条件
			if len(stack) == 0 || stack[len(stack) - 1] != pairsMap[s[i]] {
				return false
			}

			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}


