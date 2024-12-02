package utils

func IntContains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func MinMax(array []int) (int, int) {
	var max int
	var min int
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func IndexOf(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func RemoveFromArray(slice []int, s int) []int {
	index := IndexOf(slice, s)
	return append(slice[:index], slice[index+1:]...)
}

func RemoveFromArrayByKey(slice []int, key int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:key]...)
	return append(ret, slice[key+1:]...)
}