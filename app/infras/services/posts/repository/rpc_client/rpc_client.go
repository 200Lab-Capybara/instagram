package postrpc_client

import (
	"context"
	"github.com/google/uuid"
	"instagram/common"
	pb "instagram/proto"
)

type postRpcClient struct {
	client pb.UserServiceClient
}

func NewPostRpcClient(client pb.UserServiceClient) *postRpcClient {
	return &postRpcClient{
		client: client,
	}
}

func (rpc *postRpcClient) GetUserById(ctx context.Context, id uuid.UUID) (*common.SimpleUser, error) {
	res, err := rpc.client.GetUserById(ctx, &pb.GetUserByIdReq{Id: id.String()})

	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	user := common.SimpleUser{UserId: uuid.MustParse(res.User.Id), FirstName: res.User.FirstName, LastName: res.User.LastName, Role: res.User.Role, Status: res.User.Status}
	return &user, nil
}
