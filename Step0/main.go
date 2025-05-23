package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var form, name, surname, patronymic string
	var num1, num2, num3 float64

	// Получение данных
	fmt.Scanln(&form)

	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&patronymic)

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)

	// Работа с датой
	t, _ := time.Parse("02.01.2006", form)
	tAdd := t.AddDate(0, 0, 15)

	// Работа с именем
	allName := surname + " " + name + " " + patronymic

	// Работа с суммой
	allSum := num1 + num2 + num3

	Sum, dropSum := math.Modf(allSum)
	dropSum = math.Round(dropSum)

	Sum += dropSum / 100
	dropSum -= dropSum / 100

	fmt.Println("Уважаемый,", allName+",", "доводим до вашего сведения, "+
		"что бухгалтерия сформировала документы по факту выполненной вами работы."+
		"\nДата подписания договора:", tAdd.Format("02.01.2006")+".", "Просим вас подойти в офис в любое удобное для вас время в этот день."+
		"\nОбщая сумма выплат составит", math.Round(Sum), "руб.", dropSum*100, "коп.\n\nС уважением,\nГл. бух. Иванов А.Е.")
}
