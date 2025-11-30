package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_FormatConfig_nrofSlots_Enum_n2 aper.Enumerated = 0
	PUCCH_FormatConfig_nrofSlots_Enum_n4 aper.Enumerated = 1
	PUCCH_FormatConfig_nrofSlots_Enum_n8 aper.Enumerated = 2
)

type PUCCH_FormatConfig_nrofSlots struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUCCH_FormatConfig_nrofSlots) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUCCH_FormatConfig_nrofSlots", err)
	}
	return nil
}

func (ie *PUCCH_FormatConfig_nrofSlots) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUCCH_FormatConfig_nrofSlots", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
