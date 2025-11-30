package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersFRX_Diff_handoverLTE_EPC_Enum_supported aper.Enumerated = 0
)

type MeasAndMobParametersFRX_Diff_handoverLTE_EPC struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersFRX_Diff_handoverLTE_EPC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersFRX_Diff_handoverLTE_EPC", err)
	}
	return nil
}

func (ie *MeasAndMobParametersFRX_Diff_handoverLTE_EPC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersFRX_Diff_handoverLTE_EPC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
