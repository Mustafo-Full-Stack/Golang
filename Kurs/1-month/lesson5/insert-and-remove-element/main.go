package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	// ["a","b","d","e"]
	// удалить элемент с индексом 2 ("c")
	s = append(s[:2], s[3:]...)
	// ... -  означает что мы добавляем в конец массива через append несколько элементов
	// и без ... при компиляции ловим ошибку
	// cannot use s[3:] (value of type []string) as string value in argument to append
	fmt.Println(s)

	// вставить "Y" на позицию(индекс) 1
	//s = append(s[:1], append([]string{"Y"}, s[1:]...)...)
	s = append(s[:1], "fdfd")
	fmt.Println(s)
}
