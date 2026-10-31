# Как добавить новую фичу

Гайд на сквозном примере: добавляем **вакансии** — создать вакансию и получить её по id.

```
POST /api/v1/vacancies       создать (нужен access-токен)
GET  /api/v1/vacancies/{id}  получить по id
```

> Диаграммы ниже — в формате Mermaid. GitHub, GitLab, GoLand и VS Code рисуют их прямо в предпросмотре markdown.

## Архитектура

Клиент — Postman, сайт в браузере или мобильное приложение — для нас одинаков: он шлёт HTTP-запрос с JSON и получает JSON в ответ. Внутри API запрос проходит слои сверху вниз, а ответ возвращается снизу вверх.

```mermaid
flowchart TB
    subgraph clients["Клиенты"]
        direction LR
        postman["Postman / curl"]
        web["Web"]
        mobile["Mobile"]
    end

    subgraph api["Наш API на Go"]
        direction TB
        mw["<b>middleware</b><br/>Logging → Recover → Auth<br/><i>общее для всех ручек</i>"]
        handler["<b>handler/v1</b><br/>JSON ⇄ DTO<br/><i>разобрать → вызвать → ответить</i>"]
        service["<b>service</b><br/>бизнес-логика<br/><i>проверки, правила, решения</i>"]
        repo["<b>repository</b><br/>SQL, Redis, S3<br/><i>без бизнес-логики</i>"]
        mw --> handler --> service --> repo
    end

    pg[("<b>Postgres</b><br/>users, roles, posts,<br/>sessions, files")]
    redis[("<b>Redis</b><br/>хеши токенов")]
    s3[("<b>MinIO / S3</b><br/>объекты файлов")]

    clients -- "HTTP + JSON" --> mw
    repo --> pg
    repo --> redis
    repo --> s3
```

`models` на схеме нет: это общие структуры (`User`, `RegisterInput`), ошибки (`ErrNotFound`) и простые правила над ними (`CurrentUser.CanCreate`, `RegisterInput.Validate`), которыми пользуются все слои.

## Зоны ответственности

Аналогия — ресторан:

| Слой | В ресторане | Делает | НЕ делает |
|---|---|---|---|
| `middleware` | охранник на входе | пишет каждый запрос в лог, ловит паники, проверяет токен | бизнес-логику |
| `handler/v1` | официант | разбирает JSON, собирает одну входную структуру, вызывает сервис, отвечает JSON, ошибки — через `response` | решает, можно ли; проверяет лимиты; ходит в БД |
| `service` | повар | проверяет права, вызывает `in.Validate()`, применяет правила, вызывает репозитории, оборачивает ошибки | HTTP, JSON, SQL, логи |
| `repository` | кладовщик | один простой запрос в БД, перевод ошибок БД в ошибки `models` | бизнес-условия, логи |
| `models` | меню | структуры, общие ошибки, простые правила над ними | ввод-вывод, теги |

Официант не готовит и не ходит на склад, кладовщик не решает, что подать гостю. Каждый делает своё — поэтому любого можно заменить, не трогая остальных.

## Кто кого знает

```mermaid
flowchart LR
    handler["handler/v1"] -- "service.AuthService<br/>(интерфейс)" --> service["service/auth"]
    service -- "repository.UserRepository<br/>(интерфейс)" --> repo["repository/user"]
    repo --> pg[("Postgres")]

    models["models"]
    handler -.-> models
    service -.-> models
    repo -.-> models
```

- **Зависимости смотрят только вниз:** handler знает про service, service — про repository. Наоборот — никогда.
- **Слои связаны через интерфейсы** (`service.go`, `repository.go`). Сервису всё равно, Postgres внизу или фейк в тесте.
- **Собирает всё вместе `app.go`:** создаёт репозитории, передаёт их в сервисы, сервисы — в роутер.

## Жизненный цикл запроса

Пример — `GET /api/v1/users/me`: клиент с access-токеном спрашивает, кто он.

