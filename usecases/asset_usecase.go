package usecase

import (
	"context"
	"fmt"
	"testjavan/helpers/errs"
	"testjavan/model"
	"testjavan/repositories"
	"time"
)

type AssetUsecase interface {
	GetAssetByID(ctx context.Context, id int) (model.Asset, error)
	UpdateAssetByID(ctx context.Context, req model.AssetRequest, id int) error
	CreateAsset(ctx context.Context, req model.AssetRequest) error
	DeleteAssetByID(ctx context.Context, id int) error
	AddAssetToFamilyMember(ctx context.Context, id int, assetName string) error
	RemoveAssetFromFamilyMember(ctx context.Context, familyMemberID, assetID int) error
}

type asset struct {
	repo *repositories.Repository
}

func newAssetUsecase(repo *repositories.Repository) AssetUsecase {
	return &asset{repo: repo}
}

func (a *asset) GetAssetByID(ctx context.Context, id int) (model.Asset, error) {
	return a.repo.Asset.GetAssetByID(ctx, id)
}

func (a *asset) UpdateAssetByID(ctx context.Context, req model.AssetRequest, id int) error {
	return a.repo.Asset.UpdateAssetByID(ctx, model.Asset{
		AssetID:   id,
		AssetName: req.AssetName,
	})
}

func (a *asset) CreateAsset(ctx context.Context, req model.AssetRequest) error {
	_, err := a.repo.Asset.CreateAsset(ctx, model.Asset{
		AssetName: req.AssetName,
	})
	return err
}

func (a *asset) DeleteAssetByID(ctx context.Context, id int) error {
	now := time.Now()
	return a.repo.Asset.UpdateAssetByID(ctx, model.Asset{
		AssetID:   id,
		DeletedAt: &now,
	})
}

func (a *asset) AddAssetToFamilyMember(ctx context.Context, familyMemberID int, assetName string) error {
	var (
		err   error
		asset model.Asset
	)

	_, err = a.repo.Family.GetFamilyMemberByID(ctx, familyMemberID)
	if err != nil {
		return err
	}

	asset, err = a.repo.Asset.GetAssetByAssetName(ctx, assetName)
	if err != nil && err != errs.ErrRecordNotfound {
		return err
	}

	if err != nil && err == errs.ErrRecordNotfound {
		// Create new assets
		asset, err = a.repo.Asset.CreateAsset(ctx, model.Asset{AssetName: assetName})
		if err != nil {
			return err
		}

		fmt.Println(asset.AssetID)
	}

	err = a.repo.FamilyAsset.AddAssetToFamilyMemberByMemberID(ctx, model.MemberAsset{
		MemberID: familyMemberID,
		AssetID:  asset.AssetID,
	})

	return err
}

func (a *asset) RemoveAssetFromFamilyMember(ctx context.Context, familyMemberID, assetID int) error {
	now := time.Now()
	return a.repo.FamilyAsset.UpdateFamilyAssetByMemberIDandAssetID(ctx, model.MemberAsset{
		MemberID:  familyMemberID,
		AssetID:   assetID,
		DeletedAt: &now,
	})
}
