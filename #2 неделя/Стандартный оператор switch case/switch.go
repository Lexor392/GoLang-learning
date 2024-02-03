package main

import "fmt"

//примеры использования оператора switch
func main() {
	// стандартное использование оператора switch
	i := "bool"
	switch i {
	case "string":
		fmt.Println("Переменная i имеет значение string")
	case "bool":
		fmt.Println("Переменная i имеет значение bool")
	default:
		fmt.Println("Переменная i имеет другое значение")
	}

	// пример использования с несколькими условиями
	mobile := "iphone"
	switch mobile {
	case "xiaomi", "samsung":
		fmt.Println("У вас или xiaomi, или samsung")
	case "nokia":
		fmt.Println("У вас nokia")
	case "oppo":
		fmt.Println("У вас oppo")
	case "poco", "asus":
		fmt.Println("У вас или poco, или asus")
	case "iphone":
		fmt.Println("У вас iphone")
	default:
		fmt.Println("У вас не известный мне телефон")
	}

	// пример использования с пустыми ветвями
	var color = "purple"
	switch color {
	case "red":
	case "geen":
	case "blue":
		fmt.Println("Этот цвет базовый")
	case "orange":
		fmt.Println("У вас революция?")
	case "purple":
	case "pink":
		fmt.Println("Вы девочка?")
	}

	//Использование fallthrough
	class := 3
	switch class {
	case 1:
		fmt.Println("Получено достижение \"Жажда мести!\"")
	case 2:
		fmt.Println("Получено достижение \"Путь жнеца!\"")
	case 3:
		fmt.Println("Получено достижение \"Кровь воина!\"")
		fallthrough
	case 4:
		fmt.Println("Получено достижение \"Убица королей!\"")
	case 5:
		fmt.Println("Получено достижение \"Уничтожитель богов!\"")
	}
}
