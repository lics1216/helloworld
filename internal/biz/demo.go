package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Demo struct {
	Id   int64
	Name string
}

// DemoRepo 创建一个DemoRepo 是为了
type DemoRepo interface {
	Save(context.Context, *Demo) (*Demo, error)
}

// DemoUsecase 用来干嘛
type DemoUsecase struct {
	repo DemoRepo
	log *log.Helper
}

// NewDemoUsecase 这个是供wire 调用的吗？
// 第一个参数可以传入DemoRepo 的实现类型，data/demo.go 里面的type
func NewDemoUsecase(repo DemoRepo, logger log.Logger)  *DemoUsecase{
	return &DemoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// 创建一个CreatDemo 方法供service 调用
func (uc *DemoUsecase) CreateDemo(ctx context.Context, d *Demo)(*Demo, error){
	uc.log.WithContext(ctx).Infof("CreateDemo: %v", d.Name)
	// Save 是被谁实例化的，谁implement DemoRepo 接口
	return uc.repo.Save(ctx, d)
}