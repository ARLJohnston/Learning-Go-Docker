package array

func Sum(numbers []int) (sum int) {
  sum = 0
  for _, number := range numbers { //range gives Index, Value
    sum += number
  }
  return
}
