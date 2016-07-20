package reports

import (
	"fmt"
	"github.com/gustavohenrique/poupaniquel/api/database"
)

type Dao struct {}

func NewDao() *Dao {
	return &Dao{}
}

func (this *Dao) ByTag(params map[string]interface{}) (err error, result []map[string]interface{}) {
	db := database.Connect()
	tag := params["tag"].(string)
	like := "tags LIKE '%|" + tag + "|%'"
	dateFormat := "\"%Y-%m\""
	query := fmt.Sprintf(`SELECT ROUND(TOTAL(amount), 2) AS amount, strftime(%s, createdAt) AS date,
		ROUND((SELECT TOTAL(amount) FROM transactions WHERE type = :type AND createdAt >= :startDate AND createdAt <= :endDate), 2) AS total
	  	FROM transactions
	 	WHERE %s AND type = :type
	   	AND createdAt >= :startDate and createdAt <= :endDate
	 	GROUP BY strftime(%s, createdAt)`, dateFormat, like, dateFormat)

	rows, err := db.NamedQuery(query, params)
	for rows.Next() {
		var amount float32
		var month string
		var total float32
		err = rows.Scan(&amount, &month, &total)
		if err != nil {
			return err, result
		}
		result = append(result, map[string]interface{}{
			"month": month,
			"amount": amount,
			"total": total,
		})
	}
	return err, result
}
