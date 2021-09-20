package insertion

func Sort(data []int) []int {
	for i := 1; i < len(data); i++ {
		t := data[i]
		j := i
		for 0 < j && t < data[j-1] {
			data[j] = data[j-1]
			j--
		}
		data[j] = t
	}
	return data
}