package fake

type Service struct {}

func (*Service) ByTag(params map[string]interface{}) (error, []map[string]interface{}) {
	result := make([]map[string]interface{}, 1)
	result = append(result, map[string]interface{}{
		"month": "2016-01",
		"amount": 50.2,
		"total": 100.0,
	})
	result = append(result, map[string]interface{}{
		"month": "2016-02",
		"amount": 320.2,
		"total": 890.0,
	})
	return nil, result
}
