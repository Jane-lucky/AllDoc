package main

//从小到大
func maoPao(num []int) []int {
	for i := 1; i < len(num); i++ {
		for j := 0; j < len(num)-i; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	return num
}

func maoPao1(num []int) []int {
	for i := 1; i < len(num); i++ {
		flag := true
		for j := 0; j < len(num)-i; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return num
}

func quickSort(num []int) []int {
	if len(num) <= 1 {
		return num
	}

	splitData := num[0]

	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	mid = append(mid, splitData)
	for i := 1; i < len(num); i++ {
		if num[i] < splitData {
			low = append(low, num[i])
		} else if num[i] > splitData {
			high = append(high, num[i])
		} else {
			mid = append(mid, num[i])
		}
	}
	low, high = quickSort(low), quickSort(high)
	myarr := append(append(low, mid...), high...)
	return myarr
}
