package models

type User struct {
	ID   int
	Name string
}

// Поломка (не делай так!): если сюда добавить
//   import "myapp/storage"
// а storage при этом импортирует models (см. storage/storage.go),
// получим при go run .:
//   import cycle not allowed
// Правило простое: если A импортирует B, то B не может импортировать A.
//
// Как чинится:
//  1) вынести общее в третий пакет (например, myapp/types) - и models,
//     и storage импортируют только его;
//  2) либо завязаться на интерфейс, а не на конкретный тип другого пакета.
