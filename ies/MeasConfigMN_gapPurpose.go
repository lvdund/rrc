package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasConfigMN_gapPurpose_Enum_perUE  aper.Enumerated = 0
	MeasConfigMN_gapPurpose_Enum_perFR1 aper.Enumerated = 1
)

type MeasConfigMN_gapPurpose struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MeasConfigMN_gapPurpose) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MeasConfigMN_gapPurpose", err)
	}
	return nil
}

func (ie *MeasConfigMN_gapPurpose) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MeasConfigMN_gapPurpose", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
