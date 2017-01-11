package transactions

import (
	"fmt"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
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

func (*Helper) GetPageParameters(ctx *gin.Context) map[string]interface{} {
	var page, perPage = 1, 1
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	perPage, err = strconv.Atoi(ctx.Query("perPage"))
	if err != nil || perPage < 1 {
		perPage = 1
	}
	sort := ctx.Query("sort")
	if sort == "" {
		sort = "id"
	}

	return map[string]interface{}{
		"page": page,
		"perPage": perPage,
		"sort": sort,
	}
}
