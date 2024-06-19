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
	return nil, nil
}
