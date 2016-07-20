package transactions

import (
	"time"
	"errors"
)

type (
	// Raw represents the data in the database
	Raw struct {
		Id              int64
		Type            string
		CreatedAt       time.Time `db:"createdAt"`
		Description     string
		Amount          float32
		Tags            string
		ParentId        int64 `db:"parentId"`
		RecursiveNumber int64 `db:"RecursiveCallNumber"`
	}

	// Transaction represents the data should be returned in response
	Transaction struct {
		Id          int64         `json:"id"`
		ParentId    int64         `json:"parentId"`
		Type        string        `json:"type"`
		CreatedAt   time.Time     `json:"createdAt"`
		Description string        `json:"description"`
		Amount      float32       `json:"amount"`
		Tags        []string      `json:"tags"`
		Children    []Transaction `json:"children"`
	}
)

func (this *Transaction) validate() error {
	hasDescription := this.Description != "" && len(this.Description) > 0
	hasType := this.Type == "expense" || this.Type == "income"
	if hasDescription && this.Amount > 0 && hasType {
		return nil
	}
	return errors.New("Invalid request data.")
}