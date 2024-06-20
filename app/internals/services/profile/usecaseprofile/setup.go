package usecaseprofile

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/profile/model"
	"instagram/common"
)

type SetupRepository interface {
	InsertProfile(ctx context.Context, profile *model.Profile) (*uuid.UUID, error)
}

type setupUseCase struct {
	setupRepository SetupRepository
}

func NewSetupUseCase(setuprepository SetupRepository) SetupUseCase {
	return &setupUseCase{
		setupRepository: setuprepository,
	}
}

type SetupUseCase interface {
	Execute(ctx context.Context, profile *model.ProfileCreation) (*uuid.UUID, error)
}

func (p *setupUseCase) Execute(ctx context.Context, profile *model.ProfileCreation) (*uuid.UUID, error) {
	id, _ := uuid.NewV7()

	defaultGender := model.Male
	data := &model.Profile{
		ID:             id,
		DateOfBirth:    profile.DateOfBirth,
		Genders:        &defaultGender,
		Avatar:         profile.Avatar,
		CountFollowers: 0,
		CountFollowing: 0,
		CountPosts:     0,
		UserId:         common.User1UUID,
	}

	_, err := p.setupRepository.InsertProfile(ctx, data)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
