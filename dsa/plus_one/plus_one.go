package plus_one

func plusOne(digits []int) []int {

	//usng int and normal conversion will cause int overflow
	//lets directly modify digits in place

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			//just increment that part and return
			digits[i] += 1
			return digits
		}

		digits[i] = 0
	}

	newSlice := []int{1}
	return append(newSlice, digits...)
}
