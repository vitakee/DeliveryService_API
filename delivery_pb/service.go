package delivery_pb

import "vita.com/cafeshop"
import "vita.com/customer"
import (
	"vita.com/tuktuk"
	"net"
	"google.golang.org/grpc"
	pb "vita.com/grpc"
)

func Start(port string)error{
	lis,err:=net.Listen("tcp", ":"+port)
	if err!=nil{
		return err
	}

	grpcServer := grpc.NewServer()

	cafeshopService,err:=cafeshop.NewV1()
	tuktukService,err:=tuktuk.NewV1()
	customerService,err:=customer.NewV1()

	pb.RegisterCafeShopServer(grpcServer, cafeshopService)
	pb.RegisterTukTukServer(grpcServer, tuktukService)
	pb.RegisterCustomerServer(grpcServer,customerService)

	grpcServer.Serve(lis)

	return nil
}
