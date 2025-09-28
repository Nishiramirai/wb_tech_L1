package main

import "fmt"

// Собственное множество строк

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueWords := getUniqueStr(s)
	fmt.Println(uniqueWords)
}

// Возврщает слайс с уникальными словами в исходном слайсе
func getUniqueStr(s []string) []string {
	set := make(map[string]struct{})
	res := make([]string, 0)

	for _, word := range s {
		// Если слова нет в множестве, добавляем в него и итоговый слайс
		if _, ok := set[word]; !ok {
			set[word] = struct{}{}
			res = append(res, word)
		}
	}

	return res
}
