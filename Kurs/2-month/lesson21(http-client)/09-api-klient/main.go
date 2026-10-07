package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Шаг 9. Собираем свой API-клиент: структура + конструктор + метод.
// Один метод Send умеет любой запрос: GET, POST, DELETE...

// APIClient — наш клиент. Внутри — обычный http.Client с таймаутом.
type APIClient struct {
	httpClient *http.Client
}

// NewAPIClient — конструктор: создаёт клиент сразу с таймаутом.
func NewAPIClient() *APIClient {
	return &APIClient{
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

// Params — описание одного запроса. Заполняем только нужные поля.
type Params struct {
	Method   string            // http.MethodGet, http.MethodPost, ...
	URL      string            // полный адрес
	Query    map[string]string // параметры после "?"
	Headers  map[string]string // свои заголовки
	Token    string            // токен доступа: добавит Authorization: Bearer ...
	Body     any               // что отправить (структура); nil — тела нет
	Response any               // КУДА положить ответ (указатель); nil — ответ не нужен
}

// Send выполняет запрос. ctx — контекст (урок 19), всегда первым параметром.
func (c *APIClient) Send(ctx context.Context, p Params) error {
	// 1. Адрес + query-параметры (url.Values сам всё экранирует)
	u, err := url.Parse(p.URL)
	if err != nil {
		return fmt.Errorf("неверный адрес: %w", err)
	}
	if len(p.Query) > 0 {
		q := u.Query()
		for key, value := range p.Query {
			q.Set(key, value)
		}
		u.RawQuery = q.Encode()
	}

	// 2. Тело запроса: структура -> JSON
	var body io.Reader
	if p.Body != nil {
		data, err := json.Marshal(p.Body)
		if err != nil {
			return fmt.Errorf("сериализация тела: %w", err)
		}
		body = bytes.NewBuffer(data)
	}

	// 3. Собираем запрос и заголовки
	req, err := http.NewRequestWithContext(ctx, p.Method, u.String(), body)
	if err != nil {
		return fmt.Errorf("создание запроса: %w", err)
	}
	if p.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if p.Token != "" {
		req.Header.Set("Authorization", "Bearer "+p.Token)
	}
	for key, value := range p.Headers {
		req.Header.Set(key, value)
	}

	// 4. Отправляем
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("выполнение запроса: %w", err)
	}
	defer resp.Body.Close()

	// 5. Любой статус не из 2xx — ошибка. Читаем тело: сервер часто объясняет причину
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		text, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("статус %d: %s", resp.StatusCode, string(text))
	}

	// 6. Разбираем ответ, только если нас об этом попросили
	if p.Response != nil {
		if err := json.NewDecoder(resp.Body).Decode(p.Response); err != nil {
			return fmt.Errorf("разбор ответа: %w", err)
		}
	}
	return nil
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type NewPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type CreatedPost struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	api := NewAPIClient()
	ctx := context.Background() // пустой контекст: пока без отмены и дедлайна
	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 1. Простой GET: заполняем только Method, URL и Response
	var todo Todo
	err := api.Send(reqCtx, Params{
		Method:   http.MethodGet,
		URL:      "https://jsonplaceholder.typicode.com/todos/5",
		Response: &todo,
	})
	if err != nil {
		fmt.Println("GET ошибка:", err)
		return
	}
	fmt.Printf("1. GET:   %+v\n", todo)

	// 2. GET с параметрами: .../todos?userId=1&_limit=3
	var todos []Todo
	err = api.Send(ctx, Params{
		Method:   http.MethodGet,
		URL:      "https://jsonplaceholder.typicode.com/todos",
		Query:    map[string]string{"userId": "1", "_limit": "3"},
		Response: &todos,
	})
	if err != nil {
		fmt.Println("GET ошибка:", err)
		return
	}
	fmt.Println("2. Query: получено задач:", len(todos))

	// 3. POST: теперь заполнено ещё и Body
	var created CreatedPost
	err = api.Send(reqCtx, Params{
		Method:   http.MethodPost,
		URL:      "https://jsonplaceholder.typicode.com/posts",
		Body:     NewPost{Title: "Привет", Body: "Из своего клиента", UserID: 1},
		Response: &created,
	})
	if err != nil {
		fmt.Println("POST ошибка:", err)
		return
	}
	fmt.Printf("3. POST:  создан пост #%d\n", created.ID)

	// 4. Токен: httpbin.org/bearer отвечает, только если пришёл заголовок Authorization
	var auth map[string]any
	err = api.Send(ctx, Params{
		Method:   http.MethodGet,
		URL:      "https://httpbin.org/bearer",
		Token:    "my-secret-token",
		Response: &auth,
	})
	if err != nil {
		fmt.Println("Token ошибка:", err)
		return
	}
	fmt.Println("4. Token:", auth)

	// 5. Ошибка сервера: такой задачи нет. Response не указан — ответ нам не нужен
	err = api.Send(ctx, Params{
		Method: http.MethodGet,
		URL:    "https://jsonplaceholder.typicode.com/todos/99999",
	})
	fmt.Println("5. 404:  ", err)
}
