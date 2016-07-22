package transactions

import (
	"strings"
	"fmt"
	"errors"
)

type TransactionManager interface {
	FetchAll(map[string]interface{}) (error, []Transaction)
	FetchOne(int) (error, Transaction)
	Delete(int) error
	Save(Transaction) (error, int64)
}

type Service struct {}

var dao *Dao

func NewService(d *Dao) TransactionManager {
	dao = d
	return &Service{}
}

func (*Service) FetchAll(params map[string]interface{}) (err error, transactions []Transaction) {
	err, list := dao.FetchAll(params)
	if err != nil {
		return err, transactions
	}
	return err, toTransaction(list)
}

func (*Service) FetchOne(id int) (error, Transaction) {
	err, list := dao.FetchOne(map[string]interface{}{
		"id": id,
	})
	if len(list) == 0 {
		message := fmt.Sprintf(`Transaction #%d not found`, id)
		err = errors.New(message)
	}
	if err != nil {
		return err, Transaction{}
	}
	return err, toTransaction(list)[0]
}

func (this *Service) Delete(id int) error {
	err := dao.Delete(map[string]interface{}{
		"id": id,
	})
	return err
}

func (*Service) Save(transaction Transaction) (error, int64) {
	rawItem := toRaw(transaction)
	rawItem.Id = transaction.Id
	if rawItem.Id > 0 {
		return dao.Update(rawItem)
	}
	return dao.Create(rawItem)
}

func toRaw(transaction Transaction) Raw {
	tags := []string{}
	for _, tag := range transaction.Tags {
		tags = append(tags, "|" + strings.Trim(tag, " ") + "|")
	}
	return Raw{
		CreatedAt: transaction.CreatedAt,
		Type: transaction.Type,
		Description: transaction.Description,
		Amount: float32(transaction.Amount),
		Tags: strings.Join(tags, ","),
		ParentId: int64(transaction.ParentId),
	}
}

func toTransaction(list []Raw) (transactions []Transaction) {
	if len(list) == 0 {
		return transactions
	}

	result := make(map[int64][]Transaction)
	majorRecursiveNumber := list[len(list)-1].RecursiveNumber
	recursiveNumber := majorRecursiveNumber

	for i := len(list)-1; i >= 0; i-- {
		rawItem := list[i]
		transaction := Transaction{
			Id: rawItem.Id,
			ParentId: rawItem.ParentId,
			CreatedAt: rawItem.CreatedAt,
			Type: rawItem.Type,
			Description: rawItem.Description,
			Amount: rawItem.Amount,
			Tags: strings.Split(strings.Replace(rawItem.Tags, "|", "", -1), ","),
			Children: []Transaction{},
		}
		if recursiveNumber == rawItem.RecursiveNumber {
			if rawItem.RecursiveNumber < majorRecursiveNumber {
				transaction = setChildrenForTransaction(result[recursiveNumber+1], transaction)
			}
			result[recursiveNumber] = push(result[recursiveNumber], transaction)
		} else {
			recursiveNumber = rawItem.RecursiveNumber
			transaction = setChildrenForTransaction(result[recursiveNumber+1], transaction)
			result[recursiveNumber] = push(result[recursiveNumber], transaction)
		}
	}
	transactions = reverse(result[1])
	return transactions
}

func push(original []Transaction, transaction Transaction) ([]Transaction) {
	if original == nil {
		return []Transaction{transaction}
	} else {
		return append(original, transaction)
	}
}

func setChildrenForTransaction(children []Transaction, transaction Transaction) (Transaction) {
	for _, child := range children {
		if child.ParentId == transaction.Id {
			transaction.Children = append(transaction.Children, child)
			// TODO: remove child from list
		}
	}
	return transaction
}

func reverse(original []Transaction) (reversed []Transaction) {
	for i := len(original)-1; i >= 0; i-- {
		reversed = append(reversed, original[i])
	}
	return reversed
}