```mermaid
sequenceDiagram
    autonumber
    actor C as Клиент
    participant MW as middleware
    participant H as handler/v1
    participant S as service/auth
    participant R as repository
    participant DB as Postgres
    participant RDS as Redis

    C->>MW: GET /api/v1/users/me<br/>Authorization: Bearer access_token
    Note over MW: Logging засёк время<br/>Recover готов поймать панику
    MW->>S: Authenticate(token)
    Note over S: JWTManager: подпись, срок, type=access
    S->>R: TokenRepository.IsAccessValid(sessionID, token)
    R->>RDS: GET session:&lt;id&gt;:access
    RDS-->>R: sha256
    R-->>S: совпало
    S-->>MW: CurrentUser{ID, RoleID, SessionID}
    MW->>H: текущий пользователь положен в context
    H->>S: UserService.GetByID(current.ID)
    S->>R: UserRepository.GetByID(id)
    R->>DB: SELECT … FROM users WHERE id = $1
    DB-->>R: строка таблицы
    R-->>S: *models.User
    S-->>H: *models.User
    H-->>C: 200 OK {"id":1,"email":"ivan@mail.com",…}
    Note over MW: Logging: http request status=200 duration=3ms
```

По шагам:

1. **Клиент** отправляет запрос. Для API неважно, Postman это или мобильное приложение.
2. **`Logging`** запоминает время старта, **`Recover`** готовится поймать панику, если она случится ниже.
3. **`Auth`** достаёт токен и вызывает `AuthService.Authenticate`. `JWTManager` проверяет подпись, срок и `type=access`, затем `TokenRepository` сверяет хеш access с Redis. Не прошёл — сразу `401`, до ручки запрос не дойдёт.
4. **Ручка** берёт `models.CurrentUser` из контекста (`middleware.CurrentUser(ctx)`: id пользователя, id роли, текущая сессия) и вызывает сервис.
5. **Сервис** просит репозиторий найти пользователя.
6. **Репозиторий** выполняет `SELECT` и превращает строку таблицы в `models.User`.
7. Ответ поднимается обратно: сервис → ручка. **Ручка** превращает `models.User` в JSON (без `password_hash`) и отвечает `200`.
8. **`Logging`** пишет в лог одну строку: метод, путь, статус, длительность.


## Как меняются данные по пути

У каждого слоя своя структура. На границах слоёв данные перекладываются из одной в другую — поэтому смена JSON не задевает БД, а смена таблицы не задевает API.

Пример — регистрация, `POST /api/v1/auth/register`:

```mermaid
flowchart LR
    json1["<b>JSON</b><br/>{email, password,<br/>name}"]
    dto1["<b>registerRequest</b><br/>handler/v1<br/><i>json-теги</i>"]
    input["<b>RegisterInput</b><br/>models<br/><i>что прислал клиент</i>"]
    user["<b>User</b><br/>models<br/><i>email нормализован,<br/>пароль → хеш</i>"]
    sql["<b>INSERT INTO users</b><br/>repository"]

    json1 --> dto1 -- "toInput()" --> input -- "service:<br/>Validate, bcrypt" --> user --> sql
```

Обратный путь — чтение пользователя:

```mermaid
flowchart RL
    row["<b>строка users</b><br/>Postgres"]
    db["<b>dbModel</b><br/>repository<br/><i>db-теги</i>"]
    user["<b>User</b><br/>models"]
    dto["<b>userResponse</b><br/>handler/v1<br/><i>json-теги, без пароля</i>"]
    json["<b>JSON</b><br/>{id, email, name, …}"]

    row -- "RowToStructByName" --> db -- "toModel()" --> user -- "newUserResponse()" --> dto --> json
```

| Структура | Слой | Теги | Зачем |
|---|---|---|---|
| `registerRequest`, `userResponse` | handler | `json` | формат API |
| `RegisterInput`, `User` | models | нет | с ними работает бизнес-логика |
| `dbModel` | repository | `db` | формат таблицы |

## Как путешествует ошибка

