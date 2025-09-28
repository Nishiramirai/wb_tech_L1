package main

import "strings"

// Небольшой фрагмент кода — проблемы и решение

// Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?
// Приведите корректный пример реализации.
// Вопрос: что происходит с переменной justString?

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

// Ответ: взятие среза не копирует строку, а просто возвращает указатель на начало
// взятой строки и задает нужную длину. Но указывает эта строка на ту же область памяти,
// где и лежит исходная строка. Поэтому сборщик мусора не сможет очистить исходную большую строку,
// так как она используется где-то еще.

// Чтобы это исправить, надо сделать полноценную копию строки. Сделать это можно
// вручную или с помошью stringBuilder. Я воспользовался вторым способом

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	var builder strings.Builder
	builder.Grow(100)
	builder.WriteString(v[:100])
	justString = builder.String()
}

// для демонстрации реализовал эту функцию
func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func main() {
	someFunc()
}
