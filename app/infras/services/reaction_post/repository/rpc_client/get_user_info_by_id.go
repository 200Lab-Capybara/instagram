package rpc_client

import (
	"context"
	"github.com/google/uuid"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

type getUserInfoRepo struct {
	db common.SQLDatabase
}

func NewGetUserInfoRepo(db common.SQLDatabase) *getUserInfoRepo {
	return &getUserInfoRepo{db: db}
}

func (s *getUserInfoRepo) GetUserInfoById(ctx context.Context, userIDs []uuid.UUID) ([]usermodel.User, error) {
	var listInfo []usermodel.User

	db := s.db.GetConnection().Table(usermodel.User{}.TableName())

	if err := db.Where("id IN ?", userIDs).First(&listInfo).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return listInfo, nil
}
