package reports

import (
	"fmt"

	"github.com/gustavohenrique/poupaniquel/api/database"
)

type Dao struct{}

func NewDao() *Dao {
	return &Dao{}
}

func (this *Dao) ByTag(params map[string]interface{}) (err error, result []map[string]interface{}) {
	db := database.Connect()
	tag := params["tag"].(string)
	like := "tags LIKE '%|" + tag + "|%'"
	dateFormat := "\"%Y-%m\""
	query := fmt.Sprintf(`SELECT strftime(%s, dueDate) AS date, ROUND(TOTAL(amount), 2) AS amount,
		ROUND((SELECT TOTAL(amount) FROM transactions WHERE type = :type AND dueDate >= :startDate AND dueDate <= :endDate AND (parentId=0 OR parentId IS NULL)), 2) AS total
	  	FROM transactions
	 	WHERE %s AND type = :type
	   	AND dueDate >= :startDate and dueDate <= :endDate
	 	GROUP BY strftime(%s, dueDate)`, dateFormat, like, dateFormat)

	rows, err := db.NamedQuery(query, params)
	for rows.Next() {
		var month string
		var amount float64
		var total float64
		err = rows.Scan(&month, &amount, &total)
		if err != nil {
			return err, result
		}
		result = append(result, map[string]interface{}{
			"month":  month,
			"amount": amount,
			"total":  total,
		})
	}
	return err, result
}
