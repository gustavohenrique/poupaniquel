package reports_test

import (
	"testing"
	"os"

	"github.com/stretchr/testify/assert"

	"github.com/gustavohenrique/poupaniquel/api/database"
	"github.com/gustavohenrique/poupaniquel/api/reports"
)

var dao = reports.NewDao()

func TestMain(m *testing.M) {
	initial := []string {
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (1, 'Mastercard', 1000, 0, '|creditcard|', '2016-01-05')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (2, 'Superstore', 300, 1, '|superstore|,|home|', '2016-01-10')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (3, 'Travel', 700, 1, '|travel|,|fun|', '2016-01-10')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (4, 'Gas', 50, 3, '|gas|,|car|', '2016-01-10')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (5, 'Hotel', 300, 3, '|travel|,|hotel|,', '2016-01-20')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (6, 'Beans', 15, 2, '|superstore|,|food|', '2016-01-17')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (7, 'Vacation', 500, 0, '|travel|,|fun|,|creditcard|', '2016-02-05')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (8, 'Dinner with friends', 300, 7, '|friends|,|dinner|,|fun|,|creditcard|', '2016-02-05')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (9, 'Clothes', 100, 0, '|clothes|,|creditcard|', '2016-03-07')`,
		`INSERT INTO transactions (id, description, amount, parentId, tags, dueDate) VALUES (10, 'Restaurant', 90, 0, '|restaurant|,|food|,|creditcard|', '2016-03-18')`,

	}

	os.Setenv("POUPANIQUEL_DB_PATH", ":memory:")
	db := database.Create()
	database.BulkInsert(db, initial)

    os.Exit(m.Run())
}

func TestGetReportByTag(t *testing.T) {
	params := map[string]interface{}{
		"type": "expense",
		"startDate": "2016-01-01",
		"endDate": "2016-04-01",
		"tag": "creditcard",
	}
	err, result := dao.ByTag(params)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))

	assert.Equal(t, "2016-01", result[0]["month"])
	assert.Equal(t, float64(1000), result[0]["amount"])
	assert.Equal(t, float64(3355), result[0]["total"])

	assert.Equal(t, "2016-02", result[1]["month"])
	assert.Equal(t, float64(800), result[1]["amount"])
	assert.Equal(t, float64(3355), result[1]["total"])

	assert.Equal(t, "2016-03", result[2]["month"])
	assert.Equal(t, float64(190), result[2]["amount"])
	assert.Equal(t, float64(3355), result[2]["total"])
}
