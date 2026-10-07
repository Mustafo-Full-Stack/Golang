package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

const baseURL = "https://jsonplaceholder.typicode.com/todos"

// Здесь ошибки сделаны специально. Не "чините" их — это материал для показа.
// Случаи 2 и 3 запускаются сразу. Опасные 1, 4 и 5 закомментированы:
// раскомментируйте ОДИН и запустите, остальные оставьте закрытыми.
func main() {
	// --- Ошибка 1: забыт Close ---
	// Соединения не освобождаются. Через сотни запросов: "too many open files".
	//
	// for i := 0; i < 1000; i++ {
	// 	resp, _ := http.Get(baseURL + "/1")
	// 	_ = resp // Body.Close() нет!
	// }

	// --- Ошибка 2: тело читаем дважды ---
	resp, err := http.Get(baseURL + "/1")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	defer resp.Body.Close()

	body1, _ := io.ReadAll(resp.Body)
	body2, _ := io.ReadAll(resp.Body)
	fmt.Println("2. длины:", len(body1), len(body2)) // 83 0 — второй раз пусто

	// --- Ошибка 3: не проверили статус ---
	resp2, err := http.Get(baseURL + "/99999") // такой задачи нет
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	defer resp2.Body.Close()

	var todo Todo
	err = json.NewDecoder(resp2.Body).Decode(&todo)
	fmt.Println("3. статус:", resp2.Status)                     // 404 Not Found
	fmt.Printf("3. ошибка разбора: %v, todo: %+v\n", err, todo) // <nil> и пустая структура!
	// Программа "работает", а данных нет, и непонятно почему.
	// Исправление: if resp2.StatusCode != http.StatusOK { ... }

	// --- Ошибка 4: нет таймаута ---
	// http.Get без таймаута на "мёртвый" адрес висит несколько минут.
	// Исправление: свой http.Client{Timeout: 5 * time.Second}
	//
	// fmt.Println("4. ждём...")
	// http.Get("http://10.255.255.1")

	// --- Ошибка 5: defer до проверки err ---
	// Адрес неверный -> err != nil -> resp == nil -> паника на resp.Body
	// Исправление: сначала if err != nil, потом defer.
	//
	// resp5, err := http.Get("не-url")
	// defer resp5.Body.Close()
	// fmt.Println(err)
}
