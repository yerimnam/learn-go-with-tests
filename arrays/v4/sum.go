package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// numbersToSum ... -> ... : Variadic functions
// 슬라이스에서 복수의 인자를 가지고 있다면 이걸 활용 할 수 있다. func(slice ...)
// 즉 매개변수로 슬라이스인 경우인데 다수의 값이 있다면 ...을 사용하여 표현할 수 있다!
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers) //슬라이스 생성
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
