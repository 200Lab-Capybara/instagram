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

func (rpc *postRpcClient) GetUserByIds(ctx context.Context, ids []uuid.UUID) ([]common.SimpleUser, error) {
	idsStr := make([]string, len(ids))
	for i, id := range ids {
		idsStr[i] = id.String()
	}
	res, err := rpc.client.GetUserByIds(ctx, &pb.GetUserByIdsReq{Ids: idsStr})

	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	listSimpleUser := make([]common.SimpleUser, len(res.Users))

	for i, user := range res.Users {
		listSimpleUser[i] = common.SimpleUser{
			UserId:    uuid.MustParse(user.Id),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			Status:    user.Status,
		}
	}

	return listSimpleUser, nil
}
