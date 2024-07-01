package usecaseprofile

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/profile/model"
	"instagram/common"
	"time"
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

	//formDate := "2000-Jan-02"
	t, err := time.Parse("Jan 2, 2006", profile.DateOfBirth)
	if err != nil {
		return nil, err
	}
	defaultGender := model.Male
	data := &model.Profile{
		ID:             id,
		DateOfBirth:    t,
		Gender:         &defaultGender,
		Avatar:         profile.Avatar,
		CountFollowers: 0,
		CountFollowing: 0,
		CountPosts:     0,
		UserId:         common.User1UUID,
	}

	_, err = p.setupRepository.InsertProfile(ctx, data)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
