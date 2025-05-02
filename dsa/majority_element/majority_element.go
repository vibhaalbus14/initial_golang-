package majority_element

func majorityElement(nums []int) int {
	//hashmap

	hashMap := map[int]int{}
	reqCount := int(len(nums) / 2)
	//fmt.Println(reqCount)
	reqVal := 0

	for _, val := range nums {
		//check if val is present in hashMap
		_, exists := hashMap[val]

		if exists {
			hashMap[val] += 1
		} else {
			hashMap[val] = 1
		}

		if hashMap[val] > reqCount {
			reqVal = val
		}

	}
	return reqVal
}
