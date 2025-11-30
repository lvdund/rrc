package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf8   aper.Enumerated = 0
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf16  aper.Enumerated = 1
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf32  aper.Enumerated = 2
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf64  aper.Enumerated = 3
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf128 aper.Enumerated = 4
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf256 aper.Enumerated = 5
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf512 aper.Enumerated = 6
)

type PosSchedulingInfo_r16_posSI_Periodicity_r16 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *PosSchedulingInfo_r16_posSI_Periodicity_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode PosSchedulingInfo_r16_posSI_Periodicity_r16", err)
	}
	return nil
}

func (ie *PosSchedulingInfo_r16_posSI_Periodicity_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode PosSchedulingInfo_r16_posSI_Periodicity_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
