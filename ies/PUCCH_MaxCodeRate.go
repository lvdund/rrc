package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_MaxCodeRate_Enum_zeroDot08 aper.Enumerated = 0
	PUCCH_MaxCodeRate_Enum_zeroDot15 aper.Enumerated = 1
	PUCCH_MaxCodeRate_Enum_zeroDot25 aper.Enumerated = 2
	PUCCH_MaxCodeRate_Enum_zeroDot35 aper.Enumerated = 3
	PUCCH_MaxCodeRate_Enum_zeroDot45 aper.Enumerated = 4
	PUCCH_MaxCodeRate_Enum_zeroDot60 aper.Enumerated = 5
	PUCCH_MaxCodeRate_Enum_zeroDot80 aper.Enumerated = 6
)

type PUCCH_MaxCodeRate struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *PUCCH_MaxCodeRate) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode PUCCH_MaxCodeRate", err)
	}
	return nil
}

func (ie *PUCCH_MaxCodeRate) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode PUCCH_MaxCodeRate", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
