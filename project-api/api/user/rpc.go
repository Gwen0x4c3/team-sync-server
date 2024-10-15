package user

import (
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	loginServiceV1 "github.com/Gwen0x4c3/team-sync-server/project-user/pkg/service/login.service.v1"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var loginServiceClient loginServiceV1.LoginServiceClient

func InitRpcClient() {
	conn, err := grpc.NewClient("127.0.0.1:8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.LG.Error("InitRpcClient error: %v\n", err)
		panic(err)
	}
	loginServiceClient = loginServiceV1.NewLoginServiceClient(conn)
}
