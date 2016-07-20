package transactions

import (
	"fmt"
	"github.com/kataras/iris"
)

type Helper struct {}

func NewHelper() *Helper {
	return &Helper{}
}

func (*Helper) CreateErrorMap(err error) map[string]string {
	return map[string]string{
		"code": "BadRequest",
		"message": fmt.Sprintf("%v", err),
	}
}

func (*Helper) GetPageParameters(ctx *iris.Context) map[string]interface{} {
	var page, perPage = 1, 1
	page, err := ctx.URLParamInt("page")
	if err != nil || page < 1 {
		page = 1
	}
	perPage, err = ctx.URLParamInt("perPage")
	if err != nil || perPage < 1 {
		perPage = 1
	}
	sort := ctx.URLParam("sort")
	if sort == "" {
		sort = "id"
	}
	return map[string]interface{}{
		"page": page,
		"perPage": perPage,
		"sort": sort,
	}
}
