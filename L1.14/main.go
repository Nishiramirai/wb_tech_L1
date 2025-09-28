package main

import (
	"fmt"
	"reflect"
)

// Определение типа переменной в runtime

// Разработать программу, которая в runtime способна определить тип переменной,
//  переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать:
//  int, string, bool, chan (канал).

func main() {
	// Базовые типы, которые нужно распозновать (int, string, bool)
	printType(123)
	printType("Hello")
	printType(true)

	// Каналы. Определяет любой тип каналов
	chInt := make(chan int)
	chString := make(chan string)
	chBool := make(chan bool)
	chSlice := make(chan []int)

	printType(chInt)
	printType(chString)
	printType(chBool)
	printType(chSlice)

	// Пример типа, распознавание которого в задании не требовалось,
	// Следовательно напишет Unkown type
	printType(1.14)
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
		// Для каналов метод использующийся выше не подойдет, так как нужно указывать конкретный тип данных в канале.
		// Можно конечно перечислить конкретные типы, но для этого потребуется много повторяющегося кода,
		// да и все равно физически не получиться распознать каналы кастомных типов.
		// Поэтому придется воспользоваться рефлексией
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Chan {
			fmt.Printf("Chan %s\n", t.Elem())
		} else {
			fmt.Println("Unknown type")
		}
	}
}
