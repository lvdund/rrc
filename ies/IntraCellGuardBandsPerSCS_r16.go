package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IntraCellGuardBandsPerSCS_r16 struct {
	GuardBandSCS_r16        SubcarrierSpacing `madatory`
	IntraCellGuardBands_r16 []GuardBand_r16   `lb:1,ub:4,madatory`
}

func (ie *IntraCellGuardBandsPerSCS_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.GuardBandSCS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode GuardBandSCS_r16", err)
	}
	tmp_IntraCellGuardBands_r16 := utils.NewSequence[*GuardBand_r16]([]*GuardBand_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
	for _, i := range ie.IntraCellGuardBands_r16 {
		tmp_IntraCellGuardBands_r16.Value = append(tmp_IntraCellGuardBands_r16.Value, &i)
	}
	if err = tmp_IntraCellGuardBands_r16.Encode(w); err != nil {
		return utils.WrapError("Encode IntraCellGuardBands_r16", err)
	}
	return nil
}

func (ie *IntraCellGuardBandsPerSCS_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.GuardBandSCS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode GuardBandSCS_r16", err)
	}
	tmp_IntraCellGuardBands_r16 := utils.NewSequence[*GuardBand_r16]([]*GuardBand_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
	fn_IntraCellGuardBands_r16 := func() *GuardBand_r16 {
		return new(GuardBand_r16)
	}
	if err = tmp_IntraCellGuardBands_r16.Decode(r, fn_IntraCellGuardBands_r16); err != nil {
		return utils.WrapError("Decode IntraCellGuardBands_r16", err)
	}
	ie.IntraCellGuardBands_r16 = []GuardBand_r16{}
	for _, i := range tmp_IntraCellGuardBands_r16.Value {
		ie.IntraCellGuardBands_r16 = append(ie.IntraCellGuardBands_r16, *i)
	}
	return nil
}
