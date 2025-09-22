package main

import "fmt"

// Определение типа переменной в runtime

// Разработать программу, которая в runtime способна определить тип переменной,
//  переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать:
//  int, string, bool, chan (канал).

func main() {
	printType("Hello")
	printType(123)
	printType(true)
}

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Unknown type")
	}
}
