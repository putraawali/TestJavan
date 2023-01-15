package usecase

import (
	"context"
	"sync"
	"testjavan/helpers/errs"
	"testjavan/model"
	"testjavan/repositories"
)

type FamilyUsecase interface {
	GetFamilyMemberByID(ctx context.Context, id int) (model.FamilyAssets, error)
	UpdateFamilyMemberByID(ctx context.Context, req model.FamilyRequest, id int) error
	RemoveFamilyMemberByID(ctx context.Context, id int) error
}

type family struct {
	repo *repositories.Repository
}

func newFamilyUsecase(repo *repositories.Repository) FamilyUsecase {
	return &family{repo: repo}
}

func (f *family) GetFamilyMemberByID(ctx context.Context, id int) (model.FamilyAssets, error) {
	var (
		result         model.FamilyAssets
		familyMember   model.Family
		assets         []model.Asset
		assetWithPrice []model.AssetWithPrice
		err, errAsset  error
	)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		familyMember, err = f.repo.Family.GetFamilyMemberByID(ctx, id)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		assets, errAsset = f.repo.Asset.GetAssetMemberByID(ctx, id)
		if errAsset == nil {
			products, errAsset := f.repo.Product.GetAllProduct(ctx)
			if errAsset == nil {
				ch := make(chan bool, 30)
				wg.Add(len(assets))

				for _, asset := range assets {
					ch <- true

					go func(asset model.Asset) {
						defer func() {
							<-ch
							wg.Done()
						}()

						isFound := false
						for _, product := range products.Products {
							if asset.AssetName == product.Title {
								isFound = true
								assetWithPrice = append(assetWithPrice, model.AssetWithPrice{
									Price:     product.Price,
									AssetID:   asset.AssetID,
									AssetName: asset.AssetName,
									CreatedAt: asset.CreatedAt,
									UpdatedAt: asset.UpdatedAt,
									DeletedAt: asset.DeletedAt,
								})
								break
							}
						}

						if !isFound {
							// Append without price if not found from API
							assetWithPrice = append(assetWithPrice, model.AssetWithPrice{
								AssetID:   asset.AssetID,
								AssetName: asset.AssetName,
								CreatedAt: asset.CreatedAt,
								UpdatedAt: asset.UpdatedAt,
								DeletedAt: asset.DeletedAt,
							})
						}
					}(asset)
				}
			}
		}
	}()

	wg.Wait()

	if err != nil {
		return result, err
	}

	if errAsset != nil && errAsset != errs.ErrRecordNotfound {
		return result, err
	}

	result.MemberID = familyMember.MemberID
	result.MemberName = familyMember.MemberName
	result.Gender = familyMember.Gender
	result.CreatedAt = familyMember.CreatedAt
	result.UpdatedAt = familyMember.UpdatedAt
	result.DeletedAt = familyMember.DeletedAt
	result.Assets = assetWithPrice

	return result, err
}

func (f *family) UpdateFamilyMemberByID(ctx context.Context, req model.FamilyRequest, id int) error {
	return f.repo.Family.UpdateFamilyMemberByID(ctx, model.Family{
		MemberID:   id,
		MemberName: req.MemberName,
		Gender:     req.Gender,
	})
}

func (f *family) RemoveFamilyMemberByID(ctx context.Context, id int) error {
	return f.repo.Family.RemoveFamilyMemberByID(ctx, id)
}
