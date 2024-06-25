package userrpc

import (
	"context"
	"github.com/google/uuid"
	userusecase "instagram/app/internals/services/user/usecase"
	pb "instagram/proto"
)

type userRpcService struct {
	pb.UnimplementedUserServiceServer
	getUserByIdUC userusecase.GetUserByIdUseCase
}

func NewUserRpcHandler(getUserByIdUC userusecase.GetUserByIdUseCase) *userRpcService {
	return &userRpcService{getUserByIdUC: getUserByIdUC}
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
