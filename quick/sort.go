package quick

func Sort(data []int) []int {
	return quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, left, right int) []int {
	i := left + 1
	j := right
	for i < j {
		for data[i] < data[left] && i < right {
			i++
		}
		for data[left] <= data[j] && left < j {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	if data[left] > data[j] {
		data[left], data[j] = data[j], data[left]
	}
	if left < j-1 {
		quickSort(data, left, j-1)
	}
	if j+1 < right {
		quickSort(data, j+1, right)
	}
	return data
}
