package main

import "fmt"

func main() {
	var input string
	var count int
	one, two, three, four, five := "-", "-", "-", "-", "-"
	counter := 0
	for {
        input = ""
        count = 0
		fmt.Scanln(&input, &count)
		if input == "очередь" {
			fmt.Println("1.", one)
			fmt.Println("2.", two)
			fmt.Println("3.", three)
			fmt.Println("4.", four)
			fmt.Println("5.", five)
		} else if input == "конец" {
			fmt.Println("1.", one)
			fmt.Println("2.", two)
			fmt.Println("3.", three)
			fmt.Println("4.", four)
			fmt.Println("5.", five)
			break
		} else if input == "количество" {
			fmt.Println("Осталось свободных мест:", 5-counter)
			fmt.Println("Всего человек в очереди:", counter)
		} else if count == 1 && one == "-" {
			one = input
			counter++
		} else if count == 2 && two == "-" {
			two = input
			counter++
		} else if count == 3 && three == "-" {
			three = input
			counter++
		} else if count == 4 && four == "-" {
			four = input
			counter++
		} else if count == 5 && five == "-" {
			five = input
			counter++
		} else if input == "" {
			continue
		} else {
			if count > 5 || count < 1 {
				fmt.Println("Запись на место номер", count, "невозможна: некорректный ввод")
			} else if counter == 5 {
				fmt.Println("Запись на место номер", count, "невозможна: очередь переполнена")
			} else if count == 1 && one != "-" {
				fmt.Println("Запись на место номер", count, "невозможна: место уже занято")
			} else if count == 2 && two != "-" {
				fmt.Println("Запись на место номер", count, "невозможна: место уже занято")
			} else if count == 3 && three != "-" {
				fmt.Println("Запись на место номер", count, "невозможна: место уже занято")
			} else if count == 4 && four != "-" {
				fmt.Println("Запись на место номер", count, "невозможна: место уже занято")
			} else if count == 5 && five != "-" {
				fmt.Println("Запись на место номер", count, "невозможна: место уже занято")
			}
		}
	}
}
