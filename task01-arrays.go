package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	for i := 0; i < 15; i++ {
		sum += input[i]
	}
	result = sum/15
	return result
} 