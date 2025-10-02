package main

import (
	"fmt"
	"unicode"
)

// Уникальные символы в строке

// Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).
// Вывод: true, если все символы уникальны, false, если есть повторения.
// Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(checkUnique(str))
}

func checkUnique(s string) bool {
	set := make(map[rune]struct{})

	for _, v := range s {
		lowerChar := unicode.ToLower(v)
		if _, ok := set[lowerChar]; ok {
			return false
		}
		set[lowerChar] = struct{}{}
	}

	return true
}
