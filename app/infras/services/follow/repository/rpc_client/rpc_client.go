package followrpc_client

import (
	"context"
	"github.com/google/uuid"
	"instagram/common"
	pb "instagram/proto"
)

type rpcClient struct {
	client pb.UserServiceClient
}

func NewFollowRpcClient(client pb.UserServiceClient) *rpcClient {
	return &rpcClient{
		client: client,
	}
}

func (rpc *rpcClient) GetUserByIds(ctx context.Context, ids []uuid.UUID) ([]common.SimpleUser, error) {
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
