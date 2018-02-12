package bubblesort

func BubbleSort(values []int)  {
	flag := true

	for i := 0; i < len(values); i++ {
		flag = true

		for j := 0; j < len(values)-1; j++ {
			if values[j] > values[i] {
				values[j], values[j +1] = values[j +1],values[i]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}






