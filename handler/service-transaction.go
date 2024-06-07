package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "service-transaction/proto"
)

type ServiceTransaction struct{}

func (e *ServiceTransaction) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	logger.Infof("Received ServiceTransaction.Call request: %v", req)

	rsp.Data = []*pb.LoginData{
		{
			Token: "",
		},
	}

	rsp.Message = req.Username
	return nil
}

// func (e *ServiceTransaction) ClientStream(ctx context.Context, stream pb.ServiceTransaction_ClientStreamStream) error {
// 	var count int64
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			logger.Infof("Got %v pings total", count)
// 			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		logger.Infof("Got ping %v", req.Stroke)
// 		count++
// 	}
// }

// func (e *ServiceTransaction) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.ServiceTransaction_ServerStreamStream) error {
// 	logger.Infof("Received ServiceTransaction.ServerStream request: %v", req)
// 	for i := 0; i < int(req.Count); i++ {
// 		logger.Infof("Sending %d", i)
// 		if err := stream.Send(&pb.ServerStreamResponse{
// 			Count: int64(i),
// 		}); err != nil {
// 			return err
// 		}
// 		time.Sleep(time.Millisecond * 250)
// 	}
// 	return nil
// }

// func (e *ServiceTransaction) BidiStream(ctx context.Context, stream pb.ServiceTransaction_BidiStreamStream) error {
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		logger.Infof("Got ping %v", req.Stroke)
// 		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
// 			return err
// 		}
// 	}
// }
