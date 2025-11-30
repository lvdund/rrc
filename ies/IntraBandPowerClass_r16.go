package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	IntraBandPowerClass_r16_Enum_pc2    aper.Enumerated = 0
	IntraBandPowerClass_r16_Enum_pc3    aper.Enumerated = 1
	IntraBandPowerClass_r16_Enum_spare6 aper.Enumerated = 2
	IntraBandPowerClass_r16_Enum_spare5 aper.Enumerated = 3
	IntraBandPowerClass_r16_Enum_spare4 aper.Enumerated = 4
	IntraBandPowerClass_r16_Enum_spare3 aper.Enumerated = 5
	IntraBandPowerClass_r16_Enum_spare2 aper.Enumerated = 6
	IntraBandPowerClass_r16_Enum_spare1 aper.Enumerated = 7
)

type IntraBandPowerClass_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *IntraBandPowerClass_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode IntraBandPowerClass_r16", err)
	}
	return nil
}

func (ie *IntraBandPowerClass_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode IntraBandPowerClass_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
