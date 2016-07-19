package transactions

import (
	"time"
	"errors"
)

type Raw struct {
	Id int64 `db:"id"`
	Type string `db:"type"`
	CreatedAt time.Time `db:"createdAt"`
	Description string `db:"description"`
	Amount float32 `db:"amount"`
	Tags string `db:"tags"`
	ParentId int64 `db:"parentId"`
	RecursiveNumber int64 `db:"RecursiveCallNumber"`
}

type Transaction struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"`
	Type string `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	Description string `json:"description"`
	Amount float32 `json:"amount"`
	Tags []string `json:"tags"`
	Children []Transaction `json:"children"`
}

func (this *Transaction) isValid() error {
	if this.Description != "" && len(this.Description) > 0 && this.Amount > 0 && (this.Type == "expense" || this.Type == "income") {
		return nil
	}
	return errors.New("Invalid request data.")
}