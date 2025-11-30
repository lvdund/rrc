package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms0p5   aper.Enumerated = 0
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms0p625 aper.Enumerated = 1
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms1     aper.Enumerated = 2
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms1p25  aper.Enumerated = 3
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms2     aper.Enumerated = 4
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms2p5   aper.Enumerated = 5
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms5     aper.Enumerated = 6
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms10    aper.Enumerated = 7
)

type TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity", err)
	}
	return nil
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
