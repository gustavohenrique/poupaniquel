package transactions_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gustavohenrique/poupaniquel/api/database"
	"github.com/gustavohenrique/poupaniquel/api/transactions"
)

var dao = transactions.NewDao()

func TestMain(m *testing.M) {
	initial := []string{
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (1, 'Mastercard', 1000, 0, '|creditcard|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (2, 'Superstore', 300, 1, '|superstore|,|home|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (3, 'Travel', 700, 1, '|travel|,|fun|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (4, 'Gas', 50, 3, '|gas|,|car|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (5, 'Hotel', 300, 3, '|travel|,|hotel|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (6, 'Beans', 15, 2, '|superstore|,|food|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (7, 'Vacation', 500, 0, '|travel|,|fun|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (8, 'Dinner with friends', 300, 7, '|friends|,|dinner|,|fun|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (9, 'Clothes', 100, 0, '|clothes|')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags) VALUES (10, 'Restaurant', 90, 0, '|restaurant|,|food|')`,
	}

	os.Setenv("POUPANIQUEL_DB_PATH", ":memory:")
	db := database.Create()
	database.BulkInsert(db, initial)

	os.Exit(m.Run())
}

func TestFetchAll(t *testing.T) {
	err, list := dao.FetchAll(nil)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(list))
}

func TestFetchAllLimitingPerPage(t *testing.T) {
	params := make(map[string]interface{})
	params["perPage"] = 1
	err, list := dao.FetchAll(params)
	assert.Nil(t, err)

	mastercard := list[0]
	assert.Equal(t, int64(1), mastercard.Id)
	assert.Equal(t, "Mastercard", mastercard.Description)

	superstore := list[1]
	assert.Equal(t, mastercard.Id, superstore.ParentId)
	assert.Equal(t, "Superstore", superstore.Description)

	travel := list[2]
	assert.Equal(t, mastercard.Id, travel.ParentId)
	assert.Equal(t, "Travel", travel.Description)

	gas := list[3]
	assert.Equal(t, travel.Id, gas.ParentId)
	assert.Equal(t, "Gas", gas.Description)

	hotel := list[4]
	assert.Equal(t, travel.Id, hotel.ParentId)
	assert.Equal(t, "Hotel", hotel.Description)

	beans := list[5]
	assert.Equal(t, superstore.Id, beans.ParentId)
	assert.Equal(t, "Beans", beans.Description)
}

func TestFetchAllUsingPagination(t *testing.T) {
	params := make(map[string]interface{})
	params["page"] = 2
	params["perPage"] = 1
	err, list := dao.FetchAll(params)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(list))

	vacation := list[0]
	assert.Equal(t, int64(7), vacation.Id)

	dinner := list[1]
	assert.Equal(t, vacation.Id, dinner.ParentId)
}

func TestFetchAllUsingSortField(t *testing.T) {
	params := make(map[string]interface{})
	params["sort"] = "description"
	err, list := dao.FetchAll(params)
	assert.Nil(t, err)

	first := list[0]
	assert.Equal(t, int64(6), first.Id)
	assert.Equal(t, "Beans", first.Description)
}

func TestFetchOne(t *testing.T) {
	params := make(map[string]interface{})
	params["id"] = 1
	err, list := dao.FetchOne(params)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(list))
	assert.Equal(t, int64(1), list[0].Id)
	assert.Equal(t, "Mastercard", list[0].Description)
}

func TestCreate(t *testing.T) {
	raw := transactions.Raw{
		CreatedAt:   time.Now(),
		DueDate:     time.Now(),
		Type:        "expense",
		Description: "my first transaction",
		Amount:      float64(50),
		Tags:        "|t1|,|t2|",
		ParentId:    0,
	}
	err, id := dao.Create(raw)
	assert.Nil(t, err)
	assert.Equal(t, int64(11), id)
}

func TestUpdate(t *testing.T) {
	raw := transactions.Raw{
		Id:          1,
		CreatedAt:   time.Now(),
		DueDate:     time.Now(),
		Type:        "expense",
		Description: "my first transaction",
		Amount:      float64(50),
		Tags:        "|t1|,|t2|",
		ParentId:    0,
	}
	err, id := dao.Update(raw)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
}

func TestDelete(t *testing.T) {
	err := dao.Delete(map[string]interface{}{
		"id": 1,
	})
	assert.Nil(t, err)
}
