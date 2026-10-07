package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

// Шаг 12. resty — сторонняя библиотека для HTTP-запросов.
// Всё, что мы писали руками (NewRequest, Marshal, Decode, заголовки, retry),
// здесь делается одной цепочкой вызовов.
//

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Клиент создаём ОДИН раз и используем для всех запросов
	client := resty.New().
		SetBaseURL("https://jsonplaceholder.typicode.com"). // общая часть адреса
		SetTimeout(5*time.Second).                          // таймаут для каждого запроса
		SetHeader("Accept", "application/json")             // заголовок для каждого запроса

	// Хотите увидеть запрос и ответ целиком? Раскомментируйте:
	// client.SetDebug(true)

	getList(client)
	getOne(client)
	createPost(client)
	updatePost(client)
	patchPost(client)
	deletePost(client)
	notFound(client)
	withToken(client)
	withTimeout()
	withRetry()
}

// failed делает две проверки, как в наших прошлых шагах:
// 1) сам запрос не удался (err), 2) сервер ответил 4xx или 5xx (IsError).
func failed(resp *resty.Response, err error) bool {
	if err != nil {
		fmt.Println("   ошибка запроса:", err)
		return true
	}
	if resp.IsError() {
		fmt.Println("   сервер вернул:", resp.Status())
		return true
	}
	return false
}

// 1. GET со списком и query-параметром: /posts?userId=1
func getList(client *resty.Client) {
	var posts []Post

	resp, err := client.R().
		SetQueryParam("userId", "1").
		SetResult(&posts). // resty сам разберёт JSON в нашу переменную
		Get("/posts")
	if failed(resp, err) {
		return
	}

	fmt.Printf("1. GET список: статус %d, постов: %d\n", resp.StatusCode(), len(posts))
}

// 2. GET одного поста: {id} в адресе заменяется на значение
func getOne(client *resty.Client) {
	var post Post

	resp, err := client.R().
		SetPathParam("id", "1").
		SetResult(&post).
		Get("/posts/{id}")
	if failed(resp, err) {
		return
	}

	fmt.Println("2. GET один:", post.Title)
}

// 3. POST: SetBody сам превратит структуру в JSON и поставит Content-Type
func createPost(client *resty.Client) {
	var created Post

	resp, err := client.R().
		SetBody(Post{UserID: 1, Title: "Мой заголовок", Body: "Текст из Go"}).
		SetResult(&created).
		Post("/posts")
	if failed(resp, err) {
		return
	}

	fmt.Printf("3. POST: создан пост #%d, статус %d\n", created.ID, resp.StatusCode())
}

// 4. PUT: заменить пост целиком
func updatePost(client *resty.Client) {
	var updated Post

	resp, err := client.R().
		SetPathParam("id", "1").
		SetBody(Post{UserID: 1, ID: 1, Title: "Новый заголовок", Body: "Новый текст"}).
		SetResult(&updated).
		Put("/posts/{id}")
	if failed(resp, err) {
		return
	}

	fmt.Println("4. PUT:", updated.Title)
}

// 5. PATCH: изменить только часть полей — отправляем map
func patchPost(client *resty.Client) {
	var patched Post

	resp, err := client.R().
		SetPathParam("id", "1").
		SetBody(map[string]string{"title": "Изменён только заголовок"}).
		SetResult(&patched).
		Patch("/posts/{id}")
	if failed(resp, err) {
		return
	}

	fmt.Println("5. PATCH:", patched.Title)
}

// 6. DELETE: тела нет, результат нам не нужен
func deletePost(client *resty.Client) {
	resp, err := client.R().
		SetPathParam("id", "1").
		Delete("/posts/{id}")
	if failed(resp, err) {
		return
	}

	fmt.Println("6. DELETE: статус", resp.StatusCode())
}

// 7. Как и в net/http: 404 — это не ошибка Go. Статус проверяем сами.
func notFound(client *resty.Client) {
	resp, err := client.R().Get("/posts/99999")
	if err != nil {
		fmt.Println("7. ошибка запроса:", err)
		return
	}

	fmt.Println("7. 404: err =", err)          // <nil>
	fmt.Println("   статус:", resp.Status())   // 404 Not Found
	fmt.Println("   IsError:", resp.IsError()) // true
}

// 8. Токен доступа: SetAuthToken добавит заголовок Authorization: Bearer ...
// Адрес полный, поэтому BaseURL клиента не используется.
func withToken(client *resty.Client) {
	var answer map[string]any

	resp, err := client.R().
		SetAuthToken("my-secret-token").
		SetResult(&answer).
		Get("https://httpbin.org/bearer")
	if failed(resp, err) {
		return
	}

	fmt.Println("8. Токен:", answer)
}

// 9. Таймаут двумя способами. Сервер отвечает через 5 секунд — мы столько ждать не будем.
func withTimeout() {
	// Способ 1: таймаут на весь клиент
	client := resty.New().SetTimeout(2 * time.Second)

	_, err := client.R().Get("https://httpbin.org/delay/5")
	fmt.Println("9. Таймаут клиента:", err)

	// Способ 2: context (урок 19) — отмена одного конкретного запроса
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err = client.R().SetContext(ctx).Get("https://httpbin.org/delay/5")
	fmt.Println("   Таймаут context:", err)
}

// 10. Retry: resty сам повторяет запрос. Раньше мы писали это руками (шаг 11).
func withRetry() {
	client := resty.New().
		SetRetryCount(2).                         // ещё 2 попытки после первой
		SetRetryWaitTime(200 * time.Millisecond). // пауза между попытками
		SetRetryMaxWaitTime(1 * time.Second).
		AddRetryCondition(func(resp *resty.Response, err error) bool {
			// повторяем, если запрос не удался или сервер ответил 5xx
			return err != nil || resp.StatusCode() >= 500
		})

	// httpbin.org/status/503 всегда отвечает 503
	resp, err := client.R().Get("https://httpbin.org/status/503")
	if err != nil {
		fmt.Println("10. Retry: ошибка:", err)
		return
	}

	fmt.Printf("10. Retry: статус %d, попыток: %d\n", resp.StatusCode(), resp.Request.Attempt)
}