Ошибка создаётся внизу, по пути вверх обрастает контекстом, а наверху превращается в HTTP-статус. Пример — `GET /api/v1/vacancies/999`, такой вакансии нет:

```mermaid
sequenceDiagram
    autonumber
    actor C as Клиент
    participant H as handler/v1
    participant S as service
    participant R as repository
    participant DB as Postgres

    C->>H: GET /api/v1/vacancies/999
    H->>S: GetByID(999)
    S->>R: GetByID(999)
    R->>DB: SELECT … WHERE id = 999
    DB-->>R: pgx.ErrNoRows
    Note over R: ошибка БД → ошибка models
    R-->>S: models.ErrNotFound
    Note over S: добавляем контекст
    S-->>H: "get vacancy: not found"
    Note over H: response.ServiceError:<br/>errors.Is(err, ErrNotFound) → 404<br/>+ строка в лог (WARN)
    H-->>C: 404 {"error":"not found"}
```

| Где | Что делает с ошибкой |
|---|---|
| repository | ошибку драйвера → ошибка `models` (`ErrNoRows` → `ErrNotFound`, `23505` → `ErrEmailTaken`) |
| service | бизнес-отказ → ошибка `models`; ошибку снизу — обернуть: `fmt.Errorf("get vacancy: %w", err)` |
| handler | `response.ServiceError` выбирает статус и пишет лог: 4xx — WARN, 5xx — ERROR |

Неизвестная ошибка (упала БД, баг) → клиенту `500 internal server error` без подробностей, а в лог — вся цепочка: `get vacancy: select vacancy by id: connection refused`.

**Логи пишут HTTP-слой** (middleware и `response.Error`) и **планировщик фоновых задач**. Сервис и репозиторий не логируют — иначе одна ошибка попала бы в лог трижды.

## Порядок работы

| Шаг | Файл |
|---|---|
| 1. Таблица | `migrations/005_vacancies.sql` |
| 2. Модели | `internal/models/vacancy.go` |
| 3. Интерфейс репозитория + поле в `Repositories` | `internal/repository/repository.go` |
| 4. Репозиторий | `internal/repository/vacancy/db_model.go`, `repository.go` |
| 5. Интерфейс сервиса + поле в `Services` | `internal/service/service.go` |
| 6. Сервис | `internal/service/vacancy/service.go`, по файлу на метод |
| 7. DTO и ручки | `internal/handler/v1/vacancy_dto.go`, `vacancy.go` |
| 8. Маршруты | `internal/handler/router.go`, функция `NewRouter` |
| 9. Сборка | `internal/app/app.go` |
| 10. Документация | `docs/openapi.yaml`, `README.md`, `docs/commands.txt` |

---

## Если фича пишет в две таблицы

Одна бизнес-операция — одна транзакция. Пример из постов: `INSERT posts` + `UPDATE files SET post_id`. Если второй запрос не прошёл, пост тоже не должен остаться.

```go
func (r *Repository) Create(ctx context.Context, p *models.Post, fileIDs []int64) (err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		// После успешного Commit откатывать уже нечего: Rollback вернёт ErrTxClosed.
		if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			err = errors.Join(err, fmt.Errorf("rollback tx: %w", rbErr))
		}
	}()

	// ... INSERT / UPDATE через tx ...

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}
```

Правила:

- `defer` с `Rollback` пишем сразу после `Begin`. Ошибку `Rollback` не глотаем: штатный `ErrTxClosed` пропускаем, остальное добавляем к результату через именованный `err`.
- `Commit` — явный и последний.
- Бизнес-проверки, которые можно сделать до транзакции (валидация полей, права), делаем до `Begin`, чтобы не держать соединение зря.
- Если нужно проверить, что обновились все строки, сравниваем `tag.RowsAffected()` и возвращаем доменную ошибку — транзакция откатится.
- Не вызываем сервисы из репозитория и не тащим `pgx.Tx` в handler.

## Если фича должна работать в фоне

Постоянная уборка или периодическая синхронизация не должны жить в handler-е или стартовать из `init()`.

