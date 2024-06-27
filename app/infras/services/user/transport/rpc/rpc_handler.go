package userrpc

import (
	"context"
	"github.com/google/uuid"
	userusecase "instagram/app/internals/services/user/usecase"
	pb "instagram/proto"
)

type userRpcService struct {
	pb.UnimplementedUserServiceServer
	getUserByIdUC  userusecase.GetUserByIdUseCase
	getUserByIdsUC userusecase.GetUserByIdsUseCase
}

func NewUserRpcHandler(getUserByIdUC userusecase.GetUserByIdUseCase, getUserByIdsUC userusecase.GetUserByIdsUseCase) *userRpcService {
	return &userRpcService{
		getUserByIdUC:  getUserByIdUC,
		getUserByIdsUC: getUserByIdsUC,
	}
}

func (rpc *userRpcService) GetUserById(ctx context.Context, dto *pb.GetUserByIdReq) (*pb.PublicUserInfoResp, error) {
	id := uuid.MustParse(dto.Id)
	user, err := rpc.getUserByIdUC.Execute(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.PublicUserInfoResp{
		User: &pb.PublicUserInfo{
			Id:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role.String(),
			Status:    user.Status.String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}, nil
}

func (rpc *userRpcService) GetUserByIds(ctx context.Context, dto *pb.GetUserByIdsReq) (*pb.PublicUsersInfosResp, error) {
	ids := make([]uuid.UUID, len(dto.Ids))
	// Convert string ids to uuid
	for i, id := range dto.Ids {
		ids[i] = uuid.MustParse(id)
	}

	users, err := rpc.getUserByIdsUC.Execute(ctx, ids)
	if err != nil {
		return nil, err
	}

	publicUsersInfo := make([]*pb.PublicUserInfo, len(users))

	for i, user := range users {
		publicUsersInfo[i] = &pb.PublicUserInfo{
			Id:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role.String(),
			Status:    user.Status.String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		}
	}

	return &pb.PublicUsersInfosResp{Users: publicUsersInfo}, nil
}
