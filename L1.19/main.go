package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на вход строку.

// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func main() {
	var str string
	fmt.Scan(&str)

	strRev := reverseStr(str)
	fmt.Println(strRev)
}

func reverseStr(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
