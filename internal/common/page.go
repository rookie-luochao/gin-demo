package common

import (
	"github.com/go-courier/sqlx/v2/builder"
)

// PageParam 分页参数
type PageParam struct {
	PageSize   int64 `name:"pageSize,omitempty" in:"query"`   // 每页数量
	PageOffset int64 `name:"pageOffset,omitempty" in:"query"` // 页数1开始
}

func (c *PageParam) Offset() int64 {
	if c.PageOffset == 0 {
		return c.PageOffset
	}
	return c.PageOffset - 1
}

func (p *PageParam) PageAddition() builder.Addition {
	if p.PageSize != 0 {
		return builder.Limit(p.PageSize).Offset(p.Offset() * p.PageSize)
	}
	return builder.Limit(10).Offset(0)
}

func (p *PageParam) PageAll() builder.Addition {
	return builder.Limit(-1).Offset(-1)
}
