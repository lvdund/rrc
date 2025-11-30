package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_FormatConfig_additionalDMRS_Enum_true aper.Enumerated = 0
)

type PUCCH_FormatConfig_additionalDMRS struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PUCCH_FormatConfig_additionalDMRS) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PUCCH_FormatConfig_additionalDMRS", err)
	}
	return nil
}

func (ie *PUCCH_FormatConfig_additionalDMRS) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PUCCH_FormatConfig_additionalDMRS", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
