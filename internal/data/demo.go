package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type demoRepo struct {
	data *Data
	log *log.Helper
}

// NewDemoRepo 提供一个方法供自动化调用吗？返回个type
func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	// 返回指针是因为
	return  &demoRepo{
		data: data,
		log:log.NewHelper(logger),
	}
}

func (r *demoRepo) Save(ctx context.Context, d *biz.Demo) (*biz.Demo, error) {
	// @todo 把传来的数据insert 到表
	return &biz.Demo{Id: 34, Name: d.Name}, nil
}