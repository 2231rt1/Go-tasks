package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var dayOrNightErr = errors.New("исправь свой ответ, а лучше ложись поспать")

func currentDayOfTheWeek() string {
	switch time.Now().Weekday().String() {
	case "Monday":
		return "Понедельник"
	case "Tuesday":
		return "Вторник"
	case "Wednesday":
		return "Среда"
	case "Thursday":
		return "Четверг"
	case "Friday":
		return "Пятница"
	case "Saturday":
		return "Суббота"
	case "Sunday":
		return "Воскресенье"
	}
	return ""
}

func dayOrNight() string {
	timeHours := time.Now().Hour()
	if timeHours > 10 && timeHours < 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
    today := time.Now().Weekday()
    daysUntil := (int(time.Friday) - int(today) + 7) % 7
    return daysUntil
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	return strings.ToLower(answer) == strings.ToLower(currentDayOfTheWeek())
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, dayOrNightErr
	}
	if strings.ToLower(answer) == strings.ToLower(dayOrNight()) {
		return true, nil
	}
	return false, nil
}