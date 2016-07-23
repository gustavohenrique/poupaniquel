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
		DueDate         time.Time `db:"dueDate"`
		Description     string
		Amount          float64
		Tags            string
		ParentId        int64 `db:"parentId"`
		RecursiveNumber int64 `db:"RecursiveCallNumber"`
	}

	// Transaction represents the data should be returned in response
	Transaction struct {
		Id          int64         `json:"id"`
		Type        string        `json:"type"`
		CreatedAt   time.Time     `json:"createdAt"`
		DueDate     time.Time     `json:"dueDate"`
		Description string        `json:"description"`
		Amount      float64       `json:"amount"`
		Tags        []string      `json:"tags"`
		ParentId    int64         `json:"parentId"`
		Children    []Transaction `json:"children"`
	}
)

func (this *Transaction) Validate() error {
	hasDescription := this.Description != "" && len(this.Description) > 0
	hasType := this.Type == "expense" || this.Type == "income"
	if hasDescription && this.Amount > 0 && hasType {
		return nil
	}
	return errors.New("Invalid request data.")
}