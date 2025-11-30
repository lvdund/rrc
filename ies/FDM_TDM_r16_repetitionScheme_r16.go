package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FDM_TDM_r16_repetitionScheme_r16_Enum_fdmSchemeA aper.Enumerated = 0
	FDM_TDM_r16_repetitionScheme_r16_Enum_fdmSchemeB aper.Enumerated = 1
	FDM_TDM_r16_repetitionScheme_r16_Enum_tdmSchemeA aper.Enumerated = 2
)

type FDM_TDM_r16_repetitionScheme_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FDM_TDM_r16_repetitionScheme_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FDM_TDM_r16_repetitionScheme_r16", err)
	}
	return nil
}

func (ie *FDM_TDM_r16_repetitionScheme_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FDM_TDM_r16_repetitionScheme_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
