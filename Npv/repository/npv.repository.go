package repository

import (
	"npv-application/entity"
	"npv-application/helper"
)

func NewRepoNpv(npv entity.NPV) ennpv {
	return ennpv{
		enNpv: npv,
	}
}

type ennpv struct {
	enNpv entity.NPV
}

func (repo *ennpv) PostNpv(input entity.Input) error {
	SumNpv, Pv := helper.Pv(input.Cicilan, input.Persen)
	Irr := helper.Irr(input.Cicilan, input.ModalUsaha, input.Dp)
	repo.enNpv.Npv = nil
	repo.enNpv.Pv = nil
	repo.enNpv.IRR = 0
	for i := 0; i < len(SumNpv); i++ {
		repo.enNpv.Npv = append(repo.enNpv.Npv, SumNpv[i]-(input.ModalUsaha-input.Dp))
	}

	for i := 0; i < len(Pv); i++ {
		repo.enNpv.Pv = append(repo.enNpv.Pv, Pv[i])
	}

	repo.enNpv.IRR = Irr

	return nil
}

func (repo *ennpv) GetNpv() (entity.NPV, error) {
	return repo.enNpv, nil
}