1. **Бизнес-логика — в сервисе.** Например, `cleanup.Service` умеет `DeleteUnattachedFiles` и `DeleteExpiredSessions` и ничего не знает про расписание.
2. **Расписание — в `internal/cron`.** `Start()` регистрирует джоб одной строкой: `scheduler.Every(interval).SingletonMode().Do(c.cleanup)` — `SingletonMode` не даёт запускам наслаиваться. Новый шаг — ещё один вызов в `cleanup()`.
3. **Cron зависит от узкого интерфейса** `cleanupService` с нужными ему методами, объявленного прямо в пакете `cron`.
4. **Остановка.** `Stop(ctx)` ждёт текущий проход, но не дольше `HTTP_SHUTDOWN_TIMEOUT`; cron останавливается раньше, чем закроются Postgres и Redis.
5. **Сборка — в `app.go`:** `cron.NewCron(ctx, cleanupService, cfg.Cleanup)`, затем `Start()` и `defer Stop(...)`. Включается флагом `CLEANUP_ENABLED`.
6. **Тесты — на сервис, не на таймер.** Фейковые репозитории проверяют, что и в каком порядке удаляется; сам gocron отдельно тестировать не нужно.

## 1. Миграция

`migrations/005_vacancies.sql`:

```sql
CREATE TABLE IF NOT EXISTS vacancies (
    id          BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id     BIGINT      NOT NULL REFERENCES users(id),
    title       TEXT        NOT NULL,
    description TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

- Файлы выполняются по алфавиту, поэтому номер в начале имени: `001`, `002`, ...
- Миграции применяются только при **первом** создании базы. Новый файл → `make infra-reset` (данные удалятся).
- Миграции описывают **итоговую схему**: только `CREATE TABLE`, `CREATE INDEX` и справочники. Никаких `ALTER`, `DROP`, `UPDATE`, `IF NOT EXISTS`. Нужно поменять таблицу — правим её исходный файл и делаем `make infra-reset`.
- id — `BIGINT GENERATED ALWAYS AS IDENTITY`, а не устаревший `BIGSERIAL`.

## 2. Модели

Файл = сущность: сама модель и входы её сценариев рядом. Без тегов.

`internal/models/vacancy.go`:

```go
package models

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

const maxVacancyTitleRunes = 200

type Vacancy struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	CreatedAt   time.Time
}

type CreateVacancyInput struct {
	Title       string
	Description string
}

// Validate приводит поля к единому виду и проверяет их.
// Pointer receiver: метод меняет структуру (trim).
func (in *CreateVacancyInput) Validate() error {
	in.Title = strings.TrimSpace(in.Title)
	in.Description = strings.TrimSpace(in.Description)

	if n := utf8.RuneCountInString(in.Title); n == 0 || n > maxVacancyTitleRunes {
		return fmt.Errorf("%w: title must be 1 to %d characters", ErrInvalidInput, maxVacancyTitleRunes)
	}
	return nil
}
```

Правила входа:

- **Одна структура на вход.** Всё, что пришло в запросе (тело, `id` из пути, IP), — поля одной `XxxInput`. Если значений 1–2 (`id`, `limit/offset`), структура не нужна: обычные параметры, проверка — в начале метода сервиса.
- **Один `Validate()`** рядом со структурой: сначала нормализация, потом проверки. Он проверяет только то, что видно из самого входа; правила с БД, конфигом или правами — в сервисе.
- `UserID` нет в `CreateVacancyInput`: автора берём из токена (`models.CurrentUser`), а не из тела запроса — иначе можно создать вакансию от чужого имени.

Новая ошибка, которую нужно отличать от других, — в `internal/models/errors.go`. Для вакансий хватает существующих `ErrNotFound` и `ErrInvalidInput`.

## 3. Интерфейс репозитория

`internal/repository/repository.go` — интерфейс рядом с остальными и поле в `Repositories`:

```go
type Repositories struct {
	User    UserRepository
	Role    RoleRepository
	Session SessionRepository // Postgres: реестр сессий
	Token   TokenRepository   // Redis: хеши токенов
	Post    PostRepository
	File    FileRepository
	Storage ObjectStorage     // S3/MinIO
	Vacancy VacancyRepository // ← новое поле
}

