package reports

type Service struct {}

var dao = NewDao()

func NewService() *Service {
	return &Service{}
}

func (this *Service) ByTags(params map[string]interface{}) (error, []map[string]interface{}) {
	return dao.ByTags(params)
}