package main

import "fmt"

func main() {
	round()
	roundTwo()
	roundThree()
}

// Примеры циклов
func cycles() {

	// Стандартный цикл
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	/* For состоит из:
	i := 0 - инициализация;
	i < 100 - условие;
	i++ пост-оператор (либо шаг итерации) */

	// Цикл без пост-оператора
	for i := 0; i < 100; {
		fmt.Println(i)
		i++
	}

	// Цикл с условием (замена цикла while)
	i := 0
	for i < 100 {
		fmt.Println(i)
		i++
	}

	//Цикл с использованием оператора break (если кол-во итераций не известно)
	for {
		if 1 == 1 {
			break
		}
	}

	// Бесконечный цикл
	for {
		fmt.Println(1)
	}
}

//Использование стандартного цикла с ветвлением if else
func round() {
	for i := 1; i < 10; i++ {
		if i == 2 {
			fmt.Println("Hello")
		} else if i == 4 {
			fmt.Println("I'm Dory!")
		} else if i > 7 && i < 9 {
			fmt.Println("Who are you?")
		} else {
			fmt.Println("...")
		}
	}
}

//Использование цикла while с ветвлением switch
func roundTwo() {
	i := 1
	for i < 15 {
		switch i {
		case 3:
			fmt.Println("Hi")
		case 5:
			fmt.Println("I'm Alex")
		case 7:
			fmt.Println("This fine!")
		case 12:
			fmt.Println("Bye")
		default:
			fmt.Println("...")
		}
		i++
	}
}

//Использование цикла с оператором break
func roundThree() {
	a := 25

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		if i == a {
			break
		}
	}
}
