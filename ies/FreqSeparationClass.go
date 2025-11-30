package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqSeparationClass_Enum_mhz800       aper.Enumerated = 0
	FreqSeparationClass_Enum_mhz1200      aper.Enumerated = 1
	FreqSeparationClass_Enum_mhz1400      aper.Enumerated = 2
	FreqSeparationClass_Enum_mhz400_v1650 aper.Enumerated = 3
	FreqSeparationClass_Enum_mhz600_v1650 aper.Enumerated = 4
)

type FreqSeparationClass struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FreqSeparationClass) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FreqSeparationClass", err)
	}
	return nil
}

func (ie *FreqSeparationClass) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FreqSeparationClass", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
