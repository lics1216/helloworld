package service

import (
	"context"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld"
)

type DemoService struct {
	pb.UnimplementedDemoServer
	// biz 层里面定义的 usecase
	uc *biz.DemoUsecase
}

// NewDemoService 是供wire 调用的？
func NewDemoService(uc *biz.DemoUsecase) *DemoService {
	return &DemoService{uc: uc}
}

// CreateDemo
func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	// 从Request 从取出参数name, 调用biz 层的方法CreateDemo 对象，再返回对象demo
	// 确定field 的可见范围是用变量首字母大写?，写成name 会报错no export field
	d, err := s.uc.CreateDemo(ctx, &biz.Demo{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &pb.CreateDemoReply{Id: d.Id, Name: d.Name}, nil
}
