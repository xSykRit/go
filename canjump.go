package main

func CanJump(arr []uint) bool {
	if len(arr) == 1 {
		return true
	}
	if len(arr) == 0 {
		return false
	}
	for x := 0; x < uint(len(arr))-1; {
		if (len(arr)[x]) == 0 {
			return false
		}
		if x+arr[x] < uint(len(arr)) {
			x = x + arr[x]
		} else {
			return false
		}
	}
	return true
}
