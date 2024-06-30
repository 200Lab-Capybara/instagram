package userusecase

import (
	"context"
	"errors"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
	"instagram/components/hasher"
	"instagram/components/tokenprovider"
)

type loginUseCase struct {
	loginRepo LoginRepository
	hasher    hasher.Hasher
	accPro    tokenprovider.Provider
}

func NewLoginUseCase(loginRepo LoginRepository, hasher hasher.Hasher, provider tokenprovider.Provider) LoginUseCase {
	return &loginUseCase{
		loginRepo: loginRepo,
		hasher:    hasher,
		accPro:    provider,
	}
}

type LoginUseCase interface {
	Execute(ctx context.Context, user *usermodel.UserLogin) (any, error)
}

type LoginRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
}

func (l *loginUseCase) Execute(ctx context.Context, input *usermodel.UserLogin) (any, error) {
	user, err := l.loginRepo.FindUserByEmail(ctx, input.Email)

	if err != nil {
		if errors.Is(err, usermodel.ErrUserNotFound) || user.Status == common.UserDeleted {
			return nil, common.NewCustomError(usermodel.ErrInvalidEmailOrPass, usermodel.ErrInvalidEmailOrPass.Error(), "invalid_email_or_password")
		}

		return nil, err
	}

	if user.Status == common.UserBanned {
		return nil, common.NewCustomError(usermodel.ErrUserBanded, usermodel.ErrUserBanded.Error(), "user_banded")
	}

	// Compare the password
	if !l.hasher.Verify(input.Password, user.Salt, user.Password) {
		return nil, common.NewCustomError(usermodel.ErrInvalidEmailOrPass, usermodel.ErrInvalidEmailOrPass.Error(), "invalid_email_or_password")
	}

	tokenPayload := common.TokenPayload{
		Id: user.ID,
	}

	// TODO: Implement JWT token generation
	accessToken, err := l.accPro.GenerateToken(&tokenPayload, common.AccessTokenExpireDuration)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
