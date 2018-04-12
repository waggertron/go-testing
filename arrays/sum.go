package arrays

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(slices ...[]int) (tailSums []int) {
	for _, slice := range slices {
		tailSums = append(tailSums, Sum(slice[1:]))
	}
	return
}
