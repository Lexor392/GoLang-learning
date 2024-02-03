package main

import "fmt"

// Примеры использования цикла for
func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	for i := 10; i < 100; i++ {
		fmt.Println("Расчет от 10:", i)
	}

	for i := 1; i < 10; {
		fmt.Println("В сети +", i, " пользователь")
		i++
	}

	i := 1
	for i <= 20 {
		fmt.Println("В комнату заходит человек", i)
		i++
	}
}

