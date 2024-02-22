package main

import "fmt"

func main() {
	round(10)              /*Вывод в консоль определенных значений, пройдя через цикл, который имеет максимальное значение 10*/
	fmt.Println(add(2, 4)) /*Вывод суммы двох числ через return*/
}

func round(maxCount int) { /*Функция с циклом*/
	for i := 0; i < maxCount; i++ { /*Цикл, число срабатывания котрого зависит от переданного значения в функцию*/
		fmt.Println(i)
		ifElse(i, "this is ifElse - ")        /*Ветвление if else, которое выводит определенные данные*/
		switchCase(i, "this is swithCase - ") /*Ветвление switch case, которое выводит определенные данные*/
	}
}

func ifElse(i int, title string) { /*Функция, в которой срабатывает ветвление if else*/
	if i == 2 {
		fmt.Println(title, i)
	} else if i == 4 {
		fmt.Println(title, i)
	}
}

func switchCase(i int, title string) { /*Функция, в которой срабатывает ветвление swith case*/
	switch i {
	case 1:
	case 2:
		fmt.Println(title, i)
	case 4:
		fmt.Println(title, i)
	}
}

func add(a, b int) int { /*Функция, в которой суммируется два числа и возвращается через return*/
	return a + b
}

type Database struct {
	wight float32
	height float32
	length float32
}

