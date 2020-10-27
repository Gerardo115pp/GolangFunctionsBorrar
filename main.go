package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

func generadorInpares() func() uint {
	i := uint(1) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint {
		var par = i
		i += 2
		return par
	}
}

func swap(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

func main() {
	var random_number uint
	for h := 1; h <= 5; h++ {
		fmt.Printf("%d ", fibonacci(h))
	}
	fmt.Printf("\n")
	fmt.Printf("5 + 5 + 5 + 10 = %d", sum(5, 5, 5, 10))
	fmt.Printf("\n")
	generator := generadorInpares()
	random_number = generator()
	fmt.Printf("%d mod 2 = %d,", random_number, random_number%2)
	random_number = generator()
	fmt.Printf("%d mod 2 = %d,", random_number, random_number%2)
	random_number = generator()
	fmt.Printf("%d mod 2 = %d", random_number, random_number%2)
	fmt.Printf("\n")
	var a, b int = 5, 10
	fmt.Printf("a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}
