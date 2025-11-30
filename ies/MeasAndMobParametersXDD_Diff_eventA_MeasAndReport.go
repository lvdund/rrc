package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersXDD_Diff_eventA_MeasAndReport_Enum_supported aper.Enumerated = 0
)

type MeasAndMobParametersXDD_Diff_eventA_MeasAndReport struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersXDD_Diff_eventA_MeasAndReport) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersXDD_Diff_eventA_MeasAndReport", err)
	}
	return nil
}

func (ie *MeasAndMobParametersXDD_Diff_eventA_MeasAndReport) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersXDD_Diff_eventA_MeasAndReport", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
