package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX_Enum_supported aper.Enumerated = 0
)

type MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX", err)
	}
	return nil
}

func (ie *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
