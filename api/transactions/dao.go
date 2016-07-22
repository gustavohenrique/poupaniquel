package transactions

import (
	"fmt"
	"strings"

	"github.com/gustavohenrique/poupaniquel/api/database"
)

type Dao struct {}

func NewDao() *Dao {
	return &Dao{}
}

func (this *Dao) FetchAll(params map[string]interface{}) (err error, list []Raw) {
	if params == nil {
		params = make(map[string]interface{})
		params["sort"] = "id"
		params["page"] = 1
		params["perPage"] = 10
	}
	page := this.getOrDefault(params, "page", 1).(int)
	perPage := this.getOrDefault(params, "perPage", 10).(int)
	offset := (page - 1) * perPage
	sort := this.getOrDefault(params, "sort", "id").(string)
	ordering := "ASC"
	if len(sort) > 0 && sort[:1] == "-" {
		ordering = "DESC"
		sort = strings.Replace(sort, "-", "", -1)
	}

	db := database.Connect()
	query := fmt.Sprintf(`WITH CTE AS (
		SELECT *, 1 RecursiveCallNumber FROM transactions WHERE id
			IN (SELECT id FROM transactions WHERE parentId IS NULL OR parentId = 0 LIMIT %d OFFSET %d)
		UNION ALL
		SELECT  t.*, RecursiveCallNumber+1 RecursiveCallNumber FROM transactions t
		INNER JOIN CTE ON t.parentId=CTE.id)
		SELECT * FROM CTE ORDER BY RecursiveCallNumber, %s %s`, perPage, offset, sort, ordering)

	err = db.Select(&list, query)
	return err, list
}

func (this *Dao) FetchOne(params map[string]interface{}) (err error, list []Raw) {
	db := database.Connect()
	query := fmt.Sprintf(`WITH CTE AS (
		SELECT *, 1 RecursiveCallNumber FROM transactions WHERE id
			IN (SELECT id FROM transactions WHERE id = %d LIMIT 1)
		UNION ALL
		SELECT  t.*, RecursiveCallNumber+1 RecursiveCallNumber FROM transactions t
		INNER JOIN CTE ON t.parentId=CTE.id)
		SELECT * FROM CTE ORDER BY RecursiveCallNumber`, params["id"].(int))

	err = db.Select(&list, query)
	return err, list
}

func (this *Dao) Delete(params map[string]interface{}) error {
	id := params["id"].(int)
	query := fmt.Sprintf(`WITH CTE AS (
		SELECT *, 1 RecursiveCallNumber FROM transactions WHERE id
			IN (SELECT id FROM transactions WHERE id = %d LIMIT 1)
		UNION ALL
		SELECT  t.*, RecursiveCallNumber+1 RecursiveCallNumber FROM transactions t
		INNER JOIN CTE ON t.parentId=CTE.id)
		DELETE FROM transactions WHERE id IN (SELECT id FROM CTE)`, id)

	db := database.Connect()
	_, err := db.Exec(query)
	return err
}

func (this *Dao) Create(raw Raw) (error, int64) {
	db := database.Connect()
	result, err := db.NamedExec(`
		INSERT INTO transactions (type, createdAt, description, amount, tags, parentId)
		VALUES (:type, :createdAt, :description, :amount, :tags, :parentId)`, raw)
	id := int64(0)
	if err != nil {
		return err, id
	}
	id, err = result.LastInsertId()
	return err, int64(id)
}

func (this *Dao) Update(raw Raw) (error, int64) {
	db := database.Connect()
	_, err := db.NamedExec(`
		UPDATE transactions SET type=:type, createdAt=:createdAt, description=:description, amount=:amount, tags=:tags, parentId=:parentId
		WHERE id=:id`, raw)
	return err, raw.Id
}

func (this *Dao) getOrDefault (dict map[string]interface{}, key string, defaultValue interface{}) interface{} {
	var result = dict[key]
	if result == nil {
		result = defaultValue
	}
	return result
}