package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530_Enum_ms3 aper.Enumerated = 0
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530_Enum_ms4 aper.Enumerated = 1
)

type TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530", err)
	}
	return nil
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
