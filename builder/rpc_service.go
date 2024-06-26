package builder

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	usermysql "instagram/app/infras/services/user/repository/mysql"
	userrpc "instagram/app/infras/services/user/transport/rpc"
	userusecase "instagram/app/internals/services/user/usecase"
	pb "instagram/proto"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	rpcPort = os.Getenv("RPC_PORT")
)

func BuildRpcService(ctxSvr ServiceContext) {
	rpcPort, err := strconv.Atoi(rpcPort)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", rpcPort))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("GRPC Server is listening on %d ...", rpcPort))
	s := grpc.NewServer()

	userStorage := usermysql.NewMySQLStorage(ctxSvr.GetDB())
	findUserByIdUC := userusecase.NewGetUserByIdUseCase(userStorage)
	pb.RegisterUserServiceServer(s, userrpc.NewUserRpcHandler(findUserByIdUC))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func BuildUserRpcClient() pb.UserServiceClient {
	rpcPort, err := strconv.Atoi(rpcPort)
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", rpcPort), opts)

	if err != nil {
		log.Fatal(err)
	}

	return pb.NewUserServiceClient(clientConn)
}
