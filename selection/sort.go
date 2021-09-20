package selection

func Sort(data []int) []int {
	for i := 0; i < len(data); i++ {
		minIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIdx] {
				minIdx =  j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	return data
}