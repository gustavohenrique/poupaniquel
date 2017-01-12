package reports

type Reporter interface {
	ByTag(map[string]interface{}) (error, []map[string]interface{})
}

type Service struct{}

var dao *Dao

func NewService(d *Dao) Reporter {
	dao = d
	return &Service{}
}

func (*Service) ByTag(params map[string]interface{}) (error, []map[string]interface{}) {
	return dao.ByTag(params)
}
