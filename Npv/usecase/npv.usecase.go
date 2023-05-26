package usecase

import (
	"npv-application/entity"
	"npv-application/npv"

	"github.com/gin-gonic/gin"
)

func NewUsecaseNpv(repo npv.RepoNpv) npvUsecase {
	return npvUsecase{
		repoNpv: repo,
	}
}

type npvUsecase struct {
	repoNpv npv.RepoNpv
}

func (usecase npvUsecase) PostNpv(ctx *gin.Context) error {
	var input entity.Input

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return err
	}

	if err := usecase.repoNpv.PostNpv(input); err != nil {
		return err
	}

	return nil
}

func (usecase npvUsecase) GetNpv(ctx *gin.Context) (entity.NPV, error) {
	res, err := usecase.repoNpv.GetNpv()

	if err != nil {
		return entity.NPV{}, err
	}

	return res, nil
}
