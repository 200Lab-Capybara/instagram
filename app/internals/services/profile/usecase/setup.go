package usecase

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/profile/model"
)

type SetupRepository interface {
	SetupNewProfile(ctx context.Context, profile *model.Profile) (uuid.UUID, error)
}

type setupUseCase struct {
	setupRepository SetupRepository
}

// TODO compare setup vs register
func NewSetupUseCase(setuprepository SetupRepository) SetupUseCase {
	return nil
}

type SetupUseCase interface {
	Execute(ctx context.Context, profile *model.ProfileCreation) (*uuid.UUID, error)
}

// TODO compare setup vs register
func (p *setupUseCase) Execute(ctx context.Context, profile *model.ProfileCreation) (*uuid.UUID, error) {
	id, _ := uuid.NewV7()

	data := &model.Profile{
		ID:          id,
		DateOfBirth: profile.DateOfBirth,
		//TODO comongender?? or model.ProfileGender
		Genders: model.ProfileGender(),
		Avatar:  profile.Avatar,
		//TODO how to count ???
		CountFollowers: 0,
		CountFollowing: 0,
		CountPosts:     0,
		//TODO how to take userId
		UserId: 01902f82-a74f-7042-8b3b-b9c7c57c10c7,
	}

	_, err := p.setupRepository.SetupNewProfile(ctx, data)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
