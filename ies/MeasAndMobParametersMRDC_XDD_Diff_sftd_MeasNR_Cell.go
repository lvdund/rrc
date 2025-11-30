package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell_Enum_supported aper.Enumerated = 0
)

type MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell", err)
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