type VacancyRepository interface {
	Create(ctx context.Context, v *models.Vacancy) error
	GetByID(ctx context.Context, id int64) (*models.Vacancy, error)
}
```

`Repositories` — набор всех репозиториев: в `app.go` он собирается один раз, а сервисы берут из него то, что им нужно. Репозиторий не обязан быть SQL: `Token` ходит в Redis, `Storage` — в S3.

## 4. Репозиторий

Одна сущность — один пакет. Модель БД отдельно от `models`.

`internal/repository/vacancy/db_model.go`:

```go
package vacancy

import (
	"time"

	"demo-project/internal/models"
)

type dbModel struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func (m dbModel) toModel() *models.Vacancy {
	return &models.Vacancy{
		ID:          m.ID,
		UserID:      m.UserID,
		Title:       m.Title,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
	}
}
```

`internal/repository/vacancy/repository.go`:

```go
package vacancy

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"demo-project/internal/models"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, v *models.Vacancy) error {
	err := r.db.QueryRow(ctx,
		`INSERT INTO vacancies (user_id, title, description)
		 VALUES ($1, $2, $3)
		 RETURNING id, created_at`,
		v.UserID, v.Title, v.Description,
	).Scan(&v.ID, &v.CreatedAt)
	if err != nil {
		return fmt.Errorf("insert vacancy: %w", err)
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.Vacancy, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, title, description, created_at
		 FROM vacancies
		 WHERE id = $1`,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("select vacancy by id: %w", err)
	}
	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNotFound
		}
		return nil, fmt.Errorf("scan vacancy: %w", err)
	}
	return m.toModel(), nil
}
```

Правила:
- Только SQL. Никаких `WHERE status = 'open'` и «сначала SELECT, потом INSERT» — решения принимает сервис.
- `pgx.ErrNoRows` → `models.ErrNotFound`, нарушение `UNIQUE` (код `23505`) → своя ошибка из `models` (пример — `ErrEmailTaken` в `repository/user`).
- Имена колонок в `SELECT` совпадают с тегами `db`.

## 5. Интерфейс сервиса

`internal/service/service.go` — интерфейс и поле в `Services`:

```go
type Services struct {
	Auth    AuthService
	User    UserService
	Post    PostService
	File    FileService
	Vacancy VacancyService // ← новое поле
}

type VacancyService interface {
	Create(ctx context.Context, user models.CurrentUser, in models.CreateVacancyInput) (*models.Vacancy, error)
	GetByID(ctx context.Context, id int64) (*models.Vacancy, error)
}
```

Сигнатура всегда в одном порядке: `ctx`, затем `user models.CurrentUser` (если важно, кто действует), затем вход. Текущий пользователь — отдельным параметром, не внутри `XxxInput`.

## 6. Сервис

`internal/service/vacancy/service.go` — структура и конструктор:

```go
package vacancy

import (
	"demo-project/internal/repository"
	"demo-project/internal/service"
)

var _ service.VacancyService = (*Service)(nil)

type Service struct {
	vacancies repository.VacancyRepository
}

func NewService(vacancies repository.VacancyRepository) *Service {
	return &Service{vacancies: vacancies}
}
```

Сервис зависит от **интерфейса** репозитория, а не от Postgres: в тестах вместо БД подставляется фейк.

`internal/service/vacancy/create.go` — один публичный метод — один файл:

```go
package vacancy

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

// Порядок шагов всегда один: права без БД → Validate → чтение → запись.
func (s *Service) Create(ctx context.Context, user models.CurrentUser, in models.CreateVacancyInput) (*models.Vacancy, error) {
	if !user.CanCreate() {
		return nil, models.ErrForbidden
	}

	if err := in.Validate(); err != nil {
		return nil, err
	}

	v := &models.Vacancy{
		UserID:      user.ID,
		Title:       in.Title,
		Description: in.Description,
	}
	if err := s.vacancies.Create(ctx, v); err != nil {
		return nil, fmt.Errorf("create vacancy: %w", err)
	}
	return v, nil
}
```

`internal/service/vacancy/get_by_id.go`:

```go
package vacancy

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Vacancy, error) {
	v, err := s.vacancies.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get vacancy: %w", err)
	}
	return v, nil
}
```

Правила:
- Ошибка для клиента — через ошибку из `models`: `fmt.Errorf("%w: title is required", models.ErrInvalidInput)`. По ней `response.ServiceError` выберет HTTP-статус.
- Права — методами `models.CurrentUser` (`CanCreate`, `CanEditPost`), а не сравнением ролей руками.
- Ошибку снизу — всегда обернуть контекстом: `fmt.Errorf("create vacancy: %w", err)`.
- Никаких `http.*`, JSON и логов.

## 7. DTO и ручки

`internal/handler/v1/vacancy_dto.go`:

```go
package v1

import (
	"time"

	"demo-project/internal/models"
)

type createVacancyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (req createVacancyRequest) toInput() models.CreateVacancyInput {
	return models.CreateVacancyInput{
		Title:       req.Title,
		Description: req.Description,
	}
}

type vacancyResponse struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func newVacancyResponse(v *models.Vacancy) vacancyResponse {
	return vacancyResponse{
		ID:          v.ID,
		UserID:      v.UserID,
		Title:       v.Title,
		Description: v.Description,
		CreatedAt:   v.CreatedAt,
	}
}
```

`internal/handler/v1/vacancy.go` — новые методы общего `v1.Handler`. Сервис берётся из `h.services`, поэтому ни новых структур, ни конструкторов не нужно:

```go
package v1

import (
	"net/http"

	"demo-project/internal/handler/middleware"
	"demo-project/internal/handler/response"
)

func (h *Handler) CreateVacancy(w http.ResponseWriter, r *http.Request) {
	var req createVacancyRequest
	if !decodeJSON(w, r, &req) {
		return
	}
	user := middleware.CurrentUser(r.Context())

	vacancy, err := h.services.Vacancy.Create(r.Context(), user, req.toInput())
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusCreated, newVacancyResponse(vacancy))
}

func (h *Handler) GetVacancy(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r) // сам отвечает 400 на кривой id
	if !ok {
		return
	}

	vacancy, err := h.services.Vacancy.GetByID(r.Context(), id)
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newVacancyResponse(vacancy))
}
```

Ручка всегда делает три шага: **разобрать запрос → вызвать сервис → ответить**. Если в ручке появился `if` про бизнес-правило или лимит — его место в сервисе или в `Validate`.

Хелперы разбора — в `v1/request.go`: `decodeJSON`, `parseID`, `queryInt`, `clientIP`. Если вход содержит `id` из пути, передайте его в DTO: `req.toInput(id)`.

Все ручки — методы одного `v1.Handler`, поэтому имя включает сущность: `CreateVacancy`, `GetVacancy`, а не просто `Create`.

Статусы успеха: создание — `201`, чтение — `200`, изменение и удаление без тела — `204`.

## 8. Маршруты

`internal/handler/router.go`, функция `NewRouter` — добавить маршруты:

```go
func NewRouter(services *service.Services, cfg v1.Config) http.Handler {
	// ...

	mux.Handle("POST /api/v1/vacancies", requireAuth(http.HandlerFunc(api.CreateVacancy)))
	mux.HandleFunc("GET /api/v1/vacancies/{id}", api.GetVacancy)

	// ...
}
```

- Ручка только для залогиненных — обернуть в `requireAuth(...)`.
- Публичная — `mux.HandleFunc` без обёртки.
- `{id}` в пути читается в ручке через `r.PathValue("id")`.

## 9. Сборка в app

`internal/app/app.go` — по строке в `Repositories` и `Services`:

```go
import (
	// ...
	vacancyrepo "demo-project/internal/repository/vacancy"
	vacancyservice "demo-project/internal/service/vacancy"
)

	repos := &repository.Repositories{
		// ...
		Vacancy: vacancyrepo.NewRepository(postgresPool),
	}

	services := &service.Services{
		// ...
		Vacancy: vacancyservice.NewService(repos.Vacancy),
	}

	server := &http.Server{
		Handler: handler.NewRouter(services, v1.Config{ /* ... */ }),
		// ...
	}
```

Цепочка сборки всегда одна: `Repositories → Services → NewRouter → http.Server`.

Алиасы (`vacancyrepo`, `vacancyservice`) нужны, потому что пакеты repository и service одной сущности называются одинаково.

## 10. Проверка и документация

```bash
go build ./... && go vet ./...
make infra-reset   # если добавили миграцию
make run
```

Обновить `docs/openapi.yaml` (новые пути, схемы запросов и ответов, коды ошибок), `README.md` (ручки) и `docs/commands.txt` (curl-примеры).

---

## Если нужна настройка из конфига

Например, лимит длины заголовка. Без хардкода:

1. `internal/config/config.go` — поле с тегом и значением по умолчанию:
   ```go
   type Vacancy struct {
   	MaxTitleLen int `envconfig:"VACANCY_MAX_TITLE_LEN" default:"100"`
   }
   ```
   и поле `Vacancy Vacancy` в `Config`.
2. `.env.example` и `.env` — `VACANCY_MAX_TITLE_LEN=100`.
3. Передать значение в конструктор сервиса в `app.go`: `vacancyservice.NewService(repos.Vacancy, cfg.Vacancy.MaxTitleLen)`.

## Если нужна новая ошибка

1. `internal/models/errors.go`: `ErrVacancyClosed = errors.New("vacancy is closed")`.
2. Сервис возвращает её: `return nil, models.ErrVacancyClosed`.
3. `internal/handler/response/response.go`, `ServiceError` — какой статус:
   ```go
   case errors.Is(err, models.ErrVacancyClosed):
   	Error(w, r, http.StatusConflict, models.ErrVacancyClosed.Error(), err)
   ```
   Без этого шага клиент получит `500`.

## Новая версия API (v2)

Когда формат запроса или ответа нужно изменить так, что старые клиенты сломаются:

1. Новый пакет `internal/handler/v2/` со своим `Handler`, DTO и ручками.
2. В `NewRouter`: `apiV2 := v2.NewHandler(services, cfg)` и маршруты `/api/v2/...` рядом с v1.
3. v1 не трогаем: старые клиенты продолжают работать.

Сервисы, репозитории, `middleware` и `response` — общие для всех версий.

## Чек-лист

- [ ] Миграция с номером, `make infra-reset`.
- [ ] `models`: модель и `...Input` без тегов, у входа один `Validate()` (pointer receiver: нормализация + проверки), новые ошибки — в `errors.go`.
- [ ] Repository: интерфейс и поле в `Repositories` (`repository.go`), пакет сущности с `dbModel`, только SQL.
- [ ] Service: интерфейс и поле в `Services` (`service.go`), `var _ service.XxxService = (*Service)(nil)`; сигнатура `ctx, user, in`; порядок права → `in.Validate()` → чтение → запись; ошибки обёрнуты, без логов и HTTP.
- [ ] Handler: DTO с json-тегами, методы `v1.Handler` в файле сущности, ручка — разобрать → сервис → ответить, ошибки через `response`.
- [ ] Новая ошибка → `case` в `response.ServiceError`.
- [ ] Маршруты в `NewRouter`, `requireAuth` где нужно.
- [ ] Сборка в `app.go`: строка в `Repositories` и в `Services`.
- [ ] Настройки — через `config` и `.env.example`, без хардкода.
- [ ] Если пишем в две таблицы — транзакция с `defer Rollback` и явным `Commit`.
- [ ] Юнит-тесты на правила и граничные случаи.
- [ ] `go build ./... && go vet ./... && go test ./...`, ручка проверена через curl.
- [ ] `docs/openapi.yaml`, `README.md` и `docs/commands.txt` обновлены.
