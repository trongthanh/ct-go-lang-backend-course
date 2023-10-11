package usecase

import (
	"context"
	"gosocial/internal/entity"
)

func (uc *ucImplement) GetProfiles(ctx context.Context, req *entity.ProfilesRequest) (*entity.ProfilesResponse, error) {

	profileDocs, err := uc.profileStore.GetMany()

	if err != nil {
		return nil, err
	}

	var profiles []entity.Profile
	for _, profileDoc := range profileDocs {
		profiles = append(profiles, profileDoc.ToProfile())
	}

	return &entity.ProfilesResponse{
		Profiles: profiles,
	}, nil
}
