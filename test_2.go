package main

import "fmt"


func swap(a, b *int){
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	num1, num2 :=1,2

	swap(&num1, &num2)

	fmt.Println(num1, num2)
}